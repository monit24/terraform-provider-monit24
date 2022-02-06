package monit24

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNotificationAddress(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { preCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNotificationAddressConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("monit24_notification_address.test", "address", "notifications@example.com"),
				),
			},
		},
	})
}

const testAccNotificationAddressConfig = `
resource "monit24_notification_address" "test" {
  address                 = "notifications@example.com"
  notification_channel_id = "email"
  group_id                = monit24_group.test.id
  description             = "Email notification"
}

resource "monit24_group" "test" {
  name = "test group"
}
`
