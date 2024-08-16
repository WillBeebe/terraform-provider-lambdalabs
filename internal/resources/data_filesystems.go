package resources

import (
	"context"
	"time"

	lambda "github.com/WillBeebe/lambdalabs-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DataSourceFileSystensSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceFileSystemsRead,
		Schema: map[string]*schema.Schema{
			"file_systems": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"region": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFileSystemsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	resp, _, err := client.DefaultAPI.ListFileSystems(ctx).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	instanceTypes := make([]map[string]interface{}, 0)

	for _, it := range resp.Data {
		instanceType := map[string]interface{}{
			"name":   it.Name,
			"region": it.Region,
			// "gpu_description":                 it.InstanceType.GpuDescription,
			// "price_cents_per_hour":            it.InstanceType.PriceCentsPerHour,
			// "specs":                           it.InstanceType.Specs,
			// "regions_with_capacity_available": it.RegionsWithCapacityAvailable,
		}
		instanceTypes = append(instanceTypes, instanceType)
	}

	if err := d.Set("file_systems", instanceTypes); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(time.Now().UTC().String())
	return nil
}
