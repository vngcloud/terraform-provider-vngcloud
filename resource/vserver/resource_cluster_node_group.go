package vserver

//import (
//	"context"
//	"encoding/json"
//	"fmt"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
//	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
//	"github.com/vngcloud/terraform-provider-vngcloud/client"
//	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
//	"log"
//	"net/http"
//	"strings"
//	"time"
//)
//
//func ResourceClusterNodeGroup() *schema.Resource {
//	return &schema.Resource{
//		Create: resourceClusterNodeGroupCreate,
//		Read:   resourceClusterNodeGroupRead,
//		Update: UpdateClusterNodeGroupFunc,
//		Delete: DeleteClusterNodeGroupFunc,
//		Importer: &schema.ResourceImporter{
//			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
//				idParts := strings.Split(d.Id(), ":")
//				if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
//					return nil, fmt.Errorf("Unexpected format of ID (%q), expected ProjectID:NodeGroupId", d.Id())
//				}
//				projectID := idParts[0]
//				clusterId := idParts[1]
//				d.SetId(clusterId)
//				d.Set("project_id", projectID)
//				return []*schema.ResourceData{d}, nil
//			},
//		},
//		Schema: map[string]*schema.Schema{
//			"project_id": {
//				Type:     schema.TypeString,
//				Required: true,
//				ForceNew: true,
//			},
//			"name": {
//				Type:     schema.TypeString,
//				Required: true,
//				ForceNew: true,
//			},
//			"cluster_id": {
//				Type:     schema.TypeString,
//				Required: true,
//				ForceNew: true,
//			},
//			"flavor_id": {
//				Type:     schema.TypeString,
//				Required: true,
//				ForceNew: true,
//			},
//			"node_count": {
//				Type:     schema.TypeInt,
//				Required: true,
//			},
//		},
//	}
//}
//
//func resourceClusterNodeGroupCreate(d *schema.ResourceData, m interface{}) error {
//	log.Printf("Create Cluster Node Group")
//	projectId := d.Get("project_id").(string)
//	cli := m.(*client.Client)
//
//	nodeGroup := vserver.CreateNodeGroupBackendRequest{
//		ClusterId: d.Get("cluster_id").(string),
//		FlavorId:  d.Get("flavor_id").(string),
//		Name:      d.Get("name").(string),
//		NodeCount: int32(d.Get("node_count").(int)),
//	}
//
//	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.CreateClusterNodeGroupUsingPOST(context.TODO(), d.Id(), nodeGroup, projectId)
//
//	if CheckErrorResponse(httpResponse) {
//		responseBody := GetResponseBody(httpResponse)
//		errResponse := fmt.Errorf("request fail with errMsg: %s", responseBody)
//		return errResponse
//	}
//
//	respJSON, _ := json.Marshal(resp)
//	log.Printf("-------------------------------------\n")
//	log.Printf("%s\n", string(respJSON))
//	log.Printf("-------------------------------------\n")
//
//	stateConf := &resource.StateChangeConf{
//		Pending:    clusterNodeGroupCreating,
//		Target:     clusterNodeGroupCreated,
//		Refresh:    resourceClusterNodeGroupStateRefreshFunc(cli, resp.Data.ClusterId, resp.Data.Uuid, projectId),
//		Timeout:    25 * time.Minute,
//		Delay:      10 * time.Second,
//		MinTimeout: 1 * time.Second,
//	}
//
//	_, err = stateConf.WaitForState()
//	if err != nil {
//		fmt.Errorf("error waiting for create cluster node group (%s) %s", resp.Data.Uuid, err)
//	}
//	d.SetId(resp.Data.ClusterId)
//	return resourceClusterNodeGroupRead(d, m)
//}
//
//func resourceClusterNodeGroupStateRefreshFunc(cli *client.Client, clusterId string, nodegroupId string, projectId string) resource.StateRefreshFunc {
//	return func() (interface{}, string, error) {
//		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.ListClusterNodeGroupUsingGET(context.TODO(), clusterId, projectId)
//		if httpResponse.StatusCode != http.StatusOK {
//			return nil, "", fmt.Errorf("Error describing instance: %s", GetResponseBody(httpResponse))
//		}
//		respJSON, _ := json.Marshal(resp)
//		log.Printf("-------------------------------------\n")
//		log.Printf("%s\n", string(respJSON))
//		log.Printf("-------------------------------------\n")
//		clusterNodegroupList := resp
//
//		for _, nodeGroup := range clusterNodegroupList {
//			if nodeGroup.Uuid == nodegroupId {
//				return nodeGroup, nodeGroup.Status, nil
//			}
//		}
//		return nil, "", fmt.Errorf("can't not get cluster node group id %s", clusterId)
//	}
//}
//
//func resourceClusterNodeGroupRead(d *schema.ResourceData, m interface{}) error {
//	log.Printf("Read Cluster Node Group")
//	projectId := d.Get("project_id").(string)
//	nodeGroupId := d.Id()
//	cli := m.(*client.Client)
//
//	resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetClusterUsingGET(context.TODO())
//}

//
//func UpdateClusterNodeGroupFunc(d *schema.ResourceData, m interface{}) error {
//
//}
//
//func DeleteClusterNodeGroupFunc(d *schema.ResourceData, m interface{}) error {
//
//}
