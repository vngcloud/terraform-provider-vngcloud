---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "vDB"
page_title: "vngcloud_vdb_relational_backup Resource - terraform-provider-vngcloud"
description: |-
  Provides a VNG Cloud vDB Relational Backup. This can be used to import, create, modify, and delete.
  
---

# vngcloud_vdb_relational_backup (Resource)



## Example Usage

```terraform
resource "vngcloud_vdb_relational_backup" "redis_4_backup" {
  name = "tf-redis-backup"
  instance_id = "db-9ecfaeca-5cdc-4ee0-a783-8cd5c494ddc3"
  backup_type = "INCREMENTAL"
  parent_id = "bk-04627eab-ffb9-4bbb-abdb-02a8d78e06c2"
}
```

## Argument Reference

The following arguments are supported:

- **name** (String, Required, ForceNew):
  - The name of the backup. This field is required to create the backup and cannot be changed after creation.

- **instance_id** (String, Required, ForceNew):
  - The ID of the instance for which the backup is created. This field is required and cannot be changed after creation.

- **backup_type** (String, Required, ForceNew):
  - The type of backup (FULL, INCREMENTAL).

- **parent_id** (String, Optional, ForceNew):
  - The ID of the parent backup (Required when `backup_type` is `INCREMENTAL`).
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:
- **id** (String) The ID of this backup.

- **instance_name** (String, Computed):
  - The name of the instance for which the backup is created.

- **type** (String, Computed):
  - The type by which this backup is created (MANUAL, AUTO_DAILY).

- **created** (String, Computed):
  - The time of when the backup was created.

- **status** (String, Computed):
  - The current status of the backup.

- **engine_type** (String, Computed):
  - The type of the engine used by the instance.

- **engine_version** (String, Computed):
  - The version of the engine used by the instance.

- **storage_type** (String, Computed):
  - The type of storage used by the instance.

- **storage_size** (Int, Computed):
  - The size (in GB) of the storage of the instance.

- **ram** (Int, Computed):
  - The amount of RAM (in GB) allocated to the instance.

- **cpu** (Int, Computed):
  - The number of CPU cores allocated to the instance.

- **config_id** (String, Computed):
  - The ID of the configuration used for the instance.

- **config_name** (String, Computed):
  - The name of the configuration used for the instance.

- **username** (String, Computed):
  - The username associated with the instance.

- **net_ids** (List of String, Computed):
  - A list of network IDs associated with the instance.

- **net_name** (String, Computed):
  - The name of the network associated with the instance.

- **size** (Float, Computed):
  - The size of the backup in GB.

- **backup_tier** (String, Computed):
  - The backup tier (FREE, PAID).


