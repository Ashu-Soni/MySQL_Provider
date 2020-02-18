package provider_mysql

import (
	"fmt"
	client "sql/client_mysql"

	"github.com/hashicorp/terraform/helper/schema"
)

func myresourceGrant() *schema.Resource {
	return &schema.Resource{
		Create: CreateGrant,
		Read:   ReadGrant,
		//Update: UpdateUser,
		Delete: DeleteGrant,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"dbname": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func CreateGrant(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Creategrant("CREATE", d.Get("name").(string), d.Get("dbname").(string))

	if err != nil {
		fmt.Errorf("error creating database")
	}

	d.SetId(d.Get("name").(string))
	return ReadGrant(d, meta)

}

func ReadGrant(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func DeleteGrant(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Creategrant("DELETE", d.Get("name").(string), d.Get("dbname").(string))

	if err != nil {
		fmt.Errorf("error creating database")
	}

	d.SetId("")
	return nil
}
