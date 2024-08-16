package resources

import (
	"context"

	lambda "github.com/WillBeebe/lambdalabs-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceSSHKeySchema() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceSSHKeyCreate,
		ReadContext:   resourceSSHKeyRead,
		DeleteContext: resourceSSHKeyDelete,
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

func resourceSSHKeyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	name := d.Get("name").(string)
	req := lambda.AddSSHKeyRequest{
		Name: name,
	}

	if v, ok := d.GetOk("public_key"); ok {
		publicKey := v.(string)
		req.PublicKey = &publicKey
	}

	resp, _, err := client.DefaultAPI.AddSSHKey(ctx).AddSSHKeyRequest(req).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if resp.Data.Id == "" {
		return diag.Errorf("Received empty ID from API")
	}

	d.SetId(resp.Data.Id)

	if resp.Data.PrivateKey.IsSet() {
		d.Set("private_key", resp.Data.PrivateKey)
	}

	return resourceSSHKeyRead(ctx, d, meta)
}

func resourceSSHKeyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	resp, _, err := client.DefaultAPI.ListSSHKeys(ctx).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	for _, key := range resp.Data {
		if key.Id == d.Id() {
			d.Set("name", key.Name)
			d.Set("public_key", key.PublicKey)
			return nil
		}
	}

	// If we reach this point, the SSH key wasn't found
	d.SetId("")
	return diag.Errorf("SSH key with ID %s not found", d.Id())
}

func resourceSSHKeyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	_, err := client.DefaultAPI.DeleteSSHKey(ctx, d.Id()).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}
