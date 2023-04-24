data "vngcloud_vserver_volume_type_zone" "volume_type_zone" {
  name       = "SSD"
  project_id = var.project_id
}

data "vngcloud_vserver_volume_type" "volume_type" {
  name                = var.volume_type_name
  project_id          = var.project_id
  volume_type_zone_id = data.vngcloud_vserver_volume_type_zone.volume_type_zone.id
}

data "vngcloud_vserver_k8s_version" "k8sVersion1" {
  name       = "Version 1.18.7"
  project_id = var.project_id
}

data "vngcloud_vserver_k8s_network_type" "k8sNetworkType" {
  name       = "Calico"
  project_id = var.project_id
}

resource "vngcloud_vserver_k8s" "k8s" {
  project_id            = var.project_id
  ipip_mode             = "Always"
  name                  = "vng-cloud-k8s-example"
  k8s_version           = data.vngcloud_vserver_k8s_version.k8sVersion1.id
  master_count          = 1
  node_count            = 1
  network_type          = data.vngcloud_vserver_k8s_network_type.k8sNetworkType.id
  calico_cidr           = "10.46.0.0/16"
  network_id            = var.network_id
  subnet_id             = var.subnet_id
  ssh_key_id            = var.ssh_key_id
  master_flavor_id      = var.s_general_4x8
  node_flavor_id        = var.s_general_4x8
  etcd_volume_size      = 20
  etcd_volume_type_id   = data.vngcloud_vserver_volume_type.volume_type.id
  boot_volume_size      = 20
  boot_volume_type_id   = data.vngcloud_vserver_volume_type.volume_type.id
  docker_volume_size    = 20
  docker_volume_type_id = data.vngcloud_vserver_volume_type.volume_type.id
  description           = "K8S example"
  auto_scaling          = true
  min_node_count        = 1
  max_node_count        = 3
  enable_lb             = false
  auto_healing          = true
  ingress_controller = true
}

resource "vngcloud_vserver_cluster_node_group" "node_group" {
  project_id = var.project_id
  name       = "node-group-example"
  cluster_id = vngcloud_vserver_k8s.k8s.id
  flavor_id  = var.s_general_4x8
  node_count = 1
}


resource "vngcloud_vserver_change_cluster_sec_group" "change_cluster_sec_group_minion" {
  project_id        = var.project_id
  cluster_id        = vngcloud_vserver_k8s.k8s.id
  master            = false
  sec_group_id_list = var.security_group_id_list
  depends_on        = [vngcloud_vserver_cluster_node_group.node_group]

}

resource "vngcloud_vserver_change_cluster_sec_group" "change_cluster_sec_group_master" {
  project_id        = var.project_id
  cluster_id        = vngcloud_vserver_k8s.k8s.id
  master            = true
  sec_group_id_list = var.security_group_id_list
  depends_on        = [vngcloud_vserver_change_cluster_sec_group.change_cluster_sec_group_minion]
}

resource "local_file" "config" {
  content  = vngcloud_vserver_k8s.k8s.config
  filename = "config"
}


#resource "vngcloud_vserver_attach_lb_to_cluster" "attach_lb_to_cluster2" {
#  project_id = var.project_id
#  cluster_id = vngcloud_vserver_k8s.k8s.id
#  depends_on = [vngcloud_vserver_change_cluster_sec_group.change_cluster_sec_group_master]
#  load_balancers {
#    load_balancer_id = "lb-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#    pools {
#      monitor_port = 80
#      pool_id      = "pool-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#      port         = 80
#    }
#    pools {
#      monitor_port = 443
#      pool_id      = "pool-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#      port         = 443
#    }
#  }
#  load_balancers {
#    load_balancer_id = "lb-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#    pools {
#      monitor_port = 80
#      pool_id      = "pool-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
#      port         = 80
#  }
#}