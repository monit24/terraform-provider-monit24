terraform {
  required_providers {
    monit24 = {
      version = ">= 0.1.0"
      source  = "monit24/monit24"
    }
  }
}

provider "monit24" {}

resource "monit24_group" "example" {
  name = "An example group"
}

resource "monit24_notification_address" "email" {
  address                 = "notifications@example.com"
  notification_channel_id = "email"
  group_id                = monit24_group.example.id
  description             = "An example notification email address"
}

resource "monit24_service" "test_https_service" {
  type_id     = "https"
  name        = "https example.com"
  description = "An example https service"
  address     = "example.com"
  group_id    = monit24_group.example.id
}
