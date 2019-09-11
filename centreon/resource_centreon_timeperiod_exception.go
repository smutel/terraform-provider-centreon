package centreon

import (
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	pkgerrors "github.com/pkg/errors"
	"github.com/smutel/go-centreon/centreonweb"
)

func resourceCentreonTimeperiodException() *schema.Resource {
	return &schema.Resource{
		Create: resourceCentreonTimeperiodExceptionCreate,
		Read:   resourceCentreonTimeperiodExceptionRead,
		Update: resourceCentreonTimeperiodExceptionUpdate,
		Delete: resourceCentreonTimeperiodExceptionDelete,
		Exists: resourceCentreonTimeperiodExceptionExists,

		Schema: map[string]*schema.Schema{
			"timeperiod_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"days": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"timerange": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCentreonTimeperiodExceptionCreate(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	tmpID := d.Get("timeperiod_id").(string)

	if _, err := client.Timeperiods().Get(tmpID); err != nil {
		return pkgerrors.New("timeperiod with this timeperiod " +
			tmpID + " not found")
	}

	tmpExDays := d.Get("days").(string)
	tmpExTimerange := d.Get("timerange").(string)

	if err := client.Timeperiods().Setexception(tmpID, tmpExDays,
		tmpExTimerange); err != nil {
		return err
	}

	d.SetId(tmpID + "_" + tmpExDays)

	return resourceCentreonTimeperiodExceptionRead(d, m)
}

func resourceCentreonTimeperiodExceptionRead(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	tmpID, tmpExDays, err := extractTimeperiodID(d.Id())
	if err != nil {
		return err
	}

	tmp, err := client.Timeperiods().Getexception(tmpID)
	if err != nil {
		return err
	}

	for _, value := range tmp {
		if value.Days == tmpExDays {
			d.Set("days", tmpExDays)
			d.Set("timerange", value.Timerange)
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceCentreonTimeperiodExceptionUpdate(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	tmpID, tmpExDays, err := extractTimeperiodID(d.Id())
	if err != nil {
		return err
	}

	if d.HasChange("timerange") {
		if err := client.Timeperiods().Setexception(tmpID, tmpExDays,
			d.Get("timerange").(string)); err != nil {
			return err
		}
	}

	return resourceCentreonTimeperiodExceptionRead(d, m)
}

func resourceCentreonTimeperiodExceptionDelete(d *schema.ResourceData,
	m interface{}) error {

	client := m.(*centreonweb.ClientCentreonWeb)

	resourceExists, err := resourceCentreonTimeperiodExceptionExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	tmpID, tmpExDays, err := extractTimeperiodID(d.Id())
	if err != nil {
		return err
	}

	if err := client.Timeperiods().Delexception(tmpID, tmpExDays); err != nil {
		return err
	}

	return nil
}

func resourceCentreonTimeperiodExceptionExists(d *schema.ResourceData,
	m interface{}) (b bool, e error) {

	client := m.(*centreonweb.ClientCentreonWeb)
	exist := false

	tmpID, tmpExDays, err := extractTimeperiodID(d.Id())
	if err != nil {
		return exist, err
	}

	tmp, err := client.Timeperiods().Getexception(tmpID)
	if err != nil {
		return exist, err
	}

	for _, value := range tmp {
		if value.Days == tmpExDays {
			exist = true
		}
	}

	return exist, nil
}

func extractTimeperiodID(ID string) (string, string, error) {
	tmpIDSlice := strings.Split(ID, "_")
	if len(tmpIDSlice) != 2 {
		return "", "", pkgerrors.New("unable to extract timeperiod name and " +
			"timeperiod exception days from terraform Id")
	}

	tmpID := tmpIDSlice[0]
	tmpExDays := tmpIDSlice[1]
	return tmpID, tmpExDays, nil
}
