package client

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/hashicorp/terraform/helper/resource"
	"golang.org/x/net/proxy"
)

type Client struct {
	username string
	password string
	endpoint string
	Config   *mysql.Config
}

var clientImpl *Client

func initClient(user, passwd, address string) *Client {

	client := &Client{
		username: user,
		password: passwd,
		endpoint: address,
	}

	proto := "tcp"
	conf := mysql.Config{
		User:   user,
		Passwd: passwd,
		Net:    proto,
		Addr:   address,
	}

	dialer := proxy.FromEnvironment()
	mysql.RegisterDial("tcp", func(network string) (net.Conn, error) {
		return dialer.Dial("tcp", network)
	})
	client.Config = &conf

	return client

}

func GetClient(user, passwd, address string) *Client {
	if clientImpl == nil {
		clientImpl = initClient(user, passwd, address)
	}
	return clientImpl
}

func (c *Client) Createdb(process string, name string) error {
	db, err := c.ConnectDB(c.Config)
	if err != nil {
		return err
	}
	var stmt string
	if process == "CREATE" {
		stmt = "CREATE DATABASE " + name
	} else if process == "DELETE" {
		stmt = "DROP DATABASE " + name
	} else if process == "SHOW" {
		stmt = "SHOW DATABASES"
	}

	log.Println("Executing statement: ", stmt)
	reply, err1 := db.Exec(stmt)
	if err1 != nil {
		log.Println("Error exxecuting statement ")
	}
	fmt.Println(reply)
	return nil
}

func (c *Client) Createuser(process string, name string) error {
	db, err := c.ConnectDB(c.Config)
	if err != nil {
		return err
	}

	var stmt string
	if process == "CREATE" {
		stmt = "CREATE USER " + name
	} else if process == "DELETE" {
		stmt = "DROP USER " + name
	}

	log.Println("Executing statement: ", stmt)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Println("Error exxecuting statement ")
	}

	return nil
}

func (c *Client) Creategrant(process, name, database string) error {
	db, err := c.ConnectDB(c.Config)
	if err != nil {
		return err
	}

	var stmt string
	if process == "CREATE" {
		stmt = "GRANT ALL PRIVILEGES ON " + database + ".* TO " + name
	} else if process == "DELETE" {
		stmt = "REVOKE SELECT ON " + database + ".* FROM " + name
	}

	log.Println("Executing statement: ", stmt)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Println("Error exxecuting statement ")
	}

	return nil
}

func (c *Client) Createrole(process, name string) error {
	db, err := c.ConnectDB(c.Config)
	if err != nil {
		return err
	}

	var stmt string
	if process == "CREATE ROLE" {
		stmt = "CREATE ROLE " + name
	}
	if process == "DROP ROLE" {
		stmt = "DROP ROLE IF EXISTS " + name
	}

	log.Println("Executing statement: ", stmt)
	_, err = db.Exec(stmt)
	if err != nil {
		log.Println("Error exxecuting statement ")
	}

	return nil
}

func (c *Client) ConnectDB(conf *mysql.Config) (*sql.DB, error) {
	dsn := conf.FormatDSN()
	var db *sql.DB
	var err error

	retry := resource.Retry(5*time.Second, func() *resource.RetryError {
		db, err = sql.Open("mysql", dsn)

		if err != nil {
			return resource.RetryableError(err)
		}
		err = db.Ping()
		if err != nil {
			return resource.RetryableError(err)
		}
		return nil
	})

	if retry != nil {
		return nil, fmt.Errorf("could not connect to server ", retry)
	}
	return db, nil
}
