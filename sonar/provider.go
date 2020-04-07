package sonar

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider for the Sonar project
func Provider() terraform.ResourceProvider {
	p := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"auth_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SONAR_AUTH_TOKEN", nil),
				Description: "Auth token.",
			},
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SONAR_API_URL", nil),
				Description: "Full URL of the harbormaster server.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sonar_workflow": dataSourceSonarWorkflow(),
		},
	}

	p.ConfigureFunc = providerConfigure(p)
	return p
}

func providerConfigure(d *schema.Provider) schema.ConfigureFunc {
	return func(d *schema.ResourceData) (interface{}, error) {
		authToken := d.Get("auth_token").(string)
		apiURL := d.Get("api_url").(string)

		client, err := NewClient(authToken, apiURL)
		if err != nil {
			return nil, err
		}

		return client, nil
	}
}
