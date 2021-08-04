package vdb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/client/vdb"
)

func ResourceDatabase() *schema.Resource {
	return &schema.Resource{
		Create: resourceDatabaseCreate,
		Read:   resourceDatabaseRead,
		Update: resourceDatabaseUpdate,
		Delete: resourceDatabaseDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			"action": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_auto": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"backup_duration": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"backup_time": {
				Type:     schema.TypeString,
				Required: true,
			},
			"bandwidth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"config_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"database_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datastore_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"datastore_version": {
				Type:     schema.TypeString,
				Required: true,
			},
			"db_backend_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"engine_group": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"flavor_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"free_backup_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"network_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"package_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"price_key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"public_access": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"redis_password": {
				Type:     schema.TypeString,
				Required: true,
			},
			"redis_password_enabled": {
				Type:     schema.TypeBool,
				Required: true,
			},
			"replica_source_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"replicas": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"start_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"vcpus": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"volume_size": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"volume_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_type_zone_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"volume_used": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"zone_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"zone_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowed_ip_prefix": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"backup_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceDatabaseStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Println("[DEBUG] [trihm4] State refresh")

		dbResp, _, err := cli.VdbClient.VdbInstanceEndPointApi.GetDbInstanceDetailUsingGET(context.TODO(), id, cli.ProjectId)
		if err != nil {
			return nil, "", fmt.Errorf("Error when refreshing database state: %s", err)
		}
		if !dbResp.Success {
			return nil, "", fmt.Errorf("Error when refreshing database state: %s", dbResp.ErrorMsg)
		}
		log.Println("[DEBUG] [trihm4] Database status: " + dbResp.Data.Status)

		return dbResp.Data.Id, dbResp.Data.Status, nil
	}
}

func resourceDatabaseCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database create")

	cli := m.(*client.Client)

	var createDbResult vdb.DbInstanceResponse
	if d.Get("backup_id").(string) == "" {
		createRequest := generateCreateDatabaseRequest(cli.ProjectId, d)
		reqBody, _ := json.Marshal(createRequest)
		log.Println("[DEBUG] [trihm4] Body: " + string(reqBody))

		var err error
		createDbResult, _, err = cli.VdbClient.VdbInstanceEndPointApi.CreateDbInstanceUsingPOST(context.TODO(), createRequest, cli.ProjectId)
		if err != nil {
			return fmt.Errorf("Error when creating database: %s", err)
		}
		if !createDbResult.Success {
			return fmt.Errorf("Error when creating database: %s", createDbResult.ErrorMsg)
		}
		log.Println("[DEBUG] [trihm4] Create db result, id: " + createDbResult.Data.Id)
	} else {
		restoreRequest := generateRestoreBackupRequest(cli.ProjectId, d)
		reqBody, _ := json.Marshal(restoreRequest)
		log.Println("[DEBUG] [trihm4] Body: " + string(reqBody))

		createDbResult, _, _ = cli.VdbClient.VdbBackupEndPointApi.RestoreBackupUsingPOST(context.TODO(), cli.ProjectId, restoreRequest)
		log.Println("[DEBUG] [trihm4] Restore bk result, id: " + createDbResult.Data.Id)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseCreatePending,
		Target:     databaseCreateTarget,
		Refresh:    resourceDatabaseStateRefreshFunc(cli, createDbResult.Data.Id),
		Timeout:    databaseCreateTimeout,
		Delay:      databaseCreateDelay,
		MinTimeout: databaseCreateMinTimeout,
	}

	id, err := stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error when waiting for database to be created: %s", err)
	}
	idStr := id.(string)

	log.Println("[DEBUG] [trihm4] Wait for state done, id: " + idStr)

	d.SetId(idStr)

	allowedIP := d.Get("allowed_ip_prefix").([]interface{})
	log.Printf("[DEBUG] [nhannt10] allowed ip prefix: %v\n", allowedIP)

	err = resourceDatabaseRead(d, m)
	if err != nil {
		return fmt.Errorf("Error when getting database info after create: %s", err)
	}

	secgroupRulesUpdate(d, m, allowedIP)

	return nil
}

func generateCreateDatabaseRequest(projectId string, d *schema.ResourceData) vdb.CreateDbInstanceRequest {
	createRequest := vdb.CreateDbInstanceRequest{
		AutoRenewPeriod:      1,
		BackupAuto:           d.Get("backup_auto").(bool),
		BackupDuration:       int32(d.Get("backup_duration").(int)),
		BackupTime:           d.Get("backup_time").(string),
		CartItemId:           0,
		CartItemState:        0,
		ConfigId:             d.Get("config_id").(string),
		Cost:                 457600.0,
		Databases:            nil,
		DatastoreType:        d.Get("datastore_type").(string),
		DatastoreVersion:     d.Get("datastore_version").(string),
		EnableAutoRenew:      false,
		EndTime:              nil,
		EngineGroup:          int32(d.Get("engine_group").(int)),
		Extra:                nil,
		FlavorId:             d.Get("flavor_id").(string),
		Id:                   "",
		InvoiceId:            0,
		Name:                 d.Get("name").(string),
		NetIds:               []string{d.Get("network_id").(string)},
		PackageId:            d.Get("package_id").(string),
		Period:               1,
		Poc:                  false,
		PriceKey:             d.Get("price_key").(string),
		ProjectId:            projectId,
		PublicAccess:         d.Get("public_access").(bool),
		Ram:                  int32(d.Get("ram").(int)),
		RedisPassword:        d.Get("redis_password").(string),
		RedisPasswordEnabled: d.Get("redis_password_enabled").(bool),
		ReplicaSourceId:      d.Get("replica_source_id").(string),
		StartTime:            time.Now(),
		UseTrial:             false,
		User:                 nil,
		Vcpus:                int32(d.Get("vcpus").(int)),
		VolumeSize:           int32(d.Get("volume_size").(int)),
		VolumeType:           d.Get("volume_type").(string),
		VolumeTypeZoneId:     d.Get("volume_type_zone_id").(string),
		ZoneId:               int32(d.Get("zone_id").(int)),
	}

	if d.Get("replica_source_id").(string) == "" {
		createRequest.User = &vdb.UserRequest{
			Databases: []vdb.Database{{Name: d.Get("database_name").(string)}},
			Name:      d.Get("username").(string),
			Password:  d.Get("password").(string),
		}

		createRequest.Databases = []vdb.DatabaseRequest{{
			CharacterSet: "utf8",
			Collate:      "utf8_general_ci",
			Name:         d.Get("database_name").(string),
		}}
	}

	return createRequest
}

func secgroupRulesUpdate(d *schema.ResourceData, m interface{}, allowedIP []interface{}) error {
	if len(allowedIP) > 0 {
		cli := m.(*client.Client)

		instanceID := d.Id()
		projectID := cli.ProjectId
		rules := createSecurityGroupRules(&allowedIP, d.Get("port").(int))
		resp, _, err := cli.VdbClient.VdbInstanceEndPointApi.UpdateDbInstanceSecurityGroupRuleUsingPUT(context.TODO(), instanceID, projectID, rules)

		if err != nil {
			return err
		}

		if !resp.Success {
			return errors.New(resp.ErrorMsg)
		}
	}

	return secgroupRulesRead(d, m)
}

func createSecurityGroupRules(allowedIP *[]interface{}, port int) []vdb.SecurityGroupRuleInfo {
	rules := make([]vdb.SecurityGroupRuleInfo, len(*allowedIP))
	for i, rule := range *allowedIP {
		rules[i].RemoteIpPrefix = rule.(string)
		rules[i].PortRangeMin = int32(port)
		rules[i].PortRangeMax = int32(port)
	}
	return rules
}

func secgroupRulesRead(d *schema.ResourceData, m interface{}) error {
	cli := m.(*client.Client)

	instanceID := d.Id()
	projectID := cli.ProjectId
	resp, _, err := cli.VdbClient.VdbInstanceEndPointApi.GetDbInstanceSecurityGroupRuleUsingGET(context.TODO(), instanceID, projectID)

	if err != nil {
		return err
	}

	if !resp.Success {
		return errors.New(resp.ErrorMsg)
	}

	rules := make([]interface{}, len(resp.SecurityGroupRules))
	for i, securityGroupRule := range resp.SecurityGroupRules {
		rules[i] = securityGroupRule.RemoteIpPrefix
	}

	log.Printf("[DEBUG] [nhannt10] reading rules %v\n", rules)
	d.Set("allowed_ip_prefix", rules)

	return nil
}

func resourceDatabaseRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database read")

	cli := m.(*client.Client)

	dbResp, _, err := cli.VdbClient.VdbInstanceEndPointApi.GetDbInstanceDetailUsingGET(context.TODO(), d.Id(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when getting database info: %s", err)
	}
	if !dbResp.Success {
		return fmt.Errorf("Error when getting database info: %s", dbResp.ErrorMsg)
	}

	if dbResp.Data == nil {
		d.SetId("")
		return nil
	}

	d.Set("backup_auto", dbResp.Data.BackupAuto)
	d.Set("backup_duration", dbResp.Data.BackupDuration)
	d.Set("backup_time", dbResp.Data.BackupTime)
	d.Set("bandwidth", dbResp.Data.Bandwidth)
	d.Set("datastore_type", dbResp.Data.DatastoreType)
	d.Set("datastore_version", dbResp.Data.DatastoreVersion)
	d.Set("db_backend_id", dbResp.Data.DbBackendId)
	d.Set("engine_group", dbResp.Data.EngineGroup)
	d.Set("free_backup_size", dbResp.Data.FreeBackupSize)
	d.Set("ip", dbResp.Data.Ip)
	d.Set("name", dbResp.Data.Name)
	d.Set("port", dbResp.Data.Port)
	d.Set("price_key", dbResp.Data.PriceKey)
	d.Set("public_access", dbResp.Data.PublicAccess)
	d.Set("ram", dbResp.Data.Ram)
	d.Set("redis_password_enabled", dbResp.Data.RedisPasswordEnabled)
	d.Set("replica_source_id", dbResp.Data.ReplicaSourceId)
	d.Set("replicas", dbResp.Data.Replicas)
	d.Set("start_time", dbResp.Data.StartTime.String())
	d.Set("status", dbResp.Data.Status)
	d.Set("vcpus", dbResp.Data.Vcpus)
	d.Set("volume_size", dbResp.Data.VolumeSize)
	d.Set("volume_type", dbResp.Data.VolumeType)
	d.Set("volume_type_zone_id", dbResp.Data.VolumeTypeZoneId)
	d.Set("volume_used", dbResp.Data.VolumeUsed)
	d.Set("zone_uuid", dbResp.Data.ZoneUUID)

	if dbResp.Data.Configuration != nil {
		d.Set("config_id", dbResp.Data.Configuration.Id)
		d.Set("config_name", dbResp.Data.Configuration.Name)
	}

	return secgroupRulesRead(d, m)
}

func resourceDatabaseUpdate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [nhannt10] Database update")
	if d.HasChange("allowed_ip_prefix") {
		secgroupRulesUpdate(d, m, d.Get("allowed_ip_prefix").([]interface{}))
	}

	if d.HasChange("action") {
		switch d.Get("action").(string) {
		case "start":
			return resourceDatabaseStart(d, m)
		case "stop":
			return resourceDatabaseStop(d, m)
		case "reboot":
			return resourceDatabaseReboot(d, m)
		}
	}

	if d.HasChange("replica_source_id") && d.Get("replica_source_id").(string) == "" {
		return resourceDatabasePromote(d, m)
	}

	if d.HasChange("volume_size") {
		log.Println("[DEBUG] [trihm4] Database update")

		cli := m.(*client.Client)

		updateRequest := generateUpdateDatabaseRequest(cli.ProjectId, d)
		reqBody, _ := json.Marshal(updateRequest)
		log.Println("[DEBUG] [trihm4] Body: " + string(reqBody))

		cli.VdbClient.VdbInstanceEndPointApi.UpdateDbInstanceUsingPUT(context.TODO(), d.Id(), cli.ProjectId, updateRequest)
	}

	return nil
}

func resourceDatabaseDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database delete")

	cli := m.(*client.Client)

	deleteRequest := vdb.DeleteDbInstanceRequest{
		CreateFinalBackup: false,
		DbInstanceId:      d.Id(),
		DeleteAllBackup:   true,
		Extra:             nil,
		IsMetadataCreated: true,
		ProjectId:         cli.ProjectId,
	}

	deleteRequest.Extra = map[string]interface{}{
		"is_trial": false,
	}

	reqBody, _ := json.Marshal(deleteRequest)
	log.Println("[DEBUG] [trihm4] Body: " + string(reqBody))

	putDbInTrashResult, _, err := cli.VdbClient.VdbInstanceEndPointApi.DeleteDbInstanceUsingPUT(context.TODO(), d.Id(), deleteRequest, cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when putting database in trash: %s", err)
	}
	if !putDbInTrashResult.Success {
		return fmt.Errorf("Error when putting database in trash: %s", putDbInTrashResult.ErrorMsg)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseDeletePending,
		Target:     databaseDeleteTarget,
		Refresh:    resourceDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseDeleteTimeout,
		Delay:      databaseDeleteDelay,
		MinTimeout: databaseDeleteMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error when waiting for database to be put in trash: %s", err)
	}

	log.Println("[DEBUG] [trihm4] Put in trash done")

	deleteInTrashRequest := vdb.DeleteDbInstanceInTrashRequest{
		DbInstanceId: d.Id(),
		Extra:        nil,
		ProjectId:    cli.ProjectId,
	}

	deleteDbInTrashResult, _, err := cli.VdbClient.VdbInstanceEndPointApi.DeleteDbInstanceInTrashUsingDELETE(context.TODO(), d.Id(), deleteInTrashRequest, cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when deleting database in trash: %s", err)
	}
	if !deleteDbInTrashResult.Success {
		return fmt.Errorf("Error when deleting database in trash: %s", deleteDbInTrashResult.ErrorMsg)
	}

	return resource.Retry(databaseDeleteInTrashTimeout, func() *resource.RetryError {
		dbResp, _, err := cli.VdbClient.VdbInstanceEndPointApi.GetDbInstanceDetailUsingGET(context.TODO(), d.Id(), cli.ProjectId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error when refreshing database state: %s", err))
		}
		if !dbResp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error when refreshing database state: %s", dbResp.ErrorMsg))
		}

		if dbResp.Data != nil {
			return resource.RetryableError(fmt.Errorf("Database is still not deleted"))
		}

		log.Println("[DEBUG] [trihm4] Delete in trash done")

		d.SetId("")
		return nil
	})
}

func resourceDatabaseStart(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database start")

	cli := m.(*client.Client)

	startDbResult, _, err := cli.VdbClient.VdbInstanceEndPointApi.StartDbInstanceUsingPUT(context.TODO(), d.Id(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when starting database: %s", err)
	}
	if !startDbResult.Success {
		return fmt.Errorf("Error when starting database: %s", startDbResult.ErrorMsg)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseStartPending,
		Target:     databaseStartTarget,
		Refresh:    resourceDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseStartTimeout,
		Delay:      databaseStartDelay,
		MinTimeout: databaseStartMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error when waiting for database to be started: %s", err)
	}

	return resourceDatabaseRead(d, m)
}

func resourceDatabaseStop(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database stop")

	cli := m.(*client.Client)

	stopDbResult, _, err := cli.VdbClient.VdbInstanceEndPointApi.StopDbInstanceUsingPUT(context.TODO(), d.Id(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when stopping database: %s", err)
	}
	if !stopDbResult.Success {
		return fmt.Errorf("Error when stopping database: %s", stopDbResult.ErrorMsg)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseStopPending,
		Target:     databaseStopTarget,
		Refresh:    resourceDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseStopTimeout,
		Delay:      databaseStopDelay,
		MinTimeout: databaseStopMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error when waiting for database to be stopped: %s", err)
	}

	return resourceDatabaseRead(d, m)
}

func resourceDatabaseReboot(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database reboot")

	cli := m.(*client.Client)

	rebootDbResult, _, err := cli.VdbClient.VdbInstanceEndPointApi.RestartDbInstanceUsingPUT(context.TODO(), d.Id(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when rebooting database: %s", err)
	}
	if !rebootDbResult.Success {
		return fmt.Errorf("Error when rebooting database: %s", rebootDbResult.ErrorMsg)
	}

	stateConf := &resource.StateChangeConf{
		Pending:    databaseRebootPending,
		Target:     databaseRebootTarget,
		Refresh:    resourceDatabaseStateRefreshFunc(cli, d.Id()),
		Timeout:    databaseRebootTimeout,
		Delay:      databaseRebootDelay,
		MinTimeout: databaseRebootMinTimeout,
	}

	_, err = stateConf.WaitForState()
	if err != nil {
		return fmt.Errorf("Error when waiting for database to be rebooted: %s", err)
	}

	return resourceDatabaseRead(d, m)
}

func resourceDatabasePromote(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] [trihm4] Database promote")

	cli := m.(*client.Client)

	promoteDbResult, _, err := cli.VdbClient.VdbInstanceEndPointApi.DetachReplicaUsingPOST(context.TODO(), d.Id(), cli.ProjectId)
	if err != nil {
		return fmt.Errorf("Error when promoting database: %s", err)
	}
	if !promoteDbResult.Success {
		return fmt.Errorf("Error when promoting database: %s", promoteDbResult.ErrorMsg)
	}

	err = resource.Retry(databasePromoteTimeout, func() *resource.RetryError {
		dbResp, _, err := cli.VdbClient.VdbInstanceEndPointApi.GetDbInstanceDetailUsingGET(context.TODO(), d.Id(), cli.ProjectId)
		if err != nil {
			return resource.NonRetryableError(fmt.Errorf("Error when refreshing database state: %s", err))
		}
		if !dbResp.Success {
			return resource.NonRetryableError(fmt.Errorf("Error when refreshing database state: %s", dbResp.ErrorMsg))
		}

		if dbResp.Data.ReplicaSourceId != "" {
			return resource.RetryableError(fmt.Errorf("Database is still not promoted"))
		}

		log.Println("[DEBUG] [trihm4] Promote done")

		return nil
	})

	if err != nil {
		return fmt.Errorf("Error when waiting for database to be promoted: %s", err)
	}

	return resourceDatabaseRead(d, m)
}

func generateUpdateDatabaseRequest(projectId string, d *schema.ResourceData) vdb.UpdateDbInstanceRequest {
	updateRequest := vdb.UpdateDbInstanceRequest{
		BackupAuto:           false,
		BackupDuration:       2,
		BackupTime:           "00:00",
		CartItemId:           0,
		CartItemState:        0,
		ConfigId:             "",
		Cost:                 457600.0,
		DbInstanceId:         d.Id(),
		EditRedisPassword:    false,
		EndTime:              nil,
		EngineGroup:          1,
		Extra:                nil,
		FlavorId:             "9",
		InvoiceId:            0,
		IsResized:            true,
		IsUserTrial:          false,
		NewStartDate:         time.Now().Unix() * 1000,
		PackageId:            "8",
		PackageName:          "",
		Password:             "password",
		Period:               1,
		Poc:                  false,
		ProjectId:            projectId,
		PublicAccess:         false,
		Ram:                  1,
		RedisPassword:        "",
		RedisPasswordEnabled: false,
		UpdateType:           "VOLUME",
		Vcpus:                1,
		VolumeSize:           int32(d.Get("volume_size").(int)),
		VolumeType:           "ssd-iops200",
		VolumeTypeZoneId:     "11B9FF75D-00F5-49CA-842D-CCD5D04834FB",
	}

	return updateRequest
}
