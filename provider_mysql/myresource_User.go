package provider_mysql

import (
	"fmt"
	client "sql/client_mysql"

	"github.com/hashicorp/terraform/helper/schema"
)

func myresourceUser() *schema.Resource {
	return &schema.Resource{
		Create: CreateUser,
		Read:   ReadUser,
		//Update: UpdateUser,
		Delete: DeleteUser,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func CreateUser(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Createuser("CREATE", d.Get("name").(string))

	if err != nil {
		fmt.Errorf("error creating database")
	}

	d.SetId(d.Get("name").(string))
	return nil

}

func ReadUser(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func DeleteUser(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Createuser("DELETE", d.Get("name").(string))

	if err != nil {
		fmt.Errorf("error creating database")
	}

	d.SetId(d.Get("name").(string))
	return nil
}
