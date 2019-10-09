# Centreon Provider

## Resource centreon_command

Create a command in Centreon.

```hcl
resource "centreon_command" "test" {
  name = "check-host-alive"
  type = "check"
  line = "$USER1$/check_ping -H $HOSTADDRESS$ -w 3000.0,80% -c 5000.0,100% -p 1"
}
```

_Argument Reference_:
* ``name`` - (mandatory) The name of the command.
* ``type`` - (mandatory) The type of the command (check, notif, misc or discovery).
* ``line`` - (mandatory) System command line that will be run on execution.

For further information check centreon CLAPI [documentation](https://documentation-fr.centreon.com/docs/centreon/en/latest/api/clapi/objects/commands.html).

## Resource centreon_timeperiod

Create a timeperiod in Centreon.

```hcl
resource "centreon_timeperiod" "tp1" {
  name      = "workhours-custom"
  alias     = "Custom Working Hours"
  monday    = "08:00-18:00"
  tuesday   = "08:00-18:00"
  wednesday = "08:00-18:00"
  thursday  = "08:00-18:00"
  friday    = "08:00-18:00"
}
```

_Argument Reference_:
* ``name``      - (mandatory) The name of the timeperiod.
* ``alias``     - (mandatory) The alias for this timeperiod.
* ``sunday``    - (optional) A time exception (format=00:00-00:00) for Sunday.
* ``monday``    - (optional) A time exception (format=00:00-00:00) for Monday.
* ``tuesday``   - (optional) A time exception (format=00:00-00:00) for Tuesday.
* ``wednesday`` - (optional) A time exception (format=00:00-00:00) for Wednesday.
* ``thursday``  - (optional) A time exception (format=00:00-00:00) for Thursday.
* ``friday``    - (optional) A time exception (format=00:00-00:00) for Friday.
* ``saturday``  - (optional) A time exception (format=00:00-00:00) for Saturday.

For further information check centreon CLAPI [documentation](https://documentation-fr.centreon.com/docs/centreon/en/latest/api/clapi/objects/time_periods.html).

## Resource centreon_timeperiod_exception

Create a timeperiod in Centreon.

```hcl
resource "centreon_timeperiod_exception" "tp1_ex1" {
  timeperiod_id = "${centreon_timeperiod.tp1.id}"
  days          = "January 1"
  timerange     = "08:00-18:00"
}
```

_Argument Reference_:
* ``timeperiod_id`` - (mandatory) The id of the timeperiod to attach this exception.
* ``days``          - (mandatory) The day when to apply this exception.
* ``timerange``     - (mandatory) A time exception (format=00:00-00:00) for this day.

For further information check centreon CLAPI [documentation](https://documentation-fr.centreon.com/docs/centreon/en/latest/api/clapi/objects/time_periods.html).

## Resource centreon_host

Create a host in Centreon.

```hcl
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
```

_Argument Reference_:

*Host Configuration*

* ``name`` - (mandatory) Host name of the host to add into Centreon.
* ``alias`` - (mandatory) Alias of the host to add into Centreon.
* ``address`` - (mandatory) IP address of the host to monitor.
* ``snmp_community`` - (optional) SNMP community to use for SNMP request.
* ``snmp_version`` - (optional) SNMP version to use for SNMP request.
* ``instance`` - (mandatory) Instance name (poller).
* ``timezone`` - (optional) Timezone where this host is located.
* ``templates`` - (optional) List of templates to associate to this host.
* ``check_command`` - (optional) Command to check this host.
* ``check_command_arguments`` - (optional) Command args to use with the command.
* ``macro`` - (optional) List of macro associated with this host.
  * ``name`` - (mandatory) Name of the macro.
  * ``value`` - (mandatory) Value of the macro.
  * ``is_password`` - (optional) Define if the macro is a password or not.
  * ``description`` - (optional) Description of the macro.
* ``check_period`` - (optional) Host is monitoring during this period.
* ``max_check_attempts`` - (optional) Max attempts before displaying the alert.
* ``check_interval`` - (optional) Time interval between each check.
* ``retry_check_interval`` - (optional) Time interval between each check when in trouble.
* ``active_checks_enabled`` - (optional) Active check enabled or not.
* ``passive_checks_enabled`` - (optional) Passive check enabled or not.

*Notification*

* ``notifications_enabled`` - (optional) Enable or disable notifications.
* ``linked_contacts`` - (optional) Contacts to notify when host is in trouble.
* ``contact_additive_inheritance`` - (optional) Notify contacts of the template also.
* ``linked_contact_groups`` - (optional) Contact groups to notify when host is in trouble.
* ``cg_additive_inheritance`` - (optional) Notify contact groups of the template also.
* ``notification_options_none`` - (optional) Set notification options to None.
* ``notification_options`` - (optional) Set notification options
* ``notification_interval`` - (optional) Interval of time between each notifications.
* ``notification_period`` - (optional) Period of time when the notification is allowed
* ``first_notification_delay`` - (optional) Period of time before sending the first notification.
* ``recovery_notification_delay`` (optional) Period of time when host is in recovery status.

*Relations*

* ``hostgroups`` - (optional) List of hostgroups to linked to this host.
* ``parents`` - (optional) List of parents of this host.

*Data Processing*

* ``obsess_over_host`` - (optional) Whether or not obsess over host option is enabled.
* ``acknowledgement_timeout`` - (optional) Duration of acknowledgement.
* ``check_freshness`` - (optional) Enable or disable freshness check.
* ``freshness_threshold`` - (optional) Freshness threshold.
* ``flap_detection_enabled`` - (optional) Enable or disable flap detection.
* ``low_flap_threshold`` - (optional) Low flap threshold.
* ``high_flap_threshold`` - (optional) High flap threshold.
* ``retain_status_information`` - (optional) Keep status information in the history.
* ``retain_nonstatus_information`` - (optional) Keep nonstatus information in the history.
* ``stalking_options`` - (optional) Stalking options.
* ``event_handler_enabled`` - (optional) Enable or disable event handler for this host.
* ``event_handler`` - (optional) Event handler to use.
* ``event_handler_arguments`` - (optional) Arguments of the event handler command.

*Host Extended Infos*

* ``notes_url`` - (optional) URL to notes about this host.
* ``notes`` - (optional) Notes about this host.
* ``action_url`` - (optional) URL to execute an action on this host.
* ``icon_image`` - (optional) Image icon for this host.
* ``icon_image_alt`` - (optional) Comment for the image icon for this host.
* ``statusmap_image`` - (optional) Image used in statusmap.
* ``coords2d`` - (optional) 2D coordinates in status map CGI.
* ``coords3d`` - (optional) 3D coordinates in status wrl CGI.

For further information check centreon CLAPI [documentation](https://documentation-fr.centreon.com/docs/centreon/en/latest/api/clapi/objects/hosts.html).
