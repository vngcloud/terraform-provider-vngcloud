terraform {
  required_providers {
    vngcloud = {
      version = "~> 0.0.2"
      source  = "vngcloud/vngcloud"
    }
  }
}
provider "vngcloud" {
    token_url = "https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token"
    client_id = "eQZyB8nzxfz1cCIQBYMNtGh8nRGF0KV9"
    client_secret = "2fj1zljDbLCupfhsDxCkOJId1cuU1ICn"
    base_url = "https://vserverapi.vngcloud.vn/vserver-gateway"
}

data "vngcloud_vserver_project" "project" {
    name = "ZxYAwsDx1620620539811"
}
data "vngcloud_vserver_flavor_zone" "flavor_zone" {
    name = "DEV General v1 Instances"
    project_id = data.vngcloud_vserver_project.project.id
} 
data "vngcloud_vserver_flavor" "flavor" {
    name = "dev.v1.small1x1"
    project_id = data.vngcloud_vserver_project.project.id
    flavor_zone_id =  data.vngcloud_vserver_flavor_zone.flavor_zone.id
}
data "vngcloud_vserver_image" "image" {
    name = "1-Ubuntu-14.04.5x64"
    project_id = data.vngcloud_vserver_project.project.id
    flavor_zone_id =  data.vngcloud_vserver_flavor_zone.flavor_zone.id
}
data "vngcloud_vserver_volume_type_zone" "volume_type_zone" {
    name = "SSD"
    project_id = data.vngcloud_vserver_project.project.id
}
data "vngcloud_vserver_volume_type" "volume_type" {
    name = "SSD-IOPS3000"
    project_id = data.vngcloud_vserver_project.project.id
    volume_type_zone_id = data.vngcloud_vserver_volume_type_zone.volume_type_zone.id
}
resource "vngcloud_vserver_volume" "volume" {
    name = "name-volume"
    size = 20
    #volume_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
    volume_type_id = data.vngcloud_vserver_volume_type.volume_type.id
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    lifecycle {
        create_before_destroy = true
    } 
}
resource "vngcloud_vserver_sshkey" "sshkey" {
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    name = "name-sshkey"
    public_key = file("/home/name/.ssh/id_rsa.pub")
    #public_key = "ssh_public_key"
    lifecycle {
        create_before_destroy = false
    }
}
resource "vngcloud_vserver_network" "network" {
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    name = "name-network"
    cidr = "10.76.0.0/16"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vngcloud_vserver_subnet" "subnet" {
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    name = "name-subnet"
    cidr = "10.76.${count.index}.0/24"
    network_id = vngcloud_vserver_network.network.id
    #network_id = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vngcloud_vserver_secgroup" "secgroup" {
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    name = "name-secgroup"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vngcloud_vserver_secgrouprule" "secgrouprule" {
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    direction ="ingress"
    ethertype ="IPv4"
    port_range_max = 65535
    port_range_min = count.index
    protocol = "TCP"
    remote_ip_prefix = "0.0.0.0/0"
    //security_group_id = "secg-ce0669a5-28c8-4263-8a0a-01adc93fc5a3"
    security_group_id = vngcloud_vserver_secgroup.secgroup.id
    # lifecycle {
    #     create_before_destroy = true
    # }
}
resource "vngcloud_vserver_server" "server"{
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    name = "name-server-${count.index}"
    encryption_volume = false
    attach_floating = false
    #flavor_id = "flav-437f64a6-f55e-4f42-b861-65bcc62420de"
    flavor_id = data.vngcloud_vserver_flavor.flavor.id
    image_id = data.vngcloud_vserver_image.image.id
    network_id = vngcloud_vserver_network.network.id
    #network_id = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
    root_disk_size = 20
    #root_disk_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
    root_disk_type_id = data.vngcloud_vserver_volume_type.volume_type.id
    ssh_key = vngcloud_vserver_sshkey.sshkey.id
    security_group = [vngcloud_vserver_secgroup.secgroup.id]
    subnet_id = vngcloud_vserver_subnet.subnet.id
    #action = "start"
    lifecycle {
        create_before_destroy = true
    }
}

resource "vngcloud_vserver_volume_attach" "attach_volume" {
    project_id = data.vngcloud_vserver_project.project.id
    volume_id = vngcloud_vserver_volume.volume.id
    instance_id = vngcloud_vserver_server.server.id
}
