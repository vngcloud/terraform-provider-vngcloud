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
---
## Argument Reference

* `name` - (Required) The name of the cluster. Only letters (a-z, 0-9, '-') are allowed. Your input data length must be between 5 and 20.
* `config` - (Computed) This field represents the Cluster's configuration. You don't need to provide any input for this field when creating a Cluster.
* `description` - (Optional) Description of the cluster. Only letters (a-z, A-Z, 0-9, '@', '.' , '_' , '-' , ' '). Your input data length must be between 0 and 255.
* `version` - (Optional) Specifies the version you wish to use for your Cluster. You can view all available Kubernetes versions [here](https://docs.vngcloud.vn/vng-cloud-document/v/vn/vks/tham-khao-them/phien-ban-ho-tro-kubernetes). The default value is "1.29.1".
* `white_list_node_cidr` - (Optional) Specifies the IP address range that can connect to the control plane. This feature is only functional in Private Node Group mode.
* `enable_private_cluster` (Optional) - Enables the private cluster feature,
  creating a private endpoint on the cluster. The VKS public clusters refer to a type of Kubernetes cluster configuration where the Kubernetes API server endpoint is publicly accessible over the internet. In an VKS public cluster, the API server endpoint is not restricted to private access within a VPC (Virtual Private Cloud) and can be accessed over the public internet. The VKS private clusters are configured to have private access to the Kubernetes API server endpoint. This means that the API server endpoint is only accessible from within a specific Virtual Private Cloud (VPC) and is not exposed to the public internet. Private clusters provide enhanced security by restricting access to the Kubernetes API to resources within the VPC. At this time, the default value of this field is false and we only offer Public Cluster mode. The default value is "false".
* `network_type` - (Optional) The type of network for the cluster. The default value is "CALICO".
* `vpc_id` (Required) The VPC ID for the cluster. You need to create a VPC on vServer and enter the VPC's ID in this field.
* `subnet_id` (Required) The subnet ID for the cluster. You need to create a Subnet on vServer and enter the Subnet's ID in this field.
* `cidr` (Required) Specifies the CIDR block for the cluster. You can enter a private IP CIDR from the following options: 10.0.0.0 - 10.255.0.0, 172.16.0.0 - 172.24.0.0, or 192.168.0.0. The default value is "172.16.0.0/16".
* `enabled_load_balancer_plugin` (Optional) Enables the attachment of load balancers (both network and application) via Kubernetes YAML. The default value is "true".
* `enabled_block_store_csi_plugin`(Optional) Automatically deploys and manages the BlockStore Persistent Disk CSI Driver via Kubernetes YAML. The default value is "true".

