resource "centreon_command" "tf_check_host_alive" {
  name = "tf_check-host-alive"
  type = "check"
  line = "$USER1$/check_ping -H $HOSTADDRESS$ -w 3000.0,80% -c 5000.0,100% -p 1"
}

resource "centreon_timeperiod" "tf-workhours" {
  name      = "tf-workhours"
  alias     = "Timeperiod created by Terraform"
  monday    = "08:00-18:00"
  tuesday   = "08:00-18:00"
  wednesday = "08:00-18:00"
  thursday  = "08:00-18:00"
  friday    = "08:00-18:00"
}

resource "centreon_timeperiod_exception" "tf_time_exception1" {
  timeperiod_id = "${centreon_timeperiod.tf-workhours.id}"
  days          = "January 1"
  timerange     = "08:00-18:00"
}

resource "centreon_timeperiod_exception" "tf_time_exception2" {
  timeperiod_id = "${centreon_timeperiod.tf-workhours.id}"
  days          = "December 25"
  timerange     = "08:00-18:00"
}

resource "centreon_host" "tf_host" {
  # Host Configuration
  name                    = "tf_host"
  alias                   = "Host created by Terraform"
  address                 = "127.0.0.1"
  snmp_community          = "public"
  snmp_version            = "3"
  instance                = "Central"
  timezone                = "Europe/Paris"
  templates               = ["App-Monitoring-Centreon-Database-custom", "App-Monitoring-Centreon-Central-custom"]
  check_command           = centreon_command.tf_check_host_alive.id
  check_command_arguments = "!1!2"

  macro {
    name        = "WARNING"
    value       = "80"
    description = "Macro created by terraform"
  }

  macro {
    name        = "CRITICAL"
    value       = "90"
    description = "Macro created by terraform"
  }

  check_period           = centreon_timeperiod.tf-workhours.id
  max_check_attempts     = 3
  check_interval         = 1
  retry_check_interval   = 5
  active_checks_enabled  = "yes"
  passive_checks_enabled = "yes"

  # Notification
  notifications_enabled        = "no"
  linked_contacts              = ["user"]
  contact_additive_inheritance = false
  linked_contact_groups        = ["Supervisors"]
  cg_additive_inheritance      = false
  # notification_options_none    = false
  notification_options        = ["down", "unreachable"]
  notification_interval       = 5
  notification_period         = centreon_timeperiod.tf-workhours.id
  first_notification_delay    = 5
  recovery_notification_delay = 5

  # Relations
  hostgroups = ["Centreon_platform"]
  parents    = ["Centeon-central"]

  # Data Processing
  obsess_over_host             = "yes"
  acknowledgement_timeout      = 20
  check_freshness              = "yes"
  freshness_threshold          = 20
  flap_detection_enabled       = "yes"
  low_flap_threshold           = 20
  high_flap_threshold          = 50
  retain_status_information    = "yes"
  retain_nonstatus_information = "yes"
  stalking_options             = ["up", "down", "unreachable"]
  event_handler_enabled        = "yes"
  event_handler                = centreon_command.tf_check_host_alive.id
  event_handler_arguments      = "!1!2"

  # Host Extended Infos
  notes_url       = "http://www.centreon.com"
  notes           = "Test notes"
  action_url      = "http://www.terraform.com"
  icon_image      = "ppm/applications-databases-mysql-DB-MySQL-64.png"
  icon_image_alt  = "Test image"
  statusmap_image = "ppm/applications-databases-mysql-DB-MySQL-64.png"
  coords2d        = "20,20"
  coords3d        = "5.0,5.0,5.0"
}
