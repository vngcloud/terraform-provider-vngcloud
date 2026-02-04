package vks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
)

func ResourceCluster() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 3,
		//MigrateState:  resourceClusterMigrateState,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceContainerClusterResourceV1().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceClusterStateUpgradeV0,
				Version: 0,
			},
			{
				Type:    resourceContainerClusterResourceV2().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceClusterStateUpgradeV1,
				Version: 1,
			},
			{
				Type:    resourceContainerClusterResourceV3().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceClusterStateUpgradeV2,
				Version: 2,
			},
		},

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
				// DefaultFunc: func() (interface{}, error) {
				// 	return fetchByKey("k8s_version")
				// },
				Computed: true,
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
			"enable_service_endpoint": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return true, nil
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
				Default:  true,
			},
			"enabled_block_store_csi_plugin": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"secondary_subnets": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"node_netmask_size": {
				Type:         schema.TypeInt,
				Optional:     true,
				ForceNew:     true,
				Default:      25,
				ValidateFunc: validation.IntBetween(17, 26),
			},
			"node_group": {
				Type:     schema.TypeList,
				Optional: true,
				// Computed: true,
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
			"auto_upgrade_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"weekdays": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"release_channel": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"poc": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
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
	cli := m.(*client.Client)

	_, hasVersion := d.GetOk("version")
	versionKey := ""

	if cli.VksClient.Config().BasePath == "https://vks-han-1.api.vngcloud.vn" {
		versionKey = "han01_k8s_version"
	} else {
		versionKey = "k8s_version"
	}

	if !hasVersion {
		res, _ := fetchByKey(versionKey)
		imageId := res.(string)
		d.Set("version", imageId)
	}

	createNodeGroupRequests, errNodeGroup := expandNodeGroupForCreating(d.Get("node_group").([]interface{}), d, m)
	if errNodeGroup != nil {
		return errNodeGroup
	}
	secondarySubnets, errSecondarySubnets := getSecondarySubnets(d.Get("secondary_subnets").([]interface{}))
	if errSecondarySubnets != nil {
		return errSecondarySubnets
	}
	autoUpgradeConfig, errorUpgradeConfig := getAuToUpgradeConfig(d.Get("auto_upgrade_config").([]interface{}))
	if errorUpgradeConfig != nil {
		return errorUpgradeConfig
	}
	createClusterRequest := vks.CreateClusterComboDto{
		Name:                       d.Get("name").(string),
		Description:                d.Get("description").(string),
		Version:                    d.Get("version").(string),
		EnablePrivateCluster:       d.Get("enable_private_cluster").(bool),
		EnabledServiceEndpoint:     d.Get("enable_service_endpoint").(bool),
		NetworkType:                d.Get("network_type").(string),
		VpcId:                      d.Get("vpc_id").(string),
		SubnetId:                   d.Get("subnet_id").(string),
		Cidr:                       d.Get("cidr").(string),
		EnabledLoadBalancerPlugin:  d.Get("enabled_load_balancer_plugin").(bool),
		EnabledBlockStoreCsiPlugin: d.Get("enabled_block_store_csi_plugin").(bool),
		SecondarySubnets:           secondarySubnets,
		NodeNetmaskSize:            int32(d.Get("node_netmask_size").(int)),
		NodeGroups:                 createNodeGroupRequests,
		AutoUpgradeConfig:          autoUpgradeConfig,
		ReleaseChannel:             d.Get("release_channel").(string),
	}

	request := vks.V1ClusterControllerApiV1ClustersPostOpts{
		Body: optional.NewInterface(createClusterRequest),
	}

	_, hasPoc := d.GetOk("poc")
	if hasPoc && d.Get("poc").(bool) {
		request.Poc = optional.NewBool(true)
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
		clusterNodeGroupDetail, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterId, clusterNodeGroup.Id, nil)
		if httpResponse.StatusCode != http.StatusOK {
			return fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
		}
		if len(upgradeConfig) == 0 {
			upgradeConfig := []interface{}{
				map[string]interface{}{
					"strategy":        clusterNodeGroupDetail.UpgradeConfig.Strategy,
					"max_surge":       clusterNodeGroupDetail.UpgradeConfig.MaxSurge,
					"max_unavailable": clusterNodeGroupDetail.UpgradeConfig.MaxUnavailable,
				},
			}
			nodeGroup["upgrade_config"] = upgradeConfig
		}
		nodeGroup["node_group_id"] = clusterNodeGroupDetail.Id
		nodeGroup["subnet_id"] = clusterNodeGroupDetail.SubnetId
		if nodeGroup["num_nodes"] != nil && int32(nodeGroup["num_nodes"].(int)) != -1 {
			log.Printf("num_nodes !=nil\n")
		} else {
			nodeGroup["num_nodes"] = clusterNodeGroup.NumNodes
		}
		updatedNodeGroups[i] = nodeGroup
	}

	// Update the node_group field with the modified values
	if err := d.Set("node_group", updatedNodeGroups); err != nil {
		return fmt.Errorf("error setting node_group: %s", err)
	}

	return nil
}

func expandNodeGroupForCreating(node_group []interface{}, d *schema.ResourceData, m interface{}) ([]vks.CreateNodeGroupDto, error) {
	if len(node_group) == 0 {
		log.Printf("node_group 0\n")
		return []vks.CreateNodeGroupDto{}, nil
	} else if node_group[0] == nil {
		log.Printf("node_group nil\n")
		return []vks.CreateNodeGroupDto{}, nil
	}
	nodeGroupsJson, _ := json.Marshal(node_group)
	log.Printf("%s\n", string(nodeGroupsJson))
	createNodeGroupRequests := make([]vks.CreateNodeGroupDto, len(node_group))
	for i, ng := range node_group {
		nodeGroup, ok := ng.(map[string]interface{})
		if !ok {
			log.Fatalf("Element at index %d is not a map", i)
		}
		networkType := d.Get("network_type").(string)
		if networkType == "CILIUM_NATIVE_ROUTING" && (nodeGroup["secondary_subnets"] == nil || len(nodeGroup["secondary_subnets"].([]interface{})) == 0) {
			return nil, fmt.Errorf("secondary_subnets is required when cluster network type is set to CILIUM_NATIVE_ROUTING")
		}

		if nodeGroup["subnet_id"] == nil || nodeGroup["subnet_id"] == "" {
			nodeGroup["subnet_id"] = d.Get("subnet_id").(string)
		}

		setDefaultValueByZoneForNodeGroup(nodeGroup, m, d.Get("vpc_id").(string))

		createNodeGroupRequest, errNodeGroup := getCreateNodeGroupRequestForCluster(nodeGroup)
		if errNodeGroup != nil {
			return nil, errNodeGroup
		}
		createNodeGroupRequests[i] = createNodeGroupRequest
	}
	return createNodeGroupRequests, nil
}

func setDefaultValueByZoneForNodeGroup(nodeGroup map[string]interface{}, m interface{}, vpcId string) error {
	cli := m.(*client.Client)

	imageId := nodeGroup["image_id"]
	flavorId := nodeGroup["flavor_id"]
	diskType := nodeGroup["disk_type"]

	if imageId == nil || imageId.(string) == "" || flavorId == nil || flavorId.(string) == "" || diskType == nil || diskType.(string) == "" {
		workspaceRes, httpResponse, _ := cli.VksClient.V1WorkspaceControllerApi.V1WorkspaceGet(context.TODO(), nil)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
			return errResponse
		}

		subnetId := nodeGroup["subnet_id"].(string)
		subnetRes, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.GetSubnetByIdUsingGET(context.TODO(), vpcId, workspaceRes.ProjectId, subnetId)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
			return errResponse
		}
		imageIdKey := ""
		flavorIdKey := ""
		diskTypeKey := ""
		if cli.VksClient.Config().BasePath == "https://vks-han-1.api.vngcloud.vn" {
			imageIdKey = "han01_image_id"
			flavorIdKey = "han01_1a_flavor_id"
			diskTypeKey = "han01_1a_volume_type_id"
		} else {
			imageIdKey = "image_id"
			if subnetRes.Zone.Uuid == "HCM03-1A" {
				flavorIdKey = "flavor_id"
				diskTypeKey = "volume_type_id"
			} else if subnetRes.Zone.Uuid == "HCM03-1B" {
				flavorIdKey = "hcm03_1b_flavor_id"
				diskTypeKey = "hcm03_1b_volume_type_id"
			} else {
				flavorIdKey = "hcm03_1c_flavor_id"
				diskTypeKey = "hcm03_1c_volume_type_id"
			}
		}

		if imageId == nil || imageId.(string) == "" {
			res, _ := fetchByKey(imageIdKey)
			nodeGroup["image_id"] = res.(string)
		}
		if flavorId == nil || flavorId.(string) == "" {
			res, _ := fetchByKey(flavorIdKey)
			nodeGroup["flavor_id"] = res.(string)
		}
		if diskType == nil || diskType.(string) == "" {
			res, _ := fetchByKey(diskTypeKey)
			nodeGroup["disk_type"] = res.(string)
		}
	}

	return nil
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
	if resp.NetworkType == "CILIUM_NATIVE_ROUTING" {
		if !checkSecondarySubnetsSame(d, resp.SecondarySubnets) {
			d.Set("secondary_subnets", resp.SecondarySubnets)
			d.Set("node_netmask_size", resp.NodeNetmaskSize)
		}
	} else {
		d.Set("cidr", cluster.Cidr)
	}
	d.Set("vpc_id", cluster.VpcId)
	d.Set("subnet_id", cluster.SubnetId)
	d.Set("network_type", cluster.NetworkType)
	d.Set("name", cluster.Name)
	d.Set("enabled_load_balancer_plugin", cluster.EnabledLoadBalancerPlugin)
	d.Set("enabled_block_store_csi_plugin", cluster.EnabledBlockStoreCsiPlugin)
	d.Set("enable_private_cluster", cluster.EnablePrivateCluster)
	d.Set("release_channel", cluster.ReleaseChannel)
	if resp.AutoUpgradeConfig != nil {
		autoUpgradeConfig := []interface{}{
			map[string]interface{}{
				"weekdays": resp.AutoUpgradeConfig.Weekdays,
				"time":     resp.AutoUpgradeConfig.Time,
			},
		}
		d.Set("auto_upgrade_config", autoUpgradeConfig)
	} else {
		d.Set("auto_upgrade_config", nil)
	}
	log.Printf("GetConfig\n")
	configResp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdKubeconfigGet(context.TODO(), clusterID, nil)
	log.Printf("-------------------------------------\n")
	log.Printf("status %s\n", string(httpResponse.Status))
	log.Printf("-------------------------------------\n")
	aaa, _ := json.Marshal(configResp)
	log.Printf("config %s\n", string(aaa))

	if !CheckErrorResponse(httpResponse) {
		log.Printf("SetConfig\n")
		d.Set("config", configResp)
	}
	d.Set("poc", cluster.Poc)

	updateNodeGroupData(cli, d, clusterID)
	return nil
}

func resourceClusterUpdate(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("auto_upgrade_config") {
		err := updateAutoUpgradeConfig(d, m)
		if err != nil {
			return err
		}
	}
	if d.HasChange("white_list_node_cidr") || d.HasChange("version") ||
		d.HasChange("enabled_load_balancer_plugin") || d.HasChange("enabled_block_store_csi_plugin") {
		err := updateCluster(d, m)
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
	if d.HasChange("poc") {
		if d.Get("poc").(bool) {
			return fmt.Errorf("Cannot change poc from false to true")
		}
		err := stopPoc(d, m)
		if err != nil {
			return err
		}
	}
	return resourceClusterRead(d, m)
}

func stopPoc(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdStopPocPost(context.TODO(), d.Id())
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		oldPoc, _ := d.GetChange("poc")
		d.Set("poc", oldPoc)
		return errorResponse
	}
	return resourceClusterRead(d, m)
}

func updateAutoUpgradeConfig(d *schema.ResourceData, m interface{}) error {
	autoUpgradeConfig, errorUpgradeConfig := getAuToUpgradeConfig(d.Get("auto_upgrade_config").([]interface{}))
	if errorUpgradeConfig != nil {
		return errorUpgradeConfig
	}
	cli := m.(*client.Client)
	if autoUpgradeConfig != nil {
		request := vks.V1ClusterControllerApiV1ClustersClusterIdPutAutoUpgradeConfigOpts{
			Body: optional.NewInterface(autoUpgradeConfig),
		}
		_, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdPutAutoUpgradeConfig(context.TODO(), d.Id(), &request)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			oldAutoUpgradeConfig, _ := d.GetChange("auto_upgrade_config")
			d.Set("auto_upgrade_config", oldAutoUpgradeConfig)
			return errorResponse
		}
	} else {
		request := vks.V1ClusterControllerApiV1ClustersClusterIdDeleteAutoUpgradeConfigOpts{}
		_, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdDeleteAutoUpgradeConfig(context.TODO(), d.Id(), &request)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			oldAutoUpgradeConfig, _ := d.GetChange("auto_upgrade_config")
			d.Set("auto_upgrade_config", oldAutoUpgradeConfig)
			return errorResponse
		}
	}
	return resourceClusterRead(d, m)
}

func updateCluster(d *schema.ResourceData, m interface{}) error {
	whiteListCIDRsInterface := d.Get("white_list_node_cidr").([]interface{})
	var whiteListCIDR []string
	for _, s := range whiteListCIDRsInterface {
		whiteListCIDR = append(whiteListCIDR, s.(string))
	}
	if len(whiteListCIDR) == 0 {
		return fmt.Errorf(`The argument "white_list_node_cidr" must not be empty.`)
	}
	updateCluster := vks.UpdateClusterDto{
		Version:                    d.Get("version").(string),
		WhitelistNodeCIDRs:         whiteListCIDR,
		EnabledLoadBalancerPlugin:  d.Get("enabled_load_balancer_plugin").(bool),
		EnabledBlockStoreCsiPlugin: d.Get("enabled_block_store_csi_plugin").(bool),
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
		oldEnabledLoadBalancerPlugin, _ := d.GetChange("enabled_load_balancer_plugin")
		oldEnabledBlockStoreCsiPlugin, _ := d.GetChange("enabled_block_store_csi_plugin")
		d.Set("version", oldVersion)
		d.Set("white_list_node_cidr", oldWhiteListNodeCIDR)
		d.Set("enabled_load_balancer_plugin", oldEnabledLoadBalancerPlugin)
		d.Set("enabled_block_store_csi_plugin", oldEnabledBlockStoreCsiPlugin)
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
	oldNodeGroupSch, _ := d.GetChange("node_group")
	oldNodeGroups := oldNodeGroupSch.([]interface{})
	for i, ng := range nodeGroups {
		if reflect.DeepEqual(ng, oldNodeGroups[i]) {
			continue
		}
		nodeGroup := ng.(map[string]interface{})
		securityGroupsRequest := nodeGroup["security_groups"].([]interface{})
		var securityGroups []string
		for _, s := range securityGroupsRequest {
			securityGroups = append(securityGroups, s.(string))
		}
		if securityGroups == nil {
			securityGroups = make([]string, 0)
		}
		upgradeConfig := getUpgradeConfig(nodeGroup["upgrade_config"].([]interface{}))
		oldNodeGroup := oldNodeGroups[i].(map[string]interface{})

		autoScaleConfig := getAutoScaleConfig(nodeGroup["auto_scale_config"].([]interface{}))
		var numNodes *int32 = nil
		if int32(oldNodeGroup["num_nodes"].(int)) != int32(nodeGroup["num_nodes"].(int)) {
			num := int32(nodeGroup["num_nodes"].(int))
			numNodes = &num
		}
		taintsInput, ok := nodeGroup["taint"].([]interface{})
		var tains []vks.NodeGroupTaintDto
		if ok {
			tains = getTaints(taintsInput)
		} else {
			tains = nil
		}
		labels := getLabels(nodeGroup["labels"].(map[string]interface{}))
		imageId := nodeGroup["image_id"].(string)
		updateNodeGroupRequest := vks.UpdateNodeGroupDto{
			AutoScaleConfig: autoScaleConfig,
			NumNodes:        numNodes,
			UpgradeConfig:   &upgradeConfig,
			SecurityGroups:  securityGroups,
			ImageId:         imageId,
			Taints:          tains,
			Labels:          labels,
		}
		requestPutOpts := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdPutOpts{
			Body: optional.NewInterface(updateNodeGroupRequest),
		}
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdPut(context.TODO(), d.Id(), nodeGroup["node_group_id"].(string), &requestPutOpts)
		if CheckErrorResponse(httpResponse) {
			d.Set("node_group", oldNodeGroupSch)
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

func getAuToUpgradeConfig(input []interface{}) (*vks.AutoUpgradeConfigDto, error) {
	if len(input) == 0 {
		return nil, nil
	}
	autoUpgradeConfig, ok := input[0].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Both 'time' and 'weekdays' fields are required and cannot be empty.")
	}
	return &vks.AutoUpgradeConfigDto{
		Weekdays: autoUpgradeConfig["weekdays"].(string),
		Time:     autoUpgradeConfig["time"].(string),
	}, nil
}

//func resourceClusterMigrateState(
//	v int, is *terraform.InstanceState, meta interface{}) (*terraform.InstanceState, error) {
//	if is.Empty() {
//		log.Println("[DEBUG] Empty InstanceState; nothing to migrate.")
//		return is, nil
//	}
//	switch v {
//	case 1:
//		log.Println("[INFO] Found Cluster State v1 in legacy migration function; returning as non-op")
//		cli := meta.(*client.Client)
//		resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterID, nil)
//		if CheckErrorResponse(httpResponse) {
//			responseBody := GetResponseBody(httpResponse)
//			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
//			return is, errorResponse
//		}
//		is.Set("enable_endpoint", resp.EnablePrivateCluster)
//		return is, nil
//	default:
//		return is, fmt.Errorf("Unexpected schema version: %d", v)
//	}
//}

func resourceClusterStateUpgradeV0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("resourceClusterStateUpgradeV0\n")
	cli := meta.(*client.Client)
	id, ok := rawState["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is missing or not a string")
	}
	resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), id, nil)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return rawState, errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	rawState["enable_service_endpoint"] = resp.EnableServiceEndpoint
	return rawState, nil
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

func getCreateNodeGroupRequestForCluster(nodeGroup map[string]interface{}) (vks.CreateNodeGroupDto, error) {
	taintsInput, ok := nodeGroup["taint"].([]interface{})
	var tains []vks.NodeGroupTaintDto
	if ok {
		tains = getTaints(taintsInput)
	} else {
		tains = nil
	}
	secondarySubnets, errSecondarySubnets := getSecondarySubnets(nodeGroup["secondary_subnets"].([]interface{}))
	if errSecondarySubnets != nil {
		return vks.CreateNodeGroupDto{}, errSecondarySubnets
	}
	return vks.CreateNodeGroupDto{
		Name:                    nodeGroup["name"].(string),
		NumNodes:                int32(nodeGroup["num_nodes"].(int)),
		ImageId:                 nodeGroup["image_id"].(string),
		FlavorId:                nodeGroup["flavor_id"].(string),
		DiskSize:                int32(nodeGroup["disk_size"].(int)),
		DiskType:                nodeGroup["disk_type"].(string),
		EnablePrivateNodes:      nodeGroup["enable_private_nodes"].(bool),
		SshKeyId:                nodeGroup["ssh_key_id"].(string),
		Labels:                  getLabels(nodeGroup["labels"].(map[string]interface{})),
		Taints:                  tains,
		SecurityGroups:          getSecurityGroups(nodeGroup["security_groups"].([]interface{})),
		UpgradeConfig:           getUpgradeConfig(nodeGroup["upgrade_config"].([]interface{})),
		AutoScaleConfig:         getAutoScaleConfig(nodeGroup["auto_scale_config"].([]interface{})),
		EnabledEncryptionVolume: nodeGroup["enabled_encryption_volume"].(bool),
		SecondarySubnets:        secondarySubnets,
		SubnetId:                nodeGroup["subnet_id"].(string),
	}, nil
}

func resourceContainerClusterResourceV1() *schema.Resource {
	return &schema.Resource{
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

func resourceClusterStateUpgradeV1(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("resourceClusterStateUpgradeV0\n")
	cli := meta.(*client.Client)
	id, ok := rawState["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is missing or not a string")
	}
	resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), id, nil)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return rawState, errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if resp.NetworkType == "CILIUM_NATIVE_ROUTING" {
		rawState["secondary_subnets"] = resp.SecondarySubnets
		rawState["node_netmask_size"] = resp.NodeNetmaskSize
	} else {
		rawState["node_netmask_size"] = 25
	}
	return rawState, nil
}

func resourceContainerClusterResourceV2() *schema.Resource {
	return &schema.Resource{
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
			"enable_service_endpoint": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return true, nil
				},
			},
		},
	}
}

func checkSecondarySubnetsSame(d *schema.ResourceData, secondarySubnetResponse []string) bool {
	secondarySubnetsRequest, ok := d.GetOk("secondary_subnets")
	var secondarySubnets []string
	if ok {
		for _, s := range secondarySubnetsRequest.([]interface{}) {
			secondarySubnets = append(secondarySubnets, s.(string))
		}
	}
	var secondarySubnetsCluster []string
	for _, secondarySubnet := range secondarySubnetResponse {
		secondarySubnetsCluster = append(secondarySubnetsCluster, secondarySubnet)
	}
	return CheckListStringEqual(secondarySubnets, secondarySubnetsCluster)
}

func resourceClusterStateUpgradeV2(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("resourceClusterStateUpgradeV2\n")
	cli := meta.(*client.Client)
	id, ok := rawState["id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is missing or not a string")
	}
	resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), id, nil)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
		return rawState, errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if resp.AutoUpgradeConfig != nil {
		autoUpgradeConfig := map[string]interface{}{
			"weekdays": resp.AutoUpgradeConfig.Weekdays,
			"time":     resp.AutoUpgradeConfig.Time,
		}
		rawState["auto_upgrade_config"] = []interface{}{autoUpgradeConfig}
	}
	return rawState, nil
}

func resourceContainerClusterResourceV3() *schema.Resource {
	return &schema.Resource{
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
				Default:  true,
			},
			"enabled_block_store_csi_plugin": {
				Type:     schema.TypeBool,
				Optional: true,
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
			"enable_service_endpoint": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return true, nil
				},
			},
			"auto_upgrade_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"weekdays": {
							Type:     schema.TypeString,
							Required: true,
						},
						"time": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}
