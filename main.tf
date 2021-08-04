terraform {
  required_providers {
    vngcloud = {
      version = "~> 0.2"
      source  = "vngcloud.vn/terraform/vngcloud"
    }
  }
    # backend "s3" {#no locking
    #     bucket = "test"
    #     key    = "vinhph2/key"
    #     region = "HCM01"
    #     access_key = "9d04fd1d93a865758bddd94049a9c7db"
    #     secret_key = "acadd634a233d096b72e6446a7da9136"
    #     endpoint = "https://hcm01.vstorage.vngcloud.vn"
    #     skip_region_validation = true
    #     skip_credentials_validation = true
    # }
    # backend "pg" {#have locking
    #     conn_str = "postgres://vinhph2:Aa123456@103.245.249.61:5432/terraform?sslmode=disable"
    # }
    # backend "etcd" { #no locking
    #     path      = "vinhph2/terraform.tfstate"
    #     endpoints = "http://103.245.249.61:2379/"
    # }
    # backend "consul" {
    #     address = "103.245.249.61:8500"
    #     scheme  = "http"
    #     lock = true #locking
    #     path    = "full/path"
    # }
}

provider "vngcloud" {
    token_url = "https://monitoring-agent.vngcloud.vn/v1/intake/oauth2/token"
    client_id = "T18icmD6LfZNlYNHLcGprBLmFRhtmmbr"
    client_secret = "70UfRV5Mk6ss0tGyFDCzbIVKjBZQSYSW"
    base_url = "https://vserverapi.vngcloud.vn/vserver-gateway"
}

# resource "random_id" "server" {
#   byte_length = 8
# }

# resource "random_password" "password" {
#     length = 16
#     special = true
#     override_special = "!@*"
#     min_special = 1
# }
data "vngcloud_vserver_project" "project" {
    name = "ZxYAwsDx1620620539811"
}
data "vngcloud_vserver_flavor_zone" "flavor_zone" {
    name = "DEV General v1 Instances"
    project_id = data.vngcloud_vserver_project.project.id
} 
data "vngcloud_vserver_flavor" "flavor" {
    name = "dev.v1.small2x2"
    project_id = data.vngcloud_vserver_project.project.id
    flavor_zone_id =  data.vngcloud_vserver_flavor_zone.flavor_zone.id
}
data "vngcloud_vserver_image" "image" {
    name = "1-Ubuntu-18.04x64"
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
    count = 0
    name = "vinhph2-volume"
    size = 20
    #volume_type_id = "vtype-bacd68a4-8758-4fb6-a739-b047665e05d5"
    volume_type_id = data.vngcloud_vserver_volume_type.volume_type.id
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = data.vngcloud_vserver_project.project.id
    lifecycle {
        create_before_destroy = true
    } 
}
# //done
resource "vngcloud_vserver_sshkey" "sshkey" {
    #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
    name = "vinhph2-sshkey"
    public_key = file("/home/ubuntu/.ssh/id_rsa.pub")
    #public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDdKmZvGX63ftbWtwSZ7zjamQlHVuU5GpEquKNyoVqMLLG2DGAZASWm5PDtbFygGbjgLMz+zNnWWxMP+XvOE1JTkmSNRIXvW/gf9jRJzFxbgvveo55wkop4id2AX6s9MBLrqQOsF3HIiMMDpoVr5Bc/sHaPBj9IepaeiiFAIeyZxN7SZVV5M4Mz0vBB6bEsxexOJDWRsy64DrQ2O3qKqoV0zgXRqWcr7taxXZkTz+Boo1l1GpCZ4liQO4iovYlkuKj8nx2d9wI+37vWrTVvHkGHaUgmuPuFIWGl+qn17ME7HPePGEpKjfH3asHBI1mT2eoxRhii8K/rvhwywXSiM/o3 ubuntu@localhost.localdomain"
    lifecycle {
        create_before_destroy = false
    }
}
# resource "vngcloud_vserver_network" "network" {
#     count = 1
#     #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
#     project_id = data.vngcloud_vserver_project.project.id
#     name = "vinhph2-network"
#     cidr = "10.76.0.0/16"
#     lifecycle {
#         create_before_destroy = true
#     }
# }
# //done
# resource "vngcloud_vserver_subnet" "subnet" {
#     count = 1
#     #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
#     project_id = data.vngcloud_vserver_project.project.id
#     name = "vinhph2-subnet"
#     cidr = "10.76.${count.index}.0/24"
#     network_id = vngcloud_vserver_network.network[0].id
#     #network_id = "net-26afee25-75db-4ce2-8e13-8e0d65d6c31c"
#     lifecycle {
#         create_before_destroy = true
#     }
# }
# //done
# resource "vngcloud_vserver_secgroup" "secgroup" {
#     count = 1
#     #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
#     project_id = data.vngcloud_vserver_project.project.id
#     name = "vinhph2-secgroup"
#     lifecycle {
#         create_before_destroy = true
#     }
# }
# //done
# resource "vngcloud_vserver_secgrouprule" "secgrouprule" {
#     count = 0
#     #project_id = "pro-462803f3-6858-466f-bf05-df2b33faa360"
#     project_id = data.vngcloud_vserver_project.project.id
#     direction ="ingress"
#     ethertype ="IPv4"
#     port_range_max = 65535
#     port_range_min = count.index
#     protocol = "TCP"
#     remote_ip_prefix = "0.0.0.0/0"
#     //security_group_id = "secg-ce0669a5-28c8-4263-8a0a-01adc93fc5a3"
#     security_group_id = vngcloud_vserver_secgroup.secgroup[0].id
#     # lifecycle {
#     #     create_before_destroy = true
#     # }
# }
resource "vngcloud_vserver_server" "server-master" {
    encryption_volume   = false
    flavor_id           = data.vngcloud_vserver_flavor.flavor.id
    image_id            = data.vngcloud_vserver_image.image.id
    name                = "vinhph2-test-cluster-master"
    network_id          = "net-a52567c5-e840-43fa-a5d6-68bf6a9379e7"
    project_id          = data.vngcloud_vserver_project.project.id
    subnet_id           = "sub-a5e4f8e9-e99e-498c-b99a-6b4720cf5c6f"
    ssh_key = vngcloud_vserver_sshkey.sshkey.id
    root_disk_size = 40
    root_disk_type_id = data.vngcloud_vserver_volume_type.volume_type.id
    attach_floating = true
    //expire_password = true
    # https://www.terraform.io/docs/language/resources/provisioners/connection.html
    connection {
        type     = "ssh"
        user     = "stackops"
        host     = vngcloud_vserver_server.server-master.internal_interfaces[0].floating_ip
        port = 234
        # ssh key must remove passpharse before use. Or you can add the desired ssh key to the ssh-agent with ssh-add ~/.ssh/id_rsa and then set the agent field in connection stanza to:
        # agent       = true
        # https://stackoverflow.com/questions/112396/how-do-i-remove-the-passphrase-for-the-ssh-key-without-having-to-create-a-new-ke
        private_key = "${file("~/.ssh/id_rsa")}"
    }
    // to copy file form local to remote https://www.terraform.io/docs/language/resources/provisioners/file.html
    provisioner "file" {
      source      = "~/.ssh/id_rsa"
      destination = "~/.ssh/id_rsa"
    }
    provisioner "file" {
      source      = "/home/ubuntu/CLS/terraform-provider-vngcloud/k8s-create-cluster.sh"
      destination = "~/k8s-create-cluster.sh"
    }
    # https://www.terraform.io/docs/language/resources/provisioners/remote-exec.html
    provisioner "remote-exec" {
        inline = [
            "chmod u+x k8s-create-cluster.sh",
        "bash ./k8s-create-cluster.sh '${vngcloud_vserver_server.server-master.internal_interfaces[0].fixed_ip} ${vngcloud_vserver_server.server-minion[0].internal_interfaces[0].fixed_ip} ${vngcloud_vserver_server.server-minion[1].internal_interfaces[0].fixed_ip}'",
        ]
        
    }
}

resource "vngcloud_vserver_server" "server-minion" {
    count = 2
    encryption_volume   = false
    flavor_id           = data.vngcloud_vserver_flavor.flavor.id
    image_id            = data.vngcloud_vserver_image.image.id
    name                = "vinhph2-test-cluster-minion-${count.index}"
    network_id          = "net-a52567c5-e840-43fa-a5d6-68bf6a9379e7"
    project_id          = data.vngcloud_vserver_project.project.id
    subnet_id           = "sub-a5e4f8e9-e99e-498c-b99a-6b4720cf5c6f"
    ssh_key = vngcloud_vserver_sshkey.sshkey.id
    root_disk_size = 40
    root_disk_type_id = data.vngcloud_vserver_volume_type.volume_type.id
    attach_floating = true
    //expire_password = true
    # https://www.terraform.io/docs/language/resources/provisioners/connection.html
    # connection {
    #     type     = "ssh"
    #     user     = "stackops"
    #     host     = vngcloud_vserver_server.server.internal_interfaces[0].floating_ip
    #     port = 234
    #     # ssh key must remove passpharse before use. Or you can add the desired ssh key to the ssh-agent with ssh-add ~/.ssh/id_rsa and then set the agent field in connection stanza to:
    #     # agent       = true
    #     # https://stackoverflow.com/questions/112396/how-do-i-remove-the-passphrase-for-the-ssh-key-without-having-to-create-a-new-ke
    #     private_key = "${file("~/.ssh/id_rsa")}"
    # }
    # // to copy file form local to remote https://www.terraform.io/docs/language/resources/provisioners/file.html
    # provisioner "file" {
    #   source      = "~/.ssh/id_rsa"
    #   destination = "~/.ssh/id_rsa"
    # }

    # # https://www.terraform.io/docs/language/resources/provisioners/remote-exec.html
    # provisioner "remote-exec" {
    #     inline = [
    #     "mkdir vngcloud",
    #     "echo 'test remote-exec' > vngcloud/vinhph2.txt",
    #     ]
        
    # }
}

# resource "vngcloud_vserver_volume_attach" "attach_volume" {
#     count = 0
#     project_id = data.vngcloud_vserver_project.project.id
#     volume_id = vngcloud_vserver_volume.volume[0].id
#     instance_id = vngcloud_vserver_server.server[0].id
# }

# provisioner "test" {
#   connection {
#     type     = "ssh"
#     user     = "stackops"
#     password = "${var.root_password}"
#     host     = "${var.host}"
#   }
# }
