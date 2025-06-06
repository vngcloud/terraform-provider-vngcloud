---
subcategory: "Kubernetes Service"
description: |-
  Creates a Cluster on VNGCloud Kubernetes Service (VKS).
---

# vngcloud_vks_cluster

Manages a VNGCloud Kubernetes Engine (VKS) cluster.

To get more information about VKS clusters, see:

* How-to guides
    * [VKS overview](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/vks-la-gi)
    * [Getting Start with VKS](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/bat-dau-voi-vks)
* Terraform guidance
    * [Using VKS with Terraform](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/su-dung-vks-voi-terraform)

---
## Example Usage - with a separately managed node group (recommended)

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "my-vks-cluster"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  name= "my-vks-node-group"
  ssh_key_id= "ssh-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  cluster_id= vngcloud_vks_cluster.primary.id
}
```

**Important Note**: We suggest managing node groups as independent resources, as shown in this example. This approach enables you to add or remove node groups without having to rebuild the entire cluster. If you embed node groups directly within the vngcloud_vks_cluster resource, you will need to recreate the cluster to remove them.

## Example Usage - with the default node group

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "my-vks-cluster"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  node_group {
    name= "my-vks-node-group"
    ssh_key_id= "ssh-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  }
}
```

## Argument Reference

* `name` - (Required) The name of the cluster. Only letters (a-z, 0-9, '-') are allowed. Your input data length must be between 5 and 20.
* `config` - (Computed) This field represents the Cluster's configuration. You don't need to provide any input for this field when creating a Cluster.
* `description` - (Optional) Description of the cluster. Only letters (a-z, A-Z, 0-9, '@', '.' , '_' , '-' , ' '). Your input data length must be between 0 and 255.
* `version` - (Optional) Specifies the version you wish to use for your Cluster. You can view all available Kubernetes versions [here](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/tham-khao-them/phien-ban-ho-tro-kubernetes). The default value is "Version 1.29.1-vks.1724605200".
* `white_list_node_cidr` - (Optional) Specifies the IP address range that can connect to the control plane. This feature is only functional in Private Node Group mode.
* `enable_private_cluster` - (Optional) Enables the private cluster feature,
  creating a private endpoint on the cluster. The VKS public clusters refer to a type of Kubernetes cluster configuration where the Kubernetes API server endpoint is publicly accessible over the internet. In an VKS public cluster, the API server endpoint is not restricted to private access within a VPC (Virtual Private Cloud) and can be accessed over the public internet. The VKS private clusters are configured to have private access to the Kubernetes API server endpoint. This means that the API server endpoint is only accessible from within a specific Virtual Private Cloud (VPC) and is not exposed to the public internet. Private clusters provide enhanced security by restricting access to the Kubernetes API to resources within the VPC. At this time, the default value of this field is false and we only offer Public Cluster mode. The default value is "false".
* `enable_service_endpoint` - (Optional) Enables the creation and use of private service endpoints within your cluster.
* `network_type` - (Optional) The type of network for the cluster. The default value is `CALICO`. You can choose one in many options including `CALICO`, `CILIUM_OVERLAY`, `CILIUM_NATIVE_ROUTING`.
* `vpc_id` - (Required) The VPC ID for the cluster. You need to create a VPC on vServer and enter the VPC's ID in this field.
* `subnet_id` -  (Required) The subnet ID for the cluster. You need to create a Subnet on vServer and enter the Subnet's ID in this field.
* `cidr` -  (Required) Specifies the CIDR block for the cluster using `CALICO` or `CILIUM_OVERLAY` network. You can enter a private IP CIDR from the following options: 10.0.0.0 - 10.255.0.0, 172.16.0.0 - 172.24.0.0, or 192.168.0.0. The default value is "172.16.0.0/16".
* `secondary_subnets` - (Optional) Specifies additional subnets to be useds in Cilium's VPC Native Routing mode.
* `node_netmask_size` - (Optional) Specifies the node CIDR mask size used in Cilium's VPC Native Routing mode. The default value is 25. You can enter a number from the following options: 24, 25, 26.
* `enabled_load_balancer_plugin` - (Optional) Enables/ Disable the attachment of load balancers (both network and application) via Kubernetes YAML. The default value is "true".
* `enabled_block_store_csi_plugin` - (Optional) Enable/ Disable Automatically deploys and manages the BlockStore Persistent Disk CSI Driver via Kubernetes YAML. The default value is "true".
* `auto_upgrade_config` - (Optional) To configure the `auto_upgrade_config` feature for automated maintenance on your cluster, you can use the following attributes:
    * `weekdays` - A list of days of the week when maintenance should occur, e.g., sat,sun (Saturday and Sunday).
    * `time` - The specific time of the day to start maintenance, in 24-hour format, e.g., 21:00 (9 PM).
---
### **Some important notes when using VKS with Terraform:**

When using **Terraform** to create a **Cluster** and **Node Group** on the VKS system, if you modify any of the following fields, the system will automatically delete the existing Node Group/Cluster and recreate a new one with the corresponding new parameters. The deletion process will occur before the creation of the new Node Group/Cluster.

* For the resource `vngcloud_vks_cluster`, the fields that, when modified, will cause the system to delete and recreate the Cluster include:
    * `name`&#x20;
    * `description`&#x20;
    * `enable_private_cluster`&#x20;
    * `network_type`&#x20;
    * `vpc_id`&#x20;
    * `subnet_id`&#x20;
    * `cidr`&#x20;
    * `node_group`&#x20;
    * `secondary_subnets`&#x20;
    * `node_netmask_size`
* For the resource `vngcloud_vks_cluster_node_group`, the fields that, when modified, will cause the system to delete and recreate the Node Group include:
    * `cluster_id`&#x20;
    * `name`&#x20;
    * `flavor_id`&#x20;
    * `disk_size`&#x20;
    * `disk_type`&#x20;
    * `enable_private_nodes`&#x20;
    * `ssh_key_id`&#x20;
    * `secondary_subnets`&#x20;
    * `enabled_encryption_volume`&#x20;
    * `subnet_id`

To specify that the system should create a new cluster/node group before deleting the old one, you can add the parameter `lifecycle { create_before_destroy = true }`to your main.tf file. Specifically:

* For the resource `vngcloud_vks_cluster`

```
resource "vngcloud_vks_cluster" "example" {
  # ...
 
  lifecycle {
    create_before_destroy = true
  }
}
```

* For the resource `vngcloud_vks_cluster_node_group`

```
resource "vngcloud_vks_cluster_node_group" "example" {
  # ...
 
  lifecycle {
    create_before_destroy = true
  }
}
```
---

### Example Usage 1 - Create a Public Cluster with AutoScale mode enabled and a maintenance window set for every Monday to Friday at 10 PM.

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "cluster-demo"
  description = "Cluster create via terraform"
  version = "1.29.1-vks.1724605200"
  cidr      = "172.16.0.0/16"
  enable_private_cluster = false
  network_type = "CALICO"
  vpc_id    = "net-70ef12d4-d619-43fc-88f0-1c1511683123"
  subnet_id = "sub-0725ef54-a32e-404c-96f2-34745239c123"
  enabled_load_balancer_plugin = true
  enabled_block_store_csi_plugin = true
  auto_upgrade_config {
    weekdays = "mon,tue,wed,thu,fri"
    time = "22:00"
  }
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  cluster_id = vngcloud_vks_cluster.primary.id
  name = "nodegroup1"
  num_nodes = 3
  auto_scale_config {
    min_size = 0
    max_size = 5
  }
  upgrade_config {
    strategy = "SURGE"
    max_surge = 1
    max_unavailable = 0
  }
  image_id = "img-108b3a77-ab58-4000-9b3e-190d0b4b07fc"
  flavor_id = "flav-9e88cfb4-ec31-4ad4-8ba5-243459f6d123"
  disk_size = 50
  disk_type = "vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018"
  enable_private_nodes = false
  ssh_key_id= "ssh-f923c53c-cba7-4131-9f86-175d04ae2123"
  security_groups = ["secg-faf05344-fbd6-4f10-80a2-cda08d15ba5e"]
  labels = {
    "test" = "terraform"
  }
  taint {
    key    = "key1"
    value  = "value1"
    effect = "PreferNoSchedule"
  }
}
```

### Example Usage 2 - Create a Private cluster using a private endpoint on VNGCloud with AutoScale mode enabled and the network type is CILIUM OVERLAY and a maintenance window set for every Thurday at 9 AM.

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "cluster-demo"
  description = "Cluster create via terraform"
  version = "1.29.1-vks.1724605200"
  cidr      = "172.16.0.0/16"
  enable_private_cluster = true
  enable_service_endpoint = true
  network_type = "CILIUM_OVERLAY"
  vpc_id    = "net-70ef12d4-d619-43fc-88f0-1c1511683123"
  subnet_id = "sub-0725ef54-a32e-404c-96f2-34745239c123"
  enabled_load_balancer_plugin = true
  enabled_block_store_csi_plugin = true
  auto_upgrade_config {
    weekdays = "thu"
    time = "09:00"
  }
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  cluster_id = vngcloud_vks_cluster.primary.id
  name = "nodegroup1"
  num_nodes = 3
  auto_scale_config {
    min_size = 0
    max_size = 5
  }
  upgrade_config {
    strategy = "SURGE"
    max_surge = 1
    max_unavailable = 0
  }
  image_id = "img-108b3a77-ab58-4000-9b3e-190d0b4b07fc"
  flavor_id = "flav-9e88cfb4-ec31-4ad4-8ba5-243459f6d123"
  disk_size = 50
  disk_type = "vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018"
  enable_private_nodes = false
  ssh_key_id= "ssh-f923c53c-cba7-4131-9f86-175d04ae2123"
  security_groups = ["secg-faf05344-fbd6-4f10-80a2-cda08d15ba5e"]
  labels = {
    "test" = "terraform"
  }
  taint {
    key    = "key1"
    value  = "value1"
    effect = "PreferNoSchedule"
  }
}
```

### Example Usage 3 - Create a private cluster using a private endpoint on VNGCloud with AutoScale mode enabled and the network type is CILIUM VPC NATIVE ROUTING and a maintenance window set for everyday at 11 PM.

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "cluster-demo"
  description = "Cluster create via terraform"
  version = "v1.29.1"
  enable_private_cluster = false
  enable_service_endpoint = false
  network_type = "CILIUM_NATIVE_ROUTING"
  vpc_id    = "net-70ef12d4-d619-43fc-88f0-1c1511683123"
  subnet_id = "sub-0725ef54-a32e-404c-96f2-34745239c123"
  secondary_subnets = ["10.200.27.0/24", "10.200.28.0/24"]
  node_netmask_size = 25
  enabled_load_balancer_plugin = true
  enabled_block_store_csi_plugin = true
  auto_upgrade_config {
    weekdays = "mon,tue,wed,thu,fri,sat,sun"
    time = "23:00"
  }
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  cluster_id = vngcloud_vks_cluster.primary.id
  name = "nodegroup1"
  num_nodes = 3
  auto_scale_config {
    min_size = 0
    max_size = 5
  }
  upgrade_config {
    strategy = "SURGE"
    max_surge = 1
    max_unavailable = 0
  }
  image_id = "img-108b3a77-ab58-4000-9b3e-190d0b4b07fc"
  flavor_id = "flav-9e88cfb4-ec31-4ad4-8ba5-243459f6d123"
  subnet_id = "sub-cddd7ffa-be05-4698-9b3d-794e1adfcbce"
  secondary_subnets = ["10.200.27.0/24", "10.200.28.0/24"]
  disk_size = 50
  disk_type = "vtype-61c3fc5b-f4e9-45b4-8957-8aa7b6029018"
  enable_private_nodes = false
  ssh_key_id= "ssh-f923c53c-cba7-4131-9f86-175d04ae2123"
  security_groups = ["secg-faf05344-fbd6-4f10-80a2-cda08d15ba5e"]
  labels = {
    "test" = "terraform"
  }
  taint {
    key    = "key1"
    value  = "value1"
    effect = "PreferNoSchedule"
  }
}
```