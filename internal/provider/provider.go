package provider

import (
	"fmt"

	lambda "github.com/WillBeebe/lambdalabs-client"
	"github.com/WillBeebe/terraform-provider-lambdalabs/internal/resources"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("LAMBDALABS_API_KEY", nil),
				Description: "The API key for Lambda Labs Cloud API authentication",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"lambdalabs_instance": resources.ResourceInstanceSchema(),
			"lambdalabs_ssh_key":  resources.ResourceSSHKeySchema(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"lambdalabs_instance_types": resources.DataSourceInstanceTypesSchema(),
			"lambdalabs_instance":       resources.DataSourceInstanceSchema(),
			"lambdalabs_file_systems":   resources.DataSourceFileSystensSchema(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	apiKey := d.Get("api_key").(string)
	cfg := lambda.NewConfiguration()
	cfg.AddDefaultHeader("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	client := lambda.NewAPIClient(cfg)

	return client, nil
}
