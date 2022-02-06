resource "monit24_service" "https_service" {
  type_id                       = "https"
  name                          = "https example.com"
  description                   = "An example HTTPS service"
  address                       = "example.com"
  group_id                      = monit24_group.example.id
  interval                      = 600
  is_active                     = true
  is_archived                   = true
  notification_channel_ids      = ["email", "sms"]
  notification_condition_ids    = ["failure", "recovery"]
  notification_mode_id          = "default"
  recovery_notification_mode_id = "default"
  extended_settings = {
    error_tolerance        = 52
    sensors_failover       = true
    port                   = 666
    timeout                = 7000
    http_method            = "GET"
    content_regex          = "OK"
    negative_content_regex = "BAD"
    content_type           = "application/json"
    content                = "{}"
  }
}

resource "monit24_service" "ping_service" {
  type_id                       = "ping"
  name                          = "ping example.com"
  description                   = "An example ping service"
  address                       = "example.com"
  group_id                      = monit24_group.example.id
  interval                      = 900
  is_active                     = true
  is_archived                   = false
  notification_channel_ids      = ["email", "sms"]
  notification_condition_ids    = ["failure", "recovery"]
  notification_mode_id          = "default"
  recovery_notification_mode_id = "default"
  extended_settings = {
    error_tolerance  = 59
    sensors_failover = true
    bytes_to_send    = 100
    reply_timeout    = 100
  }
}