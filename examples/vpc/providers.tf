# Primary region provider. Every resource that does not set `provider`
# explicitly lands here.
provider "aws" {
  region = var.region

  default_tags {
    tags = local.tags
  }
}

# Disaster-recovery region, used for the replicated flow-log bucket.
provider "aws" {
  alias  = "dr"
  region = var.dr_region

  default_tags {
    tags = local.tags
  }
}
