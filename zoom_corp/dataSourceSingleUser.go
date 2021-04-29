package zoom_corp

import (
	"context"
	// "strconv"
	// "time"

	t "terraform-provider-zoom/Client"
	//s "terraform-provider-zoom/Server"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSingleUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceSingleUserRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			
		},
	}
}

func dataSourceSingleUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*t.Client)

	var diags diag.Diagnostics

	id := d.Get("email").(string)

	user, err := c.GetUser(id)
    if err != nil{
		return diag.FromErr(err)
	}

	d.Set("id", user.Id)
	d.Set("first_name", user.First_name)
	d.Set("last_name", user.Last_name)
	d.Set("email",user.Email)
	d.Set("type",user.Type)
	

    d.SetId(id)
	

	return diags
}
