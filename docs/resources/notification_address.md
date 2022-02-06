---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "monit24_notification_address Resource - terraform-provider-monit24"
subcategory: ""
description: |-
  
---

# monit24_notification_address (Resource)



## Example Usage

```terraform
resource "monit24_notification_address" "email" {
  address                 = "notifications@example.com"
  notification_channel_id = "email"
  group_id                = monit24_group.my_group.id
  description             = "Email notification"
}

resource "monit24_notification_address" "phone" {
  address                 = "1234567890"
  notification_channel_id = "sms"
  group_id                = monit24_group.my_group.id
  description             = "SMS notification"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **address** (String)
- **group_id** (Number)
- **notification_channel_id** (String)

### Optional

- **description** (String)
- **id** (String) The ID of this resource.

## Import

Import is supported using the following syntax:

```shell
terraform import monit24_notification_address.email 123456
```