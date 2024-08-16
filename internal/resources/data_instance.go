package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInstanceTypesSchema() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceInstanceTypesRead,
		Schema: map[string]*schema.Schema{
			"instance_types": {
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
				},
			},
		},
	}
}

func dataSourceInstanceTypesRead(d *schema.ResourceData, meta interface{}) error {
	// Implement the API call to get instance types
	return nil
}
