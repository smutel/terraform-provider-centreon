package centreon

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	pkgerrors "github.com/pkg/errors"
	"github.com/smutel/go-centreon/centreonweb"
)

var choiceToInt = map[string]string{
	"no":      "0",
	"yes":     "1",
	"default": "2",
}

var intToChoice = map[string]string{
	"0": "no",
	"1": "yes",
	"2": "default",
}

var boolToInt = map[bool]string{
	true:  "1",
	false: "0",
}

var intToBool = map[string]bool{
	"1": true,
	"0": false,
	"":  false,
}

var hostParamMap = map[string]string{
	"action_url":                   "raw",
	"activate":                     "bool",
	"active_checks_enabled":        "choice",
	"acknowledgement_timeout":      "int",
	"address":                      "raw",
	"alias":                        "raw",
	"cg_additive_inheritance":      "bool",
	"check_command":                "raw",
	"check_command_arguments":      "raw",
	"check_interval":               "int",
	"check_freshness":              "choice",
	"check_period":                 "raw",
	"contact_additive_inheritance": "bool",
	"coords2d":                     "raw",
	"coords3d":                     "raw",
	"event_handler":                "raw",
	"event_handler_arguments":      "raw",
	"event_handler_enabled":        "choice",
	"first_notification_delay":     "int",
	"flap_detection_enabled":       "choice",
	"freshness_threshold":          "int",
	"high_flap_threshold":          "int",
	"icon_image":                   "raw",
	"icon_image_alt":               "raw",
	"low_flap_threshold":           "int",
	"max_check_attempts":           "int",
	"name":                         "raw",
	"notes":                        "raw",
	"notes_url":                    "raw",
	"notifications_enabled":        "choice",
	"notification_interval":        "int",
	"notification_period":          "raw",
	"recovery_notification_delay":  "int",
	"obsess_over_host":             "choice",
	"passive_checks_enabled":       "choice",
	"retain_nonstatus_information": "choice",
	"retain_status_information":    "choice",
	"retry_check_interval":         "int",
	"snmp_community":               "raw",
	"snmp_version":                 "raw",
	"statusmap_image":              "raw",
	"timezone":                     "raw",
}

func resourceCentreonHost() *schema.Resource {
	return &schema.Resource{
		Create: resourceCentreonHostCreate,
		Read:   resourceCentreonHostRead,
		Update: resourceCentreonHostUpdate,
		Delete: resourceCentreonHostDelete,
		Exists: resourceCentreonHostExists,

		Schema: map[string]*schema.Schema{
			"action_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^http://.*$|^https://.*$|^/.*$"),
					"Must start by http://, https:// or /"),
			},
			"activate": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"active_checks_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"acknowledgement_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.SingleIP(),
			},
			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},
			"cg_additive_inheritance": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"check_command": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"check_command_arguments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"check_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"check_freshness": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"check_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"coords2d": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[0-9]+,[0-9]+$"),
					"Must be like [0-9]+,[0-9]+"),
			},
			"coords3d": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile(`^[0-9]+\.[0-9]+,[0-9]+\.[0-9]+,[0-9]+\.[0-9]+$`),
					"Must be like [0-9]+.[0-9]+,[0-9]+.[0-9]+,[0-9]+.[0-9]+"),
			},
			"contact_additive_inheritance": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"event_handler": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_handler_arguments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"event_handler_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"first_notification_delay": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"flap_detection_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"flap_detection_options": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"freshness_threshold": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"high_flap_threshold": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"hostgroups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"icon_image": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"icon_image_alt": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"instance": {
				Type:     schema.TypeString,
				Required: true,
			},
			"linked_contacts": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"linked_contact_groups": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"low_flap_threshold": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"macro": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
						"is_password": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"max_check_attempts": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"notes": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes_url": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^http://.*$|^https://.*$|^/.*$"),
					"Must start by http://, https:// or /"),
			},
			"notifications_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"notification_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"notification_none": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"notification_options": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"down",
						"unreachable", "recovery", "flapping", "downtime_scheduled"}, false),
				},
				Optional:      true,
				ConflictsWith: []string{"notification_none"},
			},
			"notification_period": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"obsess_over_host": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"parents": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"passive_checks_enabled": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"process_perf_data": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"recovery_notification_delay": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"retain_nonstatus_information": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"retain_status_information": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "default",
				ValidateFunc: validation.StringInSlice([]string{"yes", "no",
					"default"}, false),
			},
			"retry_check_interval": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      -1,
				ValidateFunc: validation.IntAtLeast(0),
			},
			"snmp_community": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_version": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{"1", "2c", "3"},
					false),
			},
			"stalking_options": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					ValidateFunc: validation.StringInSlice([]string{"up", "down",
						"unreachable"}, false),
				},
				Optional: true,
			},
			"statusmap_image": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"templates": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"timezone": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCentreonHostCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.ClientCentreonWeb)
	d.Partial(true)

	hostName := d.Get("name").(string)

	if err := centreonHostCreateMain(d, m); err != nil {
		return err
	}

	for param, paramType := range hostParamMap {
		if err := hostAddParam(d, m, hostName, param, paramType); err != nil {
			return err
		}
	}

	hostTemplates := d.Get("templates").(*schema.Set).List()
	for _, template := range hostTemplates {
		if templateStr, ok := template.(string); ok {
			if err := client.Hosts().Addtemplate(hostName, templateStr); err != nil {
				return err
			}
		}
	}

	if err := client.Hosts().Applytemplates(hostName); err != nil {
		return err
	}
	d.SetPartial("templates")

	hostMacros := d.Get("macro").(*schema.Set).List()
	for _, mRaw := range hostMacros {
		m := mRaw.(map[string]interface{})
		macro := centreonweb.HostMacro{
			Name:        m["name"].(string),
			Value:       m["value"].(string),
			IsPassword:  boolToInt[m["is_password"].(bool)],
			Description: m["description"].(string),
		}

		if err := client.Hosts().Setmacro(hostName, macro); err != nil {
			return err
		}
	}
	d.SetPartial("macro")

	hostLinkedContacts := d.Get("linked_contacts").(*schema.Set).List()
	for _, contact := range hostLinkedContacts {
		if contactStr, ok := contact.(string); ok {
			if err := client.Hosts().Addcontact(hostName, contactStr); err != nil {
				return err
			}
		}
	}
	d.SetPartial("linked_contacts")

	hostLinkedCgs := d.Get("linked_contact_groups").(*schema.Set).List()
	for _, cgs := range hostLinkedCgs {
		if cgsStr, ok := cgs.(string); ok {
			if err := client.Hosts().Addcg(hostName, cgsStr); err != nil {
				return err
			}
		}
	}
	d.SetPartial("linked_contact_groups")

	notifOptions := ""
	if hostNotifNone, ok := d.GetOk("notification_none"); ok {
		if hostNotifNone.(bool) == true {
			notifOptions = "n"
		}
	} else {
		hostNotifOptions := d.Get("notification_options").(*schema.Set).List()
		notifOptions = generateOptionString(hostNotifOptions)
	}

	if err := client.Hosts().Setparam(hostName, "notification_options",
		notifOptions); err != nil {
		return err
	}
	d.SetPartial("notification_options")
	d.SetPartial("notification_none")

	hostHgs := d.Get("hostgroups").(*schema.Set).List()
	for _, hg := range hostHgs {
		if hgStr, ok := hg.(string); ok {
			if err := client.Hosts().Addhostgroup(hostName, hgStr); err != nil {
				return err
			}
		}
	}
	d.SetPartial("hostgroups")

	hostParents := d.Get("parents").(*schema.Set).List()
	for _, parent := range hostParents {
		if parentStr, ok := parent.(string); ok {
			if err := client.Hosts().Addparent(hostName, parentStr); err != nil {
				return err
			}
		}
	}
	d.SetPartial("parents")

	hostStalkingOptions := d.Get("stalking_options").(*schema.Set).List()
	stalkingOptions := generateOptionString(hostStalkingOptions)

	if err := client.Hosts().Setparam(hostName, "stalking_options",
		stalkingOptions); err != nil {
		return err
	}
	d.SetPartial("stalking_options")

	d.Partial(false)

	return resourceCentreonHostRead(d, m)
}

func resourceCentreonHostRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.ClientCentreonWeb)

	host, err := client.Hosts().Get(d.Id())
	if err != nil {
		return err
	}

	if host.Name == d.Id() {
		hostInstance, err := client.Hosts().Getinstance(d.Id())
		if err != nil {
			return err
		}

		params := "snmp_community|snmp_version|timezone|check_command" +
			"|check_command_arguments|check_period|max_check_attempts" +
			"|check_interval|retry_check_interval|active_checks_enabled" +
			"|passive_checks_enabled|notifications_enabled" +
			"|contact_additive_inheritance|cg_additive_inheritance" +
			"|notification_options|notification_interval|notification_period" +
			"|first_notification_delay|recovery_notification_delay" +
			"|obsess_over_host|acknowledgement_timeout|check_freshness" +
			"|freshness_threshold|flap_detection_enabled|low_flap_threshold" +
			"|high_flap_threshold|retain_status_information" +
			"|retain_nonstatus_information|stalking_options|event_handler_enabled" +
			"|event_handler|event_handler_arguments|notes|icon_image" +
			"|icon_image_alt|statusmap_image|2d_coords|3d_coords"

		hostParam, err := client.Hosts().Getparam(host.Name, params)
		if err != nil {
			return err
		}

		if len(hostParam) <= 0 {
			return pkgerrors.New("Unable to get parameters of host " + host.Name)
		}

		hostTemplates, err := client.Hosts().Gettemplates(host.Name)
		if err != nil {
			return err
		}

		tpls := make([]string, len(hostTemplates))
		for i, template := range hostTemplates {
			tpls[i] = template.Name
		}

		hostMacros, err := client.Hosts().Getmacro(host.Name)
		if err != nil {
			return err
		}

		macroCount := 0
		for _, macro := range hostMacros {
			if macro.Source == "direct" {
				macroCount++
			}
		}

		macros := make([]map[string]interface{}, macroCount)
		i := 0
		for _, macro := range hostMacros {
			if macro.Source == "direct" {
				m := make(map[string]interface{})
				m["name"] = macro.Name
				m["value"] = macro.Value
				m["description"] = macro.Description
				m["is_password"] = intToBool[macro.IsPassword]
				macros[i] = m
				i++
			}
		}

		hostContacts, err := client.Hosts().Getcontacts(host.Name)
		if err != nil {
			return err
		}

		cs := make([]string, len(hostContacts))
		for i, contact := range hostContacts {
			cs[i] = contact.Name
		}

		hostCgs, err := client.Hosts().Getcgs(host.Name)
		if err != nil {
			return err
		}

		cgs := make([]string, len(hostCgs))
		for i, cg := range hostCgs {
			cgs[i] = cg.Name
		}

		hostHgs, err := client.Hosts().Gethostgroups(host.Name)
		if err != nil {
			return err
		}

		hgs := make([]string, len(hostHgs))
		for i, hg := range hostHgs {
			hgs[i] = hg.Name
		}

		hostParents, err := client.Hosts().Getparents(host.Name)
		if err != nil {
			return err
		}

		parents := make([]string, len(hostParents))
		for i, parent := range hostParents {
			parents[i] = parent.Name
		}

		d.Set("name", host.Name)
		d.Set("alias", host.Alias)
		d.Set("address", host.Address)
		d.Set("instance", hostInstance)
		d.Set("activate", intToBool[host.Activate])
		d.Set("snmp_community", hostParam[0].SnmpCommunity)
		d.Set("snmp_version", hostParam[0].SnmpVersion)
		d.Set("timezone", hostParam[0].Timezone)
		d.Set("templates", tpls)
		d.Set("check_command", hostParam[0].CheckCommand)
		d.Set("check_command_arguments", hostParam[0].CheckCommandArguments)
		d.Set("check_period", hostParam[0].CheckPeriod)
		d.Set("notification_period", hostParam[0].NotificationPeriod)
		d.Set("event_handler", hostParam[0].EventHandler)
		d.Set("event_handler_arguments", hostParam[0].EventHandlerArguments)
		d.Set("notes", hostParam[0].Notes)
		d.Set("icon_image", hostParam[0].IconImage)
		d.Set("icon_image_alt", hostParam[0].IconImageAlt)
		d.Set("statusmap_image", hostParam[0].StatusmapImage)
		d.Set("coords2d", hostParam[0].Coords2D)
		d.Set("coords3d", hostParam[0].Coords3D)

		intValue := -1
		if hostParam[0].MaxCheckAttempts != "" {
			intValue, _ = strconv.Atoi(hostParam[0].MaxCheckAttempts)
		}
		d.Set("max_check_attempts", intValue)

		intValue = -1
		if hostParam[0].CheckInterval != "" {
			intValue, _ = strconv.Atoi(hostParam[0].CheckInterval)
		}
		d.Set("check_interval", intValue)

		intValue = -1
		if hostParam[0].RetryCheckInterval != "" {
			intValue, _ = strconv.Atoi(hostParam[0].RetryCheckInterval)
		}
		d.Set("retry_check_interval", intValue)

		intValue = -1
		if hostParam[0].NotificationInterval != "" {
			intValue, _ = strconv.Atoi(hostParam[0].NotificationInterval)
		}
		d.Set("notification_interval", intValue)

		intValue = -1
		if hostParam[0].FirstNotificationDelay != "" {
			intValue, _ = strconv.Atoi(hostParam[0].FirstNotificationDelay)
		}
		d.Set("first_notification_delay", intValue)

		intValue = -1
		if hostParam[0].RecoveryNotificationDelay != "" {
			intValue, _ = strconv.Atoi(hostParam[0].RecoveryNotificationDelay)
		}
		d.Set("recovery_notification_delay", intValue)

		intValue = -1
		if hostParam[0].AcknowledgementTimeout != "" {
			intValue, _ = strconv.Atoi(hostParam[0].AcknowledgementTimeout)
		}
		d.Set("acknowledgement_timeout", intValue)

		intValue = -1
		if hostParam[0].FreshnessThreshold != "" {
			intValue, _ = strconv.Atoi(hostParam[0].FreshnessThreshold)
		}
		d.Set("freshness_threshold", intValue)

		intValue = -1
		if hostParam[0].LowFlapThreshold != "" {
			intValue, _ = strconv.Atoi(hostParam[0].LowFlapThreshold)
		}
		d.Set("low_flap_threshold", intValue)

		intValue = -1
		if hostParam[0].HighFlapThreshold != "" {
			intValue, _ = strconv.Atoi(hostParam[0].HighFlapThreshold)
		}
		d.Set("high_flap_threshold", intValue)

		d.Set("active_checks_enabled",
			intToChoice[hostParam[0].ActiveChecksEnabled])

		d.Set("passive_checks_enabled",
			intToChoice[hostParam[0].PassiveChecksEnabled])

		d.Set("notifications_enabled",
			intToChoice[hostParam[0].NotificationsEnabled])

		d.Set("check_freshness",
			intToChoice[hostParam[0].CheckFreshness])

		d.Set("obsess_over_host",
			intToChoice[hostParam[0].ObsessOverHost])

		d.Set("event_handler_enabled",
			intToChoice[hostParam[0].EventHandlerEnabled])

		d.Set("retain_status_information",
			intToBool[hostParam[0].RetainStatusInformation])

		d.Set("retain_nonstatus_information",
			intToBool[hostParam[0].RetainNonstatusInformation])

		d.Set("contact_additive_inheritance",
			intToBool[hostParam[0].ContactAdditiveInheritance])

		d.Set("flap_detection_enabled",
			intToBool[hostParam[0].FlapDetectionEnabled])

		d.Set("cg_additive_inheritance",
			intToBool[hostParam[0].CgAdditiveInheritance])

		d.Set("macro", macros)
		d.Set("linked_contacts", cs)
		d.Set("linked_contact_groups", cgs)

		notifNone := false
		if hostParam[0].NotificationOptions == "n" {
			notifNone = true
		} else {
			notifOptions := generateOptionSlice(hostParam[0].NotificationOptions)
			d.Set("notification_options", notifOptions)
		}
		d.Set("notification_none", notifNone)

		d.Set("hostgroups", hgs)
		d.Set("parents", parents)

		stalkingOptions := generateOptionSlice(hostParam[0].StalkingOptions)
		d.Set("stalking_options", stalkingOptions)

		return nil
	}

	d.SetId("")

	return nil
}

func resourceCentreonHostUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.ClientCentreonWeb)

	d.Partial(true)

	for param, paramType := range hostParamMap {
		if err := hostUpdateParam(d, m, d.Id(), param, paramType); err != nil {
			return err
		}
	}

	if d.HasChange("instance") {
		if err := client.Hosts().Setinstance(d.Id(),
			d.Get("instance").(string)); err != nil {
			return err
		}

		d.SetPartial("instance")
	}

	if d.HasChange("templates") {
		hostTemplates := d.Get("templates").(*schema.Set).List()
		for _, template := range hostTemplates {
			if templateStr, ok := template.(string); ok {
				if err := client.Hosts().Addtemplate(d.Id(), templateStr); err != nil {
					return err
				}
			}
		}

		hostTemplatesSlice := expandToStringSlice(hostTemplates)
		currentHostTemplates, err := client.Hosts().Gettemplates(d.Id())
		if err != nil {
			return err
		}

		tpls := make([]string, len(currentHostTemplates))
		for i, currentTemplate := range currentHostTemplates {
			tpls[i] = currentTemplate.Name
		}

		hostTemplatesRemove := diffSlices(tpls, hostTemplatesSlice)
		for _, hostTemplateRemove := range hostTemplatesRemove {
			if err := client.Hosts().Deltemplate(d.Id(),
				hostTemplateRemove); err != nil {
				return err
			}
		}

		if err := client.Hosts().Applytemplates(d.Id()); err != nil {
			return err
		}
		d.SetPartial("templates")
	}

	if d.HasChange("macro") {
		hostMacros := d.Get("macro").(*schema.Set).List()
		hostMacrosSlice := make([]string, len(hostMacros))
		for i, mRaw := range hostMacros {
			m := mRaw.(map[string]interface{})
			hostMacrosSlice[i] = m["name"].(string)
			macro := centreonweb.HostMacro{
				Name:        m["name"].(string),
				Value:       m["value"].(string),
				IsPassword:  boolToInt[m["is_password"].(bool)],
				Description: m["description"].(string),
			}

			if err := client.Hosts().Setmacro(d.Id(), macro); err != nil {
				return err
			}
		}

		currentHostMacros, err := client.Hosts().Getmacro(d.Id())
		if err != nil {
			return err
		}

		macros := make([]string, len(currentHostMacros))
		for i, currentMacro := range currentHostMacros {
			macros[i] = currentMacro.Name
		}

		hostMacrosRemove := diffSlices(macros, hostMacrosSlice)
		for _, hostMacroRemove := range hostMacrosRemove {
			if err := client.Hosts().Delmacro(d.Id(),
				hostMacroRemove); err != nil {
				return err
			}
		}

		d.SetPartial("macro")
	}

	if d.HasChange("linked_contacts") {
		hostContacts := d.Get("linked_contacts").(*schema.Set).List()
		for _, contact := range hostContacts {
			if contactStr, ok := contact.(string); ok {
				if err := client.Hosts().Addcontact(d.Id(), contactStr); err != nil {
					return err
				}
			}
		}

		hostContactsSlice := expandToStringSlice(hostContacts)
		currentHostContacts, err := client.Hosts().Getcontacts(d.Id())
		if err != nil {
			return err
		}

		cs := make([]string, len(currentHostContacts))
		for i, currentContact := range currentHostContacts {
			cs[i] = currentContact.Name
		}

		hostContactsRemove := diffSlices(cs, hostContactsSlice)
		for _, hostContactRemove := range hostContactsRemove {
			if err := client.Hosts().Delcontact(d.Id(),
				hostContactRemove); err != nil {
				return err
			}
		}

		d.SetPartial("linked_contact")
	}

	if d.HasChange("linked_contact_groups") {
		hostCgs := d.Get("linked_contact_groups").(*schema.Set).List()
		for _, cg := range hostCgs {
			if cgStr, ok := cg.(string); ok {
				if err := client.Hosts().Addcg(d.Id(), cgStr); err != nil {
					return err
				}
			}
		}

		hostCgsSlice := expandToStringSlice(hostCgs)
		currentHostCgs, err := client.Hosts().Getcgs(d.Id())
		if err != nil {
			return err
		}

		cgs := make([]string, len(currentHostCgs))
		for i, currentCg := range currentHostCgs {
			cgs[i] = currentCg.Name
		}

		hostCgsRemove := diffSlices(cgs, hostCgsSlice)
		for _, hostCgRemove := range hostCgsRemove {
			if err := client.Hosts().Delcg(d.Id(), hostCgRemove); err != nil {
				return err
			}
		}

		d.SetPartial("linked_contact")
	}

	if d.HasChange("notification_options") {
		hostNotifOptions := d.Get("notification_options").(*schema.Set).List()
		notifOptions := generateOptionString(hostNotifOptions)

		if err := client.Hosts().Setparam(d.Id(), "notification_options",
			notifOptions); err != nil {
			return err
		}
		d.SetPartial("notification_options")
	}

	if d.HasChange("notification_none") {
		hostNotifNone := d.Get("notification_none").(bool)
		if hostNotifNone == true {
			if err := client.Hosts().Setparam(d.Id(), "notification_options",
				"n"); err != nil {
				return err
			}
		}
		d.SetPartial("notification_none")
	}

	if d.HasChange("hostgroups") {
		hostHgs := d.Get("hostgroups").(*schema.Set).List()
		for _, hg := range hostHgs {
			if hgStr, ok := hg.(string); ok {
				if err := client.Hosts().Addhostgroup(d.Id(), hgStr); err != nil {
					return err
				}
			}
		}

		hostHgsSlice := expandToStringSlice(hostHgs)
		currentHostHgs, err := client.Hosts().Gethostgroups(d.Id())
		if err != nil {
			return err
		}

		hgs := make([]string, len(currentHostHgs))
		for i, currentHg := range currentHostHgs {
			hgs[i] = currentHg.Name
		}

		hostHgsRemove := diffSlices(hgs, hostHgsSlice)
		for _, hostHgRemove := range hostHgsRemove {
			if err := client.Hosts().Delhostgroup(d.Id(), hostHgRemove); err != nil {
				return err
			}
		}

		d.SetPartial("hostgroups")
	}

	if d.HasChange("parents") {
		hostParents := d.Get("parents").(*schema.Set).List()
		for _, parent := range hostParents {
			if parentStr, ok := parent.(string); ok {
				if err := client.Hosts().Addparent(d.Id(), parentStr); err != nil {
					return err
				}
			}
		}

		hostParentsSlice := expandToStringSlice(hostParents)
		currentHostParents, err := client.Hosts().Getparents(d.Id())
		if err != nil {
			return err
		}

		parents := make([]string, len(currentHostParents))
		for i, currentParent := range currentHostParents {
			parents[i] = currentParent.Name
		}

		hostParentsRemove := diffSlices(parents, hostParentsSlice)
		for _, hostParentRemove := range hostParentsRemove {
			if err := client.Hosts().Delparent(d.Id(),
				hostParentRemove); err != nil {
				return err
			}
		}

		d.SetPartial("parents")
	}

	if d.HasChange("stalking_options") {
		hostStalkingOptions := d.Get("stalking_options").(*schema.Set).List()
		stalkingOptions := generateOptionString(hostStalkingOptions)

		if err := client.Hosts().Setparam(d.Id(), "stalking_options",
			stalkingOptions); err != nil {
			return err
		}
		d.SetPartial("stalking_options")
	}

	d.Partial(false)

	return resourceCentreonHostRead(d, m)
}

func resourceCentreonHostDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.ClientCentreonWeb)
	resourceExists, err := resourceCentreonHostExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	if err := client.Hosts().Del(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCentreonHostExists(d *schema.ResourceData, m interface{}) (b bool,
	e error) {
	client := m.(*centreonweb.ClientCentreonWeb)
	return client.Hosts().Exists(d.Id())
}

func generateOptionString(notif []interface{}) string {
	optionsStr := ""

	for _, option := range notif {
		if option.(string) == "up" {
			if optionsStr != "" {
				optionsStr += ","
			}
			optionsStr += "o"
		} else if option.(string) == "down" {
			if optionsStr != "" {
				optionsStr += ","
			}
			optionsStr += "d"
		} else if option == "unreachable" {
			if optionsStr != "" {
				optionsStr += ","
			}
			optionsStr += "u"
		} else if option.(string) == "recovery" {
			if optionsStr != "" {
				optionsStr += ","
			}
			optionsStr += "r"
		} else if option.(string) == "flapping" {
			if optionsStr != "" {
				optionsStr += ","
			}
			optionsStr += "f"
		} else if option.(string) == "downtime_scheduled" {
			if optionsStr != "" {
				optionsStr += ","
			}
			optionsStr += "s"
		}
	}

	return optionsStr
}

func generateOptionSlice(option string) []interface{} {
	var options []interface{}

	if strings.Contains(option, "o") {
		options = append(options, "up")
	}

	if strings.Contains(option, "d") {
		options = append(options, "down")
	}

	if strings.Contains(option, "u") {
		options = append(options, "unreachable")
	}

	if strings.Contains(option, "r") {
		options = append(options, "recovery")
	}

	if strings.Contains(option, "f") {
		options = append(options, "flapping")
	}

	if strings.Contains(option, "s") {
		options = append(options, "downtime_scheduled")
	}

	return options
}

func centreonHostCreateMain(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.ClientCentreonWeb)

	hostActivateStr := "0"

	hostName := d.Get("name").(string)
	hostAlias := d.Get("alias").(string)
	hostAddress := d.Get("address").(string)
	hostInstance := d.Get("instance").(string)
	hostActivate := d.Get("activate").(bool)

	if hostActivate {
		hostActivateStr = "1"
	}

	host := centreonweb.Host{
		Name:     hostName,
		Alias:    hostAlias,
		Address:  hostAddress,
		Activate: hostActivateStr,
	}

	if err := client.Hosts().Add(host, hostInstance); err != nil {
		return err
	}

	d.SetId(hostName)
	d.SetPartial("name")
	d.SetPartial("address")
	d.SetPartial("instance")
	d.SetPartial("activate")

	return nil
}

func hostAddParam(d *schema.ResourceData, m interface{}, hostName string,
	param string, paramType string) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	hostParam := d.Get(param)

	value := ""
	if paramType == "int" {
		value = strconv.Itoa(hostParam.(int))
		if value == "-1" {
			value = ""
		}
	} else if paramType == "bool" {
		value = boolToInt[hostParam.(bool)]
	} else if paramType == "choice" {
		value = choiceToInt[hostParam.(string)]
	} else {
		value = hostParam.(string)
	}

	paramTmp := ""
	if param == "coords2d" {
		paramTmp = "2d_coords"
	} else if param == "coords3d" {
		paramTmp = "3d_coords"
	} else {
		paramTmp = param
	}

	if err := client.Hosts().Setparam(hostName, paramTmp, value); err != nil {
		return err
	}

	d.SetPartial(param)

	return nil
}

func hostUpdateParam(d *schema.ResourceData, m interface{}, hostName string,
	param string, paramType string) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	if d.HasChange(param) {
		value := ""
		if paramType == "int" {
			value = strconv.Itoa(d.Get(param).(int))
			if value == "-1" {
				value = ""
			}
		} else if paramType == "bool" {
			value = boolToInt[d.Get(param).(bool)]
		} else if paramType == "choice" {
			value = choiceToInt[d.Get(param).(string)]
		} else {
			value = d.Get(param).(string)
		}

		paramTmp := ""
		if param == "coords2d" {
			paramTmp = "2d_coords"
		} else if param == "coords3d" {
			paramTmp = "3d_coords"
		} else {
			paramTmp = param
		}

		if err := client.Hosts().Setparam(hostName, paramTmp, value); err != nil {
			return err
		}

		if param == "name" {
			d.SetId(d.Get("name").(string))
		}

		d.SetPartial(param)
	}

	return nil
}
