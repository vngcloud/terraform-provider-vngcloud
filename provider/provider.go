package provider

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/resource/vserver"
)

func Provider() *schema.Provider {
	log.SetFlags(log.Lshortfile)
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"vserver_project":          vserver.DataSourceProject(),
			"vserver_flavor_zone":      vserver.DataSourceFlavorZone(),
			"vserver_flavor":           vserver.DataSourceFlavor(),
			"vserver_image":            vserver.DataSourceImage(),
			"vserver_volume_type_zone": vserver.DataSourceVolumeTypeZone(),
			"vserver_volume_type":      vserver.DataSourceVolumeType(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vserver_volume":        vserver.ResourceVolume(),
			"vserver_server":        vserver.ResourceServer(),
			"vserver_sshkey":        vserver.ResourceSSHKey(),
			"vserver_network":       vserver.ResourceNetwork(),
			"vserver_secgroup":      vserver.ResourceSecgroup(),
			"vserver_subnet":        vserver.ResourceSubnet(),
			"vserver_secgrouprule":  vserver.ResourceSecgroupRule(),
			"vserver_volume_attach": vserver.ResourceAttachVolume(),
		},
		Schema: map[string]*schema.Schema{
			"token_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("TOKEN_ADDRESS", ""),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_SECRET", ""),
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", ""),
			},
			"base_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLIENT_ID", ""),
			},
		},
		ConfigureFunc: providerConfigvServer,
	}
}

//func providerConfigure(d *schema.ResourceData) (interface{}, error) {
//	address := d.Get("address").(string)
//	user := d.Get("user").(string)
//	accessKey := d.Get("access_key").(string)
//	return client.NewClient(address, user, accessKey), nil
//}

func providerConfigvServer(d *schema.ResourceData) (interface{}, error) {
	tokenURL := d.Get("token_url").(string)
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)
	baseURL := d.Get("base_url").(string)
	return client.NewClient(clientID, clientSecret, tokenURL, baseURL)
}
