terraform {
  required_providers {
    vngcloud = {
      source  = "vngcloud/vngcloud"
      version = "0.0.17"
    }
  }
}

provider "vngcloud" {
  token_url        = "https://iamapis.vngcloud.vn/accounts-api/v2/auth/token"
  client_id        = var.client_id
  client_secret    = var.client_secret
  vserver_base_url = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway"
}

data "vngcloud_vserver_volume_type_zone" "volume_type_zone" {
  name       = "SSD"
  project_id = var.project_id
}
data "vngcloud_vserver_volume_type" "volume_type" {
  name                = var.ssd_3000
  project_id          = var.project_id
  volume_type_zone_id = data.vngcloud_vserver_volume_type_zone.volume_type_zone.id
}

data "vngcloud_vserver_server_group_policy" "soft_affinity_policy" {
  name       = var.soft_affinity
  project_id = var.project_id
}

data "vngcloud_vserver_server_group_policy" "soft_anti_affinity_policy" {
  name       = var.soft_anti_affinity
  project_id = var.project_id
}

resource "vngcloud_vserver_server_group" "soft_affinity_server_group" {
  description = "soft-affinity"
  name        = "soft-affinity"
  policy_id   = data.vngcloud_vserver_server_group_policy.soft_affinity_policy.id
  project_id  = var.project_id
}

resource "vngcloud_vserver_server_group" "soft_anti_affinity_server_group" {
  description = "soft-anti-affinity"
  name        = "soft-anti-affinity"
  policy_id   = data.vngcloud_vserver_server_group_policy.soft_anti_affinity_policy.id
  project_id  = var.project_id
}

data "template_file" "cloud-config-bastard" {
  template = "${file("init/cloud-config.tpl")}"
}

data "template_cloudinit_config" "user_data" {
  gzip          = var.user_data_base64_encode
  base64_encode = var.user_data_base64_encode

  part {
    content_type = "text/cloud-config"
    content      = "${data.template_file.cloud-config-bastard.rendered}"
  }

  part {
    content_type = "text/x-shellscript"
    content      = "${file("init/install-apache.sh")}"
  }
}

resource "vngcloud_vserver_server" "server" {
  count             = var.server_count
  project_id        = var.project_id
  name              = "vngcloud-server-${count.index}"
  encryption_volume = false
  attach_floating   = true
  flavor_id         = var.s_general_4x8
  image_id          = var.ubuntu_20_04
  network_id        = var.network_id
  root_disk_size    = var.root_disk_size
  root_disk_type_id = data.vngcloud_vserver_volume_type.volume_type.id
  security_group    = var.security_group_id_list
  subnet_id         = var.subnet_id
  action            = "start"
  #  user_name         = "stackops"
  #  user_password     = "Vng@Cloud3030"
  #  expire_password   = false
  #  ssh_key           = var.ssh_key_id
  server_group_id         = vngcloud_vserver_server_group.soft_anti_affinity_server_group.id
  user_data_base64_encode = var.user_data_base64_encode
  user_data               = "${data.template_cloudinit_config.user_data.rendered}"
  lifecycle {
    create_before_destroy = true
  }
}

resource "vngcloud_vserver_volume" "volume" {
  count          = var.server_count
  name           = "vngcloud-volume-${count.index}"
  size           = var.data_disk_size
  volume_type_id = data.vngcloud_vserver_volume_type.volume_type.id
  project_id     = var.project_id
  lifecycle {
    create_before_destroy = true
  }
}

resource "vngcloud_vserver_volume_attach" "attach_volume" {
  count      = var.server_count
  project_id = var.project_id
  volume_id  = vngcloud_vserver_volume.volume[count.index].id
  server_id  = vngcloud_vserver_server.server[count.index].id
}
