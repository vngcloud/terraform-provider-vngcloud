terraform {
  required_providers {
    vserver = {
      version = "~> 0.2"
      source  = "vngcloud.com.vn/terraform/vserver"
    }
  }
}

provider "vserver" {
    token_url = "https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token"
    client_id = "eQZyB8nzxfz1cCIQBYMNtGh8nRGF0KV9"
    client_secret = "2fj1zljDbLCupfhsDxCkOJId1cuU1ICn"
    base_url = "https://vserverapi.vngcloud.vn/vserver-gateway"
}

resource "random_id" "server" {
  byte_length = 8
}

resource "random_password" "password" {
    length = 16
    special = true
    override_special = "!@*"
    min_special = 1
}

resource "vserver_volume_resource" "volume" {
    count = 1
    name = "vinhph2-volume"
    size = 20
    volume_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    lifecycle {
        create_before_destroy = true
    } 
}
//done
resource "vserver_sshkey_resource" "sshkey" {
    count = 1
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "vinhph2-sshkey"
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDdKmZvGX63ftbWtwSZ7zjamQlHVuU5GpEquKNyoVqMLLG2DGAZASWm5PDtbFygGbjgLMz+zNnWWxMP+XvOE1JTkmSNRIXvW/gf9jRJzFxbgvveo55wkop4id2AX6s9MBLrqQOsF3HIiMMDpoVr5Bc/sHaPBj9IepaeiiFAIeyZxN7SZVV5M4Mz0vBB6bEsxexOJDWRsy64DrQ2O3qKqoV0zgXRqWcr7taxXZkTz+Boo1l1GpCZ4liQO4iovYlkuKj8nx2d9wI+37vWrTVvHkGHaUgmuPuFIWGl+qn17ME7HPePGEpKjfH3asHBI1mT2eoxRhii8K/rvhwywXSiM/o3 ubuntu@localhost.localdomain"
    lifecycle {
        create_before_destroy = false
    }
}
resource "vserver_network_resource" "network" {
    count = 0
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "vinhph2-network"
    cidr = "10.76.0.0/16"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vserver_subnet_resource" "subnet" {
    count = 1
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "vinhph2-subnet"
    cidr = "10.79.${count.index}.0/24"
    //network_id = vserver_network_resource.network[0].id
    network_id = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vserver_secgroup_resource" "secgroup" {
    count = 1
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "vinhph2-secgroup"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vserver_secgrouprule_resource" "secgrouprule" {
    count = 1
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    direction ="ingress"
    ethertype ="IPv4"
    port_range_max = count.index 
    port_range_min = count.index
    protocol = "TCP"
    remote_ip_prefix = "169.60.${count.index}.0/24"
    //security_group_id = "secg-ce0669a5-28c8-4263-8a0a-01adc93fc5a3"
    security_group_id = vserver_secgroup_resource.secgroup[0].id
    lifecycle {
        create_before_destroy = true
    }
}
resource "vserver_server_resource" "server"{
    count = 1
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "vinhph2-server-${count.index}"
    encryption_volume = false
    flavor_id = "flav-437f64a6-f55e-4f42-b861-65bcc62420de"
    image_id = "img-1c29f7df-fa23-4dd2-bcfb-9de14dee72e7"
    #network_id = vserver_network_resource.network[0].id
    network_id = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
    root_disk_size = 20
    root_disk_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
    ssh_key = vserver_sshkey_resource.sshkey[0].id
    security_group = [vserver_secgroup_resource.secgroup[0].id]
    subnet_id = vserver_subnet_resource.subnet[0].id
    action = "start"
    lifecycle {
        create_before_destroy = true
    }
}
# resource "vserver_volume_attach_resource" "attach_volume" {
#     project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
#     volume_id = vserver_volume_resource.volume[0].id
#     instance_id = vserver_server_resource.server[0].id
# }