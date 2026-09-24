
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
  name        = "soft-affinity-groupA"
  policy_id   = data.vngcloud_vserver_server_group_policy.soft_affinity_policy.id
  project_id  = var.project_id
}

resource "vngcloud_vserver_server_group" "soft_anti_affinity_server_group" {
  description = "soft-anti-affinity"
  name        = "soft-anti-affinity-groupB"
  policy_id   = data.vngcloud_vserver_server_group_policy.soft_anti_affinity_policy.id
  project_id  = var.project_id
}

data "template_file" "cloud-config-bastard" {
  template = "${file("modules/vng-cloud-vserver/init/cloud-config.tpl")}"
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
    content      = "${file("modules/vng-cloud-vserver/init/install-apache.sh")}"
  }
}

resource "vngcloud_vserver_server" "server" {
  for_each        = toset(var.server_names)
  project_id      = var.project_id
  name            = each.value
  encryption_volume = false
  attach_floating   = true
  flavor_id         = var.s_general_4x8_zone1c
  image_id          = var.ubuntu_24_04
  network_id        = var.network_id
  root_disk_size    = var.root_disk_size
  root_disk_type_id = var.volume_type_3000_1c
  security_group    = var.security_group_id_list
  subnet_id         = var.subnet_id
  action            = "start"
  server_group_id         = vngcloud_vserver_server_group.soft_anti_affinity_server_group.id
  user_data_base64_encode = var.user_data_base64_encode
  user_data               = data.template_cloudinit_config.user_data.rendered
  zone_id =  "HCM03-1C"


  lifecycle {
    create_before_destroy = true
  }
}

resource "vngcloud_vserver_volume" "volume" {
  for_each      = toset(var.server_names)
  name          = "${each.value}-volume"
  size          = var.data_disk_size
  volume_type_id = var.volume_type_3000_1c
  project_id    = var.project_id
  zone_id = "HCM03-1C"

  lifecycle {
    create_before_destroy = true
  }
}

resource "vngcloud_vserver_volume_attach" "attach_volume" {
  for_each   = toset(var.server_names)
  project_id = var.project_id
  volume_id  = vngcloud_vserver_volume.volume[each.key].id
  server_id  = vngcloud_vserver_server.server[each.key].id
}