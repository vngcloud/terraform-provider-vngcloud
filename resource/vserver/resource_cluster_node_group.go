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

func ResourceClusterNodeGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterNodeGroupCreate,
		Read:   resourceClusterNodeGroupRead,
		Update: UpdateClusterNodeGroupFunc,
		Delete: DeleteClusterNodeGroupFunc,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				idParts := strings.Split(d.Id(), ":")
				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("unexpected format of ID (%q), expected ProjectID:NodeGroupId", d.Id())
				}
				projectID := idParts[0]
				clusterNodeGroupId := idParts[1]
				d.SetId(clusterNodeGroupId)
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"flavor_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"node_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceClusterNodeGroupCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("Create Cluster Node Group")
	projectId := d.Get("project_id").(string)
	cli := m.(*client.Client)

	nodeGroup := vserver.CreateNodeGroupBackendRequest{
		ClusterId: d.Get("cluster_id").(string),
		FlavorId:  d.Get("flavor_id").(string),
		Name:      d.Get("name").(string),
		NodeCount: int32(d.Get("node_count").(int)),
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.CreateClusterNodeGroupUsingPOST(context.TODO(), nodeGroup.ClusterId, nodeGroup, projectId)

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
		Pending:    clusterNodeGroupCreating,
		Target:     clusterNodeGroupCreated,
		Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, resp.Data.Uuid, projectId),
		Timeout:    25 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		fmt.Errorf("error waiting for create cluster node group (%s) %s", resp.Data.Uuid, err)
	}
	d.SetId(resp.Data.Uuid)
	return resourceClusterNodeGroupRead(d, m)
}

func resourceClusterNodeGroupStateRefreshFunc(cli *client.Client, nodegroupId string, projectId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterNodeGroupUsingGET(context.TODO(), nodegroupId, projectId)
		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("error describing instance: %s", GetResponseBody(httpResponse))
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		clusterNodegroup := resp
		return clusterNodegroup, clusterNodegroup.Status, nil
	}
}

func resourceClusterNodeGroupRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("Read Cluster Node Group")
	projectId := d.Get("project_id").(string)
	nodeGroupId := d.Id()
	cli := m.(*client.Client)

	resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterNodeGroupUsingGET(context.TODO(), nodeGroupId, projectId)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")
	nodeGroup := resp
	d.Set("name", nodeGroup.Name)
	d.Set("cluster_id", nodeGroup.ClusterId)
	d.Set("node_count", nodeGroup.NodeCount)
	d.Set("status", nodeGroup.Status)
	d.Set("flavor_name", nodeGroup.FlavorName)
	return nil
}

func UpdateClusterNodeGroupFunc(d *schema.ResourceData, m interface{}) error {
	if d.HasChange("node_count") {
		nodeGroupId := d.Id()
		clusterId := d.Get("cluster_id").(string)
		return ResourceK8sScalingNodeGroup(d, m, nodeGroupId, clusterId)
	}
	return resourceClusterNodeGroupRead(d, m)
}

func DeleteClusterNodeGroupFunc(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterID := d.Get("cluster_id").(string)
	cli := m.(*client.Client)

	deleterNodeGroup := vserver.DeleteNodeGroupBackendRequest{
		ClusterId:   clusterID,
		NodeGroupId: d.Id(),
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.DeleteClusterNodeGroupUsingDELETE(context.TODO(), clusterID, deleterNodeGroup, projectID)

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
		Pending:    clusterNodeGroupDeleting,
		Target:     clusterNodeGroupDeleted,
		Refresh:    resourceClusterNodeGroupDeleteStateRefreshFunc(cli, d.Id(), projectID),
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

func resourceClusterNodeGroupDeleteStateRefreshFunc(cli *client.Client, nodegroupId string, projectID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterNodeGroupUsingGET(context.TODO(), nodegroupId, projectID)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vserver.ClusterNodeGroupDto{Status: "DELETED"}, "DELETED", nil
			} else {
				return nil, "", fmt.Errorf("error describing instance: %s", GetResponseBody(httpResponse))
			}
		}
		respJSON, _ := json.Marshal(resp)
		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(respJSON))
		log.Printf("-------------------------------------\n")
		nodeGroup := resp
		return nodeGroup, nodeGroup.Status, nil
	}
}

func ResourceK8sScalingNodeGroup(d *schema.ResourceData, m interface{}, nodegroupId string, clusterId string) error {
	projectID := d.Get("project_id").(string)
	cli := m.(*client.Client)

	o, n := d.GetChange("node_count")
	scaleMinionRequest := vserver.ScaleMinionBackendRequest{
		ClusterId:   clusterId,
		NodeCount:   int32(n.(int)),
		NodeGroupId: nodegroupId,
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.ScaleMinionUsingPOST(context.TODO(), clusterId, projectID, scaleMinionRequest)
	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
		d.Set("node_count", o.(int))
		return errorResponse
	}
	scaleMinionRespJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(scaleMinionRespJSON))
	log.Printf("-------------------------------------\n")
	stateConf := &resource.StateChangeConf{
		Pending:    clusterNodeGroupScaling,
		Target:     clusterNodeGroupScaled,
		Refresh:    resourceK8sClusterStateRefreshFunc(cli, clusterId, projectID),
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for instance %s to be resizing: %s", nodegroupId, err)
	}
	return nil
}
