package zoom_corp

import (
	"context"
	"os"
	//e "errors"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	xl "terraform-provider-zoom/Client"
)

// type client struct {
// 	token string
// }

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			// "Username": &schema.Schema{
			// 	Type:     schema.TypeString,
			// 	Optional: true,
			// },
			// "Password": &schema.Schema{
			// 	Type:      schema.TypeString,
			// 	Sensitive: true,
			// 	Optional:  true,
			// },
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				DefaultFunc: schema.EnvDefaultFunc("token_zoom", nil),

			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"zoom_resource": resourceUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"zoom_corp_data": dataSourceUser(),
			"zoom_read_single_user": dataSourceSingleUser(),
			
		},
		ConfigureContextFunc: providerConfig,
	}
}

func providerConfig(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {

	//us_name := d.Get("Username").(string)
	//us_pass := d.Get("Password").(string)
	us_token := d.Get("token").(string)

	c, err := xl.NewClient(us_token)
	os.Setenv("token_zoom", us_token)
	if err != nil {
		return nil, diag.FromErr(err)
	}
	var diags diag.Diagnostics
	//c:=client{}
	//c.token = us_token
	return c, diags

}

//return c, diags
