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
		ResourcesMap: map[string]*schema.Resource{
			// "vdb_server_resource":  resourceServer(),
			// "vdb_vserver_resource": resourcevServer(),
			"vserver_volume_resource":        vserver.ResourceVolume(),
			"vserver_server_resource":        vserver.ResourceServer(),
			"vserver_sshkey_resource":        vserver.ResourceSSHKey(),
			"vserver_network_resource":       vserver.ResourceNetwork(),
			"vserver_secgroup_resource":      vserver.ResourceSecgroup(),
			"vserver_subnet_resource":        vserver.ResourceSubnet(),
			"vserver_secgrouprule_resource":  vserver.ResourceSecgroupRule(),
			"vserver_volume_attach_resource": vserver.ResourceAttachVolume(),
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
