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
