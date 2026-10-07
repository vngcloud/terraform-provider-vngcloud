# Self-contained example: VKS cluster with control-plane logging config.
#
# This directory is NOT loaded by ../run_mode.sh (which only runs terraform in
# examples/). Copy the resource block into your own config, or run terraform
# directly from inside this directory for a manual test.

terraform {
  required_providers {
    vngcloud = {
      source  = "vngcloud/vngcloud"
      version = "1.0.0"
    }
  }
}

provider "vngcloud" {
  client_id     = var.client_id
  client_secret = var.client_secret
  # vks_base_url defaults to https://vks.api.vngcloud.vn
}

variable "client_id" {
  type = string
}

variable "client_secret" {
  type      = string
  sensitive = true
}

variable "opensearch_password" {
  type      = string
  sensitive = true
  default   = ""
}

# --- OpenSearch sink ---------------------------------------------------------
resource "vngcloud_vks_cluster" "opensearch_logging" {
  name      = "cluster-logging-os"
  cidr      = "172.16.0.0/16"
  vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
  subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"

  logging_config {
    enabled               = true
    type                  = "OPENSEARCH"
    opensearch_cluster_id = "os-abc123"
    username              = "admin"
    password              = var.opensearch_password
    components            = ["AUDIT", "API_SERVER"]
  }
}

# --- Kafka sink --------------------------------------------------------------
# resource "vngcloud_vks_cluster" "kafka_logging" {
#   name      = "cluster-logging-kafka"
#   cidr      = "172.16.0.0/16"
#   vpc_id    = "net-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
#   subnet_id = "sub-xxxxxxxx-xxxx-xxxxx-xxxx-xxxxxxxxxxxx"
#
#   logging_config {
#     enabled          = true
#     type             = "KAFKA"
#     kafka_cluster_id = "kafka-abc123"
#     kafka_user_id    = "user-xyz"
#     authen_mode      = "SASL"
#     components       = ["AUDIT"]
#   }
# }
