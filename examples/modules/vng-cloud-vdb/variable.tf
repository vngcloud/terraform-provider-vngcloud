variable "client_id" {
  type = string
  default = "d459ebdf-0607-4f43-8f70-2be9b85c271f"
}
variable "client_secret" {
  type = string
  default = "3775dc35-3e8f-4343-8512-90e275cc84c2"
}

variable "relational_engine_type" {
  type = string
  default = "MySQL"
}

variable "relational_engine_version" {
  type = string
  default = "8.0"
}


variable "memmory_engine_type" {
  type = string
  default = "Redis"
}

variable "memory_engine_version" {
  type = string
  default = "4.0"
}

variable "db_name" {
  type = string
  default = "dbname"
}

variable "user_name" {
  type = string
  default = "username"
}

variable "password" {
  type = string
  default = "abcd1234"
}

variable "redis_password" {
  type = string
  default = "abcd1234abcd1234"
}

variable "subnet_id" {
  type = string
  default = "sub-c6071119-1a44-42cf-bc03-bb43e584242a"
}
