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
  name              = "my-vks-cluster"
  cidr              = "172.16.0.0/16"
  az_strategy       = "SINGLE" # or "MULTI"
  vpc_id            = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id         = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  # list_subnet_ids = ["sub-xxxxxxx", "sub-yyyyyyy"] # required if az_strategy = "MULTI" (currently HCM zone only)
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
  name          = "my-vks-cluster"
  cidr          = "172.16.0.0/16"
  az_strategy   = "SINGLE"
  vpc_id        = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id     = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  node_group {
    name        = "my-vks-node-group"
    ssh_key_id  = "ssh-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
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
* `network_type` - (Optional) The type of network for the cluster. The default value is `TIGERA`. You can choose one in many options including `TIGERA`, `CILIUM_OVERLAY`, `CILIUM_NATIVE_ROUTING`.
* `vpc_id` - (Required) The VPC ID for the cluster. You need to create a VPC on vServer and enter the VPC's ID in this field.
* `az_strategy` (Optional) Availability zone strategy: `"SINGLE"` or `"MULTI"`. Default is `"SINGLE"`. **Currently only available in HCM zone.**
* `subnet_id` - (Optional) The subnet ID for the cluster. You need create a Subnet on vServer and put the Subnet's ID on this field. **Required if `az_strategy` is `"SINGLE"`**.
* `list_subnet_ids` (Optional) List of subnet IDs, **required if `az_strategy` is `"MULTI"`**. **Currently only available in HCM zone.**
* `cidr` -  (Required) Specifies the CIDR block for the cluster using `TIGERA` or `CILIUM_OVERLAY` network. You can enter a private IP CIDR from the following options: 10.0.0.0 - 10.255.0.0, 172.16.0.0 - 172.24.0.0, or 192.168.0.0. The default value is "172.16.0.0/16".
* `secondary_subnets` - (Optional, Computed) Secondary subnets used in Cilium's VPC Native Routing mode (`CILIUM_NATIVE_ROUTING`). **The backend now auto-selects this from the cluster's primary subnet — any value you set here is ignored, and the selected CIDR is exported back into state.** It is effectively read-only; you no longer choose it. If the primary subnet has no eligible secondary subnet (none with a netmask ≤ `node_netmask_size`, or all already used by your other clusters/node groups), the create request fails with a `400` from the API.
* `node_netmask_size` - (Optional) Specifies the node CIDR mask size used in Cilium's VPC Native Routing mode. The default value is 25. You can enter a number from the following options: 24, 25, 26.
* `enabled_load_balancer_plugin` - (Optional) Enables/ Disable the attachment of load balancers (both network and application) via Kubernetes YAML. The default value is "true".
* `enabled_block_store_csi_plugin` - (Optional) Enable/ Disable Automatically deploys and manages the BlockStore Persistent Disk CSI Driver via Kubernetes YAML. The default value is "true".
* `auto_upgrade_config` - (Optional) To configure the `auto_upgrade_config` feature for automated maintenance on your cluster, you can use the following attributes:
    * `weekdays` - A list of days of the week when maintenance should occur, e.g., sat,sun (Saturday and Sunday).
    * `time` - The specific time of the day to start maintenance, in 24-hour format, e.g., 21:00 (9 PM).
* `auto_healing_config` - (Optional, Computed) Configures the auto-healing behavior for the cluster. If omitted, the server applies default settings and any previously configured values are preserved. Supports the following attributes:
    * `enable_auto_healing` - (Required) Whether auto-healing is enabled for the cluster.
    * `max_unhealthy` - (Optional) Maximum number or percentage of unhealthy nodes allowed before remediation is triggered. Accepts an absolute integer from 1 to 1000 (e.g. `"3"`) or a percentage from 1% to 100% (e.g. `"20%"`). **Mutually exclusive with `unhealthy_range`** — exactly one of the two must be specified; Terraform enforces this at plan time.
    * `unhealthy_range` - (Optional) Range (inclusive) of unhealthy nodes allowed before remediation is triggered. Format: `"[N-M]"` where N ≤ M, e.g. `"[2-5]"`. **Mutually exclusive with `max_unhealthy`** — exactly one of the two must be specified; Terraform enforces this at plan time.
    * `timeout_unhealthy` - (Optional, Computed) Time in minutes to wait before considering a node unhealthy. Valid range: 1–180. When set to `0` or omitted, the server retains the existing value.
* `poc` - (Optional) Allows the creation of POC cluster.
* `node_group` - (Optional) List of node groups to create inline with the cluster. Each `node_group` block supports:
  * `name` - (Required) Name of the node group.
  * `ssh_key_id` - (Required) The SSH Key ID to use for nodes.
  * `num_nodes` - (Optional) Desired number of nodes. Default is 1.
  * `kubernetes_version` - (Optional) Kubernetes version for the node group. At creation time, the API automatically assigns the cluster's current version — this field does not need to be set when creating. Changing this field on an existing inline node group has no effect — use the standalone `vngcloud_vks_cluster_node_group` resource for version upgrades.
  * `os` - (Optional) Operating system image for nodes (e.g. `"ubuntu-22.04"`). **Changing this forces the entire cluster to be recreated** (since `node_group` is ForceNew on the cluster).
  * `flavor_id` - (Optional) Node flavor ID.
  * `disk_size` - (Optional) Data disk size in GB. Default is 20.
  * `disk_type` - (Optional) Data disk type.
  * `enable_private_nodes` - (Optional) Whether nodes are private. Default is false.
  * `security_groups` - (Optional) List of security group IDs.
  * `auto_scale_config` - (Optional) Autoscaler configuration with `min_size` and `max_size`.
  * `upgrade_config` - (Optional) Upgrade strategy configuration.
  * `labels` - (Optional) Kubernetes labels as key/value map.
  * `taint` - (Optional) List of Kubernetes taints with `key`, `value`, `effect` — order does not matter (the provider ignores pure reordering when comparing against the server). Can be written either as repeated `taint { ... }` blocks or as a `taint = [{ key = "...", value = "...", effect = "NoSchedule" }]` list — both forms are supported (do not mix both in the same resource). Setting `taint = []` removes all taints; this is only expressible with the list form. If omitted, existing taints on the server are preserved.
  * `secondary_subnets` - (Optional) Additional subnets for CILIUM_NATIVE_ROUTING mode. Unlike the cluster-level field, this node-group-level field is still user-provided (you choose it as before).
  * `subnet_id` - (Optional) Subnet ID for nodes.
  * `enabled_encryption_volume` - (Optional) Enable volume encryption. Default is false.
---
### **Clearing all taints on the inline `node_group`**

The inline `node_group`'s `taint` supports both the classic repeated-block syntax and a list
attribute syntax (Terraform's [Attributes as Blocks](https://developer.hashicorp.com/terraform/language/attr-as-blocks)
behavior for this kind of field) — existing configs using `taint { ... }` blocks continue to work
unchanged. Reordering `taint` entries (e.g. because the server returns them in a different order
than declared) does not produce a plan diff. The list attribute form is only needed when you want
to explicitly clear all taints, since omitting every `taint { ... }` block is indistinguishable
from never having configured taints at all — `taint = []` is the only way to express "remove all
taints":

```hcl
# Existing block syntax — still valid, no changes needed
taint {
  key    = "key1"
  value  = "value1"
  effect = "PreferNoSchedule"
}

# Equivalent list attribute syntax
taint = [
  {
    key    = "key1"
    value  = "value1"
    effect = "PreferNoSchedule"
  }
]

# Explicitly clear all taints — only possible with the list form
taint = []
```

A single resource cannot mix both forms for the same `taint` argument — use either `taint { ... }`
blocks or a `taint = [...]` list, not both.

---
### **Some important notes when using VKS with Terraform:**

When using **Terraform** to create a **Cluster** and **Node Group** on the VKS system, if you modify any of the following fields, the system will automatically delete the existing Node Group/Cluster and recreate a new one with the corresponding new parameters. The deletion process will occur before the creation of the new Node Group/Cluster.

* For the resource `vngcloud_vks_cluster`, the fields that, when modified, will cause the system to delete and recreate the Cluster include:
    * `name`&#x20;
    * `description`&#x20;
    * `enable_private_cluster`&#x20;
    * `enable_service_endpoint`&#x20;
    * `network_type`&#x20;
    * `vpc_id`&#x20;
    * `az_strategy`&#x20;
    * `subnet_id`&#x20;
    * `list_subnet_ids`&#x20;
    * `cidr`&#x20;
    * `node_group`&#x20;
    * `node_netmask_size`&#x20;
    * `release_channel`&#x20;

  The following fields are updated **in-place** and do **not** trigger cluster recreation:
    * `version` — cluster upgrade
    * `auto_upgrade_config` — maintenance window
    * `auto_healing_config` — applied via a dedicated PATCH API call; the cluster continues running during the update
    * `secondary_subnets` — backend-owned (auto-selected for `CILIUM_NATIVE_ROUTING`); not settable, never forces recreation
* For the resource `vngcloud_vks_cluster_node_group`, the fields that, when modified, will cause the system to delete and recreate the Node Group include:
    * `cluster_id`&#x20;
    * `name`&#x20;
    * `os`&#x20;
    * `flavor_id`&#x20;
    * `disk_size`&#x20;
    * `disk_type`&#x20;
    * `enable_private_nodes`&#x20;
    * `ssh_key_id`&#x20;
    * `secondary_subnets`&#x20;
    * `enabled_encryption_volume`&#x20;
    * `subnet_id`

  Modifying `kubernetes_version` does **not** recreate the node group — it triggers a rolling upgrade via the upgrade-version API.

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
  network_type = "TIGERA"
  vpc_id    = "net-70ef12d4-d619-43fc-88f0-1c1511683123"
  az_strategy = "SINGLE"
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
  kubernetes_version = "v1.29.1"
  os = "ubuntu"
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

### Example Usage 2 - Create a Private cluster using a private endpoint on VNGCloud with AutoScale mode enabled and the network type is CILIUM OVERLAY, az_strategy is MULTI (**currently HCM zone only**) and a maintenance window set for every Thursday at 9 AM.

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
  az_strategy = "MULTI"
  list_subnet_ids = ["sub-xxxxxxxxxx","sub-yyyyyyyyy"]
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
  kubernetes_version = "v1.29.1"
  os = "ubuntu"
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

### Example Usage 3 - Create a cluster with auto-healing enabled (using `max_unhealthy` as a percentage).

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "cluster-demo"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"

  auto_healing_config {
    enable_auto_healing = true
    max_unhealthy       = "20%"
    timeout_unhealthy   = 10
  }
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  cluster_id = vngcloud_vks_cluster.primary.id
  name       = "nodegroup1"
  num_nodes  = 3
}
```

### Example Usage 4 - Create a cluster with auto-healing enabled (using `unhealthy_range`).

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "cluster-demo"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"

  auto_healing_config {
    enable_auto_healing = true
    unhealthy_range     = "[2-5]"
    timeout_unhealthy   = 15
  }
}

resource "vngcloud_vks_cluster_node_group" "primary" {
  cluster_id = vngcloud_vks_cluster.primary.id
  name       = "nodegroup1"
  num_nodes  = 5
}
```

> **Note:** `max_unhealthy` and `unhealthy_range` are mutually exclusive. Terraform enforces this constraint at plan time — specifying both will produce a validation error.

### Example Usage 5 - Create a private cluster using a private endpoint on VNGCloud with AutoScale mode enabled and the network type is CILIUM VPC NATIVE ROUTING and a maintenance window set for everyday at 11 PM.

```hcl
resource "vngcloud_vks_cluster" "primary" {
  name      = "cluster-demo"
  description = "Cluster create via terraform"
  version = "v1.29.1"
  enable_private_cluster = false
  enable_service_endpoint = false
  network_type = "CILIUM_NATIVE_ROUTING"
  vpc_id    = "net-70ef12d4-d619-43fc-88f0-1c1511683123"
  az_strategy = "SINGLE"
  subnet_id = "sub-0725ef54-a32e-404c-96f2-34745239c123"
  # secondary_subnets is no longer set at the cluster level: the backend auto-selects it from the
  # primary subnet for CILIUM_NATIVE_ROUTING and exports the chosen CIDR back into state.
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
  kubernetes_version = "v1.29.1"
  os = "ubuntu"
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