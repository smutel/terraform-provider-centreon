package centreonweb

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

const timeperiodObject string = "TP"

// Const with days to use in setparam function of timeperiod
const (
	TimeperiodSunday    string = "sunday"
	TimeperiodMonday    string = "monday"
	TimeperiodTuesday   string = "tuesday"
	TimeperiodWednesday string = "wednesday"
	TimeperiodThursday  string = "thursday"
	TimeperiodFriday    string = "friday"
	TimeperiodSaturday  string = "saturday"
)

// ClientTimeperiods is used to store the client to interact with the Centreon
// API
type ClientTimeperiods struct {
	CentClient *ClientCentreonWeb
}

// Timeperiods is an array of Timeperiod to store the answer from Centreon API
type Timeperiods struct {
	Tmp []Timeperiod `json:"result"`
}

// Timeperiod struct is used to store parameters of a timeperiod
type Timeperiod struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Alias     string `json:"alias"`
	Sunday    string `json:"sunday"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Saturday  string `json:"saturday"`
}

// TimeperiodExceptions is an array of Timeperiod Exception to store the answer
// from Centreon API
type TimeperiodExceptions struct {
	TmpEx []TimeperiodException `json:"result"`
}

// TimeperiodException struct is used to store parameters of a timeperiod
// exception
type TimeperiodException struct {
	Days      string `json:"days"`
	Timerange string `json:"timerange"`
}

// Show lists available timeperiods
func (c *ClientTimeperiods) Show(name string) ([]Timeperiod, error) {
	respReader, err := c.CentClient.centreonAPIRequest("show", timeperiodObject,
		name)
	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var tmps Timeperiods
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&tmps)

	return tmps.Tmp, nil
}

// Get returns a specific timeperiod
func (c *ClientTimeperiods) Get(name string) (Timeperiod, error) {
	var objFound Timeperiod

	if name == "" {
		return objFound, pkgerrors.New("name parameter cannot be empty when " +
			"calling Get function")
	}

	objs, err := c.Show(name)
	if err != nil {
		return objFound, err
	}

	for _, c := range objs {
		if c.Name == name {
			objFound = c
		}
	}

	if objFound.ID == "" {
		return objFound, pkgerrors.New("object " + name + " not found")
	}

	return objFound, nil
}

// Exists returns true if the timeperiod exists, false otherwise
func (c *ClientTimeperiods) Exists(name string) (bool, error) {
	objExists := false

	if name == "" {
		return objExists, pkgerrors.New("name parameter cannot be empty when " +
			"calling Exists function")
	}

	objs, err := c.Show(name)
	if err != nil {
		return objExists, err
	}

	for _, c := range objs {
		if c.Name == name {
			objExists = true
		}
	}

	return objExists, nil
}

// Add adds a new timeperiod
func (c *ClientTimeperiods) Add(tp Timeperiod) error {
	if tp.Name == "" || tp.Alias == "" {
		return pkgerrors.New("tp.Name or tp.Alias parameter cannot be" +
			" empty when calling Add function")
	}

	values := tp.Name + ";" + tp.Alias

	respReader, err := c.CentClient.centreonAPIRequest("add", timeperiodObject,
		values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Del removes the specified timeperiod
func (c *ClientTimeperiods) Del(name string) error {
	if name == "" {
		return pkgerrors.New("name parameter cannot be empty when calling Del " +
			"function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("del", timeperiodObject,
		name)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparam is used to change a specific parameters for a timeperiod
func (c *ClientTimeperiods) Setparam(name string, param string,
	value string) error {

	if name == "" || param == "" || value == "" {
		return pkgerrors.New("name, param or value parameters cannot be empty " +
			"when calling Setparam function")
	}

	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setparam", timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setexception is used to add or change an exception for this timeperiod
func (c *ClientTimeperiods) Setexception(name string, param string,
	value string) error {

	if name == "" || param == "" || value == "" {
		return pkgerrors.New("name, param or value parameters cannot be empty " +
			"when calling Setexception function")
	}

	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setexception",
		timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getexception is used to get all exception for this timeperiod
func (c *ClientTimeperiods) Getexception(name string) ([]TimeperiodException,
	error) {

	if name == "" {
		return nil, pkgerrors.New("name parameter cannot be empty when calling " +
			"Getexception function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("getexception",
		timeperiodObject, name)
	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var tmpExs TimeperiodExceptions
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&tmpExs)

	return tmpExs.TmpEx, nil
}

// Delexception is used to remove an exception for this timeperiod
func (c *ClientTimeperiods) Delexception(name string, param string) error {
	if name == "" || param == "" {
		return pkgerrors.New("name or param parameters cannot be empty " +
			"when calling Setexception function")
	}

	values := name + ";" + param

	respReader, err := c.CentClient.centreonAPIRequest("delexception", timeperiodObject, values)
	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}
