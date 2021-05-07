terraform {
  required_providers {
    vserver = {
      version = "~> 0.2"
      source  = "vngcloud.com.vn/terraform/vserver"
    }
  }
}

provider "vserver" {
  token_url = "https://monitoring-agent.vngcloud.tech/v1/intake/oauth2/token"
  client_id = "7ozx6Q2woO1LfCS3zqqpPLvuTGrx4KW1"
  client_secret = "WfgMgu1RIVRSJi1wD0WaTA3QzrIdWOqe"
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
    count = 0
    name = "vinhph2"
    size = 20
    volume_type_id = "vtype-6d30737c-08aa-41f8-8bb1-d486c94ccf69"
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
    lifecycle {
        create_before_destroy = true
    } 
}
//done
resource "vserver_sshkey_resource" "sshkey" {
    count = 0
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
    name = "vinhph2-public-key-${count.index}"
    public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDdKmZvGX63ftbWtwSZ7zjamQlHVuU5GpEquKNyoVqMLLG2DGAZASWm5PDtbFygGbjgLMz+zNnWWxMP+XvOE1JTkmSNRIXvW/gf9jRJzFxbgvveo55wkop4id2AX6s9MBLrqQOsF3HIiMMDpoVr5Bc/sHaPBj9IepaeiiFAIeyZxN7SZVV5M4Mz0vBB6bEsxexOJDWRsy64DrQ2O3qKqoV0zgXRqWcr7taxXZkTz+Boo1l1GpCZ4liQO4iovYlkuKj8nx2d9wI+37vWrTVvHkGHaUgmuPuFIWGl+qn17ME7HPePGEpKjfH3asHBI1mT2eoxRhii8K/rvhwywXSiM/o3 ubuntu@localhost.localdomain"
    lifecycle {
        create_before_destroy = false
    }
}
resource "vserver_network_resource" "network" {
    count = 0
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
    name = "vinhph2-terraform-2"
    cidr = "10.79.0.0/16"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vserver_subnet_resource" "subnet" {
    count = 0
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
    name = "vinhph2-subnet1"
    cidr = "10.79.${count.index}.0/24"
    network_id = vserver_network_resource.network[0].id
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vserver_secgroup_resource" "secgroup" {
    count = 0
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
    name = "vinhph2-secgroup2"
    lifecycle {
        create_before_destroy = true
    }
}
//done
resource "vserver_secgrouprule_resource" "secgrouprule" {
    count = 0
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
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
    count = 0
    project_id = "pro-26151c78-0470-4b4c-88a1-6ec41ef29492"
    name = "vinhph2-test"
    encryption_volume = false
    flavor_id = "flav-c41b827e-730a-475f-b181-c0d1c7acc6d4"
    image_id = "img-5c6639ad-8b02-4402-8be4-5bc0c6a57d32"
    network_id = vserver_network_resource.network[0].id
    root_disk_size = 20
    root_disk_type_id = "vtype-6d30737c-08aa-41f8-8bb1-d486c94ccf69"
    ssh_key = vserver_sshkey_resource.sshkey[0].id
    security_group = [vserver_secgroup_resource.secgroup[0].id]
    subnet_id = vserver_subnet_resource.subnet[0].id
    lifecycle {
        create_before_destroy = true
    }
}