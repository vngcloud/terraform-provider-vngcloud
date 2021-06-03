variable "client_id" {
  type = string
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
variable "client_secret" {
  type = string
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
variable "project_id" {
  type = string
  default = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "image_name" {
  type = string
  default = "1-Ubuntu-18.04x64"
}
variable "flavor_zone_name" {
  type = string
  default = "General v1 Instances"
}
variable "flavor_name" {
  type = string
  default = "v1.small1x1.b100"
}
variable "volume_type_name" {
  type = string
  default = "SSD-IOPS3000"
}
variable "root_disk_size" {
  type = number
  default = 20
}
variable "data_disk_size" {
  type = number
  default = 50
}
variable "network_id" {
  type = string
  default = "net-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "subnet_id" {
  type = string
  default = "sub-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "ssh_key_id" {
  type = string
  default = "ssh-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "security_group_id_list" {
  type = list(string)
  default = ["secg-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"]
}

variable "server_count" {
  type = number
  default = 0
}
