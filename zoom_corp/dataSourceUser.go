package zoom_corp

import (
	"context"
	"encoding/json"
	"strconv"
	//"os"

	//"fmt"
	"net/http"
	//"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	t "terraform-provider-zoom/Client"
	s "terraform-provider-zoom/Server"
)

func dataSourceUser() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceUserRead,
		Schema: map[string]*schema.Schema{
			"page_count": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_number": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"page_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"next_page_token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
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
							Computed: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"pmi": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"verified": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", "https://api.zoom.us/v2/users", nil)

	if err != nil {
		return diag.FromErr(err)
	}

	w := m.(*t.Client)
	// tt := os.Getenv("ttk")
	req.Header.Set("Authorization", w.Token)
	st := s.Page{}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	//info := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&st)
	if err != nil {
		return diag.FromErr(err)
	}

	us := make([]interface{}, len(st.User))

	d.Set("page_count", st.Page_count)
	d.Set("page_number", st.Page_number)
	d.Set("page_size", st.Page_size)
	d.Set("next_page_token", st.Next_page_tk)

	for i, item := range st.User {
		oi := make(map[string]interface{})

		oi["id"] = item.Id
		oi["first_name"] = item.First_name
		oi["last_name"] = item.Last_name
		oi["email"] = item.Email
		oi["type"] = item.Type
		oi["pmi"] = item.Pmi
		oi["verified"] = item.Verified
		oi["status"] = item.Status

		us[i] = oi

	}

	d.Set("users", us)

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
