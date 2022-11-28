variable "client_id" {
  type    = string
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
variable "client_secret" {
  type    = string
  default = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
variable "project_id" {
  type    = string
  default = "pro-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"
}
variable "flavor_id" {
  type    = string
  default = "flav-eb756bc4-b820-49d7-97bf-cfc8335ad518"
}
variable "image_id" {
  type    = string
  default = "img-a34d639b-e070-46ff-8b91-addf4fac45b4"
}
variable "volume_type_name" {
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
  default = ["secg-xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"]
}

variable "server_count" {
  type    = number
  default = 0
}

variable "server_group_policy_name" {
  type    = string
  default = "SOFT AFFINITY"
}
