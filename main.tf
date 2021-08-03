terraform {
    required_providers {
        vdb = {
            version = "~> 0.2"
            source = "vngcloud.vn/terraform/vdb"
        }
    }
}

provider "vdb" {
    user_id = "11279"
    project_id = "pro-5c16c28a-07fa-4836-bc83-daba5a907d86"
    base_url = "https://vcmc-internal.vngcloud.tech/vdb-gateway"
    token_url = "https://vcmc.vngcloud.tech/user-api/v1/user/oauth2/token"
    client_id = "5V23JsOSOwM7puYnMPqxc6ibBqOlpbPI"
    client_secret = "Q4DRR473i2weaE2xmWHR5mR9fBwp8X9K"
}

data "vdb_backup_storage_package" "backup_storage_package_1"{
    engine_group = 1
    name = "db.backup.quota.2"
}

/*resource "vdb_backup_storage" "backup_storage_1" {
    backup_storage_package_id = data.vdb_backup_storage_package.backup_storage_package_1.id
    backup_storage_package_name = data.vdb_backup_storage_package.backup_storage_package_1.name
    engine_group = data.vdb_backup_storage_package.backup_storage_package_1.engine_group
    period = 0
    monthly_cost = 0
    name = "relational_backup_storage"
    quota = data.vdb_backup_storage_package.backup_storage_package_1.quota
}*/

data "vdb_backup_storage_package" "backup_storage_package_2"{
    engine_group = 2
    name = "db.backup.quota.2"
}

/*resource "vdb_backup_storage" "backup_storage_2" {
    backup_storage_package_id = data.vdb_backup_storage_package.backup_storage_package_2.id
    backup_storage_package_name = data.vdb_backup_storage_package.backup_storage_package_2.name
    engine_group = data.vdb_backup_storage_package.backup_storage_package_2.engine_group
    period = 0
    monthly_cost = 0
    name = "memory_store_backup_storage"
    quota = data.vdb_backup_storage_package.backup_storage_package_2.quota
}*/
