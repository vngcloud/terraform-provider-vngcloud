package provider

import (
	"github.com/vngcloud/terraform-provider-vngcloud/resource/vdbv2"
	"github.com/vngcloud/terraform-provider-vngcloud/resource/vks"
	"log"

	"github.com/vngcloud/terraform-provider-vngcloud/resource/vdb"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/resource/vloadbalancing"
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
			"vngcloud_vserver_k8s_version":         vserver.DataSourceK8sVersion(),
			"vngcloud_vserver_k8s_network_type":    vserver.DataSourceK8sNetworkType(),
			"vngcloud_vdb_network":                 vdb.DataSourceNetwork(),
			"vngcloud_vdb_package":                 vdb.DataSourcePackage(),
			"vngcloud_vdb_subnet":                  vdb.DataSourceSubnet(),
			"vngcloud_vdb_volume_type":             vdb.DataSourceVolumeType(),
			"vngcloud_vdb_backup_storage_package":  vdb.DataSourceBackupStoragePackage(),
			"vngcloud_vlb_lb_packages":             vloadbalancing.DataSourceLBPackages(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"vngcloud_vserver_volume":                            vserver.ResourceVolume(),
			"vngcloud_vserver_server":                            vserver.ResourceServer(),
			"vngcloud_vserver_sshkey":                            vserver.ResourceSSHKey(),
			"vngcloud_vserver_network":                           vserver.ResourceNetwork(),
			"vngcloud_vserver_vip":                               vserver.ResourceVip(),
			"vngcloud_vserver_route_table":                       vserver.ResourceRouteTable(),
			"vngcloud_vserver_address_pair_interface_attachment": vserver.ResourceAddressPairInterfaceAttachment(),
			"vngcloud_vserver_secgroup":                          vserver.ResourceSecgroup(),
			"vngcloud_vserver_subnet":                            vserver.ResourceSubnet(),
			"vngcloud_vserver_secgrouprule":                      vserver.ResourceSecgroupRule(),
			"vngcloud_vserver_volume_attach":                     vserver.ResourceAttachVolume(),
			"vngcloud_vserver_server_group":                      vserver.ResourceServerGroup(),
			"vngcloud_vlb_pool":                                  vloadbalancing.ResourcePool(),
			"vngcloud_vlb_load_balancer":                         vloadbalancing.ResourceLoadBalancer(),
			"vngcloud_vlb_listener":                              vloadbalancing.ResourceListener(),
			"vngcloud_vlb_l7policy":                              vloadbalancing.ResourceListenerL7Policy(),
			"vngcloud_vlb_certificate":                           vloadbalancing.ResourceCA(),
			"vngcloud_vdb_database":                              vdb.ResourceDatabase(),
			"vngcloud_vdb_backup":                                vdb.ResourceBackup(),
			"vngcloud_vdb_configuration_group":                   vdb.ResourceConfigurationGroup(),
			"vngcloud_vdb_backup_storage":                        vdb.ResourceBackupStorage(),
			"vngcloud_vdb_relational_config_group":               vdbv2.ResourceRelationalConfigurationGroup(),
			"vngcloud_vserver_k8s":                               vserver.ResourceK8s(),
			"vngcloud_vserver_cluster_node_group":                vserver.ResourceClusterNodeGroup(),
			"vngcloud_vserver_attach_lb_to_cluster":              vserver.ResourceAttachLb(),
			"vngcloud_vserver_change_cluster_sec_group":          vserver.ResourceChangeClusterSecGroup(),
			"vngcloud_vserver_network_interface":                 vserver.ResourceNetworkInterface(),
			"vngcloud_vserver_external_interface_attach":         vserver.ResourceAttachExternalInterface(),
			"vngcloud_vserver_internal_interface_attach":         vserver.ResourceAttachInternalInterface(),
			"vngcloud_vks_cluster":                               vks.ResourceCluster(),
			"vngcloud_vks_cluster_node_group":                    vks.ResourceClusterNodeGroup(),
		},
		Schema: map[string]*schema.Schema{
			"token_url": {
				Type:        schema.TypeString,
				Optional:    true,
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
			"vserver_base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VSERVER_BASE_URL", ""),
				Description: "endpoint to connection with provider resource",
			},
			"vlb_base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VLB_BASE_URL", ""),
				Description: "endpoint to connection with provider resource",
			},
			"vks_base_url": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("VKS_BASE_URL", ""),
				Description: "endpoint to connection with provider resource",
			},
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	vserverBaseURL := "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway"
	_, hasVserverBaseUrl := d.GetOk("vserver_base_url")
	if hasVserverBaseUrl {
		vserverBaseURL = d.Get("vserver_base_url").(string)
	}
	_, hasVlbBaseUrl := d.GetOk("vlb_base_url")
	vlbBaseURL := "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway"
	if hasVlbBaseUrl {
		vlbBaseURL = d.Get("vlb_base_url").(string)
	}
	_, hasTokenUrl := d.GetOk("token_url")
	tokenURL := "https://iamapis.vngcloud.vn/accounts-api/v2/auth/token"
	if hasTokenUrl {
		tokenURL = d.Get("token_url").(string)
	}
	clientID := d.Get("client_id").(string)
	clientSecret := d.Get("client_secret").(string)

	vksBaseURL := "https://vks.api.vngcloud.vn"
	_, hasVksBaseUrl := d.GetOk("vks_base_url")
	if hasVksBaseUrl {
		vksBaseURL = d.Get("vks_base_url").(string)
	}
	return client.NewClientV2(vserverBaseURL, vlbBaseURL, vksBaseURL, clientID, clientSecret, tokenURL)
}
