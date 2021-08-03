package vdb

/*import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/resource/vdb"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"vdb_db":                     vdb.DataSourceDb(),
			"vdb_network":                vdb.DataSourceNetwork(),
			"vdb_package":                vdb.DataSourcePackage(),
			"vdb_subnet":                 vdb.DataSourceSubnet(),
			"vdb_volume_type":            vdb.DataSourceVolumeType(),
			"vdb_backup_storage_package": vdb.DataSourceBackupStoragePackage(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vdb_database":            vdb.ResourceDatabase(),
			"vdb_backup":              vdb.ResourceBackup(),
			"vdb_configuration_group": vdb.ResourceConfigurationGroup(),
			"vdb_backup_storage":      vdb.ResourceBackupStorage(),
		},
		Schema: map[string]*schema.Schema{
			"user_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"base_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token_url": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"client_secret": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ConfigureFunc: providerConfigvDB,
	}
}

func providerConfigvDB(d *schema.ResourceData) (interface{}, error) {
	baseURL := d.Get("base_url").(string)
	projectId := d.Get("project_id").(string)
	userId := d.Get("user_id").(string)
	tokenURL := d.Get("token_url").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	return client.NewVDBClient(baseURL, projectId, userId, clientID, clientSecret, tokenURL)
}*/
