package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func ResourceInstanceSchema() *schema.Resource {
	return &schema.Resource{
		Create: resourceInstanceCreate,
		Read:   resourceInstanceRead,
		Delete: resourceInstanceDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringLenBetween(1, 64),
				Description:  "User-provided name for the instance",
			},
			"region_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Short name of the region where the instance will be launched",
			},
			"instance_type_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the instance type to launch",
			},
			"ssh_key_names": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Names of the SSH keys to allow access to the instance",
			},
			"file_system_names": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Names of the file systems to attach to the instance",
			},
			"ip": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address of the instance",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The current status of the instance",
			},
			"hostname": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hostname assigned to this instance",
			},
			"jupyter_token": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Secret token used to log into the jupyter lab server hosted on the instance",
			},
			"jupyter_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that opens a jupyter lab notebook on the instance",
			},
		},
	}
}

func resourceInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to launch an instance
	return resourceInstanceRead(d, meta)
}

func resourceInstanceRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to get instance details
	return nil
}

func resourceInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to terminate an instance
	return nil
}
