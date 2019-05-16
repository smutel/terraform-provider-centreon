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
