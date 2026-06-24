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

func ResourceKafkaUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaUserCreate,
		Read:   resourceKafkaUserRead,
		Update: resourceKafkaUserUpdate,
		Delete: resourceKafkaUserDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				parts := strings.Split(d.Id(), "/")
				if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
					return nil, fmt.Errorf("expected import id in format <cluster_id>/<user_id>")
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
			"mtls_authen": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"sasl_authen": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"produce_topic_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"produce_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"consume_topic_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"consume_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"produce_consume_topic_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"produce_consume_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"admin_topic_names": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"admin_all": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"old_status": {
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

func resourceKafkaUserCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka user create")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)
	defer lockKafkaCluster(clusterID)()

	body := vdbv2.UserCreateRequest{
		Name:                     d.Get("name").(string),
		MtlsAuthen:               d.Get("mtls_authen").(bool),
		SaslAuthen:               d.Get("sasl_authen").(bool),
		ProduceTopicNames:        expandStringList(d.Get("produce_topic_names").([]interface{})),
		ProduceAll:               d.Get("produce_all").(bool),
		ConsumeTopicNames:        expandStringList(d.Get("consume_topic_names").([]interface{})),
		ConsumeAll:               d.Get("consume_all").(bool),
		ProduceConsumeTopicNames: expandStringList(d.Get("produce_consume_topic_names").([]interface{})),
		ProduceConsumeAll:        d.Get("produce_consume_all").(bool),
		AdminTopicNames:          expandStringList(d.Get("admin_topic_names").([]interface{})),
		AdminAll:                 d.Get("admin_all").(bool),
	}

	resp, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.CreateUser(context.TODO(), body, clusterID)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("create kafka user: %s", GetResponseBody(httpResponse))
	}
	if resp.Id == "" {
		return fmt.Errorf("create kafka user: empty user id in response")
	}
	d.SetId(resp.Id)
	time.Sleep(10 * time.Second)

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaUserCreatePending,
		Target:     kafkaUserCreateTarget,
		Refresh:    resourceKafkaUserStateRefreshFunc(cli, clusterID, resp.Id),
		Timeout:    kafkaUserCreateTimeout,
		Delay:      kafkaUserCreateDelay,
		MinTimeout: kafkaUserCreateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka user to be created: %s", err)
	}

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterTopicUserPending, kafkaClusterTopicUserTarget, kafkaClusterTopicUserTimeout, kafkaClusterTopicUserDelay, kafkaClusterTopicUserMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be ACTIVE: %s", err)
	}

	return resourceKafkaUserRead(d, m)
}

func resourceKafkaUserRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka user read")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)

	user, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetUserById(context.TODO(), clusterID, d.Id())
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("get kafka user: %s", GetResponseBody(httpResponse))
	}

	d.Set("cluster_id", user.ClusterId)
	d.Set("name", user.Name)
	d.Set("mtls_authen", user.MtlsAuthen)
	d.Set("sasl_authen", user.SaslAuthen)
	d.Set("produce_topic_names", user.ProduceTopicNames)
	d.Set("produce_all", user.ProduceAll)
	d.Set("consume_topic_names", user.ConsumeTopicNames)
	d.Set("consume_all", user.ConsumeAll)
	d.Set("produce_consume_topic_names", user.ProduceConsumeTopicNames)
	d.Set("produce_consume_all", user.ProduceConsumeAll)
	d.Set("admin_topic_names", user.AdminTopicNames)
	d.Set("admin_all", user.AdminAll)
	d.Set("status", user.Status)
	d.Set("old_status", user.OldStatus)
	d.Set("portal_user_id", user.PortalUserId)
	d.Set("created_at", user.CreatedAt)
	return nil
}

func resourceKafkaUserUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka user update")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)

	permissionFields := []string{
		"produce_topic_names", "produce_all",
		"consume_topic_names", "consume_all",
		"produce_consume_topic_names", "produce_consume_all",
		"admin_topic_names", "admin_all",
	}
	hasPermissionChange := false
	for _, f := range permissionFields {
		if d.HasChange(f) {
			hasPermissionChange = true
			break
		}
	}
	if !hasPermissionChange {
		return resourceKafkaUserRead(d, m)
	}
	defer lockKafkaCluster(clusterID)()

	body := vdbv2.UserUpdatePermissionsRequest{
		MtlsAuthen:               d.Get("mtls_authen").(bool),
		SaslAuthen:               d.Get("sasl_authen").(bool),
		ProduceTopicNames:        expandStringList(d.Get("produce_topic_names").([]interface{})),
		ProduceAll:               d.Get("produce_all").(bool),
		ConsumeTopicNames:        expandStringList(d.Get("consume_topic_names").([]interface{})),
		ConsumeAll:               d.Get("consume_all").(bool),
		ProduceConsumeTopicNames: expandStringList(d.Get("produce_consume_topic_names").([]interface{})),
		ProduceConsumeAll:        d.Get("produce_consume_all").(bool),
		AdminTopicNames:          expandStringList(d.Get("admin_topic_names").([]interface{})),
		AdminAll:                 d.Get("admin_all").(bool),
	}

	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateUser(context.TODO(), body, clusterID, d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka user: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaUserUpdatePending,
		Target:     kafkaUserUpdateTarget,
		Refresh:    resourceKafkaUserStateRefreshFunc(cli, clusterID, d.Id()),
		Timeout:    kafkaUserUpdateTimeout,
		Delay:      kafkaUserUpdateDelay,
		MinTimeout: kafkaUserUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka user to be updated: %s", err)
	}

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterTopicUserPending, kafkaClusterTopicUserTarget, kafkaClusterTopicUserTimeout, kafkaClusterTopicUserDelay, kafkaClusterTopicUserMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be ACTIVE: %s", err)
	}

	return resourceKafkaUserRead(d, m)
}

func resourceKafkaUserDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka user delete")
	cli := m.(*client.Client)
	clusterID := d.Get("cluster_id").(string)
	defer lockKafkaCluster(clusterID)()

	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.DeleteUser(context.TODO(), clusterID, d.Id())
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("delete kafka user: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaUserDeletePending,
		Target:     kafkaUserDeleteTarget,
		Refresh:    resourceKafkaUserDeleteStateRefreshFunc(cli, clusterID, d.Id()),
		Timeout:    kafkaUserDeleteTimeout,
		Delay:      kafkaUserDeleteDelay,
		MinTimeout: kafkaUserDeleteMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka user to be deleted: %s", err)
	}

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterTopicUserPending, kafkaClusterTopicUserTarget, kafkaClusterTopicUserTimeout, kafkaClusterTopicUserDelay, kafkaClusterTopicUserMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be ACTIVE: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceKafkaUserStateRefreshFunc(cli *client.Client, clusterID, userID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		user, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetUserById(context.TODO(), clusterID, userID)
		if CheckErrorResponse(httpResponse) {
			return nil, "", fmt.Errorf("error when refreshing kafka user state: %s", GetResponseBody(httpResponse))
		}
		log.Println("[DEBUG] Kafka user status: " + user.Status)
		return user, user.Status, nil
	}
}

func resourceKafkaUserDeleteStateRefreshFunc(cli *client.Client, clusterID, userID string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		user, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetUserById(context.TODO(), clusterID, userID)
		if httpResponse != nil && httpResponse.StatusCode == http.StatusBadRequest {
			return vdbv2.UserDto{Status: "DELETED"}, "DELETED", nil
		}
		if CheckErrorResponse(httpResponse) {
			return nil, "", fmt.Errorf("error when refreshing kafka user state: %s", GetResponseBody(httpResponse))
		}
		return user, user.Status, nil
	}
}

func expandStringList(input []interface{}) []string {
	res := make([]string, 0, len(input))
	for _, v := range input {
		res = append(res, v.(string))
	}
	return res
}
