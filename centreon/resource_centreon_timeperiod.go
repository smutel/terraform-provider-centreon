package centreon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/smutel/go-centreon/centreonweb"
)

func resourceCentreonTimeperiod() *schema.Resource {
	return &schema.Resource{
		Create: resourceCentreonTimeperiodCreate,
		Read:   resourceCentreonTimeperiodRead,
		Update: resourceCentreonTimeperiodUpdate,
		Delete: resourceCentreonTimeperiodDelete,
		Exists: resourceCentreonTimeperiodExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"alias": {
				Type:     schema.TypeString,
				Required: true,
			},

			centreonweb.TimeperiodSunday: {
				Type:     schema.TypeString,
				Optional: true,
			},

			centreonweb.TimeperiodMonday: {
				Type:     schema.TypeString,
				Optional: true,
			},

			centreonweb.TimeperiodTuesday: {
				Type:     schema.TypeString,
				Optional: true,
			},

			centreonweb.TimeperiodWednesday: {
				Type:     schema.TypeString,
				Optional: true,
			},

			centreonweb.TimeperiodThursday: {
				Type:     schema.TypeString,
				Optional: true,
			},

			centreonweb.TimeperiodFriday: {
				Type:     schema.TypeString,
				Optional: true,
			},

			centreonweb.TimeperiodSaturday: {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceCentreonTimeperiodCreate(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	d.Partial(true)

	tmpName := d.Get("name").(string)
	tmpAlias := d.Get("alias").(string)
	tmp := centreonweb.Timeperiod{
		Name:  tmpName,
		Alias: tmpAlias,
	}

	if err := client.Timeperiods().Add(tmp); err != nil {
		return err
	}
	d.SetPartial("name")
	d.SetPartial("alias")

	tmpSunday := d.Get(centreonweb.TimeperiodSunday).(string)
	if tmpSunday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodSunday, tmpSunday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodSunday)
	}

	tmpMonday := d.Get(centreonweb.TimeperiodMonday).(string)
	if tmpMonday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodMonday, tmpMonday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodMonday)
	}

	tmpTuesday := d.Get(centreonweb.TimeperiodTuesday).(string)
	if tmpTuesday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodTuesday, tmpTuesday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodTuesday)
	}

	tmpWednesday := d.Get(centreonweb.TimeperiodWednesday).(string)
	if tmpWednesday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodWednesday, tmpWednesday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodWednesday)
	}

	tmpThursday := d.Get(centreonweb.TimeperiodThursday).(string)
	if tmpThursday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodThursday, tmpThursday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodThursday)
	}

	tmpFriday := d.Get(centreonweb.TimeperiodFriday).(string)
	if tmpFriday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodFriday, tmpFriday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodFriday)
	}

	tmpSaturday := d.Get(centreonweb.TimeperiodSaturday).(string)
	if tmpSaturday != "" {
		if err := client.Timeperiods().Setparam(tmpName,
			centreonweb.TimeperiodSaturday, tmpSaturday); err != nil {
			return err
		}
		d.SetPartial(centreonweb.TimeperiodSaturday)
	}

	d.Partial(false)
	d.SetId(tmpName)

	return resourceCentreonTimeperiodRead(d, m)
}

func resourceCentreonTimeperiodRead(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	tmp, err := client.Timeperiods().Get(d.Id())
	if err != nil {
		return err
	}

	if tmp.Name == d.Id() {
		d.Set("name", tmp.Name)
		d.Set("alias", tmp.Alias)
		d.Set("sunday", tmp.Sunday)
		d.Set("monday", tmp.Monday)
		d.Set("tuesday", tmp.Tuesday)
		d.Set("wednesday", tmp.Wednesday)
		d.Set("thursday", tmp.Thursday)
		d.Set("friday", tmp.Friday)
		d.Set("saturday", tmp.Saturday)
		return nil
	}

	d.SetId("")
	return nil
}

func resourceCentreonTimeperiodUpdate(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)
	d.Partial(true)

	if d.HasChange("name") {
		if err := client.Timeperiods().Setparam(d.Id(), "name",
			d.Get("name").(string)); err != nil {
			return err
		}

		d.SetId(d.Get("name").(string))
		d.SetPartial("name")
	}

	if d.HasChange("alias") {
		if err := client.Timeperiods().Setparam(d.Id(), "alias",
			d.Get("alias").(string)); err != nil {
			return err
		}

		d.SetPartial("alias")
	}

	if d.HasChange("sunday") {
		if err := client.Timeperiods().Setparam(d.Id(), "sunday",
			d.Get("sunday").(string)); err != nil {
			return err
		}

		d.SetPartial("sunday")
	}

	if d.HasChange("monday") {
		if err := client.Timeperiods().Setparam(d.Id(), "monday",
			d.Get("monday").(string)); err != nil {
			return err
		}

		d.SetPartial("monday")
	}

	if d.HasChange("tuesday") {
		if err := client.Timeperiods().Setparam(d.Id(), "tuesday",
			d.Get("tuesday").(string)); err != nil {
			return err
		}

		d.SetPartial("tuesday")
	}

	if d.HasChange("wednesday") {
		if err := client.Timeperiods().Setparam(d.Id(), "wednesday",
			d.Get("wednesday").(string)); err != nil {
			return err
		}

		d.SetPartial("wednesday")
	}

	if d.HasChange("thursday") {
		if err := client.Timeperiods().Setparam(d.Id(), "thursday",
			d.Get("thursday").(string)); err != nil {
			return err
		}

		d.SetPartial("thursday")
	}

	if d.HasChange("friday") {
		if err := client.Timeperiods().Setparam(d.Id(), "friday",
			d.Get("friday").(string)); err != nil {
			return err
		}

		d.SetPartial("friday")
	}

	if d.HasChange("saturday") {
		if err := client.Timeperiods().Setparam(d.Id(), "saturday",
			d.Get("saturday").(string)); err != nil {
			return err
		}

		d.SetPartial("saturday")
	}

	d.Partial(false)

	return resourceCentreonTimeperiodRead(d, m)
}

func resourceCentreonTimeperiodDelete(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)
	resourceExists, err := resourceCentreonTimeperiodExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	if err := client.Timeperiods().Del(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCentreonTimeperiodExists(d *schema.ResourceData,
	m interface{}) (b bool, e error) {

	client := m.(*centreonweb.ClientCentreonWeb)
	return client.Timeperiods().Exists(d.Id())
}
