package resources

import (
	"context"
	"time"

	lambda "github.com/WillBeebe/lambdalabs-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceInstanceSchema() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceInstanceCreate,
		ReadContext:   resourceInstanceRead,
		DeleteContext: resourceInstanceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"region_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_type_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_key_names": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MaxItems: 1,
			},
			"file_system_names": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				MaxItems: 1,
			},
			"ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jupyter_token": {
				Type:      schema.TypeString,
				Computed:  true,
				Sensitive: true,
			},
			"jupyter_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceInstanceCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	sshKeyNames := make([]string, 0)
	for _, v := range d.Get("ssh_key_names").([]interface{}) {
		sshKeyNames = append(sshKeyNames, v.(string))
	}

	fileSystemNames := make([]string, 0)
	if v, ok := d.GetOk("file_system_names"); ok {
		for _, name := range v.([]interface{}) {
			fileSystemNames = append(fileSystemNames, name.(string))
		}
	}

	req := lambda.LaunchInstanceRequest{
		RegionName:       d.Get("region_name").(string),
		InstanceTypeName: d.Get("instance_type_name").(string),
		SshKeyNames:      sshKeyNames,
		FileSystemNames:  fileSystemNames,
		Quantity:         lambda.PtrInt32(1),
	}

	if v, ok := d.GetOk("name"); ok {
		req.Name = *lambda.NewNullableString(lambda.PtrString(v.(string)))
	}

	resp, _, err := client.DefaultAPI.LaunchInstance(ctx).LaunchInstanceRequest(req).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	if len(resp.Data.InstanceIds) == 0 {
		return diag.Errorf("No instances were launched")
	}

	d.SetId(resp.Data.InstanceIds[0])

	// Wait for the instance to be in the "active" state
	err = waitForInstanceStatus(ctx, client, d.Id(), "active", 5*time.Minute)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceInstanceRead(ctx, d, meta)
}

func resourceInstanceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	instance, _, err := client.DefaultAPI.GetInstance(ctx, d.Id()).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("name", instance.Data.Name)
	d.Set("region_name", instance.Data.Region.Name)
	d.Set("instance_type_name", instance.Data.InstanceType.Name)
	d.Set("ssh_key_names", instance.Data.SshKeyNames)
	d.Set("file_system_names", instance.Data.FileSystemNames)

	if instance.Data.Ip.IsSet() {
		d.Set("ip", instance.Data.Ip.Get())
	}
	d.Set("status", instance.Data.Status)
	if instance.Data.Hostname.IsSet() {
		d.Set("hostname", instance.Data.Hostname.Get())
	}
	if instance.Data.JupyterToken.IsSet() {
		d.Set("jupyter_token", instance.Data.JupyterToken.Get())
	}
	if instance.Data.JupyterUrl.IsSet() {
		d.Set("jupyter_url", instance.Data.JupyterUrl.Get())
	}

	return nil
}

func resourceInstanceDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	client := meta.(*lambda.APIClient)

	req := lambda.TerminateInstanceRequest{
		InstanceIds: []string{d.Id()},
	}

	_, _, err := client.DefaultAPI.TerminateInstance(ctx).TerminateInstanceRequest(req).Execute()
	if err != nil {
		return diag.FromErr(err)
	}

	// Wait for the instance to be in the "terminated" state
	err = waitForInstanceStatus(ctx, client, d.Id(), "terminated", 5*time.Minute)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId("")
	return nil
}

func waitForInstanceStatus(ctx context.Context, client *lambda.APIClient, id string, targetStatus string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"booting", "unhealthy", "terminating"},
		Target:  []string{targetStatus},
		Refresh: func() (interface{}, string, error) {
			instance, _, err := client.DefaultAPI.GetInstance(ctx, id).Execute()
			if err != nil {
				return nil, "", err
			}
			return instance, instance.Data.Status, nil
		},
		Timeout:    timeout,
		Delay:      10 * time.Second,
		MinTimeout: 5 * time.Second,
	}

	_, err := stateConf.WaitForStateContext(ctx)
	return err
}
