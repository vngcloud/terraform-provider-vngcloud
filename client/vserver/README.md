# Go API client for vserver

Api Documentation

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```golang
import "./vserver"
```

## Documentation for API Endpoints

All URIs are relative to *https://api.vngcloud.tech*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*FlavorRestControllerApi* | [**GetFlavorUsingGET**](docs/FlavorRestControllerApi.md#getflavorusingget) | **Get** /v1/{project_id}/flavors/{flavor_id} | getFlavor
*FlavorRestControllerApi* | [**ListFlavorUsingGET**](docs/FlavorRestControllerApi.md#listflavorusingget) | **Get** /v1/{project_id}/{flavor_zone_id}/flavors | listFlavor
*FlavorZoneRestControllerApi* | [**GetFlavorZoneUsingGET**](docs/FlavorZoneRestControllerApi.md#getflavorzoneusingget) | **Get** /v1/{project_id}/flavor_zones/{flavor_zone_id} | getFlavorZone
*FlavorZoneRestControllerApi* | [**ListFlavorZoneUsingGET**](docs/FlavorZoneRestControllerApi.md#listflavorzoneusingget) | **Get** /v1/{project_id}/flavor_zones | listFlavorZone
*ImageRestControllerApi* | [**ListOSImageUsingGET**](docs/ImageRestControllerApi.md#listosimageusingget) | **Get** /v1/{project_id}/images/os | listOSImage
*NetworkAclRestControllerApi* | [**CreateNetworkAclUsingPOST**](docs/NetworkAclRestControllerApi.md#createnetworkaclusingpost) | **Post** /v1/{project_id}/policies | Create network-acl
*NetworkAclRestControllerApi* | [**DeleteNetworkAclUsingDELETE**](docs/NetworkAclRestControllerApi.md#deletenetworkaclusingdelete) | **Delete** /v1/{project_id}/policies/{uuid} | Delete Network-acl
*NetworkAclRestControllerApi* | [**GetNetworkAclUsingGET**](docs/NetworkAclRestControllerApi.md#getnetworkaclusingget) | **Get** /v1/{project_id}/policies/{uuid} | Get network-acl by uuid
*NetworkAclRestControllerApi* | [**ListNetworkAclUsingGET**](docs/NetworkAclRestControllerApi.md#listnetworkaclusingget) | **Get** /v1/{project_id}/policies | List network-acl
*NetworkAclRestControllerApi* | [**UpdateAssociatedSubnetsUsingPUT**](docs/NetworkAclRestControllerApi.md#updateassociatedsubnetsusingput) | **Put** /v1/{project_id}/policies/{uuid}/subnets | Update ACL Rule of Network-acl
*NetworkAclRestControllerApi* | [**UpdateRulesUsingPUT**](docs/NetworkAclRestControllerApi.md#updaterulesusingput) | **Put** /v1/{project_id}/policies/{uuid}/rules | Update Inbound/Outbound Rules of Network-acl
*NetworkRestControllerApi* | [**CreateNetworkUsingPOST**](docs/NetworkRestControllerApi.md#createnetworkusingpost) | **Post** /v1/{project_id}/networks | createNetwork
*NetworkRestControllerApi* | [**DeleteNetworkUsingDELETE**](docs/NetworkRestControllerApi.md#deletenetworkusingdelete) | **Delete** /v1/{project_id}/networks | deleteNetwork
*NetworkRestControllerApi* | [**EditNetworkUsingPUT**](docs/NetworkRestControllerApi.md#editnetworkusingput) | **Put** /v1/{project_id}/networks | editNetwork
*NetworkRestControllerApi* | [**GetNetworkUsingGET**](docs/NetworkRestControllerApi.md#getnetworkusingget) | **Get** /v1/{project_id}/networks/{network_id} | getNetwork
*NetworkRestControllerApi* | [**ListNetworkUsingGET**](docs/NetworkRestControllerApi.md#listnetworkusingget) | **Get** /v1/{project_id}/networks | listNetwork
*ProjectRestControllerApi* | [**GetProjectUsingGET**](docs/ProjectRestControllerApi.md#getprojectusingget) | **Get** /v1/projects/{project_id} | getProject
*ProjectRestControllerApi* | [**ListProjectUsingGET**](docs/ProjectRestControllerApi.md#listprojectusingget) | **Get** /v1/projects | listProject
*QuotaRestControllerApi* | [**InitUsedUsingPUT**](docs/QuotaRestControllerApi.md#initusedusingput) | **Put** /v1/{project_id}/quota/init-used | initUsed
*QuotaRestControllerApi* | [**ListByUserUsingGET**](docs/QuotaRestControllerApi.md#listbyuserusingget) | **Get** /v1/{project_id}/quota/user/{user_id} | listByUser
*QuotaRestControllerApi* | [**ListQuotaUsedUsingGET**](docs/QuotaRestControllerApi.md#listquotausedusingget) | **Get** /v1/{project_id}/quote-used | listQuotaUsed
*QuotaRestControllerApi* | [**ListUsingGET**](docs/QuotaRestControllerApi.md#listusingget) | **Get** /v1/{project_id}/quota | list
*QuotaRestControllerApi* | [**UpdateUsingPUT**](docs/QuotaRestControllerApi.md#updateusingput) | **Put** /v1/{project_id}/quota | update
*RouteTableControllerApi* | [**CreateRouteTableUsingPOST**](docs/RouteTableControllerApi.md#createroutetableusingpost) | **Post** /v1/{project_id}/route-table | Create route-table
*RouteTableControllerApi* | [**DeleteRouteTableUsingDELETE**](docs/RouteTableControllerApi.md#deleteroutetableusingdelete) | **Delete** /v1/{project_id}/route-table/{uuid} | Delete Route Table
*RouteTableControllerApi* | [**DeleteRouteTablesUsingDELETE**](docs/RouteTableControllerApi.md#deleteroutetablesusingdelete) | **Delete** /v1/{project_id}/route-table | Delete all route-tables
*RouteTableControllerApi* | [**GetRouteTableUsingGET**](docs/RouteTableControllerApi.md#getroutetableusingget) | **Get** /v1/{project_id}/route-table/{uuid} | Get route-table by uuid
*RouteTableControllerApi* | [**ListRouteTablesUsingGET**](docs/RouteTableControllerApi.md#listroutetablesusingget) | **Get** /v1/{project_id}/route-table | List route-tables
*RouteTableControllerApi* | [**UpdateRouteTableRoutesUsingPUT**](docs/RouteTableControllerApi.md#updateroutetableroutesusingput) | **Put** /v1/{project_id}/route-table/{uuid}/routes | Add/remove routes of route-table
*RouteTableControllerApi* | [**UpdateRouteTableSubnetsUsingPUT**](docs/RouteTableControllerApi.md#updateroutetablesubnetsusingput) | **Put** /v1/{project_id}/route-table/{uuid}/subnets | Add/remove subnets of route-table
*SecgroupRestControllerApi* | [**CreateSecgroupUsingPOST**](docs/SecgroupRestControllerApi.md#createsecgroupusingpost) | **Post** /v1/{project_id}/secgroups | createSecgroup
*SecgroupRestControllerApi* | [**DeleteSecgroupUsingDELETE**](docs/SecgroupRestControllerApi.md#deletesecgroupusingdelete) | **Delete** /v1/{project_id}/secgroups | deleteSecgroup
*SecgroupRestControllerApi* | [**EditSecgroupUsingPUT**](docs/SecgroupRestControllerApi.md#editsecgroupusingput) | **Put** /v1/{project_id}/secgroups | editSecgroup
*SecgroupRestControllerApi* | [**GetSecgroupUsingGET**](docs/SecgroupRestControllerApi.md#getsecgroupusingget) | **Get** /v1/{project_id}/secgroups/{secgroup_id} | getSecgroup
*SecgroupRestControllerApi* | [**ListSecgroupByInstanceUsingGET**](docs/SecgroupRestControllerApi.md#listsecgroupbyinstanceusingget) | **Get** /v1/{project_id}/secgroups/vm-ids/{vm_id} | listSecgroupByInstance
*SecgroupRestControllerApi* | [**ListSecgroupUsingGET**](docs/SecgroupRestControllerApi.md#listsecgroupusingget) | **Get** /v1/{project_id}/secgroups | listSecgroup
*SecgroupRuleRestControllerApi* | [**CreateSecgroupRuleUsingPOST**](docs/SecgroupRuleRestControllerApi.md#createsecgroupruleusingpost) | **Post** /v1/{project_id}/secgroup-rules | createSecgroupRule
*SecgroupRuleRestControllerApi* | [**DeleteSecgroupRuleUsingDELETE**](docs/SecgroupRuleRestControllerApi.md#deletesecgroupruleusingdelete) | **Delete** /v1/{project_id}/secgroup-rules | deleteSecgroupRule
*SecgroupRuleRestControllerApi* | [**GetSecgroupRuleUsingGET**](docs/SecgroupRuleRestControllerApi.md#getsecgroupruleusingget) | **Get** /v1/{project_id}/secgroup-rules/{secgroup_rule_id} | getSecgroupRule
*SecgroupRuleRestControllerApi* | [**ListSecgroupRuleBySecgroupUsingGET**](docs/SecgroupRuleRestControllerApi.md#listsecgrouprulebysecgroupusingget) | **Get** /v1/{project_id}/secgroup-rules/secgroup-ids/{secgroup_id} | listSecgroupRuleBySecgroup
*ServerGroupRestControllerApi* | [**CreateServerGroupUsingPOST**](docs/ServerGroupRestControllerApi.md#createservergroupusingpost) | **Post** /v1/{project_id}/server-groups | Create server group
*ServerGroupRestControllerApi* | [**DeleteServerGroupUsingDELETE**](docs/ServerGroupRestControllerApi.md#deleteservergroupusingdelete) | **Delete** /v1/{project_id}/server-groups | Delete server group
*ServerGroupRestControllerApi* | [**GetServerGroupUsingGET**](docs/ServerGroupRestControllerApi.md#getservergroupusingget) | **Get** /v1/{project_id}/server-groups/{server_group_id} | Get server group
*ServerGroupRestControllerApi* | [**ListServerGroupPolicyUsingGET**](docs/ServerGroupRestControllerApi.md#listservergrouppolicyusingget) | **Get** /v1/{project_id}/server-groups/policies | List server group policy
*ServerGroupRestControllerApi* | [**ListServerGroupUsingGET**](docs/ServerGroupRestControllerApi.md#listservergroupusingget) | **Get** /v1/{project_id}/server-groups | List server group
*ServerGroupRestControllerApi* | [**UpdateServerGroupUsingPUT**](docs/ServerGroupRestControllerApi.md#updateservergroupusingput) | **Put** /v1/{project_id}/server-groups | Update server group
*ServerRestControllerApi* | [**CreateServerUsingPOST**](docs/ServerRestControllerApi.md#createserverusingpost) | **Post** /v1/{project_id}/servers | Create server
*ServerRestControllerApi* | [**DeleteServerUsingDELETE**](docs/ServerRestControllerApi.md#deleteserverusingdelete) | **Delete** /v1/{project_id}/servers | Delete Server
*ServerRestControllerApi* | [**GetServerUsingGET**](docs/ServerRestControllerApi.md#getserverusingget) | **Get** /v1/{project_id}/servers/{server_id} | Get server by id
*ServerRestControllerApi* | [**ListServerUsingGET**](docs/ServerRestControllerApi.md#listserverusingget) | **Get** /v1/{project_id}/servers | List server
*ServerRestControllerApi* | [**RebootServerUsingPUT**](docs/ServerRestControllerApi.md#rebootserverusingput) | **Put** /v1/{project_id}/servers/reboot | Reboot server
*ServerRestControllerApi* | [**ResizeServerUsingPUT**](docs/ServerRestControllerApi.md#resizeserverusingput) | **Put** /v1/{project_id}/servers/resize | Change flavor of server
*ServerRestControllerApi* | [**StartServerUsingPUT**](docs/ServerRestControllerApi.md#startserverusingput) | **Put** /v1/{project_id}/servers/start | Start server
*ServerRestControllerApi* | [**StopServerUsingPUT**](docs/ServerRestControllerApi.md#stopserverusingput) | **Put** /v1/{project_id}/servers/stop | Stop server
*ServerRestControllerApi* | [**UpdateSecGroupServerUsingPUT**](docs/ServerRestControllerApi.md#updatesecgroupserverusingput) | **Put** /v1/{project_id}/servers/update_sec_group | Update SecGroups of server
*SshKeyRestControllerApi* | [**CreateSSHKeyUsingPOST**](docs/SshKeyRestControllerApi.md#createsshkeyusingpost) | **Post** /v1/{project_id}/ssh_keys | createSSHKey
*SshKeyRestControllerApi* | [**DeleteSSHKeyUsingDELETE**](docs/SshKeyRestControllerApi.md#deletesshkeyusingdelete) | **Delete** /v1/{project_id}/ssh_keys | deleteSSHKey
*SshKeyRestControllerApi* | [**GetSSHKeyUsingGET**](docs/SshKeyRestControllerApi.md#getsshkeyusingget) | **Get** /v1/{project_id}/ssh_keys/{ssh_key_id} | getSSHKey
*SshKeyRestControllerApi* | [**ImportSSHKeyUsingPOST**](docs/SshKeyRestControllerApi.md#importsshkeyusingpost) | **Post** /v1/{project_id}/ssh_keys/import | importSSHKey
*SshKeyRestControllerApi* | [**ListSSHKeyUsingGET**](docs/SshKeyRestControllerApi.md#listsshkeyusingget) | **Get** /v1/{project_id}/ssh_keys | listSSHKey
*SubnetRestControllerApi* | [**CreateSubnetUsingPOST**](docs/SubnetRestControllerApi.md#createsubnetusingpost) | **Post** /v1/{project_id}/subnets | createSubnet
*SubnetRestControllerApi* | [**DeleteSubnetUsingDELETE**](docs/SubnetRestControllerApi.md#deletesubnetusingdelete) | **Delete** /v1/{project_id}/subnets | deleteSubnet
*SubnetRestControllerApi* | [**GetMpPublicInterfaceUsingGET**](docs/SubnetRestControllerApi.md#getmppublicinterfaceusingget) | **Get** /v1/{project_id}/subnets/mp-public-interfaces | getMpPublicInterface
*SubnetRestControllerApi* | [**GetSubnetUsingGET**](docs/SubnetRestControllerApi.md#getsubnetusingget) | **Get** /v1/{project_id}/subnets/{subnet_id} | getSubnet
*SubnetRestControllerApi* | [**ListSubnetsByNetworkUsingGET**](docs/SubnetRestControllerApi.md#listsubnetsbynetworkusingget) | **Get** /v1/{project_id}/subnets/networks/{network_id} | listSubnetsByNetwork
*VolumeRestControllerApi* | [**AttachVolumeUsingPUT**](docs/VolumeRestControllerApi.md#attachvolumeusingput) | **Put** /v1/{project_id}/volumes/attach | Attach volume into the server
*VolumeRestControllerApi* | [**CreateVolumeUsingPOST**](docs/VolumeRestControllerApi.md#createvolumeusingpost) | **Post** /v1/{project_id}/volumes | Create volume
*VolumeRestControllerApi* | [**DeleteVolumeUsingDELETE**](docs/VolumeRestControllerApi.md#deletevolumeusingdelete) | **Delete** /v1/{project_id}/volumes | Delete volume
*VolumeRestControllerApi* | [**DetachVolumeUsingPUT**](docs/VolumeRestControllerApi.md#detachvolumeusingput) | **Put** /v1/{project_id}/volumes/detach | Detach volume into the server
*VolumeRestControllerApi* | [**GetVolumeUsingGET**](docs/VolumeRestControllerApi.md#getvolumeusingget) | **Get** /v1/{project_id}/volumes/{volume_id} | Get volume by id
*VolumeRestControllerApi* | [**ListVolumeUsingGET**](docs/VolumeRestControllerApi.md#listvolumeusingget) | **Get** /v1/{project_id}/volumes | List volume
*VolumeRestControllerApi* | [**ResizeVolumeUsingPUT**](docs/VolumeRestControllerApi.md#resizevolumeusingput) | **Put** /v1/{project_id}/volumes/resize | Resize volume
*VolumeTypeRestControllerApi* | [**GetVolumeTypeUsingGET**](docs/VolumeTypeRestControllerApi.md#getvolumetypeusingget) | **Get** /v1/{project_id}/volume_types/{volume_type_id} | getVolumeType
*VolumeTypeRestControllerApi* | [**ListVolumeTypeUsingGET**](docs/VolumeTypeRestControllerApi.md#listvolumetypeusingget) | **Get** /v1/{project_id}/{volume_type_zone_id}/volume_types | listVolumeType
*VolumeTypeZoneRestControllerApi* | [**GetVolumeTypeZoneUsingGET**](docs/VolumeTypeZoneRestControllerApi.md#getvolumetypezoneusingget) | **Get** /v1/{project_id}/volume_type_zones/{volume_type_zone_id} | getVolumeTypeZone
*VolumeTypeZoneRestControllerApi* | [**ListVolumeTypeZoneUsingGET**](docs/VolumeTypeZoneRestControllerApi.md#listvolumetypezoneusingget) | **Get** /v1/{project_id}/volume_type_zones | listVolumeTypeZone


## Documentation For Models

 - [AclPolicyRule](docs/AclPolicyRule.md)
 - [AttachVolumeRequest](docs/AttachVolumeRequest.md)
 - [BaseResponse](docs/BaseResponse.md)
 - [ChangeSecGroupRequest](docs/ChangeSecGroupRequest.md)
 - [CreateNetworkAclRequest](docs/CreateNetworkAclRequest.md)
 - [CreateNetworkRequest](docs/CreateNetworkRequest.md)
 - [CreateRouteTableRequest](docs/CreateRouteTableRequest.md)
 - [CreateSecurityGroupRequest](docs/CreateSecurityGroupRequest.md)
 - [CreateSecurityGroupRuleRequest](docs/CreateSecurityGroupRuleRequest.md)
 - [CreateServerGroupRequest](docs/CreateServerGroupRequest.md)
 - [CreateServerRequest](docs/CreateServerRequest.md)
 - [CreateSshKeyRequest](docs/CreateSshKeyRequest.md)
 - [CreateSubnetRequest](docs/CreateSubnetRequest.md)
 - [CreateVolumeRequest](docs/CreateVolumeRequest.md)
 - [DeleteNetworkRequest](docs/DeleteNetworkRequest.md)
 - [DeleteSecurityGroupRequest](docs/DeleteSecurityGroupRequest.md)
 - [DeleteSecurityGroupRuleRequest](docs/DeleteSecurityGroupRuleRequest.md)
 - [DeleteServerGroupRequest](docs/DeleteServerGroupRequest.md)
 - [DeleteServerRequest](docs/DeleteServerRequest.md)
 - [DeleteSubnetRequest](docs/DeleteSubnetRequest.md)
 - [DeleteVolumeRequest](docs/DeleteVolumeRequest.md)
 - [DetachVolumeRequest](docs/DetachVolumeRequest.md)
 - [EditSecurityGroupRequest](docs/EditSecurityGroupRequest.md)
 - [ElasticIpEntity](docs/ElasticIpEntity.md)
 - [Flavor](docs/Flavor.md)
 - [FlavorResponse](docs/FlavorResponse.md)
 - [FlavorZone](docs/FlavorZone.md)
 - [FlavorZoneResponse](docs/FlavorZoneResponse.md)
 - [ImportSshKeyRequest](docs/ImportSshKeyRequest.md)
 - [InterfaceNetworkInterface](docs/InterfaceNetworkInterface.md)
 - [InterfaceProjectQuota](docs/InterfaceProjectQuota.md)
 - [InterfaceProjectQuotaResponse](docs/InterfaceProjectQuotaResponse.md)
 - [InterfacePublicIp](docs/InterfacePublicIp.md)
 - [InterfacePublicSubnet](docs/InterfacePublicSubnet.md)
 - [ListRouteTablesResponse](docs/ListRouteTablesResponse.md)
 - [Network](docs/Network.md)
 - [NetworkAclListResponse](docs/NetworkAclListResponse.md)
 - [NetworkAclModel](docs/NetworkAclModel.md)
 - [NetworkAclResponse](docs/NetworkAclResponse.md)
 - [NetworkResponse](docs/NetworkResponse.md)
 - [OsImage](docs/OsImage.md)
 - [OsImageBaseResponse](docs/OsImageBaseResponse.md)
 - [ProjectInfo](docs/ProjectInfo.md)
 - [ProjectResponse](docs/ProjectResponse.md)
 - [Quota](docs/Quota.md)
 - [QuotaResponse](docs/QuotaResponse.md)
 - [ResizeServerRequest](docs/ResizeServerRequest.md)
 - [ResizeVolumeRequest](docs/ResizeVolumeRequest.md)
 - [RouteModel](docs/RouteModel.md)
 - [RouteTableModel](docs/RouteTableModel.md)
 - [RouteTableResponse](docs/RouteTableResponse.md)
 - [RouteTableUpdateRouteRequest](docs/RouteTableUpdateRouteRequest.md)
 - [RouteTableUpdateSubnetRequest](docs/RouteTableUpdateSubnetRequest.md)
 - [SdnSshKeyDeleteRequest](docs/SdnSshKeyDeleteRequest.md)
 - [Secgroup](docs/Secgroup.md)
 - [SecgroupResponse](docs/SecgroupResponse.md)
 - [SecgroupRule](docs/SecgroupRule.md)
 - [SecgroupRuleResponse](docs/SecgroupRuleResponse.md)
 - [Server](docs/Server.md)
 - [ServerGroup](docs/ServerGroup.md)
 - [ServerGroupPolicy](docs/ServerGroupPolicy.md)
 - [ServerGroupPolicyResponse](docs/ServerGroupPolicyResponse.md)
 - [ServerGroupResponse](docs/ServerGroupResponse.md)
 - [ServerResponse](docs/ServerResponse.md)
 - [SshKey](docs/SshKey.md)
 - [SshKeyResponse](docs/SshKeyResponse.md)
 - [Subnet](docs/Subnet.md)
 - [SubnetResponse](docs/SubnetResponse.md)
 - [Timestamp](docs/Timestamp.md)
 - [UpdateNetworkAclRulesRequest](docs/UpdateNetworkAclRulesRequest.md)
 - [UpdateNetworkAclSubnetsRequest](docs/UpdateNetworkAclSubnetsRequest.md)
 - [UpdateNetworkRequest](docs/UpdateNetworkRequest.md)
 - [UpdateQuota](docs/UpdateQuota.md)
 - [UpdateQuotaVsgRequest](docs/UpdateQuotaVsgRequest.md)
 - [UpdateServerGroupRequest](docs/UpdateServerGroupRequest.md)
 - [UpdateServerRequest](docs/UpdateServerRequest.md)
 - [UpdatedAclPolicyRule](docs/UpdatedAclPolicyRule.md)
 - [UpdatedRoute](docs/UpdatedRoute.md)
 - [UserService](docs/UserService.md)
 - [Volume](docs/Volume.md)
 - [VolumeResponse](docs/VolumeResponse.md)
 - [VolumeType](docs/VolumeType.md)
 - [VolumeTypeResponse](docs/VolumeTypeResponse.md)
 - [VolumeTypeZone](docs/VolumeTypeZone.md)
 - [VolumeTypeZoneResponse](docs/VolumeTypeZoneResponse.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author



