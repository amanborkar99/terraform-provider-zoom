package zoom_corp

import (
	"os"
	"testing"
	//"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	//"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string] *schema.Provider
var testAccProvider *schema.Provider

func init(){
	testAccProvider = Provider()
	testAccProviders = map[string] *schema.Provider{
		"zoom": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T){
	 
	 os.Setenv("TOKEN_ZOOM", "Bearer token")
	 token := os.Getenv("TOKEN_ZOOM")
	 if  token == "" {
		 t.Fatalf("TOKEN CANNOT BE EMPTY")
	 }
}


