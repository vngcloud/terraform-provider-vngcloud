---
# generated by https://github.com/hashicorp/terraform-plugin-docs
subcategory: "VPC (Virtual Private Cloud)"
page_title: "vngcloud_vserver_subnet Resource - terraform-provider-vngcloud"
description: |-
  Provides a VNG Cloud VServer Subnet resource. This can be used to import, create, and delete.
---

# vngcloud_vserver_subnet (Resource)



## Example Usage

```terraform
resource "vngcloud_vserver_subnet" "subnet" {
    project_id = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    name = "example-subnet"
    cidr = "10.76.1.0/24"
    network_id = "net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    zone_id = "HCM03-1A"
}
```
```terraform
resource "vngcloud_vserver_subnet" "subnet" {
    project_id = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    name = "example-subnet"
    cidr = "10.76.1.0/24"
    network_id = "net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    zone_id = "HCM03-1B"
}
```
```terraform
resource "vngcloud_vserver_subnet" "subnet" {
    project_id = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    name = "example-subnet"
    cidr = "10.76.1.0/24"
    network_id = "net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
    zone_id = "HCM03-1C"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String, Required) Name of Subnet.
- **network_id** (String, Required) ID of Network.
- **cidr** (String, Required) Classless Inter-Domain Routing of Subnet. It must be contained in its network.
- **project_id** (String, Required) ID of Project
- **interface_acl_policy_uuid** (String, Optional) ID of Interface ACL Policy.
- **route_table_uuid** (String, Optional) ID of route table. 
- **zone_id** (String, Optional) ID of Zone. If not specified, the default zone will be used

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

- **id** (String) The ID of this resource.


