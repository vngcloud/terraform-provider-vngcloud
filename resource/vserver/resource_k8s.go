package vserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"log"
	"net/http"
	"strings"
	"time"
)

func ResourceK8s() *schema.Resource {
	return &schema.Resource{
		Create: resourceK8sCreate,
		Read:   resourceK8sRead,
		Update: resourceK8sUpdate,
		Delete: resourceK8sDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf(
						"unexpected format of ID (%q), expected ProjectID:VolumeID", d.Id())
				}
				projectID := idParts[0]
				clusterID := idParts[1]
				d.SetId(clusterID)
				d.Set("project_id", projectID)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ipip_mode": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"k8s_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"master_count": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			//Update node group default
			"node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"master_flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"node_flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"etcd_volume_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"etcd_volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"boot_volume_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"boot_volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"docker_volume_size": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"docker_volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"calico_cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"ssh_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"security_group": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
			"min_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"max_node_count": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
			},
			"enable_lb": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"auto_scaling": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"auto_healing": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"ingress_controller": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			"auto_monitoring": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
			//"node_group_list": {
			//	Type:     schema.TypeList,
			//	Optional: true,
			//	Elem: map[string]*schema.Schema{
			//		"name": {
			//			Type:     schema.TypeString,
			//			ForceNew: true,
			//		},
			//		"node_amount": {
			//			Type: schema.TypeInt,
			//		},
			//		"flavor_id": {
			//			Type:     schema.TypeString,
			//			ForceNew: true,
			//		},
			//	},
			//},
			"k8s_network_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssh_key_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_flavor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"master_flavor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_point": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_group_default_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secgroup_default_master": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ethertype": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port_range_max": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"port_range_min": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_group_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_ip_prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sec_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"secgroup_default_minion": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"direction": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"ethertype": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"port_range_max": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"port_range_min": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"protocol": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_group_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remote_ip_prefix": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"sec_group_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

//func buildCreateNodeGroupRequests(d *schema.ResourceData) []vserver.NodeGroupRequestModel {
//	nodeGroupList := d.Get("node_group_list").([]interface{})
//	NodeGroups := make([]vserver.NodeGroupRequestModel, len(nodeGroupList))
//
//	for _, nodeGroup := range nodeGroupList {
//		nodeGroup := nodeGroup.(map[string]interface{})
//		var request vserver.NodeGroupRequestModel
//		request.Name = nodeGroup["name"].(string)
//		request.FlavorId = nodeGroup["flavor_id"].(string)
//		request.NodeCount = nodeGroup["node_amount"].(int)
//		NodeGroups = append(NodeGroups, request)
//	}
//	return NodeGroups
//}

func resourceK8sCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Create K8s")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	var secGroupIdList []string

	cluster := vserver.CreateClusterRequest{
		AutoHealingEnabled:       d.Get("auto_healing").(bool),
		AutoMonitoringEnabled:    d.Get("auto_monitoring").(bool),
		AutoScalingEnabled:       d.Get("auto_scaling").(bool),
		BootVolumeSize:           int32(d.Get("boot_volume_size").(int)),
		BootVolumeTypeId:         d.Get("boot_volume_type_id").(string),
		CalicoCidr:               d.Get("calico_cidr").(string),
		DockerVolumeSize:         int32(d.Get("docker_volume_size").(int)),
		DockerVolumeTypeId:       d.Get("docker_volume_type_id").(string),
		Description:              d.Get("description").(string),
		EnabledLb:                d.Get("enable_lb").(bool),
		EtcdVolumeSize:           int32(d.Get("etcd_volume_size").(int)),
		EtcdVolumeTypeId:         d.Get("etcd_volume_type_id").(string),
		IngressControllerEnabled: d.Get("ingress_controller").(bool),
		IpipMode:                 d.Get("ipip_mode").(string),
		K8sVersion:               d.Get("k8s_version").(string),
		MasterCount:              int32(d.Get("master_count").(int)),
		MasterFlavorId:           d.Get("master_flavor_id").(string),
		MaxNodeCount:             int32(d.Get("max_node_count").(int)),
		MinNodeCount:             int32(d.Get("min_node_count").(int)),
		Name:                     d.Get("name").(string),
		NetworkId:                d.Get("network_id").(string),
		NetworkType:              d.Get("network_type").(string),
		NodeCount:                int32(d.Get("node_count").(int)),
		NodeFlavorId:             d.Get("node_flavor_id").(string),
		SecGroupIds:              secGroupIdList,
		SshKeyId:                 d.Get("ssh_key_id").(string),
		SubnetId:                 d.Get("subnet_id").(string),
		//NodeGroupRequestModelList: d.Get("node_group_list").([]vserver.NodeGroupRequestModel),
		//NodeGroupRequestModelList: buildCreateNodeGroupRequests(d),
		NodeGroupRequestModelList: []vserver.NodeGroupRequestModel{},
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.CreateClusterUsingPOST(context.TODO(), cluster, projectId)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return errResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	stateConf := &resource.StateChangeConf{
		Pending:    k8sClusterCreating,
		Target:     k8sClusterCreated,
		Refresh:    resourceK8sClusterStateRefreshFunc(cli, resp.Data.Uuid, projectId),
		Timeout:    180 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for create cluster (%s) %s", resp.Data.Uuid, err)
	}

	d.SetId(resp.Data.Uuid)

	return resourceK8sRead(d, m)
}

func resourceK8sRead(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	clusterId := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO(), clusterId, projectId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}
	respSecGroup, httpResponseSecGroup, _ := cli.VserverClient.K8SClusterRestControllerApi.ListSecgroupDefaultUsingGET(context.TODO(), clusterId, projectId)
	if CheckErrorResponse(httpResponseSecGroup) {
		responseBody := GetResponseBody(httpResponseSecGroup)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	respConfig, httpResponseConfig, _ := cli.VserverClient.K8SClusterRestControllerApi.GetConfigUsingGET(context.TODO(), clusterId, projectId)
	if CheckErrorResponse(httpResponseConfig) {
		responseBody := GetResponseBody(httpResponseConfig)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	respSecGroupJSON, _ := json.Marshal(respSecGroup)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respSecGroupJSON))
	log.Printf("-------------------------------------\n")

	respConfigJSON, _ := json.Marshal(respConfig)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respConfigJSON))
	log.Printf("-------------------------------------\n")

	cluster := resp.Data
	masterSecgroupDefault := respSecGroup[0].Master
	minionSecgroupDefault := respSecGroup[0].Minion
	config := respConfig.Configuration
	if cluster.AutoHealingEnabled {
		d.Set("auto_healing", cluster.AutoHealingEnabled)
	}

	if cluster.AutoMonitoringEnabled {
		d.Set("auto_monitoring", cluster.AutoMonitoringEnabled)
	}

	if cluster.AutoScalingEnabled {
		d.Set("auto_scaling", cluster.AutoScalingEnabled)
	}

	if cluster.EnabledLb {
		d.Set("enable_lb", cluster.EnabledLb)
	}

	d.Set("boot_volume_size", cluster.BootVolumeSize)
	d.Set("boot_volume_type_id", cluster.BootVolumeTypeId)
	d.Set("calico_cidr", cluster.CalicoCidr)
	d.Set("description", cluster.Description)
	d.Set("docker_volume_size", cluster.DockerVolumeSize)
	d.Set("docker_volume_type_id", cluster.DockerVolumeTypeId)
	d.Set("end_point", cluster.Endpoint)
	d.Set("etcd_volume_size", cluster.EtcdVolumeSize)
	d.Set("etcd_volume_type_id", cluster.EtcdVolumeTypeId)
	d.Set("id", cluster.ClusterId)

	if cluster.IngressControllerEnabled {
		d.Set("ingress_controller", cluster.IngressControllerEnabled)
	}

	d.Set("ipip_mode", "Always")
	d.Set("k8s_network_type", cluster.K8sNetworkType)
	d.Set("network_type", cluster.K8sNetworkTypeId)
	d.Set("k8s_version_name", cluster.K8sVersion)
	d.Set("k8s_version", cluster.K8sVersionId)
	d.Set("master_count", cluster.MasterCount)
	d.Set("master_flavor_name", cluster.MasterFlavorName)
	d.Set("master_flavor_id", cluster.MasterFlavorId)
	if cluster.AutoScalingEnabled {
		d.Set("max_node_count", cluster.MaxNodeCount)
	}

	if cluster.AutoScalingEnabled {
		d.Set("min_node_count", cluster.MinNodeCount)
	}

	d.Set("name", cluster.Name)
	d.Set("network_id", cluster.NetworkId)
	d.Set("node_count", cluster.NodeCount)
	d.Set("node_flavor_name", cluster.NodeFlavorName)
	d.Set("node_flavor_id", cluster.NodeFlavorId)
	d.Set("node_group_default_id", cluster.NodegroupDefaultId)
	d.Set("ssh_key_name", cluster.SshKeyName)
	d.Set("ssh_key_id", cluster.SshKeyId)
	d.Set("subnet_id", cluster.SubnetId)
	d.Set("config", config)
	d.Set("secgroup_default_master", flattenClusterSecGroupDefault(masterSecgroupDefault))
	d.Set("secgroup_default_minion", flattenClusterSecGroupDefault(minionSecgroupDefault))
	return nil
}

func resourceK8sUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("node_count") {
		nodeGroupDefaultId := d.Get("node_group_default_id").(string)
		clusterId := d.Id()

		if d.Get("auto_scaling") != nil && d.Get("auto_scaling").(bool) == true {
			log.Printf("You can not scale manual node group default in cluster with auto-scale option enable.")
			return resourceK8sRead(d, m)
		}

		return ResourceK8sScalingNodeGroup(d, m, nodeGroupDefaultId, clusterId)
	}

	if d.HasChange("description") {
		clusterId := d.Id()
		return ResourceK8sUpdateDescription(d, m, clusterId)
	}

	return resourceK8sRead(d, m)
}

func resourceK8sDelete(d *schema.ResourceData, m interface{}) error {
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.DeleteClusterUsingDELETE(context.TODO(), d.Id(), projectId)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsh: %s", responseBody)
		return errorResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    k8sClusterDeleting,
		Target:     k8sClusterDeleted,
		Refresh:    resourceK8sDeleteStateRefreshFunc(cli, d.Id(), projectId),
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for delete cluster (%s) : %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func ResourceK8sUpdateDescription(d *schema.ResourceData, m interface{}, clusterId string) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)

	_, n := d.GetChange("description")
	updateClusterDescriptionRequest := vserver.UpdateClusterDescriptionRequest{
		ClusterId:   clusterId,
		Description: n.(string),
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.UpdateClusterDescriptionUsingPut(context.TODO(), projectID, clusterId, updateClusterDescriptionRequest)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return errResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if err != nil {
		return fmt.Errorf("error updating description for cluster (%s) %s", clusterId, err)
	}

	d.SetId(clusterId)
	return nil
}

func resourceK8sClusterStateRefreshFunc(cli *client.Client, clusterId string, projectId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO(), clusterId, projectId)

		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("error describing: %s", GetResponseBody(httpResponse))
		}

		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		cluster := resp.Data
		return cluster, cluster.Status, nil
	}
}

func resourceK8sDeleteStateRefreshFunc(cli *client.Client, clusterId string, projectId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO(), clusterId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.ClusterDto{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("error describing instance: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		cluster := resp.Data
		return cluster, cluster.Status, nil
	}
}

func flattenClusterSecGroupDefault(secgroupDefaults *vserver.SecGroupDefault) []map[string]interface{} {
	if secgroupDefaults == nil {
		return []map[string]interface{}{}
	}

	m := map[string]interface{}{
		"rules":          flattenRules(secgroupDefaults.Rules),
		"sec_group_name": secgroupDefaults.SecgroupName,
	}
	fmt.Printf("[DEBUG]: %s", m)

	return []map[string]interface{}{m}
}

func flattenRules(rules []vserver.SecGroupRuleDefault) []map[string]interface{} {
	if len(rules) == 0 {
		return []map[string]interface{}{}
	}

	l := make([]map[string]interface{}, 0, len(rules))

	for _, rule := range rules {
		m := map[string]interface{}{
			"direction":         rule.Direction,
			"ethertype":         rule.Ethertype,
			"port_range_max":    int64(rule.PortRangeMax),
			"port_range_min":    int64(rule.PortRangeMin),
			"protocol":          rule.Protocol,
			"remote_group_name": rule.RemoteGroupName,
			"remote_ip_prefix":  rule.RemoteIpPrefix,
		}

		l = append(l, m)
	}

	return l
}
