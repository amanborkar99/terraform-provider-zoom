package zoom_corp

import (
	"fmt"

	"testing"

	//"github.com/hashicorp/terraform/helper/resource"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccUser_dataSourceSingleUser(t *testing.T){
	resource.Test(t, resource.TestCase{
		PreCheck: func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
			  
              Config: testAccReadSingleUser(),
			  Check: resource.ComposeTestCheckFunc(
				  testAccCheckZoomUserExists("data.zoom_read_single_user.user"),
				  resource.TestCheckResourceAttr(
					"data.zoom_read_single_user.user", "email", "kunalp.gohire@gmail.com"),

				  resource.TestCheckResourceAttr(
					  "data.zoom_read_single_user.user", "first_name", "KUNAL" ),

				  resource.TestCheckResourceAttr( 
					  "data.zoom_read_single_user.user", "last_name", "GOHIRE"),

				  resource.TestCheckResourceAttr( 
						"data.zoom_read_single_user.user", "type", "1"),
				  
			  ),
			},
		},
	})
	
}
func testAccReadSingleUser() string {
	return fmt.Sprintf(`
data "zoom_read_single_user" "user" {
	email = "kunalp.gohire@gmail.com"
}  
`)
}



