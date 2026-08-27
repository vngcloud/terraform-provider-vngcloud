package vdbv2

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/vngcloud/terraform-provider-vngcloud/client"
	"github.com/vngcloud/terraform-provider-vngcloud/client/vdbv2"
)

func ResourcePostgreSQLCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourcePostgreSQLClusterCreate,
		Read:   resourcePostgreSQLClusterRead,
		Update: resourcePostgreSQLClusterUpdate,
		Delete: resourcePostgreSQLClusterDelete,
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
			"engine_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_version": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"subnet_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"zone_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"package_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"volume_type_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"number_of_nodes": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"public_access": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"password": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"db_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_poc": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"backup_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"backup_policy_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"backup_location_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"action": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"port_ro": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"private_rw_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_rw_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_ro_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_ro_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bandwidth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"free_backup_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deploy_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"volume_used": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"secgroup_rule": {
				Type:     schema.TypeSet,
				Optional: true,
				// Computed: true,
				Set: postgresSecgroupRuleHash,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"remote_ip_prefix": {
							Type:     schema.TypeString,
							Required: true,
						},
						"port": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntInSlice([]int{5432, 15432}),
						},
					},
				},
			},
		},
	}
}

func postgresSecgroupRuleHash(v interface{}) int {
	m := v.(map[string]interface{})
	return schema.HashString(fmt.Sprintf("%s|%d", m["remote_ip_prefix"].(string), m["port"].(int)))
}

func resourcePostgreSQLClusterStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Println("[DEBUG]  Postgres cluster state refresh")

		dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), id)
		if CheckErrorResponse(httpResponse) {
			responseBody := GetResponseBody(httpResponse)
			return nil, "", fmt.Errorf("error when refreshing postgres cluster state: %s", responseBody)
		}
		if dbResp.Data == nil {
			return nil, "", fmt.Errorf("error when refreshing postgres cluster state: data is empty")
		}
		log.Println("[DEBUG]  Postgres cluster status: " + dbResp.Data.Status)

		return dbResp.Data.Id, dbResp.Data.Status, nil
	}
}

func resourcePostgreSQLClusterDeleteStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), id)
		if httpResponse.StatusCode != http.StatusOK {
			if httpResponse.StatusCode == http.StatusNotFound {
				return vdbv2.DbInstanceInfo{Status: "DELETED"}, "DELETED", nil
			}
			return nil, "", fmt.Errorf("error refreshing postgres cluster delete state: %s", GetResponseBody(httpResponse))
		}
		if dbResp.Data == nil {
			return vdbv2.DbInstanceInfo{Status: "DELETED"}, "DELETED", nil
		}
		return dbResp.Data, dbResp.Data.Status, nil
	}
}

func resourcePostgreSQLClusterCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Postgres cluster create")

	cli := m.(*client.Client)

	createRequest := vdbv2.CreatePostgreClusterRequest{
		Name:             d.Get("name").(string),
		LocateZoneId:     d.Get("zone_id").(string),
		PackageId:        d.Get("package_id").(string),
		VolumeTypeId:     d.Get("volume_type_id").(string),
		VolumeSize:       int32(d.Get("volume_size").(int)),
		NumberOfNodes:    int32(d.Get("number_of_nodes").(int)),
		DatastoreVersion: d.Get("engine_version").(string),
		NetIds:           []string{d.Get("subnet_id").(string)},
		ConfigId:         d.Get("config_id").(string),
		PublicAccess:     d.Get("public_access").(bool),
		IsPoc:            d.Get("is_poc").(bool),
		BackupPolicyId:   d.Get("backup_policy_id").(string),
		BackupLocationId: d.Get("backup_location_id").(string),
	}

	if d.Get("backup_id").(string) == "" {
		createRequest.User = &vdbv2.UserRequest2{
			Name:     d.Get("username").(string),
			Password: d.Get("password").(string),
		}
		createRequest.Databases = []vdbv2.DatabaseRequest2{{Name: d.Get("db_name").(string)}}
	} else {
		createRequest.BackupPointId = d.Get("backup_id").(string)
	}

	reqBody, _ := json.Marshal(createRequest)
	log.Println("[DEBUG]  Body: " + string(reqBody))

	resp, httpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.CreateOrderPostgreCluster(context.TODO(), createRequest, nil)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("create postgres cluster: %s", GetResponseBody(httpResponse))
	}
	if len(resp.Data) == 0 || resp.Data[0].ResourceId == "" {
		return fmt.Errorf("create postgres cluster: empty resource id in response")
	}

	clusterID := resp.Data[0].ResourceId
	log.Println("[DEBUG]  Created postgres cluster id: " + clusterID)

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterCreatePending,
		Target:     postgresClusterCreateTarget,
		Refresh:    resourcePostgreSQLClusterStateRefreshFunc(cli, clusterID),
		Timeout:    postgresClusterCreateTimeout,
		Delay:      postgresClusterCreateDelay,
		MinTimeout: postgresClusterCreateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for postgres cluster to be created: %s", err)
	}

	d.SetId(clusterID)

	if rules := d.Get("secgroup_rule").(*schema.Set).List(); len(rules) > 0 {
		if err := updatePostgreSQLClusterSecgroupRules(d, m, rules); err != nil {
			return err
		}
	}

	return resourcePostgreSQLClusterRead(d, m)
}

func resourcePostgreSQLClusterRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Postgres cluster read")

	cli := m.(*client.Client)

	dbResp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetDatabaseInstancesById(context.TODO(), d.Id())
	if httpResponse != nil && httpResponse.StatusCode == http.StatusNotFound {
		d.SetId("")
		return nil
	}
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("get postgres cluster: %s", GetResponseBody(httpResponse))
	}
	if dbResp.Data == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", dbResp.Data.Name)
	d.Set("engine_type", dbResp.Data.DatastoreType)
	d.Set("engine_version", dbResp.Data.DatastoreVersion)
	d.Set("subnet_id", dbResp.Data.SubnetId)
	d.Set("zone_id", dbResp.Data.ZoneId)
	d.Set("package_id", dbResp.Data.QuotaPackageId)
	d.Set("volume_type_id", dbResp.Data.VolumeTypeId)
	d.Set("volume_type", dbResp.Data.VolumeType)
	d.Set("volume_size", dbResp.Data.VolumeSize)
	d.Set("number_of_nodes", dbResp.Data.NumberOfNodes)
	d.Set("public_access", dbResp.Data.PublicAccess)
	d.Set("status", dbResp.Data.Status)
	d.Set("ram", dbResp.Data.Ram)
	d.Set("cpu", dbResp.Data.Vcpus)
	d.Set("port", dbResp.Data.Port)
	d.Set("port_ro", dbResp.Data.PortRo)
	d.Set("private_rw_ip", dbResp.Data.PrivateRwIp)
	d.Set("public_rw_ip", dbResp.Data.PublicRwIp)
	d.Set("private_ro_ip", dbResp.Data.PrivateRoIp)
	d.Set("public_ro_ip", dbResp.Data.PublicRoIp)
	d.Set("domain_name", dbResp.Data.DomainName)
	d.Set("bandwidth", dbResp.Data.Bandwidth)
	d.Set("free_backup_size", dbResp.Data.FreeBackupSize)
	d.Set("project_id", dbResp.Data.ProjectId)
	d.Set("deploy_type", dbResp.Data.DeployType)

	if dbResp.Data.Configuration != nil {
		d.Set("config_id", dbResp.Data.Configuration.Id)
		d.Set("config_name", dbResp.Data.Configuration.Name)
	}

	volResp, volHttpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.GetPostgreClusterVolumeUsed(context.TODO(), d.Id())
	if volHttpResponse != nil && !CheckErrorResponse(volHttpResponse) {
		d.Set("volume_used", volResp.Data)
	}

	backupResp, backupHttpResponse, _ := cli.Vdbv2Client.PostgreSQLBackupAPIApi.GetBackupVDB(context.TODO(), d.Id())
	if backupHttpResponse != nil && !CheckErrorResponse(backupHttpResponse) && backupResp.Data != nil {
		d.Set("backup_policy_id", backupResp.Data.BackupPolicyId)
		d.Set("backup_location_id", backupResp.Data.BackupDestinationId)
	}

	return readPostgreSQLClusterSecgroupRules(d, m)
}

func resourcePostgreSQLClusterUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Postgres cluster update")

	if d.HasChange("secgroup_rule") {
		if err := updatePostgreSQLClusterSecgroupRules(d, m, d.Get("secgroup_rule").(*schema.Set).List()); err != nil {
			return err
		}
	}

	if d.HasChange("config_id") {
		if err := resourcePostgreSQLClusterUpdateConfigGroup(d, m); err != nil {
			return err
		}
	}

	if d.HasChange("action") {
		if d.Get("action").(string) == "reboot" {
			if err := resourcePostgreSQLClusterReboot(d, m); err != nil {
				return err
			}
		}
	}

	if d.HasChange("password") || d.HasChange("public_access") {
		if err := resourcePostgreSQLClusterUpdateSettings(d, m); err != nil {
			return err
		}
	}

	if d.HasChange("number_of_nodes") {
		if err := resourcePostgreSQLClusterResize(d, m, "NUMBER-OF-NODES"); err != nil {
			return err
		}
	}

	if d.HasChange("volume_type_id") {
		if err := resourcePostgreSQLClusterResize(d, m, "VOLUME-TYPE"); err != nil {
			return err
		}
	}

	if d.HasChange("volume_size") {
		if err := resourcePostgreSQLClusterResize(d, m, "VOLUME-SIZE"); err != nil {
			return err
		}
	}

	return resourcePostgreSQLClusterRead(d, m)
}

func resourcePostgreSQLClusterResize(d *schema.ResourceData, m interface{}, resizeType string) error {
	cli := m.(*client.Client)

	body := vdbv2.ResizePostgreClusterRequest{
		Type_: resizeType,
	}
	switch resizeType {
	case "NUMBER-OF-NODES":
		body.NumberOfNodes = int32(d.Get("number_of_nodes").(int))
	case "VOLUME-TYPE":
		body.VolumeTypeId = d.Get("volume_type_id").(string)
	case "VOLUME-SIZE":
		body.VolumeSize = int32(d.Get("volume_size").(int))
	}

	reqBody, _ := json.Marshal(body)
	log.Println("[DEBUG]  Resize postgres cluster body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.ResizePostgreCluster(context.TODO(), body, d.Id(), nil)
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("resize postgres cluster (%s): %s", resizeType, GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterUpdatePending,
		Target:     postgresClusterUpdateTarget,
		Refresh:    resourcePostgreSQLClusterStateRefreshFunc(cli, d.Id()),
		Timeout:    postgresClusterUpdateTimeout,
		Delay:      postgresClusterUpdateDelay,
		MinTimeout: postgresClusterUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for postgres cluster resize (%s): %s", resizeType, err)
	}

	return nil
}

func resourcePostgreSQLClusterUpdateConfigGroup(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	body := vdbv2.UpdatePostgreClusterConfigGroupRequest{
		ConfigGroupId: d.Get("config_id").(string),
	}

	reqBody, _ := json.Marshal(body)
	log.Println("[DEBUG]  Update postgres cluster config group body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.UpdateConfigGroupPostgreCluster(context.TODO(), body, d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update postgres cluster config group: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterUpdateConfigGroupPending,
		Target:     postgresClusterUpdateConfigGroupTarget,
		Refresh:    resourcePostgreSQLClusterStateRefreshFunc(cli, d.Id()),
		Timeout:    postgresClusterUpdateTimeout,
		Delay:      postgresClusterUpdateDelay,
		MinTimeout: postgresClusterUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for postgres cluster config group update: %s", err)
	}

	return nil
}

func resourcePostgreSQLClusterUpdateSettings(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	body := vdbv2.UpdatePostgreClusterSettingsRequest{
		Password:     d.Get("password").(string),
		PublicAccess: d.Get("public_access").(bool),
	}

	reqBody, _ := json.Marshal(body)
	log.Println("[DEBUG]  Update postgres cluster settings body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.PostgreSQLClusterAPIApi.UpdateSettingsPostgreCluster(context.TODO(), body, d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update postgres cluster settings: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterUpdatePending,
		Target:     postgresClusterUpdateTarget,
		Refresh:    resourcePostgreSQLClusterStateRefreshFunc(cli, d.Id()),
		Timeout:    postgresClusterUpdateTimeout,
		Delay:      postgresClusterUpdateDelay,
		MinTimeout: postgresClusterUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for postgres cluster settings update: %s", err)
	}

	return nil
}

func resourcePostgreSQLClusterReboot(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "reboot")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG]  Reboot postgres cluster body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.RestartDatabaseInstances(context.TODO(), string(reqBody), d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("reboot postgres cluster: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterRebootPending,
		Target:     postgresClusterRebootTarget,
		Refresh:    resourcePostgreSQLClusterStateRefreshFunc(cli, d.Id()),
		Timeout:    postgresClusterUpdateTimeout,
		Delay:      postgresClusterUpdateDelay,
		MinTimeout: postgresClusterUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for postgres cluster reboot: %s", err)
	}

	return nil
}

func resourcePostgreSQLClusterDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  Postgres cluster delete")

	cli := m.(*client.Client)

	actionRequest := generateActionRequest(d, "delete")
	reqBody, _ := json.Marshal(actionRequest)
	log.Println("[DEBUG]  Delete postgres cluster body: " + string(reqBody))

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.DeleteDatabaseInstances(context.TODO(), string(reqBody), d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("delete postgres cluster: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterDeletePending,
		Target:     postgresClusterDeleteTarget,
		Refresh:    resourcePostgreSQLClusterDeleteStateRefreshFunc(cli, d.Id()),
		Timeout:    postgresClusterDeleteTimeout,
		Delay:      postgresClusterDeleteDelay,
		MinTimeout: postgresClusterDeleteMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error when waiting for postgres cluster delete: %s", err)
	}

	d.SetId("")
	return nil
}

func updatePostgreSQLClusterSecgroupRules(d *schema.ResourceData, m interface{}, ruleList []interface{}) error {
	cli := m.(*client.Client)

	rules := make([]vdbv2.SecurityGroupRuleEntity, 0, len(ruleList))
	for _, raw := range ruleList {
		ruleMap := raw.(map[string]interface{})
		port := int32(ruleMap["port"].(int))
		rules = append(rules, vdbv2.SecurityGroupRuleEntity{
			RemoteIpPrefix: ruleMap["remote_ip_prefix"].(string),
			PortRangeMin:   port,
			PortRangeMax:   port,
		})
	}

	_, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.UpdateSecurityGroupRules(context.TODO(), rules, d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("update postgres cluster security group rules: %s", GetResponseBody(httpResponse))
	}

	stateConf := &resource.StateChangeConf{
		Pending:    postgresClusterUpdatePending,
		Target:     postgresClusterUpdateTarget,
		Refresh:    resourcePostgreSQLClusterStateRefreshFunc(cli, d.Id()),
		Timeout:    postgresClusterUpdateTimeout,
		Delay:      postgresClusterUpdateDelay,
		MinTimeout: postgresClusterUpdateMinTimeout,
	}
	if _, err := stateConf.WaitForState(); err != nil {
		return fmt.Errorf("error waiting for postgres cluster security group rule update: %s", err)
	}

	return readPostgreSQLClusterSecgroupRules(d, m)
}

func readPostgreSQLClusterSecgroupRules(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	resp, httpResponse, _ := cli.Vdbv2Client.RelationalDatabaseAPIApi.GetSecurityRules(context.TODO(), d.Id())
	if CheckErrorResponse(httpResponse) {
		return fmt.Errorf("get postgres cluster security group rules: %s", GetResponseBody(httpResponse))
	}

	if resp.Data == nil {
		d.Set("secgroup_rule", []interface{}{})
		return nil
	}

	flat := make([]interface{}, 0, len(resp.Data))
	for _, r := range resp.Data {
		flat = append(flat, map[string]interface{}{
			"id":               r.Id,
			"remote_ip_prefix": r.RemoteIpPrefix,
			"port":             int(r.PortRangeMin),
		})
	}
	d.Set("secgroup_rule", flat)
	return nil
}
