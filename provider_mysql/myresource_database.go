package provider_mysql

import (
	"fmt"
	client "sql/client_mysql"

	"github.com/hashicorp/terraform/helper/schema"
)

func myresourceDatabase() *schema.Resource {
	return &schema.Resource{
		Create: CreateDatabase,
		//Update: UpdateDatabase,
		Read:   ReadDatabase,
		Delete: DeleteDatabase,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}
func CreateDatabase(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Createdb("CREATE", d.Get("name").(string))
	if err != nil {
		fmt.Errorf("error creating database")
	}

	d.SetId(d.Get("name").(string))
	return nil
}

// func UpdateDatabase(d *schema.ResourceData,meta interface) error{
// 	return ReadDatabase(d,meta)
// }

func ReadDatabase(d *schema.ResourceData, meta interface{}) error {

	return nil
}

func DeleteDatabase(d *schema.ResourceData, meta interface{}) error {
	sqlconnect := meta.(*client.Client)
	err := sqlconnect.Createdb("DELETE", d.Get("name").(string))
	if err != nil {
		fmt.Errorf("Error deleting database")
	}
	d.SetId("")

	err = sqlconnect.Createdb("SHOW", "")
	return nil
}
