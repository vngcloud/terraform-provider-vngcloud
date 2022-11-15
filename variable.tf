variable "client_id" {
  type = string
  default = "9e734abc-7171-4b91-aab1-f652d2147f83"
}
variable "client_secret" {
  type = string
  default = "c945f0d3-b1bf-4343-920c-89f6c36f23a8"
}
variable "project_id" {
  type = string
  default = "pro-11dd35af-8349-4441-b34e-0a969f6aeb1a"
}
variable "image_id" {
  type = string
  default = "img-2aba43f5-281a-4978-81ee-824e68cb0fb4"
}
variable "flavor_id" {
  type = string
  default = "flav-15dba7ef-500a-4d6c-89f9-bc951dead4bd"
}

variable "volume_type_name" {
  type = string
  default = "200"
}
variable "root_disk_size" {
  type = number
  default = 20
}
variable "data_disk_size" {
  type = number
  default = 20
}
variable "network_id" {
  type = string
  default = "net-44c1098c-2271-47bc-bfba-9b504ee167b2"
}
variable "subnet_id" {
  type = string
  default = "sub-1e55d876-cdf8-489f-9000-7ffee722d5aa"
}
variable "ssh_key_id" {
  type = string
  default = "ssh-614594fa-e31a-4a36-ad9c-fc18db327e43"
}
variable "security_group_id_list" {
  type = list(string)
  default = ["secg-dc43929e-4c24-4044-be9b-2d57d4435a55"]
}

variable "server_count" {
  type = number
  default = 1
}
