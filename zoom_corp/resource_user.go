package zoom_corp

import (
	"context"
	"fmt"
	//"time"

	xl "terraform-provider-zoom/Client"
	sl "terraform-provider-zoom/Server"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUser() *schema.Resource {

	return &schema.Resource{
		ReadContext:   resourceUserRead,
		CreateContext: resourceUserCreate,
		UpdateContext: resourceUserUpdate,
		DeleteContext: resourceUserDelete,
		Schema: map[string]*schema.Schema{
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"first_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": &schema.Schema{
				Type: schema.TypeString,
				Computed: true,
			},
		},
	}

}

func resourceUserCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//var diags diag.Diagnostics
	c := m.(*xl.Client)
    fmt.Printf("USER CREATED\n")
	t := sl.Uz{
		Email:      d.Get("email").(string),
		First_name: d.Get("first_name").(string),
		Last_name:  d.Get("last_name").(string),
		Type:       d.Get("type").(int),
	}
	j_req := sl.Create{
		Action:    "create",
		User_info: t,
	}

	data, err := c.CreateUser(j_req)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.Email)

	
    
	return nil
}

func resourceUserRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*xl.Client)

	id := d.Id()
    fmt.Println("USER READ")
	_, err := c.GetUser(id)
	if err != nil {
		return diag.FromErr(err)
	}
        Email:=      d.Get("email").(string)
		First_name:= d.Get("first_name").(string)
		Last_name:=  d.Get("last_name").(string)
		Type:=       d.Get("type").(int)
	if err:= d.Set("email", Email); err!= nil{
		return diag.FromErr(err)
	}
	if err:= d.Set("first_name", First_name); err != nil{
		return diag.FromErr(err)
	}
	if err:= d.Set("last_name", Last_name); err != nil{
		return diag.FromErr(err)
	}
	if err:= d.Set("type", Type); err != nil{
		return diag.FromErr(err)
	}
	

	d.SetId(d.Id())

	return diags

}

func resourceUserUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*xl.Client)
	st := d.Get("status").(string)
    fmt.Println("USER UPDATED")
	
	if st == "pending"{
	   return 	diag.Errorf("User is not active")
	}
	
	t := sl.Uz{
		Email:      d.Get("email").(string),
		First_name: d.Get("first_name").(string),
		Last_name:  d.Get("last_name").(string),
		Type:       d.Get("type").(int),
	}
	id := d.Id()
	if d.HasChange("email") {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to change the values",
			Detail:   "Unable to find email",
		})
		return diags
	} else {
		if d.HasChange("first_name") || d.HasChange("last_name") || d.HasChange("type") {
		    err := c.UpdateUser(id, t)
			if err != nil {
				return diag.FromErr(err)
			}
			
		}
		d.SetId(t.Email)
	}
	return resourceUserRead(ctx, d, m)
}

func resourceUserDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	

	c := m.(*xl.Client)
	id := d.Id()
    fmt.Println("DELETE WORKING")
	st := d.Get("status").(string)

	err := c.DeleteUser(id, st)
	if err != nil{
		return diag.FromErr(err)
	}
    d.SetId("")
	

	return diags
}
