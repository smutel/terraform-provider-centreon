package centreonweb

import (
	"encoding/json"

	pkgerrors "github.com/pkg/errors"
)

const hostObject string = "HOST"

// ClientHosts is used to store the client to interact with the Centreon API
type ClientHosts struct {
	CentClient *ClientCentreonWeb
}

// Hosts is an array of Host to store the answer from Centreon API
type Hosts struct {
	Host []Host `json:"result"`
}

// Host struct is used to store parameters of a host
type Host struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	Address  string `json:"address"`
	Activate string `json:"activate"`
}

// HostParams is an array of HostParam to store the answer from Centreon API
type HostParams struct {
	HostParam []HostParam `json:"result"`
}

// HostParam struct is used to store result of function getparam
type HostParam struct {
	ActionURL                  string `json:"action_url"`
	Activate                   string `json:"activate"`
	ActiveChecksEnabled        string `json:"active_checks_enabled"`
	AcknowledgementTimeout     string `json:"acknowledgement_timeout"`
	Address                    string `json:"address"`
	Alias                      string `json:"alias"`
	CheckCommand               string `json:"check_command"`
	CheckCommandArguments      string `json:"check_command_arguments"`
	CheckInterval              string `json:"check_interval"`
	CheckFreshness             string `json:"check_freshness"`
	CheckPeriod                string `json:"check_period"`
	Coords2D                   string `json:"2d_coords"`
	Coords3D                   string `json:"3d_coords"`
	ContactAdditiveInheritance string `json:"contact_additive_inheritance"`
	CgAdditiveInheritance      string `json:"cg_additive_inheritance"`
	EventHandler               string `json:"event_handler"`
	EventHandlerArguments      string `json:"event_handler_arguments"`
	EventHandlerEnabled        string `json:"event_handler_enabled"`
	FirstNotificationDelay     string `json:"first_notification_delay"`
	FlapDetectionEnabled       string `json:"flap_detection_enabled"`
	FlapDetectionOptions       string `json:"flap_detection_options"`
	FreshnessThreshold         string `json:"freshness_threshold"`
	HighFlapThreshold          string `json:"high_flap_threshold"`
	IconImage                  string `json:"icon_image"`
	IconImageAlt               string `json:"icon_image_alt"`
	LowFlapThreshold           string `json:"low_flap_threshold"`
	MaxCheckAttempts           string `json:"max_check_attempts"`
	Name                       string `json:"name"`
	Notes                      string `json:"notes"`
	NotesURL                   string `json:"notes_url"`
	NotificationsEnabled       string `json:"notifications_enabled"`
	NotificationInterval       string `json:"notification_interval"`
	NotificationOptions        string `json:"notification_options"`
	NotificationPeriod         string `json:"notification_period"`
	RecoveryNotificationDelay  string `json:"recovery_notification_delay"`
	ObsessOverHost             string `json:"obsess_over_host"`
	PassiveChecksEnabled       string `json:"passive_checks_enabled"`
	ProcessPerfData            string `json:"process_perf_data"`
	RetainNonstatusInformation string `json:"retain_nonstatus_information"`
	RetainStatusInformation    string `json:"retain_status_information"`
	RetryCheckInterval         string `json:"retry_check_interval"`
	SnmpCommunity              string `json:"snmp_community"`
	SnmpVersion                string `json:"snmp_version"`
	StalkingOptions            string `json:"stalking_options"`
	StatusmapImage             string `json:"statusmap_image"`
	Timezone                   string `json:"timezone"`
}

// Instances is an array of Instance to store the answer from Centreon API
type Instances struct {
	Instance []Instance `json:"result"`
}

// Instance struct is used to store parameters of an instance
type Instance struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HostMacros is an array of HostMacro to store the answer from Centreon API
type HostMacros struct {
	HostMacro []HostMacro `json:"result"`
}

// HostMacro struct is used to store parameters of a macro
type HostMacro struct {
	Name        string `json:"macro name"`
	Value       string `json:"macro value"`
	IsPassword  string `json:"is_password"`
	Description string `json:"description"`
	Source      string `json:"source"`
}

// HostTemplates is an array of HostTemplate to store the answer from
// Centreon API
type HostTemplates struct {
	HostTemplate []HostTemplate `json:"result"`
}

// HostTemplate struct is used to store parameters of a macro
type HostTemplate struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HostParents is an array of HostParent to store the answer from
// Centreon API
type HostParents struct {
	HostParent []HostParent `json:"result"`
}

// HostParent struct is used to store parameters of a parent
type HostParent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HostCgs is an array of HostCg to store the answer from
// Centreon API
type HostCgs struct {
	HostCg []HostCg `json:"result"`
}

// HostCg struct is used to store parameters of a contactgroup
type HostCg struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HostContacts is an array of HostContact to store the answer from
// Centreon API
type HostContacts struct {
	HostContact []HostContact `json:"result"`
}

// HostContact struct is used to store parameters of a contact
type HostContact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// HostGroups is an array of HostGroup to store the answer from
// Centreon API
type HostGroups struct {
	HostGroup []HostGroup `json:"result"`
}

// HostGroup struct is used to store parameters of a contact
type HostGroup struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Show lists available hosts
func (c *ClientHosts) Show(name string) ([]Host, error) {
	respReader, err := c.CentClient.centreonAPIRequest("show", hostObject,
		name)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hosts Hosts
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hosts)

	return hosts.Host, nil
}

// Get returns a specific host
func (c *ClientHosts) Get(name string) (Host, error) {
	var objFound Host

	if name == "" {
		return objFound, pkgerrors.New("name parameter cannot be empty when " +
			"calling Get function")
	}

	hosts, err := c.Show(name)
	if err != nil {
		return objFound, err
	}

	for _, h := range hosts {
		if h.Name == name {
			objFound = h
		}
	}

	if objFound.ID == "" {
		return objFound, pkgerrors.New("host " + name + " not found")
	}

	return objFound, nil
}

// Exists returns true if the host exists, false otherwise
func (c *ClientHosts) Exists(name string) (bool, error) {
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

// Add adds a new host
func (c *ClientHosts) Add(host Host, instance string) error {
	values := host.Name + ";" + host.Alias + ";" + host.Address + ";;" + instance + ";"

	respReader, err := c.CentClient.centreonAPIRequest("add", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Del removes the specified host
func (c *ClientHosts) Del(name string) error {
	respReader, err := c.CentClient.centreonAPIRequest("del", hostObject,
		name)

	if name == "" {
		return pkgerrors.New("name parameter cannot be empty when calling Del " +
			"function")
	}

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparam is used to change a specific parameters for a host
func (c *ClientHosts) Setparam(name string, param string,
	value string) error {

	if name == "" || param == "" {
		return pkgerrors.New("name or param parameters cannot be empty when " +
			"calling Setparam function")
	}

	values := name + ";" + param + ";" + value

	respReader, err := c.CentClient.centreonAPIRequest("setparam", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getparam is used to retrieve a specific parameters for a host
func (c *ClientHosts) Getparam(name string, param string) ([]HostParam, error) {

	if name == "" || param == "" {
		return nil, pkgerrors.New("name or param parameters cannot be empty when " +
			"calling Getparam function")
	}

	// Workaround to be sure we have an array as return value
	values := name + ";" + "name|" + param

	respReader, err := c.CentClient.centreonAPIRequest("getparam", hostObject,
		values)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hps HostParams
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hps)

	return hps.HostParam, nil
}

// Setinstance is used to set the instance (poller) for a host
func (c *ClientHosts) Setinstance(name string, instance string) error {

	if name == "" || instance == "" {
		return pkgerrors.New("name or instance parameters cannot be empty " +
			"when calling Setinstance function")
	}

	values := name + ";" + instance

	respReader, err := c.CentClient.centreonAPIRequest("setinstance", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getinstance is used to retrieve the instance (poller) of a host
func (c *ClientHosts) Getinstance(name string) (string, error) {
	if name == "" {
		return "", pkgerrors.New("name parameter cannot be empty when calling " +
			"Getinstance function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("showinstance",
		hostObject, name)

	if err != nil {
		return "", err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var is Instances
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&is)

	if len(is.Instance) != 1 {
		return "", pkgerrors.New("too much results from API (!=1) when calling " +
			"Getinstance function")
	}

	return is.Instance[0].Name, nil
}

// Setmacro is used to add or to update a macro linked to a host
func (c *ClientHosts) Setmacro(hostName string, macro HostMacro) error {
	if hostName == "" || macro.Name == "" {
		return pkgerrors.New("hostName or macro.Name parameters cannot be empty " +
			"when calling Setmacro function")
	}

	if macro.IsPassword != "0" && macro.IsPassword != "1" {
		return pkgerrors.New("macro.IsPassword parameters should be equal to 0 " +
			"or 1 in Setmacro function")
	}

	values := hostName + ";" + macro.Name + ";" + macro.Value + ";" +
		macro.IsPassword + ";" + macro.Description

	respReader, err := c.CentClient.centreonAPIRequest("setmacro", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getmacro is used to get a macro linked to a host
func (c *ClientHosts) Getmacro(hostName string) ([]HostMacro, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Getmacro function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("getmacro",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var ms HostMacros
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&ms)

	return ms.HostMacro, nil
}

// Delmacro is used to delete a macro linked to a host
func (c *ClientHosts) Delmacro(hostName string, macroName string) error {

	if hostName == "" || macroName == "" {
		return pkgerrors.New("hostName or macroName parameter cannot be " +
			"empty when calling Delmacro function")
	}

	values := hostName + ";" + macroName

	respReader, err := c.CentClient.centreonAPIRequest("delmacro", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Gettemplates is used to get a template linked to a host
func (c *ClientHosts) Gettemplates(hostName string) ([]HostTemplate, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Gettemplates function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("gettemplate",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var ms HostTemplates
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&ms)

	return ms.HostTemplate, nil
}

// Addtemplate is used to link a template to a host
func (c *ClientHosts) Addtemplate(hostName string, template string) error {
	if hostName == "" || template == "" {
		return pkgerrors.New("hostName or template parameters cannot be empty " +
			"when calling Addtemplate function")
	}

	values := hostName + ";" + template

	respReader, err := c.CentClient.centreonAPIRequest("addtemplate", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Deltemplate is used to delete a template linked to a host
func (c *ClientHosts) Deltemplate(hostName string, templateName string) error {

	if hostName == "" || templateName == "" {
		return pkgerrors.New("hostName or templateName parameter cannot be " +
			"empty when calling Deltemplate function")
	}

	values := hostName + ";" + templateName

	respReader, err := c.CentClient.centreonAPIRequest("deltemplate", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Applytemplates is used to apply templates linked to a host
func (c *ClientHosts) Applytemplates(hostName string) error {

	if hostName == "" {
		return pkgerrors.New("hostName parameter cannot be empty when calling " +
			"Applytemplates function")
	}

	values := hostName

	respReader, err := c.CentClient.centreonAPIRequest("applytpl", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getparents is used to get a parent linked to a host
func (c *ClientHosts) Getparents(hostName string) ([]HostParent, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Getparents function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("getparent",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hp HostParents
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hp)

	return hp.HostParent, nil
}

// Addparent is used to link a parent to a host
func (c *ClientHosts) Addparent(hostName string, parent string) error {
	if hostName == "" || parent == "" {
		return pkgerrors.New("hostName or parent parameters cannot be empty " +
			"when calling Addparent function")
	}

	values := hostName + ";" + parent

	respReader, err := c.CentClient.centreonAPIRequest("addparent", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setparent is used to replace parents linked to a host
func (c *ClientHosts) Setparent(hostName string, parents string) error {
	if hostName == "" || parents == "" {
		return pkgerrors.New("hostName or parent parameters cannot be empty " +
			"when calling Setparent function")
	}

	values := hostName + ";" + parents

	respReader, err := c.CentClient.centreonAPIRequest("setparent", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Delparent is used to delete a parent linked to a host
func (c *ClientHosts) Delparent(hostName string, parentName string) error {

	if hostName == "" || parentName == "" {
		return pkgerrors.New("hostName or parentName parameter cannot be " +
			"empty when calling Delparent function")
	}

	values := hostName + ";" + parentName

	respReader, err := c.CentClient.centreonAPIRequest("delparent", hostObject,
		values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getcgs is used to get contactgroups linked to a host
func (c *ClientHosts) Getcgs(hostName string) ([]HostCg, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Getcgs function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("getcontactgroup",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hcg HostCgs
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hcg)

	return hcg.HostCg, nil
}

// Addcg is used to link a contactgroup to a host
func (c *ClientHosts) Addcg(hostName string, cg string) error {
	if hostName == "" || cg == "" {
		return pkgerrors.New("hostName or cg parameters cannot be empty " +
			"when calling Addcg function")
	}

	values := hostName + ";" + cg

	respReader, err := c.CentClient.centreonAPIRequest("addcontactgroup",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setcg is used to replace contactgroups linked to a host
func (c *ClientHosts) Setcg(hostName string, cgs string) error {
	if hostName == "" || cgs == "" {
		return pkgerrors.New("hostName or cg parameters cannot be empty " +
			"when calling Setcg function")
	}

	values := hostName + ";" + cgs

	respReader, err := c.CentClient.centreonAPIRequest("setcontactgroup",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Delcg is used to delete a contactgroup linked to a host
func (c *ClientHosts) Delcg(hostName string, cgName string) error {

	if hostName == "" || cgName == "" {
		return pkgerrors.New("hostName or cgName parameter cannot be " +
			"empty when calling Delcg function")
	}

	values := hostName + ";" + cgName

	respReader, err := c.CentClient.centreonAPIRequest("delcontactgroup",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Getcontacts is used to get contacts linked to a host
func (c *ClientHosts) Getcontacts(hostName string) ([]HostContact, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Getcontacts function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("getcontact",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hc HostContacts
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hc)

	return hc.HostContact, nil
}

// Addcontact is used to link a contact to a host
func (c *ClientHosts) Addcontact(hostName string, contact string) error {
	if hostName == "" || contact == "" {
		return pkgerrors.New("hostName or contact parameters cannot be empty " +
			"when calling Addcontact function")
	}

	values := hostName + ";" + contact

	respReader, err := c.CentClient.centreonAPIRequest("addcontact",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setcontact is used to replace contactgroups linked to a host
func (c *ClientHosts) Setcontact(hostName string, contacts string) error {
	if hostName == "" || contacts == "" {
		return pkgerrors.New("hostName or contact parameters cannot be empty " +
			"when calling Setcontact function")
	}

	values := hostName + ";" + contacts

	respReader, err := c.CentClient.centreonAPIRequest("setcontact",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Delcontact is used to delete a contact linked to a host
func (c *ClientHosts) Delcontact(hostName string, contact string) error {

	if hostName == "" || contact == "" {
		return pkgerrors.New("hostName or contact parameter cannot be " +
			"empty when calling Delcontact function")
	}

	values := hostName + ";" + contact

	respReader, err := c.CentClient.centreonAPIRequest("delcontact",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Gethostgroups is used to get hostgroups linked to a host
func (c *ClientHosts) Gethostgroups(hostName string) ([]HostGroup, error) {
	if hostName == "" {
		return nil, pkgerrors.New("hostName parameter cannot be empty when " +
			"calling Getcontacts function")
	}

	respReader, err := c.CentClient.centreonAPIRequest("gethostgroup",
		hostObject, hostName)

	if err != nil {
		return nil, err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	var hgs HostGroups
	decoder := json.NewDecoder(respReader)
	decoder.Decode(&hgs)

	return hgs.HostGroup, nil
}

// Addhostgroup is used to link a contact to a host
func (c *ClientHosts) Addhostgroup(hostName string, hostgroup string) error {
	if hostName == "" || hostgroup == "" {
		return pkgerrors.New("hostName or hostgroup parameters cannot be empty " +
			"when calling Addhostgroup function")
	}

	values := hostName + ";" + hostgroup

	respReader, err := c.CentClient.centreonAPIRequest("addhostgroup",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Sethostgroup is used to replace hostgroups linked to a host
func (c *ClientHosts) Sethostgroup(hostName string, hostgroups string) error {
	if hostName == "" || hostgroups == "" {
		return pkgerrors.New("hostName or hostgroups parameters cannot be empty " +
			"when calling Sethostgroup function")
	}

	values := hostName + ";" + hostgroups

	respReader, err := c.CentClient.centreonAPIRequest("sethostgroup",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Delhostgroup is used to delete a hostgroup linked to a host
func (c *ClientHosts) Delhostgroup(hostName string, hostgroup string) error {

	if hostName == "" || hostgroup == "" {
		return pkgerrors.New("hostName or hostgroup parameter cannot be " +
			"empty when calling Delhostgroup function")
	}

	values := hostName + ";" + hostgroup

	respReader, err := c.CentClient.centreonAPIRequest("delhostgroup",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Setseverity is used to set the severity of a host
func (c *ClientHosts) Setseverity(hostName string, severity string) error {

	if hostName == "" {
		return pkgerrors.New("hostName parameter cannot be empty when calling " +
			"Setseverity function")
	}

	values := hostName + ";" + severity

	respReader, err := c.CentClient.centreonAPIRequest("setseverity",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Unsetseverity is used to unset the severity of a host
func (c *ClientHosts) Unsetseverity(hostName string) error {

	if hostName == "" {
		return pkgerrors.New("hostName parameter cannot be empty when calling " +
			"Unsetseverity function")
	}

	values := hostName

	respReader, err := c.CentClient.centreonAPIRequest("unsetseverity",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Enable is used to enable a host
func (c *ClientHosts) Enable(hostName string) error {

	if hostName == "" {
		return pkgerrors.New("hostName parameter cannot be empty when calling " +
			"Enable function")
	}

	values := hostName

	respReader, err := c.CentClient.centreonAPIRequest("enable",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}

// Disable is used to disable a host
func (c *ClientHosts) Disable(hostName string) error {

	if hostName == "" {
		return pkgerrors.New("hostName parameter cannot be empty when calling " +
			"Disable function")
	}

	values := hostName

	respReader, err := c.CentClient.centreonAPIRequest("disable",
		hostObject, values)

	if err != nil {
		return err
	}

	if respReader != nil {
		defer respReader.Close()
	}

	return nil
}
