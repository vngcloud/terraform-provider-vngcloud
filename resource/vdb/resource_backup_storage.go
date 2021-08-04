package vdb

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vngcloud/terraform/client"
	"github.com/vngcloud/terraform/client/vdb"
	"log"
	"time"
)

func ResourceBackupStorage() *schema.Resource {
	return &schema.Resource{
		Create: resourceBackupStorageCreate,
		Read:   resourceBackupStorageRead,
		Update: resourceBackupStorageUpdate,
		Delete: resourceBackupStorageDelete,
		Importer: &schema.ResourceImporter{
			State: func(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
				d.SetId(d.Id())
				return []*schema.ResourceData{d}, nil
			},
		},
		Schema: map[string]*schema.Schema{
			/*"action": {
				Type:     schema.TypeString,
				Required: true,
			},*/
			"backup_storage_package_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"backup_storage_package_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"period": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"engine_group": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"monthly_cost": {
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"quota": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"usage": {
				Type:     schema.TypeFloat,
				Computed: true,
			},
		},
	}
}

func resourceBackupStorageRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage read")

	cli := m.(*client.Client)

	backupStorageResult, _, _ := cli.VdbClient.VdbBackupStorageEndPointApi.GetBackupStorageByIdUsingGET(context.TODO(), d.Id(), cli.ProjectId)
	if backupStorageResult.Success == false {
		d.SetId("")
		return nil
	}

	backupStorageInfo := backupStorageResult.Data[0]

	d.Set("backup_storage_package_id", backupStorageInfo.BackupPackageId)
	d.Set("backup_storage_package_name", backupStorageInfo.BackupPackageName)
	d.Set("name", backupStorageInfo.Name)
	d.Set("quota", backupStorageInfo.Quota)
	d.Set("period", backupStorageInfo.Period)
	d.Set("engine_group", backupStorageInfo.EngineGroup)
	d.Set("usage", backupStorageInfo.Usage)

	return nil
}

func resourceBackupStorageCreate(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage create")

	cli := m.(*client.Client)

	createBackupStorageRequest := generateCreateBackupStorageRequest(cli.ProjectId, d)
	reqBody, _ := json.Marshal(createBackupStorageRequest)
	log.Println("[DEBUG] Request body " + string(reqBody))

	engineGroup := d.Get("engine_group").(int)
	createBackupStorageResult, _, _ := cli.VdbClient.VdbBackupStorageEndPointApi.CreateBackupStorageUsingPOST(context.TODO(), int32(engineGroup), cli.ProjectId, createBackupStorageRequest)
	if createBackupStorageResult.Success == false {
		return errors.New(createBackupStorageResult.ErrorMsg)
	}

	idStr := createBackupStorageResult.Data[0].Id

	d.SetId(idStr)
	return nil
}

func generateCreateBackupStorageRequest(projectId string, d *schema.ResourceData) vdb.CreateBackupStorageRequest {
	engineGr := int32(d.Get("engine_group").(int))
	var backupStorageName = ""
	if engineGr == 1 {
		backupStorageName = "relational_backup_storage"
	} else if engineGr == 2 {
		backupStorageName = "memory_store_backup_storage"
	}
	createRequest := vdb.CreateBackupStorageRequest{
		BackupPackageId:   d.Get("backup_storage_package_id").(string),
		BackupPackageName: d.Get("backup_storage_package_name").(string),
		StartDate:         time.Now(),
		EndDate:           time.Now(),
		Period:            0,
		EngineGroup:       engineGr,
		Extra:             map[string]string{},
		MonthlyCost:       0,
		Name:              backupStorageName,
		Quota:             int32(d.Get("quota").(int)),
		ProjectId:         projectId,
	}
	return createRequest
}

func resourceBackupStorageUpdate(d *schema.ResourceData, m interface{}) error {
	/*if d.HasChange("action") {
		switch d.Get("action").(string) {
		case "resize":
			return resourceBackupStorageResizeQuota(d, m)
			break
		}
	}
	return nil*/
	return resourceBackupStorageResizeQuota(d, m)
}

func resourceBackupStorageResizeQuota(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage resize quota")

	cli := m.(*client.Client)

	resizeQuotaRequest := vdb.ResizeBackupStorageRequest{
		BackupStorageId:   d.Id(),
		BackupStorageName: d.Get("name").(string),
		BackupPackageId:   d.Get("backup_storage_package_id").(string),
		BackupPackageName: d.Get("backup_storage_package_name").(string),
		EngineGroup:       int32(d.Get("engine_group").(int)),
		Extra:             map[string]string{},
		MonthlyCost:       d.Get("monthly_cost").(float64),
		ProjectId:         cli.ProjectId,
		StartDate:         time.Now(),
		Quota:             int32(d.Get("quota").(int)),
	}
	reqBody, _ := json.Marshal(resizeQuotaRequest)
	log.Println("[DEBUG] Resize quota request body " + string(reqBody))

	resizeQuotaResult, _, _ := cli.VdbClient.VdbBackupStorageEndPointApi.ResizeBackupStorageUsingPUT(context.TODO(), cli.ProjectId, resizeQuotaRequest)
	if resizeQuotaResult.Success == false {
		return errors.New(resizeQuotaResult.ErrorMsg)
	}
	return nil
}

func resourceBackupStorageDelete(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG] Backup storage delete")

	cli := m.(*client.Client)

	deleteBackupStorageRequest := vdb.DeleteBackupStorageRequest{
		BackupStorageId:   d.Id(),
		BackupStorageName: d.Get("name").(string),
		EngineGroup:       int32(d.Get("engine_group").(int)),
		ProjectId:         cli.ProjectId,
		Extra:             map[string]string{},
	}
	reqBody, _ := json.Marshal(deleteBackupStorageRequest)
	log.Println("[DEBUG] Delete request body " + string(reqBody))

	deleteBackupStorageResult, _, _ := cli.VdbClient.VdbBackupStorageEndPointApi.DeleteBackupStorageUsingPUT(context.TODO(), cli.ProjectId, deleteBackupStorageRequest)
	if deleteBackupStorageResult.Success == false {
		return errors.New(deleteBackupStorageResult.ErrorMsg)
	}

	deleteInTrashBackupStorageRequest := vdb.DeleteInTrashBackupStorageRequest{
		BackupStorageId:   deleteBackupStorageRequest.BackupStorageId,
		BackupStorageName: deleteBackupStorageRequest.BackupStorageName,
		EngineGroup:       deleteBackupStorageRequest.EngineGroup,
		ProjectId:         deleteBackupStorageRequest.ProjectId,
		Extra:             deleteBackupStorageRequest.Extra,
	}
	delInTrashReqBody, _ := json.Marshal(deleteBackupStorageRequest)
	log.Println("[DEBUG] Delete in trash request body " + string(delInTrashReqBody))

	deleteInTrashBackupStorageResult, _, _ := cli.VdbClient.VdbBackupStorageEndPointApi.DeleteInTrashBackupStorageUsingDELETE(context.TODO(), cli.ProjectId, deleteInTrashBackupStorageRequest)
	if deleteInTrashBackupStorageResult.Success == false {
		return errors.New(deleteInTrashBackupStorageResult.ErrorMsg)
	}
	d.SetId("")
	return nil
}
