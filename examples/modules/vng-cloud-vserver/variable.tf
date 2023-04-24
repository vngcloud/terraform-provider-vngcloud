variable "project_id" {
  type    = string
  default = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "s_general_4x8" {
  type    = string
  default = "flav-05f97524-0410-46a4-87a8-af92aa759231"
}
variable "ubuntu_20_04" {
  type    = string
  default = "img-a34d639b-e070-46ff-8b91-addf4fac45b4"
}
variable "ssd_3000" {
  type    = string
  default = "3000"
}
variable "root_disk_size" {
  type    = number
  default = 20
}
variable "data_disk_size" {
  type    = number
  default = 50
}
variable "network_id" {
  type    = string
  default = "net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "subnet_id" {
  type    = string
  default = "sub-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "ssh_key_id" {
  type    = string
  default = "ssh-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "security_group_id_list" {
  type    = list(string)
  default = [
    "secg-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
  ]
}

variable "server_count" {
  type    = number
  default = 1
}

variable "soft_affinity" {
  type    = string
  default = "SOFT AFFINITY"
}

variable "soft_anti_affinity" {
  type    = string
  default = "SOFT ANTI AFFINITY"
}

variable "user_data_base64_encode" {
  type    = bool
  default = false
}
