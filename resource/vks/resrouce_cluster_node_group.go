package vks

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
	"log"
	"net/http"
)

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
		Computed: true,
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
		DefaultFunc: func() (interface{}, error) {
			return []interface{}{
				map[string]interface{}{
					"max_size": 10,
					"min_size": 1,
				},
			}, nil
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
			return "img-108b3a77-ab58-4000-9b3e-190d0b4b07fc", nil
		},
	},
	"flavor_id": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
		DefaultFunc: func() (interface{}, error) {
			return "flav-152e00ea-5412-44c5-868f-0b30e37e4f90", nil
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
			return "vtype-4b437071-da9c-4e1b-a2c0-839fb89e5182", nil
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
		ForceNew:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: `The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.`,
	},
	"taint": {
		Type:        schema.TypeList,
		Optional:    true,
		ForceNew:    true,
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
}

func getCreateNodeGroupReuqest(nodeGroup map[string]interface{}) vks.CreateNodeGroupDto {
	taintsInput, ok := nodeGroup["taints"].([]interface{})
	var tains []vks.NodeGroupTaintDto
	if ok {
		tains = getTaints(taintsInput)
	} else {
		tains = nil
	}
	return vks.CreateNodeGroupDto{
		Name:               nodeGroup["name"].(string),
		NumNodes:           int32(nodeGroup["num_nodes"].(int)),
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
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdNodesGet(context.TODO(), clusterID, clusterNodeGroupID, nil)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("Error : %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		clusterNodeGroup := resp
		return clusterNodeGroup, clusterNodeGroup.Items[0].Status, nil
	}
}
