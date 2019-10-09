package centreonweb

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

const commandObject string = "CMD"

// ClientCommands is used to store the client to interact with the Centreon API
type ClientCommands struct {
	CentClient *ClientCentreonWeb
}

// Commands is an array of Command to store the answer from Centreon API
type Commands struct {
	Cmd []Command `json:"result"`
}

// Command struct is used to store parameters of a command
type Command struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Line string `json:"line"`
}

// Show lists available commands
func (c *ClientCommands) Show(name string) ([]Command, error) {
	respReader, err := c.CentClient.centreonAPIRequest("show", commandObject,
		name)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var cmds Commands
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&cmds)

	return cmds.Cmd, nil
}

// Get returns a specific command
func (c *ClientCommands) Get(name string) (Command, error) {
	var objFound Command

	if name == "" {
		return objFound, pkgerrors.New("name parameter cannot be empty when " +
			"calling Get function")
	}

	cmds, err := c.Show(name)
	if err != nil {
		return objFound, err
	}

	for _, c := range cmds {
		if c.Name == name {
			objFound = c
		}
	}

	if objFound.ID == "" {
		return objFound, pkgerrors.New("command " + name + " not found")
	}

	return objFound, nil
}

// Exists returns true if the command exists, false otherwise
func (c *ClientCommands) Exists(name string) (bool, error) {
	objExists := false

	if name == "" {
		return objExists, pkgerrors.New("name parameter cannot be empty when " +
			"calling Exists function")
	}

	cmds, err := c.Show(name)
	if err != nil {
		return objExists, err
	}

	for _, c := range cmds {
		if c.Name == name {
			objExists = true
		}
	}

	return objExists, nil
}

// Add adds a new command
func (c *ClientCommands) Add(cmd Command) error {
	if cmd.Name == "" || cmd.Type == "" || cmd.Line == "" {
		return pkgerrors.New("cmd.Name or cmd.Type or cmd.Line parameter cannot" +
			" be empty when calling Add function")
	}

	values := cmd.Name + ";" + cmd.Type + ";" + cmd.Line

	respReader, err := c.CentClient.centreonAPIRequest("add", commandObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Del removes the specified command
func (c *ClientCommands) Del(name string) error {
	if name == "" {
		return pkgerrors.New("name parameter cannot be empty when calling Del " +
			"function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("del", commandObject,
		name)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparam is used to change a specific parameters for a command
func (c *ClientCommands) Setparam(name string, param string,
	value string) error {

	if name == "" || param == "" {
		return pkgerrors.New("name or param parameters cannot be empty when " +
			"calling Setparam function")
	}

	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setparam", commandObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}
