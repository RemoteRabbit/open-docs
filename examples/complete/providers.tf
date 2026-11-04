provider "aws" {
  region = "us-east-1"
}

# OpenTofu 1.9+ multi-instance provider.
provider "aws" {
  alias    = "by_region"
  for_each = var.regions
  region   = each.key
}

provider "aws" {
  alias  = "audit"
  region = "us-west-2"

  assume_role {
    role_arn = "arn:aws:iam::123456789012:role/audit"
  }
}

provider "random" {}

provider "http" {}
