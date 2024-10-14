terraform {
  required_providers {
    vngcloud = {
      source  = "vngcloud/vngcloud"
      version = "1.1.3"
    }
  }
  #  backend "s3" {
  #    skip_credentials_validation = true
  #    skip_metadata_api_check = true
  #    skip_region_validation = true
  #    bucket = "bucket-name"
  #    endpoint = "https://hcm01.vstorage.vngcloud.vn/"
  #    key = "terraform.tfstate"
  #    region = "HCM01"
  #    access_key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  #    secret_key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  #  }
}

provider "vngcloud" {
  token_url        = "https://iamapis.vngcloud.vn/accounts-api/v2/auth/token"
  client_id        = var.client_id
  client_secret    = var.client_secret
  vserver_base_url = "https://hcm-3.api.vngcloud.vn/vserver/vserver-gateway"
  vlb_base_url = "https://hcm-3.api.vngcloud.vn/vserver/vlb-gateway"
}

module "vserver" {
  source = "./modules/vng-cloud-vserver"
}

module "k8s" {
  source = "./modules/vng-cloud-k8s"
}

module "vlb" {
  source = "./modules/vng-cloud-vlb"
}
