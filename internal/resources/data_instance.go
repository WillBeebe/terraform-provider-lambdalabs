package resources

import (
	"context"

	lambda "github.com/WillBeebe/lambdalabs-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInstanceSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInstanceRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gpu_description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"price_cents_per_hour": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"specs": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
			"regions_with_capacity_available": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	client := meta.(*lambda.APIClient)

	instance, _, err := client.DefaultAPI.GetInstance(ctx, d.Id()).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", instance.Data.Name)
	d.Set("region_name", instance.Data.Region.Name)
	d.Set("instance_type_name", instance.Data.InstanceType.Name)
	// d.Set("ssh_key_names", instance.Data.SshKeyNames)
	// d.Set("file_system_names", instance.Data.FileSystemNames)

	// if instance.Data.Ip.IsSet() {
	// 	d.Set("ip", instance.Data.Ip.Get())
	// }
	// d.Set("status", instance.Data.Status)
	// if instance.Data.Hostname.IsSet() {
	// 	d.Set("hostname", instance.Data.Hostname.Get())
	// }
	// if instance.Data.JupyterToken.IsSet() {
	// 	d.Set("jupyter_token", instance.Data.JupyterToken.Get())
	// }
	// if instance.Data.JupyterUrl.IsSet() {
	// 	d.Set("jupyter_url", instance.Data.JupyterUrl.Get())
	// }
	return nil
}
