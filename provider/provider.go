package provider

import (
	"log"

	"github.com/vngcloud/terraform-provider-vngcloud/resource/vdb"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/resource/vserver"
)

func Provider() *schema.Provider {
	log.SetFlags(log.Lshortfile)
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"vngcloud_vserver_project":             vserver.DataSourceProject(),
			"vngcloud_vserver_flavor_zone":         vserver.DataSourceFlavorZone(),
			"vngcloud_vserver_flavor":              vserver.DataSourceFlavor(),
			"vngcloud_vserver_image":               vserver.DataSourceImage(),
			"vngcloud_vserver_volume_type_zone":    vserver.DataSourceVolumeTypeZone(),
			"vngcloud_vserver_volume_type":         vserver.DataSourceVolumeType(),
			"vngcloud_vserver_server_group_policy": vserver.DataSourceServerGroupPolicy(),
			"vngcloud_vdb_network":                 vdb.DataSourceNetwork(),
			"vngcloud_vdb_package":                 vdb.DataSourcePackage(),
			"vngcloud_vdb_subnet":                  vdb.DataSourceSubnet(),
			"vngcloud_vdb_volume_type":             vdb.DataSourceVolumeType(),
			"vngcloud_vdb_backup_storage_package":  vdb.DataSourceBackupStoragePackage(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vngcloud_vserver_volume":                    vserver.ResourceVolume(),
			"vngcloud_vserver_server":                    vserver.ResourceServer(),
			"vngcloud_vserver_sshkey":                    vserver.ResourceSSHKey(),
			"vngcloud_vserver_network":                   vserver.ResourceNetwork(),
			"vngcloud_vserver_secgroup":                  vserver.ResourceSecgroup(),
			"vngcloud_vserver_subnet":                    vserver.ResourceSubnet(),
			"vngcloud_vserver_secgrouprule":              vserver.ResourceSecgroupRule(),
			"vngcloud_vserver_network_interface":         vserver.ResourceNetworkInterface(),
			"vngcloud_vserver_external_interface_attach": vserver.ResourceAttachExternalInterface(),
			"vngcloud_vserver_volume_attach":             vserver.ResourceAttachVolume(),
			"vngcloud_vserver_server_group":              vserver.ResourceServerGroup(),
			"vngcloud_vserver_load_balancer":             vserver.ResourceLoadBalancer(),
			"vngcloud_vserver_listener":                  vserver.ResourceListener(),
			"vngcloud_vserver_pool":                      vserver.ResourcePool(),
			"vngcloud_vdb_database":                      vdb.ResourceDatabase(),
			"vngcloud_vdb_backup":                        vdb.ResourceBackup(),
			"vngcloud_vdb_configuration_group":           vdb.ResourceConfigurationGroup(),
			"vngcloud_vdb_backup_storage":                vdb.ResourceBackupStorage(),
		},
		Schema: map[string]*schema.Schema{
			//"user_id": {
			//	Type:     schema.TypeString,
			//	Required: true,
			//},
			//"project_id": {
			//	Type:     schema.TypeString,
			//	Required: true,
			//},
			"token_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOKEN_ADDRESS", ""),
				Description: "endpoint for terraform request token",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", ""),
				Description: "client secret for auth get access token",
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", ""),
				Description: "client id for auth get access token",
			},
			//"vdb_base_url": {
			//	Type:        schema.TypeString,
			//	Required:    true,
			//	DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", ""),
			//},
			"vserver_base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSERVER_BASE_URL", ""),
				Description: "endpoint to connection with provider resource",
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

//func providerConfigure(d *schema.ResourceData) (interface{}, error) {
//	address := d.Get("address").(string)
//	user := d.Get("user").(string)
//	accessKey := d.Get("access_key").(string)
//	return client.NewVDBClient(address, user, accessKey), nil
//}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	//vdbBaseURL := d.Get("vdb_base_url").(string)
	vdbBaseURL := ""
	vserverBaseURL := d.Get("vserver_base_url").(string)
	//projectId := d.Get("project_id").(string)
	//userId := d.Get("user_id").(string)
	projectId := ""
	userId := "0"
	tokenURL := d.Get("token_url").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	return client.NewClient(vdbBaseURL, vserverBaseURL, projectId, userId, clientID, clientSecret, tokenURL)
}
