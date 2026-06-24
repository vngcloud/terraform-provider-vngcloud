package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
)

func ResourceKafkaCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceKafkaClusterCreate,
		Read:   resourceKafkaClusterRead,
		Update: resourceKafkaClusterUpdate,
		Delete: resourceKafkaClusterDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kafka_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"server_flavor_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"kafka_broker_count": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"kafka_storage_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"kafka_storage_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"vserver_project_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"mtls_authen": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"sasl_authen": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"public_access": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"config_group_version_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"encryption_volume": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Default:  false,
			},
			"auto_rebalance_topics": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Automatically rebalance topics across brokers when scaling out (kafka_broker_count increases). Ignored when scaling in.",
			},
			"security_group_rules": {
				Type:     schema.TypeSet,
				Optional: true,
				//Computed: true,
				Set: kafkaSecurityGroupRuleHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_ip": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"status": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_message": {
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
			"fixed_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floating_ips": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"kafka_storage_usage": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"iops": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"volume_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_type_zone_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"instance_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// kafkaClusterCreateMu serialize hóa pha "accept" của request create kafka cluster.
// Terraform chạy các callback Create song song; nếu nhiều resource kafka cùng gọi
// CreateOrderCluster một lúc sẽ trúng lock theo user ở phía server và bị lỗi. Mutex
// này đảm bảo mỗi lần chỉ một request accept được gửi đi; phần chờ cluster ACTIVE
// (kéo dài) chạy ngoài mutex nên các cluster vẫn provision song song.
var kafkaClusterCreateMu sync.Mutex

func kafkaSecurityGroupRuleHash(v interface{}) int {
	m := v.(map[string]interface{})
	return schema.HashString(fmt.Sprintf("%s|%d", m["remote_ip"].(string), m["port"].(int)))
}

func resourceKafkaClusterCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka cluster create")
	cli := m.(*client.Client)

	createRequest := vdbv2.CreateKafkaClusterRequest{
		Name:                 d.Get("name").(string),
		KafkaVersion:         d.Get("kafka_version").(string),
		ServerFlavorId:       d.Get("server_flavor_id").(string),
		KafkaBrokerCount:     int32(d.Get("kafka_broker_count").(int)),
		KafkaStorageType:     d.Get("kafka_storage_type").(string),
		KafkaStorageSize:     int32(d.Get("kafka_storage_size").(int)),
		VserverProjectId:     d.Get("vserver_project_id").(string),
		NetworkId:            d.Get("network_id").(string),
		SubnetId:             d.Get("subnet_id").(string),
		MtlsAuthen:           d.Get("mtls_authen").(bool),
		SaslAuthen:           d.Get("sasl_authen").(bool),
		ConfigGroupVersionId: d.Get("config_group_version_id").(string),
		EncryptionVolume:     d.Get("encryption_volume").(bool),
		Tags:                 map[string]string{},
	}

	// Serialize hóa pha accept: chỉ một request CreateOrderCluster được gửi tại một
	// thời điểm, và giữ thêm kafkaClusterCreateLockHold để chờ server nhả lock theo user
	// trước khi resource kafka kế tiếp gửi request. Phần wait ACTIVE bên dưới nằm ngoài
	// mutex nên không chặn các cluster khác.
	kafkaClusterCreateMu.Lock()
	resp, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.CreateOrderCluster(context.TODO(), createRequest, nil)
	if CheckErrorResponse(httpResponse) {
		kafkaClusterCreateMu.Unlock()
		return fmt.Errorf("create kafka cluster: %s", GetResponseBody(httpResponse))
	}
	time.Sleep(kafkaClusterCreateLockHold)
	kafkaClusterCreateMu.Unlock()
	if len(resp.Data) == 0 || resp.Data[0].ResourceId == "" {
		return fmt.Errorf("create kafka cluster: empty resource id in response")
	}

	clusterID := resp.Data[0].ResourceId
	log.Println("[DEBUG] Created kafka cluster id: " + clusterID)
	d.SetId(clusterID)
	time.Sleep(10 * time.Second)

	if err := waitKafkaClusterState(cli, clusterID, kafkaClusterCreatePending, kafkaClusterCreateTarget, kafkaClusterCreateTimeout, kafkaClusterCreateDelay, kafkaClusterCreateMinTimeout); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be created: %s", err)
	}

	if d.Get("public_access").(bool) {
		if err := updateKafkaClusterPublicAccess(cli, clusterID, true); err != nil {
			return err
		}
	}

	rules := d.Get("security_group_rules").(*schema.Set).List()
	if len(rules) > 0 {
		if err := addKafkaClusterSecRules(cli, clusterID, rules); err != nil {
			return err
		}
	}

	return resourceKafkaClusterRead(d, m)
}

func resourceKafkaClusterRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka cluster read")
	cli := m.(*client.Client)

	cluster, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetClusterById(context.TODO(), d.Id())
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("get kafka cluster: %s", GetResponseBody(httpResponse))
	}

	d.Set("name", cluster.Name)
	d.Set("kafka_version", cluster.KafkaVersion)
	d.Set("server_flavor_id", cluster.ServerFlavorId)
	d.Set("kafka_broker_count", cluster.KafkaBrokerCount)
	d.Set("kafka_storage_type", cluster.KafkaStorageType)
	d.Set("kafka_storage_size", cluster.KafkaStorageSize)
	d.Set("vserver_project_id", cluster.VserverProjectId)
	d.Set("network_id", cluster.NetworkId)
	d.Set("subnet_id", cluster.SubnetId)
	d.Set("mtls_authen", cluster.MtlsAuthen)
	d.Set("sasl_authen", cluster.SaslAuthen)
	d.Set("public_access", cluster.PublicAccess)
	d.Set("config_group_version_id", cluster.ConfigGroupVersionId)
	d.Set("encryption_volume", cluster.EncryptionVolume)
	d.Set("status", cluster.Status)
	d.Set("error_message", cluster.ErrorMessage)
	d.Set("portal_user_id", cluster.PortalUserId)
	d.Set("created_at", cluster.CreatedAt)
	d.Set("fixed_ips", cluster.FixedIps)
	d.Set("floating_ips", cluster.FloatingIps)
	d.Set("kafka_storage_usage", flattenKafkaStorageUsage(cluster.KafkaStorageUsage))
	d.Set("iops", cluster.Iops)
	d.Set("volume_type", cluster.VolumeType)
	d.Set("volume_type_zone_id", cluster.VolumeTypeZoneId)
	d.Set("ram", cluster.Ram)
	d.Set("vcpus", cluster.Vcpus)
	d.Set("instance_type", cluster.InstanceType)
	d.Set("security_group_rules", flattenKafkaSecurityGroupRules(cluster.SecurityGroupRules))

	return nil
}

func resourceKafkaClusterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka cluster update")
	cli := m.(*client.Client)
	clusterID := d.Id()

	if d.HasChange("mtls_authen") || d.HasChange("sasl_authen") {
		if err := updateKafkaClusterAuthentication(cli, clusterID, d.Get("mtls_authen").(bool), d.Get("sasl_authen").(bool)); err != nil {
			return err
		}
	}
	if d.HasChange("public_access") {
		if err := updateKafkaClusterPublicAccess(cli, clusterID, d.Get("public_access").(bool)); err != nil {
			return err
		}
	}
	if d.HasChange("config_group_version_id") {
		if err := updateKafkaClusterConfigGroup(cli, clusterID, d.Get("config_group_version_id").(string)); err != nil {
			return err
		}
	}
	if d.HasChange("kafka_storage_type") {
		if err := updateKafkaClusterStorageType(cli, clusterID, d.Get("kafka_storage_type").(string)); err != nil {
			return err
		}
	}
	if d.HasChange("kafka_storage_size") {
		if err := updateKafkaClusterStorageSize(cli, clusterID, int32(d.Get("kafka_storage_size").(int))); err != nil {
			return err
		}
	}
	if d.HasChange("kafka_broker_count") {
		oldVal, newVal := d.GetChange("kafka_broker_count")
		rebalance := false
		if newVal.(int) > oldVal.(int) {
			rebalance = d.Get("auto_rebalance_topics").(bool)
		}
		if err := updateKafkaClusterBrokerCount(cli, clusterID, int32(newVal.(int)), rebalance); err != nil {
			return err
		}
	}
	if d.HasChange("security_group_rules") {
		oldRaw, newRaw := d.GetChange("security_group_rules")
		if err := updateKafkaClusterSecRules(cli, clusterID, oldRaw.(*schema.Set), newRaw.(*schema.Set)); err != nil {
			return err
		}
	}

	return resourceKafkaClusterRead(d, m)
}

func resourceKafkaClusterDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Kafka cluster delete")
	cli := m.(*client.Client)

	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.DeleteCluster(context.TODO(), d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("delete kafka cluster: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    kafkaClusterDeletePending,
		Target:     kafkaClusterDeleteTarget,
		Refresh:    resourceKafkaClusterDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    kafkaClusterDeleteTimeout,
		Delay:      kafkaClusterDeleteDelay,
		MinTimeout: kafkaClusterDeleteMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster to be deleted: %s", err)
	}

	d.SetId("")
	return nil
}

func resourceKafkaClusterStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		cluster, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetClusterById(context.TODO(), id)
		if CheckErrorResponse(httpResponse) {
			return nil, "", fmt.Errorf("error when refreshing kafka cluster state: %s", GetResponseBody(httpResponse))
		}
		resp, _ := json.Marshal(cluster)
		log.Println("[DEBUG] Kafka cluster status: " + string(resp))
		return cluster, cluster.Status, nil
	}
}

func resourceKafkaClusterDeleteStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		cluster, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.GetClusterById(context.TODO(), id)
		if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
			return vdbv2.KafkaCluster{Status: "DELETED"}, "DELETED", nil
		}
		if CheckErrorResponse(httpResponse) {
			return nil, "", fmt.Errorf("error when refreshing kafka cluster state: %s", GetResponseBody(httpResponse))
		}
		return cluster, cluster.Status, nil
	}
}

func waitKafkaClusterState(cli *client.Client, id string, pending, target []string, timeout, delay, minTimeout time.Duration) error {
	log.Println("[DEBUG] wait kafka cluster state")
	stateConf := &resource.StateChangeConf{
		Pending:    pending,
		Target:     target,
		Refresh:    resourceKafkaClusterStateRefreshFunc(cli, id),
		Timeout:    timeout,
		Delay:      delay,
		MinTimeout: minTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for kafka cluster state: %s", err)
	}
	return nil
}

func updateKafkaClusterAuthentication(cli *client.Client, id string, mtls, sasl bool) error {
	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateAuthentication(context.TODO(), id, mtls, sasl)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka authentication: %s", GetResponseBody(httpResponse))
	}
	return waitKafkaClusterState(cli, id, kafkaClusterUpdateAuthPending, kafkaClusterUpdateAuthTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout)
}

func updateKafkaClusterPublicAccess(cli *client.Client, id string, enable bool) error {
	enableStr := "false"
	if enable {
		enableStr = "true"
	}
	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdatePublicAccess(context.TODO(), id, enableStr)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka public access: %s", GetResponseBody(httpResponse))
	}
	return waitKafkaClusterState(cli, id, kafkaClusterUpdatePublicAccessPending, kafkaClusterUpdatePublicAccessTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout)
}

func updateKafkaClusterConfigGroup(cli *client.Client, id string, configGroupVersionId string) error {
	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateConfigGroup(context.TODO(), id, configGroupVersionId)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka config group: %s", GetResponseBody(httpResponse))
	}
	return waitKafkaClusterState(cli, id, kafkaClusterUpdateConfigGroupPending, kafkaClusterUpdateConfigGroupTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout)
}

func updateKafkaClusterStorageType(cli *client.Client, id, storageType string) error {
	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateStorageType(context.TODO(), id, storageType, nil)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka storage type: %s", GetResponseBody(httpResponse))
	}
	return waitKafkaClusterState(cli, id, kafkaClusterUpdateStorageTypePending, kafkaClusterUpdateStorageTypeTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout)
}

func updateKafkaClusterStorageSize(cli *client.Client, id string, size int32) error {
	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateStorageSize(context.TODO(), id, size, nil)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka storage size: %s", GetResponseBody(httpResponse))
	}
	return waitKafkaClusterState(cli, id, kafkaClusterUpdateStorageSizePending, kafkaClusterUpdateStorageSizeTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout)
}

func updateKafkaClusterBrokerCount(cli *client.Client, id string, count int32, rebalance bool) error {
	_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.UpdateBrokerCount(context.TODO(), id, count, rebalance, nil)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update kafka broker count: %s", GetResponseBody(httpResponse))
	}
	return waitKafkaClusterState(cli, id, kafkaClusterUpdateBrokerCountPending, kafkaClusterUpdateBrokerCountTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout)
}

func addKafkaClusterSecRules(cli *client.Client, id string, rules []interface{}) error {
	for _, r := range rules {
		rm := r.(map[string]interface{})
		body := vdbv2.SecurityGroupRuleCreateRequest{
			RemoteIp: rm["remote_ip"].(string),
			Port:     int32(rm["port"].(int)),
		}
		_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.CreateSecRule(context.TODO(), body, id)
		if CheckErrorResponse(httpResponse) {
			return fmt.Errorf("create kafka security group rule: %s", GetResponseBody(httpResponse))
		}
		if err := waitKafkaClusterState(cli, id, kafkaClusterUpdateSecRulesPending, kafkaClusterUpdateSecRulesTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout); err != nil {
			return err
		}
	}
	return nil
}

func updateKafkaClusterSecRules(cli *client.Client, id string, oldSet, newSet *schema.Set) error {
	toRemove := oldSet.Difference(newSet).List()
	toAdd := newSet.Difference(oldSet).List()

	oldByKey := make(map[string]string, oldSet.Len())
	for _, r := range oldSet.List() {
		rm := r.(map[string]interface{})
		key := fmt.Sprintf("%s|%d", rm["remote_ip"].(string), rm["port"].(int))
		if ruleID, ok := rm["id"].(string); ok && ruleID != "" {
			oldByKey[key] = ruleID
		}
	}

	for _, r := range toRemove {
		rm := r.(map[string]interface{})
		ruleID, _ := rm["id"].(string)
		if ruleID == "" {
			ruleID = oldByKey[fmt.Sprintf("%s|%d", rm["remote_ip"].(string), rm["port"].(int))]
		}
		if ruleID == "" {
			continue
		}
		_, httpResponse, _ := cli.Vdbv2Client.KafkaClusterAPIApi.DeleteSecRule(context.TODO(), id, ruleID)
		if CheckErrorResponse(httpResponse) {
			return fmt.Errorf("delete kafka security group rule: %s", GetResponseBody(httpResponse))
		}
		if err := waitKafkaClusterState(cli, id, kafkaClusterUpdateSecRulesPending, kafkaClusterUpdateSecRulesTarget, kafkaClusterUpdateTimeout, kafkaClusterUpdateDelay, kafkaClusterUpdateMinTimeout); err != nil {
			return err
		}
	}

	return addKafkaClusterSecRules(cli, id, toAdd)
}

func flattenKafkaStorageUsage(values []int64) []interface{} {
	res := make([]interface{}, len(values))
	for i, v := range values {
		res[i] = int(v)
	}
	return res
}

func flattenKafkaSecurityGroupRules(rules []vdbv2.KafkaSecurityGroupRule) []interface{} {
	res := make([]interface{}, len(rules))
	for i, r := range rules {
		res[i] = map[string]interface{}{
			"id":        r.Id,
			"remote_ip": r.RemoteIp,
			"port":      int(r.Port),
			"status":    r.Status,
		}
	}
	return res
}
