package vdbv2

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
)

func ResourceKafkaTopic() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaTopicCreate,
		Read:   resourceKafkaTopicRead,
		Update: resourceKafkaTopicUpdate,
		Delete: resourceKafkaTopicDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
					return nil, fmt.Errorf("expected import id in format <cluster_id>/<topic_id>")
				}
				d.Set("cluster_id", parts[0])
				d.SetId(parts[1])
				return []*schema.ResourceData{d}, nil
			},
		},

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
			"partitions": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"replicas": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"retention_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"retention_bytes": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_user_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_at": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceKafkaTopicCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka topic create")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)
	defer lockKafkaCluster(clusterID)()

	body := vdbv2.TopicCreateRequest{
		Name:             d.Get("name").(string),
		Partitions:       int32(d.Get("partitions").(int)),
		Replicas:         int32(d.Get("replicas").(int)),
		RetentionSeconds: int64(d.Get("retention_seconds").(int)),
		RetentionBytes:   int64(d.Get("retention_bytes").(int)),
	}

	resp, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.CreateTopic(context.TODO(), body, clusterID)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("create kafka topic: %s", GetResponseBody(httpResponse))
	}
	if resp.Id == "" {
		return fmt.Errorf("create kafka topic: empty topic id in response")
	}
	d.SetId(resp.Id)
	time.Sleep(10 * time.Second)

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaTopicCreatePending,
		Target:     kafkaTopicCreateTarget,
		Refresh:    resourceKafkaTopicStateRefreshFunc(cli, clusterID, resp.Id),
		Timeout:    kafkaTopicCreateTimeout,
		Delay:      kafkaTopicCreateDelay,
		MinTimeout: kafkaTopicCreateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka topic to be created: %s", err)
	}

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterTopicUserPending, kafkaClusterTopicUserTarget, kafkaClusterTopicUserTimeout, kafkaClusterTopicUserDelay, kafkaClusterTopicUserMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be ACTIVE: %s", err)
	}

	return resourceKafkaTopicRead(d, m)
}

func resourceKafkaTopicRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka topic read")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)

	topic, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetTopicById(context.TODO(), clusterID, d.Id())
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("get kafka topic: %s", GetResponseBody(httpResponse))
	}

	d.Set("cluster_id", topic.ClusterId)
	d.Set("name", topic.Name)
	d.Set("partitions", topic.Partitions)
	d.Set("replicas", topic.Replicas)
	d.Set("retention_seconds", topic.RetentionSeconds)
	d.Set("retention_bytes", topic.RetentionBytes)
	d.Set("status", topic.Status)
	d.Set("portal_user_id", topic.PortalUserId)
	d.Set("created_at", topic.CreatedAt)
	return nil
}

func resourceKafkaTopicUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka topic update")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)

	if !d.HasChange("partitions") && !d.HasChange("replicas") && !d.HasChange("retention_seconds") && !d.HasChange("retention_bytes") {
		return resourceKafkaTopicRead(d, m)
	}
	defer lockKafkaCluster(clusterID)()

	body := vdbv2.TopicUpdateRequest{
		Partitions:       int32(d.Get("partitions").(int)),
		Replicas:         int32(d.Get("replicas").(int)),
		RetentionSeconds: int64(d.Get("retention_seconds").(int)),
		RetentionBytes:   int64(d.Get("retention_bytes").(int)),
	}

	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateTopic(context.TODO(), body, clusterID, d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka topic: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaTopicUpdatePending,
		Target:     kafkaTopicUpdateTarget,
		Refresh:    resourceKafkaTopicStateRefreshFunc(cli, clusterID, d.Id()),
		Timeout:    kafkaTopicUpdateTimeout,
		Delay:      kafkaTopicUpdateDelay,
		MinTimeout: kafkaTopicUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka topic to be updated: %s", err)
	}

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterTopicUserPending, kafkaClusterTopicUserTarget, kafkaClusterTopicUserTimeout, kafkaClusterTopicUserDelay, kafkaClusterTopicUserMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be ACTIVE: %s", err)
	}

	return resourceKafkaTopicRead(d, m)
}

func resourceKafkaTopicDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka topic delete")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)
	defer lockKafkaCluster(clusterID)()

	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.DeleteTopic(context.TODO(), clusterID, d.Id())
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("delete kafka topic: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaTopicDeletePending,
		Target:     kafkaTopicDeleteTarget,
		Refresh:    resourceKafkaTopicDeleteStateRefreshFunc(cli, clusterID, d.Id()),
		Timeout:    kafkaTopicDeleteTimeout,
		Delay:      kafkaTopicDeleteDelay,
		MinTimeout: kafkaTopicDeleteMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka topic to be deleted: %s", err)
	}

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterTopicUserPending, kafkaClusterTopicUserTarget, kafkaClusterTopicUserTimeout, kafkaClusterTopicUserDelay, kafkaClusterTopicUserMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be ACTIVE: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceKafkaTopicStateRefreshFunc(cli *client.Client, clusterID, topicID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		topic, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetTopicById(context.TODO(), clusterID, topicID)
		if CheckErrorResponse(httpResponse) {
			return nil, "", fmt.Errorf("error when refreshing kafka topic state: %s", GetResponseBody(httpResponse))
		}
		log.Println("[DEBUG] Kafka topic status: " + topic.Status)
		return topic, topic.Status, nil
	}
}

func resourceKafkaTopicDeleteStateRefreshFunc(cli *client.Client, clusterID, topicID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		topic, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetTopicById(context.TODO(), clusterID, topicID)
		if httpResponse != nil && httpResponse.StatusCode == http.StatusBadRequest {
			return vdbv2.TopicDto{Status: "DELETED"}, "DELETED", nil
		}
		if CheckErrorResponse(httpResponse) {
			return nil, "", fmt.Errorf("error when refreshing kafka topic state: %s", GetResponseBody(httpResponse))
		}
		return topic, topic.Status, nil
	}
}
