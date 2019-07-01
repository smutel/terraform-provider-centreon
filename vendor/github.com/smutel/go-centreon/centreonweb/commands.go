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
	respReader, err := c.CentClient.centreonAPIRequest("show", commandObject, name)
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
	var cmdFound Command
	cmds, err := c.Show(name)
	if err != nil {
		return cmdFound, err
	}

	for _, c := range cmds {
		if c.Name == name {
			cmdFound = c
		}
	}

	if cmdFound.ID == "" {
		return cmdFound, pkgerrors.New("command " + name + " not found")
	}

	return cmdFound, nil
}

// Exists returns true if the command exists, false otherwise
func (c *ClientCommands) Exists(name string) (bool, error) {
	cmdExists := false

	cmds, err := c.Show(name)
	if err != nil {
		return cmdExists, err
	}

	for _, c := range cmds {
		if c.Name == name {
			cmdExists = true
		}
	}

	return cmdExists, nil
}

// Add adds a new command
func (c *ClientCommands) Add(cmd Command) error {
	values := cmd.Name + ";" + cmd.Type + ";" + cmd.Line

	respReader, err := c.CentClient.centreonAPIRequest("add", commandObject, values)
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
	respReader, err := c.CentClient.centreonAPIRequest("del", commandObject, name)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparam is used to change a specific parameters for a command
func (c *ClientCommands) Setparam(name string, param string, value string) error {
	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setparam", commandObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}
