package monit24

import (
	"context"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/monit24/terraform-provider-monit24/client"
)

func resourceNotificationAddress() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceNotificationAddressCreate,
		ReadContext:   resourceNotificationAddressRead,
		UpdateContext: resourceNotificationAddressUpdate,
		DeleteContext: resourceNotificationAddressDelete,
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notification_channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"group_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func notificationAddressFromResourceData(d *schema.ResourceData, c client.Client) client.NotificationAddress {
	address := client.NotificationAddress{
		Address:               d.Get("address").(string),
		NotificationChannelID: d.Get("notification_channel_id").(string),
		GroupID:               d.Get("group_id").(int),
		OwnerID:               c.OwnerID(),
	}

	if v, ok := d.GetOk("description"); ok {
		address.Description = v.(string)
	}

	return address
}

func resourceNotificationAddressCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.Client)

	address := notificationAddressFromResourceData(d, c)

	id, err := c.CreateNotificationAddress(ctx, address)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.Itoa(id))

	return diags
}

func resourceNotificationAddressRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	address, err := c.ReadNotificationAddress(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("address", address.Address); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("notification_channel_id", address.NotificationChannelID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("group_id", address.GroupID); err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("description", address.Description); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNotificationAddressUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.Client)

	address := notificationAddressFromResourceData(d, c)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.UpdateNotificationAddress(ctx, id, address)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceNotificationAddressDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(client.Client)

	id, err := strconv.Atoi(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	err = c.DeleteNotificationAddress(ctx, id)
	if err != nil {
		return diag.FromErr(err)
	}

	return diags
}
