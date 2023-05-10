terraform {
  required_providers {
    jumpcloud = {
      source = "vincenttjia/jumpcloud"
    }
  }
}

provider "jumpcloud" {
  api_key = "test"
  org_id  = "test"
}