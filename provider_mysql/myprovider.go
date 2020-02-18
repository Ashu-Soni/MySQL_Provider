package provider_mysql

import (
	"fmt"
	client "sql/client_mysql"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Myprovider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},

			"username": {
				Type:     schema.TypeString,
				Required: true,
			},

			"password": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"sql_database": myresourceDatabase(),
			"sql_user":     myresourceUser(),
			"sql_grant":    myresourceGrant(),
		},

		ConfigureFunc: providerConfigure,
	}
}
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	conf := Config{
		User:     d.Get("username").(string),
		Password: d.Get("password").(string),
		Address:  d.Get("endpoint").(string),
	}

	if err := conf.Valid(); err != nil {
		return nil, err
	}
	return conf.getClient(), nil
}

func (c Config) getClient() interface{} {
	return client.GetClient(c.User, c.Password, c.Address)
}

func (c Config) Valid() error {
	if c.User == "" {
		return fmt.Errorf("Invalid Username")
	}

	if c.Password == "" {
		return fmt.Errorf("Invalid Password")
	}

	if c.Address == "" {
		return fmt.Errorf("Invalid Address")
	}
	return nil
}

type Config struct {
	User     string
	Password string
	Address  string
}
