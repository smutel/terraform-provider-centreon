resource "centreon_command" "check-host-alive" {
  name = "check-host-alive"
  type = "check"
  line = "$USER1$/check_ping -H $HOSTADDRESS$ -w 3000.0,80% -c 5000.0,100% -p 1"
}

resource "centreon_timeperiod" "workhours-custom" {
  name = "workhours-custom"
  alias = "Custom Working Hours"
  monday = "08:00-18:00"
  tuesday = "08:00-18:00"
  wednesday = "08:00-18:00"
  thursday = "08:00-18:00"
  friday = "08:00-18:00"
}

resource "centreon_timeperiod_exception" "time_exception1" {
  timeperiod_id = "${centreon_timeperiod.workhours-custom.id}"
  days = "January 2"
  timerange = "11:00-15:10"
}

resource "centreon_timeperiod_exception" "time_exception2" {
  timeperiod_id = "${centreon_timeperiod.workhours-custom.id}"
  days = "January 2"
  timerange = "10:00-15:00"
}
