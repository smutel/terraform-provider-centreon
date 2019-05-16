package centreonweb

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

const command_object string = "CMD"

type commandsClient struct {
	CentClient *CentreonwebClient
}

type Commands struct {
	Cmd []Command `json:"result"`
}

type Command struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Line string `json:"line"`
}

func (c *commandsClient) Show(name string) ([]Command, error) {
	respReader, err := c.CentClient.centreonApiRequest("show", command_object, name)
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

func (c *commandsClient) Get(name string) (Command, error) {
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

func (c *commandsClient) Exists(name string) (bool, error) {
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

func (c *commandsClient) Add(cmd Command) error {
	values := cmd.Name + ";" + cmd.Type + ";" + cmd.Line

	respReader, err := c.CentClient.centreonApiRequest("add", command_object, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

func (c *commandsClient) Del(name string) error {
	respReader, err := c.CentClient.centreonApiRequest("del", command_object, name)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

func (c *commandsClient) Setparam(name string, param string, value string) error {
	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonApiRequest("setparam", command_object, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}
