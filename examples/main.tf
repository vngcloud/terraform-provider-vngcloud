terraform {
  required_providers {
    vngcloud = {
      source  = "vngcloud/vngcloud"
      version = "1.3.9"
    }
  }

  backend "s3" {
    bucket   = "terraform-test-bucket"
    key      = "terraform.tfstate"
    region   = "hcm04"  # Required dummy value (ignored by vStorage)
 
    # vStorage specific settings
    endpoints = {
      s3 = "https://hcm04.vstorage.vngcloud.vn"
    }
    skip_region_validation      = true
    skip_credentials_validation = true
    skip_metadata_api_check     = true
    skip_requesting_account_id  = true
    use_path_style              = true  # Replaces deprecated force_path_style

    # Credentials are read from environment variables:
    # export AWS_ACCESS_KEY_ID="your-access-key"
    # export AWS_SECRET_ACCESS_KEY="your-secret-key"

  }


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

# module "vlb" {
#   source = "./modules/vng-cloud-vlb"
# }
