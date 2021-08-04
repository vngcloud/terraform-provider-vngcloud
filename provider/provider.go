package provider

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/resource/vserver"
)

func Provider() *schema.Provider {
	log.SetFlags(log.Lshortfile)
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"vngcloud_vserver_project":          vserver.DataSourceProject(),
			"vngcloud_vserver_flavor_zone":      vserver.DataSourceFlavorZone(),
			"vngcloud_vserver_flavor":           vserver.DataSourceFlavor(),
			"vngcloud_vserver_image":            vserver.DataSourceImage(),
			"vngcloud_vserver_volume_type_zone": vserver.DataSourceVolumeTypeZone(),
			"vngcloud_vserver_volume_type":      vserver.DataSourceVolumeType(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vngcloud_vserver_volume":        vserver.ResourceVolume(),
			"vngcloud_vserver_server":        vserver.ResourceServer(),
			"vngcloud_vserver_sshkey":        vserver.ResourceSSHKey(),
			"vngcloud_vserver_network":       vserver.ResourceNetwork(),
			"vngcloud_vserver_secgroup":      vserver.ResourceSecgroup(),
			"vngcloud_vserver_subnet":        vserver.ResourceSubnet(),
			"vngcloud_vserver_secgrouprule":  vserver.ResourceSecgroupRule(),
			"vngcloud_vserver_volume_attach": vserver.ResourceAttachVolume(),
		},
		Schema: map[string]*schema.Schema{
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
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BASE_URL", ""),
				Description: "endpoint to connection with provider resource",
			},
		},
		ConfigureFunc: providerConfigvServer,
	}
}

func providerConfigvServer(d *schema.ResourceData) (interface{}, error) {
	tokenURL := d.Get("token_url").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	baseURL := d.Get("base_url").(string)
	return client.NewClient(clientID, clientSecret, tokenURL, baseURL)
}
