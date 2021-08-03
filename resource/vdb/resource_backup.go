package vdb

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"git.vngcloud.tech/vdb/vdb-terraform/client"
	"git.vngcloud.tech/vdb/vdb-terraform/client/vdb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupCreate,
		Read:   resourceBackupRead,
		Delete: resourceBackupDelete,
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
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"instance_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"backup_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datastore_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datastore_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ram": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vcpu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"config_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"config_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"net_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"net_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"price_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
			"backup_tier": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"engine_group": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: true,
			},
			"backup_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func resourceBackupRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup read")

	cli := m.(*client.Client)

	dbResp, _, _ := cli.VdbClient.VdbBackupEndPointApi.GetBackupByIdUsingGET(context.TODO(), d.Id(), cli.ProjectId)

	if dbResp.Data == nil {
		d.SetId("")
		return nil
	}

	d.Set("name", dbResp.Data.Name)
	d.Set("description", dbResp.Data.Description)
	d.Set("instance_id", dbResp.Data.DbInstanceId)
	d.Set("instance_name", dbResp.Data.InstanceName)
	d.Set("type", dbResp.Data.Type_)
	d.Set("parent_id", dbResp.Data.Parent)
	d.Set("parent_name", dbResp.Data.ParentName)
	d.Set("created", dbResp.Data.Created)
	d.Set("backup_type", dbResp.Data.BackupType)
	d.Set("status", dbResp.Data.Status)
	d.Set("datastore_type", dbResp.Data.DatastoreType)
	d.Set("datastore_version", dbResp.Data.DatastoreVersion)
	d.Set("storage_type", dbResp.Data.StorageType)
	d.Set("storage_size", dbResp.Data.StorageSize)
	d.Set("ram", dbResp.Data.Ram)
	d.Set("vcpu", dbResp.Data.Vcpu)
	d.Set("config_id", dbResp.Data.ConfigId)
	d.Set("config_name", dbResp.Data.ConfigName)
	d.Set("username", dbResp.Data.Username)
	d.Set("net_ids", dbResp.Data.NetIds)
	d.Set("net_name", dbResp.Data.NetName)
	d.Set("price_key", dbResp.Data.PriceKey)
	d.Set("size", dbResp.Data.Size)
	d.Set("backup_tier", dbResp.Data.BackupTier)
	d.Set("engine_group", dbResp.Data.EngineGroup)
	d.Set("backup_duration", dbResp.Data.BackupDuration)

	return nil
}

func resourceBackupCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup create")

	cli := m.(*client.Client)

	createRequest := generateCreateBackupRequest(cli.ProjectId, d)
	reqBody, _ := json.Marshal(createRequest)
	log.Println("[DEBUG] Body: " + string(reqBody))

	createDbResult, _, _ := cli.VdbClient.VdbBackupEndPointApi.CreateBackupUsingPOST(context.TODO(), cli.ProjectId, createRequest)
	log.Println("[DEBUG] Create backup result, id: " + createDbResult.BackupId)

	stateConf := &resource.StateChangeConf{
		Pending:    []string{"NEW", "BUILDING", "SAVING"},
		Target:     []string{"COMPLETED"},
		Refresh:    resourceBackupStateRefreshFunc(cli, createDbResult.BackupId),
		Timeout:    10 * time.Minute, //TODO update config
		Delay:      60 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	id, _ := stateConf.WaitForStateContext(context.TODO())
	idStr := id.(string)

	log.Println("[DEBUG] Wait for state done, id: " + idStr)

	d.SetId(idStr)

	resourceBackupRead(d, m)

	return nil
}

func generateCreateBackupRequest(projectId string, d *schema.ResourceData) vdb.BackupRequest {
	notifResTypeMap := make(map[int]string)
	notifResTypeMap[1] = "relational_dbaas_backup_v3"
	notifResTypeMap[2] = "memory_store_dbaas_backup_v3"

	extraMap := make(map[string]string)
	extraMap["notif_resource_type"] = notifResTypeMap[d.Get("engine_group").(int)]

	createRequest := vdb.BackupRequest{
		ProjectId:    projectId,
		DbInstanceId: d.Get("instance_id").(string),
		Name:         d.Get("name").(string),
		Description:  d.Get("description").(string),
		BackupType:   d.Get("backup_type").(string),
		ParentId:     d.Get("parent_id").(string),
		EngineGroup:  int32(d.Get("engine_group").(int)),
		Extra:        extraMap,
	}

	return createRequest
}

func resourceBackupStateRefreshFunc(cli *client.Client, id string) resource.StateRefreshFunc {
	return func() (interface{}, string, error) {
		log.Println("[DEBUG] State refresh")

		dbResp, _, _ := cli.VdbClient.VdbBackupEndPointApi.GetBackupByIdUsingGET(context.TODO(), id, cli.ProjectId)
		log.Println("[DEBUG] Database status: " + dbResp.Data.Status)

		return dbResp.Data.Id, dbResp.Data.Status, nil
	}
}

func resourceBackupDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup delete")

	cli := m.(*client.Client)
	backupId := d.Id()
	projectID := cli.ProjectId

	resp, _, _ := cli.VdbClient.VdbBackupEndPointApi.DeleteBackupUsingDELETE(context.TODO(), backupId, projectID)

	if !resp.Success {
		return errors.New(resp.ErrorMsg)
	}

	log.Println("[DEBUG] Backup deleted " + d.Id())
	d.SetId("")

	return nil
}

func generateRestoreBackupRequest(projectId string, d *schema.ResourceData) vdb.RestoreRequest {
	restoreRequest := vdb.RestoreRequest{
		AutoRenewPeriod:      1,
		BackupAuto:           d.Get("backup_auto").(bool),
		BackupDuration:       int32(d.Get("backup_duration").(int)),
		BackupId:             d.Get("backup_id").(string),
		BackupTime:           d.Get("backup_time").(string),
		CartItemId:           0,
		CartItemState:        0,
		ConfigId:             d.Get("config_id").(string),
		Cost:                 0,
		DatastoreType:        d.Get("datastore_type").(string),
		DatastoreVersion:     d.Get("datastore_version").(string),
		EnableAutoRenew:      false,
		EndTime:              nil,
		EngineGroup:          int32(d.Get("engine_group").(int)),
		Extra:                nil,
		FlavorId:             d.Get("flavor_id").(string),
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
		StartTime:            time.Now(),
		UseTrial:             false,
		Vcpus:                int32(d.Get("vcpus").(int)),
		VolumeSize:           int32(d.Get("volume_size").(int)),
		VolumeType:           d.Get("volume_type").(string),
		VolumeTypeZoneId:     d.Get("volume_type_zone_id").(string),
		ZoneId:               int32(d.Get("zone_id").(int)),
	}

	return restoreRequest
}
