package vserver

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vserver"
	"golang.org/x/net/context"
	"log"
	"net/http"
	"time"
)

func ResourceAttachLb() *schema.Resource {
	return &schema.Resource{
		Create: resourceClusterLbAttach,
		Update: resourceClusterLbAttach,
		Delete: resourceClusterLbDetach,
		Read:   resourceClusterLbRead,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"cluster_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"load_balancers": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"load_balancer_id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"pools": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"monitor_port": {
										Type:     schema.TypeInt,
										Required: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"pool_id": {
										Type:     schema.TypeString,
										Required: true,
									},
									"port": {
										Type:     schema.TypeInt,
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceClusterLbAttach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterID := d.Get("cluster_id").(string)
	cli := m.(*client.Client)

	loadBalancers := d.Get("load_balancers").([]interface{})

	loadBalancerList := make([]vserver.LoadBalancerItem, 0, len(loadBalancers))

	for _, lb := range loadBalancers {
		loadBalancer := lb.(map[string]interface{})
		var loadBalancerItem vserver.LoadBalancerItem
		loadBalancerItem.LoadBalancerId = loadBalancer["load_balancer_id"].(string)
		loadBalancerItem.Pools = buildPoolItemRequest(loadBalancer["pools"].([]interface{}))
		loadBalancerList = append(loadBalancerList, loadBalancerItem)
	}

	req := vserver.AttachToLoadBalancerBackendRequest{
		ClusterId:     clusterID,
		LoadBalancers: loadBalancerList,
	}
	log.Printf("[DEBUG] %v\n", req)
	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.AttachToLbUsingPOST(context.TODO(), req, clusterID, projectID)

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
		Pending:    k8sClusterAttachingLB,
		Target:     k8sClusterAttachedLB,
		Refresh:    resourceServerTaskStateRefreshFunc(cli, clusterID, projectID, resp.Data.Id),
		Timeout:    60 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}
	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for attach load balancer to cluster (%s) %s", clusterID, err)
	}

	d.SetId(clusterID)

	return resourceClusterLbRead(d, m)
}

func buildPoolItemRequest(ps []interface{}) []vserver.PoolItem {
	poolItemList := make([]vserver.PoolItem, 0, len(ps))
	for _, p := range ps {
		pool := p.(map[string]interface{})
		var poolItem vserver.PoolItem
		poolItem.MonitorPort = int32(pool["monitor_port"].(int))
		poolItem.Name = pool["name"].(string)
		poolItem.PoolId = pool["pool_id"].(string)
		poolItem.Port = int32(pool["port"].(int))
		poolItemList = append(poolItemList, poolItem)
	}
	return poolItemList
}

func resourceClusterLbDetach(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterID := d.Id()
	cli := m.(*client.Client)
	loadBalancerList := make([]vserver.LoadBalancerItem, 0)

	req := vserver.AttachToLoadBalancerBackendRequest{
		ClusterId:     clusterID,
		LoadBalancers: loadBalancerList,
	}

	resp, httpResponse, err := cli.VserverClient.K8SClusterRestControllerApi.AttachToLbUsingPOST(context.TODO(), req, clusterID, projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		errorResponse := fmt.Errorf("request faile with errMsg: %s", responseBody)
		return errorResponse
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	stateConf := &resource.StateChangeConf{
		Pending:    k8sClusterAttachingLB,
		Target:     k8sClusterAttachedLB,
		Refresh:    resourceServerTaskStateRefreshFunc(cli, clusterID, projectID, resp.Data.Id),
		Timeout:    50 * time.Minute,
		Delay:      10 * time.Second,
		MinTimeout: 1 * time.Second,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("error waiting for attach load balancer to cluster (%s) %s", clusterID, err)
	}

	d.SetId("")

	return nil
}

func resourceClusterLbRead(d *schema.ResourceData, m interface{}) error {
	projectID := d.Get("project_id").(string)
	clusterID := d.Id()
	cli := m.(*client.Client)
	resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.ListClusterPoolByClusterUsingGET(context.TODO(), clusterID, projectID)

	if CheckErrorResponse(httpResponse) {
		responseBody := GetResponseBody(httpResponse)
		err := fmt.Errorf("request fail with errMsg: %s", responseBody)
		return err
	}

	respJSON, _ := json.Marshal(resp)
	log.Printf("-------------------------------------\n")
	log.Printf("%s\n", string(respJSON))
	log.Printf("-------------------------------------\n")

	if len(resp) > 0 {
		d.Set("load_balancer_id", resp[0].LoadBalancerId)
		d.Set("cluster_id", resp[0].ClusterId)
		//d.Set("cluster_pool", flattenPools(resp))
	} else {
		d.Set("load_balancer_id", "")
		d.Set("cluster_id", clusterID)
		//d.Set("cluster_pool", "")
	}
	return nil
}

func flattenPools(poolLists []vserver.ClusterPoolDto) []map[string]interface{} {
	if len(poolLists) == 0 {
		return []map[string]interface{}{}
	}

	l := make([]map[string]interface{}, 0, len(poolLists))

	for _, poolList := range poolLists {
		m := map[string]interface{}{
			"monitor_port": poolList.MonitorPort,
			"name":         poolList.PoolName,
			"pool_id":      poolList.PoolId,
			"port":         int(poolList.Port),
		}

		l = append(l, m)
	}
	return l
}

func resourceServerTaskStateRefreshFunc(cli *client.Client, clusterId string, projectId string, serverTaskId string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		resp, httpResponse, _ := cli.VserverClient.K8SClusterRestControllerApi.GetServerTaskStatusUsingGET(context.TODO(), clusterId, projectId, serverTaskId)

		if httpResponse.StatusCode != http.StatusOK {
			return nil, "", fmt.Errorf("error describing: %s", GetResponseBody(httpResponse))
		}

		log.Printf("-------------------------------------\n")
		log.Printf("%s\n", string(resp))
		log.Printf("-------------------------------------\n")

		return resp, resp, nil
	}
}
