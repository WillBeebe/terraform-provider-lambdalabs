package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSSHKeySchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceSSHKeyCreate,
		Read:   resourceSSHKeyRead,
		Delete: resourceSSHKeyDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"private_key": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
		},
	}
}

func resourceSSHKeyCreate(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to add an SSH key
	return resourceSSHKeyRead(d, meta)
}

func resourceSSHKeyRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to get SSH key details
	return nil
}

func resourceSSHKeyDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to delete an SSH key
	return nil
}
