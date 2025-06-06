---
subcategory: "Kubernetes Service"
description: |-
  Manages a VKS NodeGroup resource.
---

# vngcloud_vks_cluster_node_group

Manages a VNGCloud Kubernetes Engine (VKS) Node Group.

To get more information about VKS Node Group, see:

* How-to guides
    * [VKS overview](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/vks-la-gi)
    * [Getting Start with VKS](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/bat-dau-voi-vks)
* Terraform guidance
    * [Using VKS with Terraform](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/su-dung-vks-voi-terraform)

Manages a node group in a VNGCloud Kubernetes Service (VKS) cluster separately from
the cluster control plane. For more information see [Node Group](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/node-groups).

---

# vngcloud_vks_cluster_node_group

## Example Usage - using a separately managed node group (recommended)

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

## Example Usage - 2 node pools, 1 separately managed + the default node group

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "my-vks-cluster"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  node_group {
    name= "my-vks-default"
    ssh_key_id= "ssh-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  }
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  name= "my-vks-node-group"
  ssh_key_id= "ssh-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  cluster_id= vngcloud_vks_cluster.primary.id
}
```

## Argument Reference

* `cluster_id` - (Required) The ID of the Cluster into which you want to create one or more node groups.
* `name` - (Required) The name of the node group. Only letters (a-z), numbers (0-9), and hyphens ('-') are allowed. The length of your input must be between 5 and 15 characters.
* `num_nodes` - (Optional) The number of nodes to create in this cluster's node group. The number of nodes are between 1 and 10 nodes. If `auto_scale_config` is set, `num_nodes` need to be greater than or equal to Minimum node, and less than or equal to Maximum node. If `auto_scale_config` is not set, `num_nodes` need to be greater than or equal to 0, and less than or equal to 10.
* `auto_scale_config` - (Optional) Configuration required by the cluster autoscaler to adjust the size of the node group based on current cluster usage.
    * `min_size` - (Optional) Minimum number of nodes in the Node Group. `min_size` need to be greater than or equal to 0 and less than or equal to 10 and less than or equal to `max_size`.
    * `max_size` - (Optional) Maximum number of nodes in the Node Group. `max_size` need to be greater than or equal to 0 and less than or equal to 10 and greater than or equal to `min_size`.
* `upgrade_config` - (Optional) Specifies node upgrade settings to control how VKS upgrades nodes. The maximum number of nodes upgraded simultaneously is limited to 20.
    * `strategy` - (Optional) The strategy used for node group updates. The only available strategy is SURGE.
    * `max_surge` - (Optional) The number of additional nodes that can be added to the node pool during an upgrade. Increasing `max_surge` allows more nodes to be upgraded simultaneously. It can be set to 0 or greater. By default, an extra temporary node is created during each node upgrade. To minimize costs (though with a higher risk of disruption), consider setting  `max_surge` to 1 and `max_unavailable` to 0.
    * `max_unvailable` - (Optional) The number of nodes that can be unavailable simultaneously during an upgrade. Increasing `max_unavailable` allows more nodes to be upgraded in parallel. It can be set to 0 or greater. To reduce risk for workloads sensitive to disruptions, this approach involves creating a new node pool while temporarily retaining the old nodes. It offers flexible upgrade pacing through batch requests and straightforward rollbacks but comes with higher costs compared to surge upgrades.
* `image_id` - (Optional) Specifies the image you want to use for your node group. You can obtain the Image ID from the VKS Portal or from this [link](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/tham-khao-them/danh-sach-system-image-dang-ho-tro) and enter it in this field.
* `flavor_id` - (Optional) Specifies the flavor you want to use for your node in the node group. You can obtain the Flavor ID from this [link](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/tham-khao-them/danh-sach-flavor-dang-ho-tro) and enter it in this field.
* `subnet_id` (Required) The subnet ID for the node group. You need to create a Subnet on vServer and enter the Subnet's ID in this field. Remember, you subnet that you use for each node group must be in the same VPC.
* `secondary_subnets` (Optional) Specifies additional subnets to be useds in Cilium's VPC Native Routing mode.
* `disk_size` - (Optional) - Specifies the data disk size for new nodes in this node group. Must be between 20 GB and 1000 GB. The default value is "100".
* `disk_type` - (Optional) - Specifies the type of data disk for new nodes in this node group. Currently, SSD disks and NVME disks are available.
* `enabled_encryption_volume` - (Optional) - Enables volume encryption. The default value is false.
* `enable_private_nodes` - (Optional) You can choose the mode that you want your node group works.  The VKS public node groups include worker nodes deployed in public subnets within a VPC.  These worker nodes have public IP addresses and CAN communicate directly with the public internet. The private node groups configuration involves deploying worker nodes within subnets of a VPC, ensuring they cannot directly access the public internet. All outbound traffic from these nodes is routed exclusively through a NAT gateway service. The default value is "false".
* `security_groups` - (Optional) - Specifies the security group for your cluster. A security group acts as a virtual firewall, controlling inbound and outbound traffic for associated resources. You can find the Security Group ID on the vServer Portal and input it here.
* `ssh_key_id` - (Required) - Specifies the SSH key for secure credentials to prove your identity when connecting to the server. You can import a key and get the SSH Key ID on the vServer Portal to input here
* `labels` - (Optional) - Key/value pairs attached to objects like Pods. They specify identifying attributes meaningful to users but do not imply semantics to the core system.
* `taint` - (Optional) - A taint consists of a key, value, and effect, expressed as key=value:effect.
    * `key`- (Required) - The key for the taint. Must be 63 characters or less, using letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (_), and periods (.). Must start and end with a letter, number, or underscore.
    * `value` - (Required) - The value for the taint. Must be 63 characters or less, using letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (_), and periods (.). Must start and end with a letter, number, or underscore.
    * `effect` - (Optional) - The effect for the taint. Accepted values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
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