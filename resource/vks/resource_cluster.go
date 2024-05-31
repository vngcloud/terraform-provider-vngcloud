package vks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
	"log"
	"net/http"
	"time"
)

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 0,
		//MigrateState:  resourceClusterMigrateState,
		//StateUpgraders: []schema.StateUpgrader{
		//	{
		//		Type:    resourceContainerClusterResourceV1().CoreConfigSchema().ImpliedType(),
		//		Upgrade: ResourceContainerClusterUpgradeV1,
		//		Version: 0,
		//	},
		//},

		Create: resourceClusterCreate,
		Read:   resourceClusterRead,
		Update: resourceClusterUpdate,
		Delete: resourceClusterDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				cli := m.(*client.Client)
				_, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), d.Id(), nil)
				if CheckErrorResponse(httpResponse) {
					responseBody := GetResponseBody(httpResponse)
					errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
					return nil, errResponse
				}
				resourceClusterRead(d, m)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"version": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return fetchByKey("k8s_version")
				},
			},
			"white_list_node_cidr": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enable_private_cluster": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return false, nil
				},
			},
			"network_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CALICO",
				ForceNew: true,
			},
			"vpc_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cidr": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"enabled_load_balancer_plugin": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"enabled_block_store_csi_plugin": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  true,
			},
			"node_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: MergeSchemas(
						schemaNodeGroup,
						map[string]*schema.Schema{
							"node_group_id": {
								Type:     schema.TypeString,
								Computed: true,
							},
						}),
				},
			},
		},
	}
}
func resourceClusterStateRefreshFunc(cli *client.Client, clusterID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterID, nil)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		cluster := resp
		return cluster, cluster.Status, nil
	}
}
func resourceClusterCreate(d *schema.ResourceData, m interface{}) error {
	createNodeGroupRequests := expandNodeGroupForCreating(d.Get("node_group").([]interface{}))
	createClusterRequest := vks.CreateClusterComboDto{
		Name:                       d.Get("name").(string),
		Description:                d.Get("description").(string),
		Version:                    d.Get("version").(string),
		EnablePrivateCluster:       d.Get("enable_private_cluster").(bool),
		NetworkType:                d.Get("network_type").(string),
		VpcId:                      d.Get("vpc_id").(string),
		SubnetId:                   d.Get("subnet_id").(string),
		Cidr:                       d.Get("cidr").(string),
		EnabledLoadBalancerPlugin:  d.Get("enabled_load_balancer_plugin").(bool),
		EnabledBlockStoreCsiPlugin: d.Get("enabled_block_store_csi_plugin").(bool),
		NodeGroups:                 createNodeGroupRequests,
	}
	cli := m.(*client.Client)
	request := vks.V1ClusterControllerApiV1ClustersPostOpts{
		Body: optional.NewInterface(createClusterRequest),
	}
	resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersPost(context.TODO(), &request)

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
		Pending:    CREATING,
		Target:     ACTIVE,
		Refresh:    resourceClusterStateRefreshFunc(cli, resp.Id),
		Timeout:    180 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for create cluster (%s) %s", resp.Id, err)
	}
	if len(createNodeGroupRequests) > 0 {
		stateConf = &resource.StateChangeConf{
			Pending:    CREATING,
			Target:     ACTIVE,
			Refresh:    resourceNodeGroupForClusterStateRefreshFunc(cli, resp.Id),
			Timeout:    180 * time.Minute,
			Delay:      10 * time.Second,
			MinTimeout: 1 * time.Second,
		}
		log.Printf("Node group size >0.................")
		_, err := stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf("error waiting for create cluster (%s) %s", resp.Id, err)
		}
		updateNodeGroupData(cli, d, resp.Id)
	}
	d.SetId(resp.Id)

	return resourceClusterRead(d, m)
}

func updateNodeGroupData(cli *client.Client, d *schema.ResourceData, clusterId string) error {
	nodeGroups := d.Get("node_group").([]interface{})
	updatedNodeGroups := make([]interface{}, len(nodeGroups))
	resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsGet(context.TODO(), clusterId, nil)
	if httpResponse.StatusCode != http.StatusOK {
		return fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	clusterNodeGroups := resp.Items
	for i, ng := range nodeGroups {
		nodeGroup := ng.(map[string]interface{})
		clusterNodeGroup := clusterNodeGroups[len(clusterNodeGroups)-i-1]
		// Set the value for a specific field
		upgradeConfig := nodeGroup["upgrade_config"].([]interface{})
		if len(upgradeConfig) == 0 {
			clusterNodeGroupDetail, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterId, clusterNodeGroup.Id, nil)
			if httpResponse.StatusCode != http.StatusOK {
				return fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
			}
			upgradeConfig := []interface{}{
				map[string]interface{}{
					"strategy":        clusterNodeGroupDetail.UpgradeConfig.Strategy,
					"max_surge":       clusterNodeGroupDetail.UpgradeConfig.MaxSurge,
					"max_unavailable": clusterNodeGroupDetail.UpgradeConfig.MaxUnavailable,
				},
			}
			nodeGroup["upgrade_config"] = upgradeConfig
			nodeGroup["node_group_id"] = clusterNodeGroupDetail.Id
		}
		//if nodeGroup["num_nodes"] != nil && int32(nodeGroup["num_nodes"].(int)) != -1 {
		//	nodeGroup["num_nodes"] = clusterNodeGroup.NumNodes
		//}
		updatedNodeGroups[i] = nodeGroup
	}

	// Update the node_group field with the modified values
	if err := d.Set("node_group", updatedNodeGroups); err != nil {
		return fmt.Errorf("error setting node_group: %s", err)
	}

	return nil
}

func expandNodeGroupForCreating(node_group []interface{}) []vks.CreateNodeGroupDto {
	if len(node_group) == 0 {
		log.Printf("node_group 0\n")
		return []vks.CreateNodeGroupDto{}
	} else if node_group[0] == nil {
		log.Printf("node_group nil\n")
		return []vks.CreateNodeGroupDto{}
	}
	nodeGroupsJson, _ := json.Marshal(node_group)
	log.Printf("aaaaa-------------------------------------\n")
	log.Printf("%s\n", string(nodeGroupsJson))
	createNodeGroupRequests := make([]vks.CreateNodeGroupDto, len(node_group))
	for i, ng := range node_group {
		nodeGroup, ok := ng.(map[string]interface{})
		if !ok {
			log.Fatalf("Element at index %d is not a map", i)
		}
		nodeGroupsJson, _ := json.Marshal(nodeGroup)
		log.Printf("bbbbb-------------------------------------\n")
		log.Printf("%s\n", string(nodeGroupsJson))
		createNodeGroupRequest := getCreateNodeGroupRequestForCluster(nodeGroup)
		createNodeGroupRequests[i] = createNodeGroupRequest
	}
	return createNodeGroupRequests
}

func resourceClusterRead(d *schema.ResourceData, m interface{}) error {
	clusterID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterID, nil)
	if httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	cluster := resp
	d.Set("version", cluster.Version)
	whiteListNodeCIDRRequest := d.Get("white_list_node_cidr").([]interface{})
	var whiteListNodeCIDR []string
	for _, s := range whiteListNodeCIDRRequest {
		whiteListNodeCIDR = append(whiteListNodeCIDR, s.(string))
	}
	var whiteListCIDRCluster []string
	for _, whiteList := range cluster.WhitelistNodeCIDRs {
		whiteListCIDRCluster = append(whiteListCIDRCluster, whiteList)
	}
	if !CheckListStringEqual(whiteListNodeCIDR, whiteListCIDRCluster) {
		d.Set("white_list_node_cidr", whiteListCIDRCluster)
	}
	d.Set("cidr", cluster.Cidr)
	d.Set("vpc_id", cluster.VpcId)
	d.Set("subnet_id", cluster.SubnetId)
	d.Set("network_type", cluster.NetworkType)
	d.Set("name", cluster.Name)
	d.Set("enabled_load_balancer_plugin", cluster.EnabledLoadBalancerPlugin)
	d.Set("enabled_block_store_csi_plugin", cluster.EnabledBlockStoreCsiPlugin)
	d.Set("enable_private_cluster", cluster.EnablePrivateCluster)
	log.Printf("GetConfig\n")
	configResp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdKubeconfigGet(context.TODO(), clusterID, nil)
	if !CheckErrorResponse(httpResponse) {
		log.Printf("SetConfig\n")
		d.Set("config", configResp)
	}
	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("node_group") {
		err := checkRequestNodeGroup(d)
		if err != nil {
			return err
		}
	}
	if d.HasChange("white_list_node_cidr") || d.HasChange("version") {
		err := changeWhiteListNodeOrVersion(d, m)
		if err != nil {
			return err
		}
	}
	if d.HasChange("node_group") {
		err := changeNodeGroup(d, m)
		if err != nil {
			return err
		}
	}
	return resourceClusterRead(d, m)
}

func checkRequestNodeGroup(d *schema.ResourceData) error {
	nodeGroups := d.Get("node_group").([]interface{})
	for _, ng := range nodeGroups {
		nodeGroup := ng.(map[string]interface{})
		autoScaleConfig := getAutoScaleConfig(nodeGroup["auto_scale_config"].([]interface{}))
		var numNodes *int32
		if value, ok := nodeGroup["num_nodes"]; ok {
			num := int32(value.(int))
			if num != -1 {
				numNodes = &num
			}
		}
		var err error
		if autoScaleConfig != nil && numNodes != nil {
			err = fmt.Errorf("If auto_scale_config is set then num_nodes must be -1\n")
		}
		if autoScaleConfig == nil && numNodes == nil {
			err = fmt.Errorf("If auto_scale_config is not set then num_nodes must be different from -1\n")
		}
		if err != nil {
			oldNodeGroup, _ := d.GetChange("node_group")
			d.Set("node_group", oldNodeGroup)
			return err
		}
	}
	return nil
}

func changeWhiteListNodeOrVersion(d *schema.ResourceData, m interface{}) error {
	whiteListCIDRsInterface := d.Get("white_list_node_cidr").([]interface{})
	var whiteListCIDR []string
	for _, s := range whiteListCIDRsInterface {
		whiteListCIDR = append(whiteListCIDR, s.(string))
	}
	if len(whiteListCIDR) == 0 {
		return fmt.Errorf(`The argument "white_list_node_cidr" must not be empty.`)
	}
	updateCluster := vks.UpdateClusterDto{
		Version:            d.Get("version").(string),
		WhitelistNodeCIDRs: whiteListCIDR,
	}
	cli := m.(*client.Client)
	request := vks.V1ClusterControllerApiV1ClustersClusterIdPutOpts{
		Body: optional.NewInterface(updateCluster),
	}
	resp, httpResponse, err := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdPut(context.TODO(), d.Id(), &request)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		oldVersion, _ := d.GetChange("version")
		oldWhiteListNodeCIDR, _ := d.GetChange("white_list_node_cidr")
		d.Set("version", oldVersion)
		d.Set("white_list_node_cidr", oldWhiteListNodeCIDR)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    UPDATING,
		Target:     ACTIVE,
		Refresh:    resourceClusterStateRefreshFunc(cli, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutUpdate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", d.Id(), err)
	}
	return resourceClusterRead(d, m)
}

func changeNodeGroup(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	nodeGroups := d.Get("node_group").([]interface{})
	for _, ng := range nodeGroups {
		nodeGroup := ng.(map[string]interface{})

		securityGroupsRequest := nodeGroup["security_groups"].([]interface{})
		var securityGroups []string
		for _, s := range securityGroupsRequest {
			securityGroups = append(securityGroups, s.(string))
		}
		if securityGroups == nil {
			securityGroups = make([]string, 0)
		}
		autoScaleConfig := getAutoScaleConfig(nodeGroup["auto_scale_config"].([]interface{}))
		upgradeConfig := getUpgradeConfig(nodeGroup["upgrade_config"].([]interface{}))
		var numNodes *int32
		if value, ok := nodeGroup["num_nodes"]; ok {
			num := int32(value.(int))
			if num != -1 {
				numNodes = &num
			}
		}
		imageId := nodeGroup["image_id"].(string)
		updateNodeGroupRequest := vks.UpdateNodeGroupDto{
			AutoScaleConfig: autoScaleConfig,
			NumNodes:        numNodes,
			UpgradeConfig:   &upgradeConfig,
			SecurityGroups:  securityGroups,
			ImageId:         imageId,
		}
		requestPutOpts := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdPutOpts{
			Body: optional.NewInterface(updateNodeGroupRequest),
		}
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdPut(context.TODO(), d.Id(), nodeGroup["node_group_id"].(string), &requestPutOpts)
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
			Pending:    UPDATING,
			Target:     ACTIVE,
			Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, d.Id(), nodeGroup["node_group_id"].(string)),
			Timeout:    180 * time.Minute,
			Delay:      10 * time.Second,
			MinTimeout: 1 * time.Second,
		}
		_, err := stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf("error waiting for update cluster node group (%s) %s", resp.Id, err)
		}
	}
	return resourceClusterRead(d, m)
}

func resourceClusterDelete(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdDelete(context.TODO(), d.Id(), nil)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    DELETING,
		Target:     DELETED,
		Refresh:    resourceClusterDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error waiting for instance (%s) to be created: %s", d.Id(), err)
	}
	d.SetId("")
	return nil
}

func resourceClusterDeleteStateRefreshFunc(cli *client.Client, clusterId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterId, nil)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vks.ClusterDto{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		return resp, resp.Status, nil
	}
}

func resourceClusterMigrateState(
	v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
	if is.Empty() {
		log.Println("[DEBUG] Empty InstanceState; nothing to migrate.")
		return is, nil
	}
	switch v {
	case 1:
		log.Println("[INFO] Found Cluster State v1 in legacy migration function; returning as non-op")
		return is, nil
	default:
		return is, fmt.Errorf("Unexpected schema version: %d", v)
	}
}

func resourceNodeGroupForClusterStateRefreshFunc(cli *client.Client, clusterID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsGet(context.TODO(), clusterID, nil)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		clusterNodeGroups := resp.Items
		var status string
		for _, clusterNodeGroup := range clusterNodeGroups {
			status = clusterNodeGroup.Status
			if clusterNodeGroup.Status != "ACTIVE" {
				break
			}
		}
		return clusterNodeGroups, status, nil
	}
}

func getCreateNodeGroupRequestForCluster(nodeGroup map[string]interface{}) vks.CreateNodeGroupDto {
	taintsInput, ok := nodeGroup["taints"].([]interface{})
	var tains []vks.NodeGroupTaintDto
	if ok {
		tains = getTaints(taintsInput)
	} else {
		tains = nil
	}
	return vks.CreateNodeGroupDto{
		Name:               nodeGroup["name"].(string),
		NumNodes:           int32(nodeGroup["initial_node_count"].(int)),
		ImageId:            nodeGroup["image_id"].(string),
		FlavorId:           nodeGroup["flavor_id"].(string),
		DiskSize:           int32(nodeGroup["disk_size"].(int)),
		DiskType:           nodeGroup["disk_type"].(string),
		EnablePrivateNodes: nodeGroup["enable_private_nodes"].(bool),
		SshKeyId:           nodeGroup["ssh_key_id"].(string),
		Labels:             getLabels(nodeGroup["labels"].(map[string]interface{})),
		Taints:             tains,
		SecurityGroups:     getSecurityGroups(nodeGroup["security_groups"].([]interface{})),
		UpgradeConfig:      getUpgradeConfig(nodeGroup["upgrade_config"].([]interface{})),
		AutoScaleConfig:    getAutoScaleConfig(nodeGroup["auto_scale_config"].([]interface{})),
	}
}
