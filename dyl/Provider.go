package dyl

import (
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const (
	PROVIDER_SECRET_ID  = "CDS_SECRET_ID"
	PROVIDER_SECRET_KEY = "CDS_SECRET_KEY"
	PROVIDER_REGION     = "CDS_REGION"
)

func Provider() terraform.ResourceProvider {
	// The actual Provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"secret_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_ID, nil),
				Description: "Secret ID of CDS",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc(PROVIDER_SECRET_KEY, nil),
				Description: "Secret key of CDS",
				Sensitive:   true,
			},
			"region": {
				Type:         schema.TypeString,
				Required:     true,
				DefaultFunc:  schema.EnvDefaultFunc(PROVIDER_REGION, nil),
				Description:  "Region of CDS",
				InputDefault: "CN_Beijing_A",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"dyl_data_source_vdc":            dataSourceCdsVdc(),
			"dyl_data_source_security_group": dataSourceCdsSecurityGroup(),
			"dyl_data_source_instance":       dataSourceCdsInstance(),
			"dyl_data_source_private_subnet": dataSourceCdsPrivateSubnet(),
			"dyl_data_source_haproxy":        dataSourceHaproxy(),
			"dyl_data_source_certificate":    dataSourceHaproxyCertificate(),
			"dyl_data_source_mysql":          dataSourceCdsMySQL(),
			"dyl_data_source_redis":          dataSourceCdsRedis(),
			"dyl_data_source_mongodb":        dataSourceCdsMongodb(),
			"dyl_data_source_public_network": dataSourceCdsPublicNetwork(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"dyl_vdc":            resourceCdsVdc(),
			"dyl_instance":       resourceCdsCcsInstance(),
			"dyl_security_group": resourceCdsSecurityGroup(),
			"dyl_private_subnet": resourceCdsPrivateSubnet(),
			"dyl_haproxy":        resourceCdsHaproxy(),
			"dyl_certificate":    resourceCdsCert(),
			"dyl_mysql":          resourceCdsMySQL(),
			"dyl_mysql_readonly": resourceCdsMySQLReadonly(),
			"dyl_mysql_account":  resourceCdsMySQLAccount(),
			"dyl_redis":          resourceCdsRedis(),
			"dyl_mongodb":        resourceCdsMongodb(),
			"dyl_public_network": resourceCdsPublicNetwork(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	secretId, ok := d.GetOk("secret_id")
	if !ok {
		secretId = os.Getenv(PROVIDER_SECRET_ID)
	}
	secretKey, ok := d.GetOk("secret_key")
	if !ok {
		secretKey = os.Getenv(PROVIDER_SECRET_KEY)
	}
	region, ok := d.GetOk("region")
	if !ok {
		region = os.Getenv(PROVIDER_REGION)
	}
	config := Config{
		SecretId:  secretId.(string),
		SecretKey: secretKey.(string),
		Region:    region.(string),
	}
	return config.Client()
}
