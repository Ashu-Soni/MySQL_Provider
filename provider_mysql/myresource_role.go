package provider_mysql

import (
	"fmt"
	client "sql/client_mysql"

	"github.com/hashicorp/terraform/helper/schema"
)

func myresourceRole() *schema.Resource {
	return &schema.Resource{
		Create: CreateRole,
		//Update: UpdateDatabase,
		Read:   ReadRole,
		Delete: DeleteRole,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
func CreateRole(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Createrole("CREATE ROLE", d.Get("name").(string))
	if err != nil {
		fmt.Errorf("error creating role")
	}

	d.SetId(d.Get("name").(string))
	return nil
}

// func UpdateDatabase(d *schema.ResourceData,meta interface) error{
//  return ReadDatabase(d,meta)
// }

func ReadRole(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func DeleteRole(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Createrole("DROP ROLE", d.Get("name").(string))
	if err != nil {
		fmt.Errorf("Error deleting database")
	}
	d.SetId(d.Get("name").(string))
	return nil
}
