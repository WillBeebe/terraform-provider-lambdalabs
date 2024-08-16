package resources

import (
	"context"
	"time"

	lambda "github.com/WillBeebe/lambdalabs-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceInstanceTypesSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceInstanceTypesRead,
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
						// "specs": {
						// 	Type:     schema.TypeMap,
						// 	Computed: true,
						// 	Elem: &schema.Schema{
						// 		Type: schema.TypeInt,
						// 	},
						// },
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

func dataSourceInstanceTypesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	resp, _, err := client.DefaultAPI.InstanceTypes(ctx).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceTypes := make([]map[string]interface{}, 0)

	for _, it := range resp.Data {
		instanceType := map[string]interface{}{
			"name":                 it.InstanceType.Name,
			"description":          it.InstanceType.Description,
			"gpu_description":      it.InstanceType.GpuDescription,
			"price_cents_per_hour": it.InstanceType.PriceCentsPerHour,
			// "specs":                           it.InstanceType.Specs,
			// "regions_with_capacity_available": it.RegionsWithCapacityAvailable,
		}
		regions := make([]map[string]interface{}, 0)
		for _, it2 := range it.RegionsWithCapacityAvailable {
			region := map[string]interface{}{
				"name":        it2.Name,
				"description": it2.Description,
			}
			regions = append(regions, region)
		}
		instanceType["regions_with_capacity_available"] = regions
		instanceTypes = append(instanceTypes, instanceType)
	}

	if err := d.Set("instance_types", instanceTypes); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(time.Now().UTC().String())
	return nil

}
