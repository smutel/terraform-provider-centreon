package centreon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/smutel/go-centreon/centreonweb"
)

func resourceCentreonCommand() *schema.Resource {
	return &schema.Resource{
		Create: resourceCentreonCommandCreate,
		Read:   resourceCentreonCommandRead,
		Update: resourceCentreonCommandUpdate,
		Delete: resourceCentreonCommandDelete,
		Exists: resourceCentreonCommandExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"line": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCentreonCommandCreate(d *schema.ResourceData, m interface{}) error {
	cmd_name := d.Get("name").(string)
	cmd_type := d.Get("type").(string)
	cmd_line := d.Get("line").(string)
	client := m.(*centreonweb.CentreonwebClient)

	cmd := centreonweb.Command{
		Name: cmd_name,
		Type: cmd_type,
		Line: cmd_line,
	}

	if err := client.Commands().Add(cmd); err != nil {
		return err
	}

	d.SetId(cmd_name)

	return resourceCentreonCommandRead(d, m)
}

func resourceCentreonCommandRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.CentreonwebClient)

	cmd, err := client.Commands().Get(d.Id())
	if err != nil {
		return err
	}

	if cmd.Name == d.Id() {
		d.Set("name", cmd.Name)
		d.Set("type", cmd.Type)
		d.Set("line", cmd.Line)
		return nil
	}

	d.SetId("")
	return nil
}

func resourceCentreonCommandUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.CentreonwebClient)
	d.Partial(true)

	if d.HasChange("name") {
		if err := client.Commands().Setparam(d.Id(), "name", d.Get("name").(string)); err != nil {
			return err
		}

		d.SetId(d.Get("name").(string))
		d.SetPartial("type")
	}

	if d.HasChange("type") {
		if err := client.Commands().Setparam(d.Id(), "type", d.Get("type").(string)); err != nil {
			return err
		}

		d.SetPartial("type")
	}

	if d.HasChange("line") {
		if err := client.Commands().Setparam(d.Id(), "line", d.Get("line").(string)); err != nil {
			return err
		}

		d.SetPartial("line")
	}

	d.Partial(false)

	return resourceCentreonCommandRead(d, m)
}

func resourceCentreonCommandDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*centreonweb.CentreonwebClient)
	resourceExists, err := resourceCentreonCommandExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	if err := client.Commands().Del(d.Id()); err != nil {
		return err
	}

	return nil
}

func resourceCentreonCommandExists(d *schema.ResourceData, m interface{}) (b bool, e error) {
	client := m.(*centreonweb.CentreonwebClient)
	return client.Commands().Exists(d.Id())
}
