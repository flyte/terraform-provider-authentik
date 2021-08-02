package provider

import (
	"context"

	"github.com/goauthentik/terraform-provider-authentik/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceStageAuthenticatorTOTP() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceStageAuthenticatorTOTPCreate,
		ReadContext:   resourceStageAuthenticatorTOTPRead,
		UpdateContext: resourceStageAuthenticatorTOTPUpdate,
		DeleteContext: resourceStageAuthenticatorTOTPDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"configure_flow": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"digits": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  6,
			},
		},
	}
}

func resourceStageAuthenticatorTOTPSchemaToProvider(d *schema.ResourceData) (*api.AuthenticatorTOTPStageRequest, diag.Diagnostics) {
	r := api.AuthenticatorTOTPStageRequest{
		Name:   d.Get("name").(string),
		Digits: api.DigitsEnum((d.Get("digits").(int))),
	}

	if h, hSet := d.GetOk("configure_flow"); hSet {
		r.ConfigureFlow.Set(stringToPointer(h.(string)))
	}
	return &r, nil
}

func resourceStageAuthenticatorTOTPCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	r, diags := resourceStageAuthenticatorTOTPSchemaToProvider(d)
	if diags != nil {
		return diags
	}

	res, hr, err := c.client.StagesApi.StagesAuthenticatorTotpCreate(ctx).AuthenticatorTOTPStageRequest(*r).Execute()
	if err != nil {
		return httpToDiag(hr)
	}

	d.SetId(res.Pk)
	return resourceStageAuthenticatorTOTPRead(ctx, d, m)
}

func resourceStageAuthenticatorTOTPRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*APIClient)

	res, hr, err := c.client.StagesApi.StagesAuthenticatorTotpRetrieve(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr)
	}

	d.Set("name", res.Name)
	d.Set("digits", res.Digits)
	if res.ConfigureFlow.IsSet() {
		d.Set("configure_flow", res.ConfigureFlow.Get())
	}
	return diags
}

func resourceStageAuthenticatorTOTPUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)

	app, di := resourceStageAuthenticatorTOTPSchemaToProvider(d)
	if di != nil {
		return di
	}

	res, hr, err := c.client.StagesApi.StagesAuthenticatorTotpUpdate(ctx, d.Id()).AuthenticatorTOTPStageRequest(*app).Execute()
	if err != nil {
		return httpToDiag(hr)
	}

	d.SetId(res.Pk)
	return resourceStageAuthenticatorTOTPRead(ctx, d, m)
}

func resourceStageAuthenticatorTOTPDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*APIClient)
	hr, err := c.client.StagesApi.StagesAuthenticatorTotpDestroy(ctx, d.Id()).Execute()
	if err != nil {
		return httpToDiag(hr)
	}
	return diag.Diagnostics{}
}