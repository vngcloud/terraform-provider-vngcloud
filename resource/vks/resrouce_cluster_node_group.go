package vks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vks"
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

// inlineNodeGroupTaintSchema overrides schemaNodeGroup["taint"] (TypeSet) for the INLINE
// node_group block on vngcloud_vks_cluster only (see ResourceCluster's node_group Elem, which
// merges this in after schemaNodeGroup). TypeSet + ConfigModeAttr, while correct for the
// standalone vngcloud_vks_cluster_node_group resource (a top-level attribute), is silently
// decoded as empty by terraform-plugin-sdk/v2 (confirmed on v2.6.1 through v2.40.1) when nested
// inside another block the way node_group is — any add/remove/edit of a specific taint on the
// inline node_group would go undetected or misapplied. TypeList does not have this bug when
// nested. Order-independence for the inline case is instead handled in Go at apply time (see
// taintDtoSetsEqual in resource_cluster.go), the same pattern already used for security_groups
// via checkSecurityGroupsSame — not by the schema type.
//
// WARNING: resourceClusterCustomizeDiff (resource_cluster.go) calls d.Clear on a path ending in
// ".taint", and the vendored SDK's Clear matches diff keys by plain string prefix (no delimiter
// boundary check). Do not add another field to this schema whose name starts with "taint" (e.g.
// "taint_policy") without checking that call site — it would also get silently cleared.
var inlineNodeGroupTaintSchema = &schema.Schema{
	Type:     schema.TypeList,
	Optional: true,
	Computed: true,
	// Elem is a *Resource, so SDK v2's Auto ConfigMode would default to block syntax
	// (taint { ... }) regardless of Optional/Computed; ConfigModeAttr forces attribute
	// syntax (taint = [...]) instead, which is required so taint = [] can explicitly
	// clear all taints (block syntax can never represent "explicitly empty").
	ConfigMode:  schema.SchemaConfigModeAttr,
	Description: `List of Kubernetes taints to be applied to each node, e.g. taint = [{ key = "...", value = "...", effect = "NoSchedule" }]. Omit to leave existing taints unchanged; set taint = [] to explicitly remove all taints.`,
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
		Computed: true,
		DiffSuppressFunc: func(k, oldVal, newVal string, d *schema.ResourceData) bool {
			return true
		},
	},
	"kubernetes_version": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
	},
	"os": {
		Type:     schema.TypeString,
		Optional: true,
		Computed: true,
		ForceNew: true,
	},
	"flavor_id": {
		Type:     schema.TypeString,
		Optional: true,
		ForceNew: true,
		Computed: true,
		// DefaultFunc: func() (interface{}, error) {
		// 	return fetchByKey("flavor_id")
		// },
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
		Computed: true,
		// DefaultFunc: func() (interface{}, error) {
		// 	return fetchByKey("volume_type_id")
		// },
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
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: `The map of Kubernetes labels (key/value pairs) to be applied to each node. These will added in addition to any default label(s) that Kubernetes may apply to the node.`,
	},
	"taint": {
		Type:     schema.TypeSet,
		Optional: true,
		Computed: true,
		// Elem is a *Resource, so SDK v2's Auto ConfigMode would default to block syntax
		// (taint { ... }) regardless of Optional/Computed; ConfigModeAttr forces attribute
		// syntax (taint = [...]) instead, which is required so taint = [] can explicitly
		// clear all taints (block syntax can never represent "explicitly empty").
		// TypeSet (rather than TypeList) makes the diff order-independent: taints are
		// compared by content (hashed on key+value+effect), not by position, so the API
		// returning them in a different order than the config doesn't trigger a spurious update.
		ConfigMode:  schema.SchemaConfigModeAttr,
		Description: `List of Kubernetes taints to be applied to each node, e.g. taint = [{ key = "...", value = "...", effect = "NoSchedule" }]. Omit to leave existing taints unchanged; set taint = [] to explicitly remove all taints. Order does not matter, and duplicate entries with identical key/value/effect are treated as a single taint.`,
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
	"tags": {
		Type:        schema.TypeMap,
		Optional:    true,
		Computed:    true,
		Elem:        &schema.Schema{Type: schema.TypeString},
		Description: `Key-value pairs of cloud tags to apply to all VMs and volumes in the node group.`,
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
	d.Set("kubernetes_version", resp.KubernetesVersion)
	d.Set("os", resp.ImageOS)
	if !checkSecurityGroupsSame(d, resp) {
		d.Set("security_groups", resp.SecurityGroups)
	}
	respCluster, httpResponseCluster, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterID, nil)
	if CheckErrorResponse(httpResponseCluster) {
		responseBodyCluster := GetResponseBody(httpResponseCluster)
		errorResponseCluster := fmt.Errorf("request cluster fail with errMsg : %s", responseBodyCluster)
		return errorResponseCluster
	}
	respJSONCluster, _ := json.Marshal(respCluster)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSONCluster))
	log.Printf("-------------------------------------\n")
	if respCluster.NetworkType == "CILIUM_NATIVE_ROUTING" && !checkSecondarySubnetsSame(d, resp.SecondarySubnets) {
		d.Set("secondary_subnets", resp.SecondarySubnets)
	}
	d.Set("disk_size", resp.DiskSize)
	d.Set("disk_type", resp.DiskType)
	d.Set("enable_private_nodes", resp.EnablePrivateNodes)
	d.Set("flavor_id", resp.FlavorId)
	d.Set("name", resp.Name)
	d.Set("ssh_key_id", resp.SshKeyId)
	d.Set("enabled_encryption_volume", resp.EnabledEncryptionVolume)
	// d.Set("subnet_id", resp.SubnetId)
	_, hasSubnetId := d.GetOk("subnet_id")
	if !hasSubnetId {
		d.Set("subnet_id", resp.SubnetId)
	}

	// Import labels
	if resp.Labels != nil {
		d.Set("labels", resp.Labels)
	}

	// Import taints
	if resp.Taints != nil {
		taints := make([]interface{}, len(resp.Taints))
		for i, taint := range resp.Taints {
			taints[i] = map[string]interface{}{
				"key":    taint.Key,
				"value":  taint.Value,
				"effect": taint.Effect,
			}
		}
		d.Set("taint", taints)
	}

	// Import tags
	if resp.Tags != nil {
		d.Set("tags", resp.Tags)
	}

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

	cli := m.(*client.Client)
	clusterId := d.Get("cluster_id").(string)

	respCluster, httpResponseCluster, _ := cli.VksClient.V1ClusterControllerApi.V1ClustersClusterIdGet(context.TODO(), clusterId, nil)
	if CheckErrorResponse(httpResponseCluster) {
		responseBodyCluster := GetResponseBody(httpResponseCluster)
		errResponseCluster := fmt.Errorf("request get cluster fail with errMsg: %s", responseBodyCluster)
		return errResponseCluster
	}
	respJSONCluster, _ := json.Marshal(respCluster)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSONCluster))
	log.Printf("-------------------------------------\n")

	_, hasSubnetId := d.GetOk("subnet_id")
	if !hasSubnetId {
		d.Set("subnet_id", respCluster.SubnetId)
	}

	if respCluster.NetworkType == "CILIUM_NATIVE_ROUTING" && (d.Get("secondary_subnets") == nil || len(d.Get("secondary_subnets").([]interface{})) == 0) {
		return fmt.Errorf("secondary_subnets is required when cluster network type is set to CILIUM_NATIVE_ROUTING")
	}

	setDefaultValueByZone(d, m, respCluster.VpcId)

	createNodeGroupRequest, errorNodeGroup := getCreateNodeGroupRequest(d)
	if errorNodeGroup != nil {
		return errorNodeGroup
	}

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

func setDefaultValueByZone(d *schema.ResourceData, m interface{}, vpcId string) error {
	cli := m.(*client.Client)

	_, hasFlavorId := d.GetOk("flavor_id")
	_, hasDiskType := d.GetOk("disk_type")

	if !hasFlavorId || !hasDiskType {
		workspaceRes, httpResponse, _ := cli.VksClient.V1WorkspaceControllerApi.V1WorkspaceGet(context.TODO(), nil)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return fmt.Errorf("request fail with errMsg: %s", responseBody)
		}

		subnetId := d.Get("subnet_id").(string)
		subnetRes, httpResponse, _ := cli.VserverClient.SubnetRestControllerApi.GetSubnetByIdUsingGET(context.TODO(), vpcId, workspaceRes.ProjectId, subnetId)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return fmt.Errorf("request fail with errMsg: %s", responseBody)
		}
		flavorIdKey := ""
		diskTypeKey := ""
		if cli.VksClient.Config().BasePath == "https://vks-han-1.api.vngcloud.vn" {
			flavorIdKey = "han01_1a_flavor_id"
			diskTypeKey = "han01_1a_volume_type_id"
		} else {
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
		if !hasFlavorId {
			res, _ := fetchByKey(flavorIdKey)
			d.Set("flavor_id", res.(string))
		}
		if !hasDiskType {
			res, _ := fetchByKey(diskTypeKey)
			d.Set("disk_type", res.(string))
		}
	}

	return nil
}

func getSecondarySubnets(input []interface{}) ([]string, error) {

	secondarySubnets := make([]string, len(input))

	for i, v := range input {
		str, ok := v.(string)
		if !ok {
			return nil, fmt.Errorf("secondary_subnets is required when cluster network type is set to CILIUM_NATIVE_ROUTING")
		}
		secondarySubnets[i] = str
	}

	return secondarySubnets, nil
}

func getCreateNodeGroupRequest(d *schema.ResourceData) (vks.CreateNodeGroupDto, error) {
	var tains []vks.NodeGroupTaintDto
	if taintSet, ok := d.Get("taint").(*schema.Set); ok {
		tains = getTaints(taintSet.List())
	} else {
		tains = nil
	}
	secondarySubnets, errSecondarySubnets := getSecondarySubnets(d.Get("secondary_subnets").([]interface{}))
	if errSecondarySubnets != nil {
		return vks.CreateNodeGroupDto{}, errSecondarySubnets
	}
	return vks.CreateNodeGroupDto{
		Name:                    d.Get("name").(string),
		NumNodes:                int32(d.Get("num_nodes").(int)),
		Os:                      d.Get("os").(string),
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
		SecondarySubnets:        secondarySubnets,
		SubnetId:                d.Get("subnet_id").(string),
		Tags:                    getLabels(d.Get("tags").(map[string]interface{})),
	}, nil
}

func resourceClusterNodeGroupUpdate(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)
	clusterId := d.Get("cluster_id").(string)
	clusterNodeGroupId := d.Id()

	// Block 1: PUT — handles num_nodes, auto_scale_config, upgrade_config, security_groups.
	// labels and taints are intentionally omitted (zero values); server ignores them per API spec.
	hasChangeOtherField := false
	if d.HasChange("security_groups") {
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdGet(context.TODO(), clusterId, clusterNodeGroupId, nil)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return fmt.Errorf("request fail with errMsg : %s", responseBody)
		}
		if checkSecurityGroupsSame(d, resp) {
			return resourceClusterRead(d, m)
		}
		hasChangeOtherField = true
	}

	if hasChangeOtherField || d.HasChange("auto_scale_config") || d.HasChange("num_nodes") || d.HasChange("upgrade_config") {
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
		updateNodeGroupRequest := vks.UpdateNodeGroupDto{
			AutoScaleConfig: autoScaleConfig,
			NumNodes:        numNodes,
			UpgradeConfig:   &upgradeConfig,
			SecurityGroups:  securityGroups,
		}
		requestPutOpts := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdPutOpts{
			Body: optional.NewInterface(updateNodeGroupRequest),
		}
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdPut(context.TODO(), clusterId, clusterNodeGroupId, &requestPutOpts)
		if CheckErrorResponse(httpResponse) {
			autoScaleConfig, _ := d.GetChange("auto_scale_config")
			numNodes, _ := d.GetChange("num_nodes")
			upgradeConfig, _ := d.GetChange("upgrade_config")
			securityGroups, _ := d.GetChange("security_groups")
			d.Set("auto_scale_config", autoScaleConfig)
			d.Set("num_nodes", numNodes)
			d.Set("upgrade_config", upgradeConfig)
			d.Set("security_groups", securityGroups)
			responseBody := GetResponseBody(httpResponse)
			return fmt.Errorf("request fail with errMsg: %s", responseBody)
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

	// Block 2: PATCH /metadata — handles labels, taints, tags.
	// Always sends all 3 fields when triggered so portal-side drift is corrected in the same call.
	if d.HasChange("labels") || d.HasChange("taint") || d.HasChange("tags") {
		labels := getLabels(d.Get("labels").(map[string]interface{}))
		var taints []vks.NodeGroupTaintDto
		if taintSet, ok := d.Get("taint").(*schema.Set); ok {
			taints = getTaints(taintSet.List())
		} else {
			taints = nil
		}
		tags := getLabels(d.Get("tags").(map[string]interface{}))

		patchRequest := vks.PatchNodeGroupMetadataDto{
			Labels: &labels,
			Taints: &taints,
			Tags:   &tags,
		}
		patchOpts := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdMetadataPatchOpts{
			Body: optional.NewInterface(patchRequest),
		}
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdMetadataPatch(context.TODO(), clusterId, clusterNodeGroupId, &patchOpts)
		if CheckErrorResponse(httpResponse) {
			labels, _ := d.GetChange("labels")
			taint, _ := d.GetChange("taint")
			tags, _ := d.GetChange("tags")
			d.Set("labels", labels)
			d.Set("taint", taint)
			d.Set("tags", tags)
			responseBody := GetResponseBody(httpResponse)
			return fmt.Errorf("request fail with errMsg: %s", responseBody)
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")

		stateConf := &resource.StateChangeConf{
			Pending:    UPDATING,
			Target:     ACTIVE,
			Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, clusterId, clusterNodeGroupId),
			Timeout:    180 * time.Minute,
			Delay:      10 * time.Second,
			MinTimeout: 1 * time.Second,
		}
		_, err := stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf("error waiting for patch cluster node group metadata (%s) %s", clusterNodeGroupId, err)
		}
	}

	// Block 3: upgradeVersion — handles kubernetes_version.
	if d.HasChange("kubernetes_version") {
		newVersion := d.Get("kubernetes_version").(string)
		upgradeVersionRequest := vks.UpgradeNodeGroupVersionDto{
			KubernetesVersion: newVersion,
		}
		upgradeVersionOpts := vks.V1NodeGroupControllerApiV1ClustersClusterIdNodeGroupsNodeGroupIdUpgradeVersionPostOpts{
			Body: optional.NewInterface(upgradeVersionRequest),
		}
		resp, httpResponse, _ := cli.VksClient.V1NodeGroupControllerApi.V1ClustersClusterIdNodeGroupsNodeGroupIdUpgradeVersionPost(
			context.TODO(), clusterId, clusterNodeGroupId, &upgradeVersionOpts)
		if CheckErrorResponse(httpResponse) {
			oldVersion, _ := d.GetChange("kubernetes_version")
			d.Set("kubernetes_version", oldVersion)
			responseBody := GetResponseBody(httpResponse)
			return fmt.Errorf("request fail with errMsg: %s", responseBody)
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		stateConf := &resource.StateChangeConf{
			Pending:    UPGRADING,
			Target:     ACTIVE,
			Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, clusterId, clusterNodeGroupId),
			Timeout:    180 * time.Minute,
			Delay:      10 * time.Second,
			MinTimeout: 1 * time.Second,
		}
		_, err := stateConf.WaitForState()
		if err != nil {
			return fmt.Errorf("error waiting for upgrade cluster node group version (%s) %s", clusterNodeGroupId, err)
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
	if resp.NetworkType == "CILIUM_NATIVE_ROUTING" {
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
