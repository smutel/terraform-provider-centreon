package centreon

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/smutel/go-centreon/centreonweb"
)

// Provider exports the actual provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CENTREON_URL", nil),
				Description: "URL to reach centreon web application.",
			},
			"allow_unverified_ssl": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CENTREON_ALLOW_UNVERIFIED_SSL", nil),
				Description: "To disable checking the SSL cert.",
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CENTREON_USER", nil),
				Description: "The user name for centreon API operations.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CENTREON_PASSWORD", nil),
				Description: "The user password for centreon API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"centreon_command": resourceCentreonCommand(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	ssl := d.Get("allow_unverified_ssl").(bool)
	user := d.Get("user").(string)
	password := d.Get("password").(string)

	return centreonweb.New(url, ssl, user, password)
}