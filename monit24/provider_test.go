package monit24

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var providerFactories = map[string]func() (*schema.Provider, error){
	"monit24": func() (*schema.Provider, error) {
		return Provider(), nil
	},
}

func preCheck(t *testing.T) {
	vars := []string{
		"MONIT24_USER",
		"MONIT24_PASSWORD",
	}

	for _, v := range vars {
		if os.Getenv(v) == "" {
			t.Fatalf("Missing %v variable for acceptance tests", v)
		}
	}
}
