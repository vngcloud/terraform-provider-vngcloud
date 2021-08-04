terraform {
    required_providers {
        vngcloud = {
            version = "~> 0.2"
            source = "vngcloud.vn/terraform/vngcloud"
        }
    }
}

provider "vngcloud" {
    user_id = "11317"
    project_id = "pro-4a3fcd8d-2459-4a5d-b602-db6552205f78"
    #base_url = "http://localhost:8101"
    vdb_base_url = "https://vcmc-internal.vngcloud.tech/vdb-gateway"
    vserver_base_url = "https://vserverapi.vngcloud.vn/vserver-gateway"
    token_url = "https://vcmc.vngcloud.tech/user-api/v1/user/oauth2/token"
    client_id = "5V23JsOSOwM7puYnMPqxc6ibBqOlpbPI"
    client_secret = "Q4DRR473i2weaE2xmWHR5mR9fBwp8X9K"
}

data "vngcloud_vdb_package" "package" {
    engine_type = "MySQL"
    engine_version = "5.7"
    name = "v1.small1x1"
}