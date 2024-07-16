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