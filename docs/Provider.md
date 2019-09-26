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
