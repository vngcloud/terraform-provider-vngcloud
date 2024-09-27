package vks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
	"log"
	"net/http"
	"strings"
	"time"
)

func ResourceClusterNodeGroup() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,
		//MigrateState:  resourceClusterNodeGroupMigrateState,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    resourceContainerClusterNodeGroupResourceV1().CoreConfigSchema().ImpliedType(),
				Upgrade: resourceClusterNodeGroupStateUpgradeV0,
				Version: 0,
			},
		},

		Create: resourceClusterNodeGroupCreate,
		Read:   resourceClusterNodeGroupRead,
		Update: resourceClusterNodeGroupUpdate,
		Delete: resourceClusterNodeGroupDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ClusterID:NodeGroupId", d.Id())
				}
				clusterID := idParts[0]
				nodeGroupID := idParts[1]
				cli := m.(*client.Client)
				_, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterID, nodeGroupID, nil)
				if CheckErrorResponse(httpResponse) {
					responseBody := GetResponseBody(httpResponse)
					errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
					return nil, errResponse
				}
				d.SetId(nodeGroupID)
				d.Set("cluster_id", clusterID)
				resourceClusterNodeGroupRead(d, m)
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: MergeSchemas(
			schemaNodeGroup,
			map[string]*schema.Schema{
				"cluster_id": {
					Type:     schema.TypeString,
					Required: true,
					ForceNew: true,
				},
			},
		),
	}
}

var schemaNodeGroup = map[string]*schema.Schema{
	"name": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"num_nodes": {
		Type:     schema.TypeInt,
		Optional: true,
		Default:  1,
	},
	"auto_scale_config": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"min_size": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  1,
				},
				"max_size": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  10,
				},
			},
		},
	},
	"upgrade_config": {
		Type:     schema.TypeList,
		MaxItems: 1,
		Optional: true,
		Computed: true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"strategy": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "SURGE",
				},
				"max_surge": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  1,
				},
				"max_unavailable": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  0,
				},
			},
		},
		DefaultFunc: func() (interface{}, error) {
			return []interface{}{
				map[string]interface{}{
					"strategy":        "SURGE",
					"max_surge":       1,
					"max_unavailable": 0,
				},
			}, nil
		},
	},
	"image_id": {
		Type:     schema.TypeString,
		Optional: true,
		DefaultFunc: func() (interface{}, error) {
			return fetchByKey("image_id")
		},
	},
	"flavor_id": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
		DefaultFunc: func() (interface{}, error) {
			return fetchByKey("flavor_id")
		},
	},
	"disk_size": {
		Type:     schema.TypeInt,
		Optional: true,
		ForceNew: true,
		DefaultFunc: func() (interface{}, error) {
			return 20, nil
		},
	},
	"disk_type": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
		DefaultFunc: func() (interface{}, error) {
			return fetchByKey("volume_type_id")
		},
	},
	"enable_private_nodes": {
		Type:     schema.TypeBool,
		Optional: true,
		ForceNew: true,
		DefaultFunc: func() (interface{}, error) {
			return false, nil
		},
	},
	"security_groups": {
		Type:     schema.TypeList,
		Optional: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"ssh_key_id": {
		Type:     schema.TypeString,
		Required: true,
		ForceNew: true,
	},
	"labels": {
		Type:        schema.TypeMap,
		Optional:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: `The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.`,
	},
	"taint": {
		Type:        schema.TypeList,
		Optional:    true,
		Description: `List of Kubernetes taints to be applied to each node.`,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"key": {
					Type:        schema.TypeString,
					Required:    true,
					Description: `Key for taint.`,
				},
				"value": {
					Type:        schema.TypeString,
					Required:    true,
					Description: `Value for taint.`,
				},
				"effect": {
					Type:         schema.TypeString,
					Optional:     true,
					Default:      "NoSchedule",
					ValidateFunc: validation.StringInSlice([]string{"NoSchedule", "PreferNoSchedule", "NoExecute"}, false),
					Description:  `Effect for taint.`,
				},
			},
		},
	},
	"secondary_subnets": {
		Type:     schema.TypeList,
		Optional: true,
		ForceNew: true,
		Elem: &schema.Schema{
			Type: schema.TypeString,
		},
	},
	"enabled_encryption_volume": {
		Type:     schema.TypeBool,
		Optional: true,
		ForceNew: true,
		DefaultFunc: func() (interface{}, error) {
			return false, nil
		},
	},
	"subnet_id": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
}

func getSecurityGroups(input []interface{}) []string {

	securityGroups := make([]string, len(input))

	for i, v := range input {
		str, ok := v.(string)
		if !ok {
			return []string{}
		}
		securityGroups[i] = str
	}

	return securityGroups
}

func getUpgradeConfig(input []interface{}) vks.NodeGroupUpgradeConfigDto {
	if len(input) == 0 {
		return vks.NodeGroupUpgradeConfigDto{
			MaxUnavailable: 0,
			MaxSurge:       1,
			Strategy:       "SURGE",
		}
	}
	upgradeConfig, ok := input[0].(map[string]interface{})
	if !ok {
		log.Fatalf("Element at index %d is not a map", 0)
	}

	return vks.NodeGroupUpgradeConfigDto{
		Strategy:       upgradeConfig["strategy"].(string),
		MaxSurge:       int32(upgradeConfig["max_surge"].(int)),
		MaxUnavailable: int32(upgradeConfig["max_unavailable"].(int)),
	}
}

func getAutoScaleConfig(input []interface{}) *vks.NodeGroupAutoScaleConfigDto {
	if len(input) == 0 {
		return nil
	}
	autoScaleConfig := input[0].(map[string]interface{})

	return &vks.NodeGroupAutoScaleConfigDto{
		MaxSize: int32(autoScaleConfig["max_size"].(int)),
		MinSize: int32(autoScaleConfig["min_size"].(int)),
	}
}

func getLabels(input map[string]interface{}) map[string]string {
	labels := make(map[string]string, len(input))

	for k, v := range input {
		str, ok := v.(string)
		if !ok {
			return nil
		}
		labels[k] = str
	}

	return labels
}

func getTaints(input []interface{}) []vks.NodeGroupTaintDto {
	taints := make([]vks.NodeGroupTaintDto, len(input))

	for i, rawTaint := range input {
		taint := rawTaint.(map[string]interface{})
		taintDto := vks.NodeGroupTaintDto{
			Key:    taint["key"].(string),
			Value:  taint["value"].(string),
			Effect: taint["effect"].(string),
		}
		taints[i] = taintDto
	}

	return taints
}

func resourceClusterNodeGroupStateRefreshFunc(cli *client.Client, clusterID string, clusterNodeGroupID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterID, clusterNodeGroupID, nil)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		clusterNodeGroup := resp
		return clusterNodeGroup, clusterNodeGroup.Status, nil
	}
}

func resourceClusterNodeGroupRead(d *schema.ResourceData, m interface{}) error {
	clusterID := d.Get("cluster_id").(string)
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterID, d.Id(), nil)
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
	upgradeConfig := []interface{}{
		map[string]interface{}{
			"strategy":        resp.UpgradeConfig.Strategy,
			"max_surge":       resp.UpgradeConfig.MaxSurge,
			"max_unavailable": resp.UpgradeConfig.MaxUnavailable,
		},
	}
	d.Set("upgrade_config", upgradeConfig)
	if resp.AutoScaleConfig != nil {
		autoScaleConfig := []interface{}{
			map[string]interface{}{
				"min_size": resp.AutoScaleConfig.MinSize,
				"max_size": resp.AutoScaleConfig.MaxSize,
			},
		}
		d.Set("auto_scale_config", autoScaleConfig)
	} else {
		d.Set("auto_scale_config", nil)
	}
	d.Set("image_id", resp.ImageId)
	if !checkSecurityGroupsSame(d, resp) {
		d.Set("security_groups", resp.SecurityGroups)
	}
	if !checkSecondarySubnetsSame(d, resp.SecondarySubnets) {
		d.Set("secondary_subnets", resp.SecondarySubnets)
	}
	d.Set("disk_size", resp.DiskSize)
	d.Set("disk_type", resp.DiskType)
	d.Set("enable_private_nodes", resp.EnablePrivateNodes)
	d.Set("flavor_id", resp.FlavorId)
	d.Set("name", resp.Name)
	d.Set("ssh_key_id", resp.SshKeyId)
	d.Set("enabled_encryption_volume", resp.EnabledEncryptionVolume)
	d.Set("subnet_id", resp.SubnetId)
	return nil
}

//func resourceClusterNodeGroupReadForCreate(d *schema.ResourceData, m interface{}) error {
//	clusterID := d.Get("cluster_id").(string)
//	cli := m.(*client.Client)
//	resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterID, d.Id(), nil)
//	if httpResponse.StatusCode == http.StatusNotFound {
//		d.SetId("")
//		return nil
//	}
//	if CheckErrorResponse(httpResponse) {
//		responseBody := GetResponseBody(httpResponse)
//		errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
//		return errorResponse
//	}
//	respJSON, _ := json.Marshal(resp)
//	log.Printf("-------------------------------------\n")
//	log.Printf("%s\n", string(respJSON))
//	log.Printf("-------------------------------------\n")
//	upgradeConfig := []interface{}{
//		map[string]interface{}{
//			"strategy":        resp.UpgradeConfig.Strategy,
//			"max_surge":       resp.UpgradeConfig.MaxSurge,
//			"max_unavailable": resp.UpgradeConfig.MaxUnavailable,
//		},
//	}
//	d.Set("upgrade_config", upgradeConfig)
//	if resp.AutoScaleConfig != nil {
//		autoScaleConfig := []interface{}{
//			map[string]interface{}{
//				"min_size": resp.AutoScaleConfig.MinSize,
//				"max_size": resp.AutoScaleConfig.MaxSize,
//			},
//		}
//		d.Set("auto_scale_config", autoScaleConfig)
//	} else {
//		d.Set("auto_scale_config", nil)
//	}
//	d.Set("image_id", resp.ImageId)
//	if !checkSecurityGroupsSame(d, resp) {
//		d.Set("security_groups", resp.SecurityGroups)
//	}
//	d.Set("disk_size", resp.DiskSize)
//	d.Set("disk_type", resp.DiskType)
//	d.Set("enable_private_nodes", resp.EnablePrivateNodes)
//	d.Set("flavor_id", resp.FlavorId)
//	d.Set("name", resp.Name)
//	d.Set("ssh_key_id", resp.SshKeyId)
//
//	return nil
//}

func resourceClusterNodeGroupCreate(d *schema.ResourceData, m interface{}) error {

	createNodeGroupRequest := getCreateNodeGroupRequest(d)
	cli := m.(*client.Client)
	request := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsPostOpts{
		Body: optional.NewInterface(createNodeGroupRequest),
	}
	resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsPost(context.TODO(), d.Get("cluster_id").(string), &request)

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
		Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, d.Get("cluster_id").(string), resp.Id),
		Timeout:    180 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for create cluster node group (%s) %s", resp.Id, err)
	}
	d.SetId(resp.Id)
	return resourceClusterNodeGroupRead(d, m)
}

func getSecondarySubnets(input []interface{}) []string {

	secondarySubnets := make([]string, len(input))

	for i, v := range input {
		str, ok := v.(string)
		if !ok {
			return []string{}
		}
		secondarySubnets[i] = str
	}

	return secondarySubnets
}

func getCreateNodeGroupRequest(d *schema.ResourceData) vks.CreateNodeGroupDto {
	taintsInput, ok := d.Get("taint").([]interface{})
	var tains []vks.NodeGroupTaintDto
	if ok {
		tains = getTaints(taintsInput)
	} else {
		tains = nil
	}
	return vks.CreateNodeGroupDto{
		Name:                    d.Get("name").(string),
		NumNodes:                int32(d.Get("num_nodes").(int)),
		ImageId:                 d.Get("image_id").(string),
		FlavorId:                d.Get("flavor_id").(string),
		DiskSize:                int32(d.Get("disk_size").(int)),
		DiskType:                d.Get("disk_type").(string),
		EnablePrivateNodes:      d.Get("enable_private_nodes").(bool),
		SshKeyId:                d.Get("ssh_key_id").(string),
		Labels:                  getLabels(d.Get("labels").(map[string]interface{})),
		Taints:                  tains,
		SecurityGroups:          getSecurityGroups(d.Get("security_groups").([]interface{})),
		UpgradeConfig:           getUpgradeConfig(d.Get("upgrade_config").([]interface{})),
		AutoScaleConfig:         getAutoScaleConfig(d.Get("auto_scale_config").([]interface{})),
		EnabledEncryptionVolume: d.Get("enabled_encryption_volume").(bool),
		SecondarySubnets:        getSecondarySubnets(d.Get("secondary_subnets").([]interface{})),
		SubnetId:                d.Get("subnet_id").(string),
	}
}

func resourceClusterNodeGroupUpdate(d *schema.ResourceData, m interface{}) error {
	hasChangeOtherField := false
	cli := m.(*client.Client)
	clusterId := d.Get("cluster_id").(string)
	clusterNodeGroupId := d.Id()
	if d.HasChange("security_groups") {
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterId, clusterNodeGroupId, nil)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			errorResponse := fmt.Errorf("request fail with errMsg : %s", responseBody)
			return errorResponse
		}
		if checkSecurityGroupsSame(d, resp) {
			return resourceClusterRead(d, m)
		} else {
			hasChangeOtherField = true
		}
	}
	var tains []vks.NodeGroupTaintDto
	var labels map[string]string
	if d.HasChange("taint") || d.HasChange("security_groups") || d.HasChange("labels") {
		hasChangeOtherField = true
		taintsInput, ok := d.Get("taint").([]interface{})
		if ok {
			tains = getTaints(taintsInput)
		} else {
			tains = nil
		}
		labels = getLabels(d.Get("labels").(map[string]interface{}))
	}
	if hasChangeOtherField || d.HasChange("auto_scale_config") || d.HasChange("num_nodes") || d.HasChange("upgrade_config") || d.HasChange("image_id") {
		securityGroupsRequest := d.Get("security_groups").([]interface{})
		var securityGroups []string
		for _, s := range securityGroupsRequest {
			securityGroups = append(securityGroups, s.(string))
		}
		if securityGroups == nil {
			securityGroups = make([]string, 0)
		}
		autoScaleConfig := getAutoScaleConfig(d.Get("auto_scale_config").([]interface{}))
		upgradeConfig := getUpgradeConfig(d.Get("upgrade_config").([]interface{}))
		var numNodes *int32 = nil
		if d.HasChange("num_nodes") {
			num := int32(d.Get("num_nodes").(int))
			numNodes = &num
		}
		imageId := d.Get("image_id").(string)
		updateNodeGroupRequest := vks.UpdateNodeGroupDto{
			AutoScaleConfig: autoScaleConfig,
			NumNodes:        numNodes,
			UpgradeConfig:   &upgradeConfig,
			SecurityGroups:  securityGroups,
			ImageId:         imageId,
			Labels:          labels,
			Taints:          tains,
		}
		requestPutOpts := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdPutOpts{
			Body: optional.NewInterface(updateNodeGroupRequest),
		}
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdPut(context.TODO(), clusterId, clusterNodeGroupId, &requestPutOpts)
		if CheckErrorResponse(httpResponse) {
			autoScaleConfig, _ := d.GetChange("auto_scale_config")
			numNodes, _ := d.GetChange("num_nodes")
			upgradeConfig, _ := d.GetChange("upgrade_config")
			imageId, _ := d.GetChange("image_id")
			securityGroups, _ := d.GetChange("security_groups")
			d.Set("auto_scale_config", autoScaleConfig)
			d.Set("num_nodes", numNodes)
			d.Set("upgrade_config", upgradeConfig)
			d.Set("image_id", imageId)
			d.Set("security_groups", securityGroups)
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
			Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, clusterId, d.Id()),
			Timeout:    180 * time.Minute,
			Delay:      10 * time.Second,
			MinTimeout: 1 * time.Second,
		}
		_, err := stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf("error waiting for update cluster node group (%s) %s", resp.Id, err)
		}
	}
	return resourceClusterNodeGroupRead(d, m)
}

func resourceClusterNodeGroupDelete(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	resp, httpResponse, err := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdDelete(context.TODO(), d.Get("cluster_id").(string), d.Id(), nil)
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
		Refresh:    resourceClusterNodeGroupDeleteStateRefreshFunc(cli, d.Get("cluster_id").(string), d.Id()),
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

func resourceClusterNodeGroupDeleteStateRefreshFunc(cli *client.Client, clusterId string, clusterNodeGroupId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterId, clusterNodeGroupId, nil)
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

func checkSecurityGroupsSame(d *schema.ResourceData, clusterNodeGroup vks.NodeGroupDetailDto) bool {
	securityGroupsRequest := d.Get("security_groups").([]interface{})
	var securityGroups []string
	for _, s := range securityGroupsRequest {
		securityGroups = append(securityGroups, s.(string))
	}
	var securityGroupsCluster []string
	for _, securityGroup := range clusterNodeGroup.SecurityGroups {
		securityGroupsCluster = append(securityGroupsCluster, securityGroup)
	}
	return CheckListStringEqual(securityGroups, securityGroupsCluster)
}

func resourceContainerClusterNodeGroupResourceV1() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"num_nodes": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"auto_scale_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"min_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},
						"max_size": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  10,
						},
					},
				},
			},
			"upgrade_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"strategy": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "SURGE",
						},
						"max_surge": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
						},
						"max_unavailable": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
					},
				},
				DefaultFunc: func() (interface{}, error) {
					return []interface{}{
						map[string]interface{}{
							"strategy":        "SURGE",
							"max_surge":       1,
							"max_unavailable": 0,
						},
					}, nil
				},
			},
			"image_id": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					return fetchByKey("image_id")
				},
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return fetchByKey("flavor_id")
				},
			},
			"disk_size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return 20, nil
				},
			},
			"disk_type": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return fetchByKey("volume_type_id")
				},
			},
			"enable_private_nodes": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				DefaultFunc: func() (interface{}, error) {
					return false, nil
				},
			},
			"security_groups": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ssh_key_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"labels": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: `The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.`,
			},
			"taint": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `List of Kubernetes taints to be applied to each node.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Key for taint.`,
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: `Value for taint.`,
						},
						"effect": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "NoSchedule",
							ValidateFunc: validation.StringInSlice([]string{"NoSchedule", "PreferNoSchedule", "NoExecute"}, false),
							Description:  `Effect for taint.`,
						},
					},
				},
			},
		},
	}
}

func resourceClusterNodeGroupStateUpgradeV0(ctx context.Context, rawState map[string]interface{}, meta interface{}) (map[string]interface{}, error) {
	log.Printf("resourceClusterNodeGroupStateUpgradeV0\n")
	cli := meta.(*client.Client)
	id, ok := rawState["id"].(string)
	clusterId, _ := rawState["cluster_id"].(string)
	if !ok {
		return nil, fmt.Errorf("id is missing or not a string")
	}

	resp, httpResponse, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterId, nil)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request cluster fail with errMsg : %s", responseBody)
		return rawState, errorResponse
	}
	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	if *resp.NetworkType == vks.CILIUM_NATIVE_ROUTING_NetworkType {
		nodeGroupResponse, httpNodeGroupResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterId, id, nil)

		if CheckErrorResponse(httpNodeGroupResponse) {
			responseBodyNodeGroup := GetResponseBody(httpNodeGroupResponse)
			errorResponseNodeGroup := fmt.Errorf("request cluster node group fail with errMsg : %s", responseBodyNodeGroup)
			return rawState, errorResponseNodeGroup
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		rawState["secondary_subnets"] = nodeGroupResponse.SecondarySubnets
	}

	return rawState, nil
}
