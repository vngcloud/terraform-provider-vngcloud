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

* `cluster_id` - (Required) The Cluster ID which you want to create one or more node group into.
* `name` - (Required) The name of the node group. Only letters (a-z, 0-9, '-') are allowed. Your input data length must be between 5 and 15.
* `num_nodes` - (Optional) The desired number of nodes that the group should launch with initially. The number of nodes are between 1 and 10 nodes. If `auto_scale_config` is set then `num_nodes` must be -1. If `auto_scale_config` is not set then `num_nodes` must be different from -1
* `auto_scale_config` - (Optional) Configuration required by cluster autoscaler to adjust the size of the node group to the current cluster usage.
  * `min_size` - (Optional) Minimum number of nodes in the Node Group. Must be >=0 and <= 10 and <= max_size.
  * `max_size` - (Optional) Maximum number of nodes in the Node Group. Must be >=0 and <= 10 and >= min_size.
* `upgrade_config` - (Optional) Specify node upgrade settings to change how VKS upgrades nodes. The maximum number of nodes upgraded simultaneously is limited to 20.
  * `strategy` - (Optional) Strategy used for node group update. Strategy can only be SURGE.
  * `max_surge` - (Optional) The number of additional nodes that can be added to the node pool during
    an upgrade. Increasing `max_surge` raises the number of nodes that can be upgraded simultaneously.
    Can be set to 0 or greater.  By default, an extra temporary node is generated during each node upgrade. To minimize expenses (although with a higher risk of disruption), consider configuring Max Surge to 1 and Max Unavailable to 0.
  * `max_unvailable` - (Optional) The number of nodes that can be simultaneously unavailable during
    an upgrade. Increasing `max_unavailable` raises the number of nodes that can be upgraded in
    parallel. Can be set to 0 or greater.  To mitigate risk for workloads that are sensitive to disruptions, this approach involves creating a fresh node pool while retaining the old nodes temporarily. It offers flexible upgrade pacing through batch requests and straightforward rollbacks. However, it comes with a higher cost compared to surge upgrades.
* `image_id` - (Optional) The image that you want to use for your node group. You can get the Image's ID on VKS Portal or [here](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/tham-khao-them/danh-sach-system-image-dang-ho-tro) and input to this field.
* `flavor_id` - (Optional) The flavor that you want to use for your node on your node group. You can get the Flavor's ID in [here](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/tham-khao-them/danh-sach-flavor-dang-ho-tro) and input to this field.
* `disk_size` - (Optional) The Size of Data Disk will be used when new nodes are created using this node group. Disk size must be greater than or equal 20 GB and the largest is 5000 GB.
* `disk_type` - (Optional) The Type of Data Disk will be used when new nodes are created using this node group. At that time, we only provide one type SSD Disk.
* `enable_private_nodes` - (Optional) You can choose the mode that you want your node group works.  The VKS public node groups include worker nodes deployed in public subnets within a VPC.  These worker nodes have public IP addresses and CAN communicate directly with the public internet.  The private node groups configuration involves deploying worker nodes within subnets of a VPC, ensuring they cannot directly access the public internet. All outbound traffic from these nodes is routed exclusively through a NAT gateway service.
* `security_groups` - (Optional) The Security Group that you want to use to your Cluster. A Security Group acts as a virtual firewall, controlling the traffic that is allowed to enter and leave the resources associated with it. For example, after you associate a security group with any server, it controls the Inbound and Outbound traffic for that Server. You can get the Security Group ID on vServer Portal and input to this field.
* `ssh_key_id` - (Required) The SSH Key that you want to set of secure credentials that you use to prove your identity when connecting to our Server. You can import a key and get the SSH Key ID on vServer Portal and input to this field.
* `labels` - (Optional) Labels are key/value pairs that are attached to objects such as Pods. Labels are intended to be used to specify identifying attributes of objects that are meaningful and relevant to users, but do not directly imply semantics to the core system.
* `taint` - (Optional) A taint consists of a key, value, and effect. As an argument here, it is expressed as key=value:effect.
  * `key`- (Required) The Key for taint. The string must be 63 characters or less and consist only of letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (_), and periods (.). It must also begin and end with a letter, number, or underscore.
  * `value` - (Required) The Value for taint. The string must be 63 characters or less and consist only of letters (a-z, A-Z), numbers (0-9), hyphens (-), underscores (_), and periods (.). It must also begin and end with a letter, number, or underscore.
  * `effect` - (Optional) The Effect for taint. Accepted values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.