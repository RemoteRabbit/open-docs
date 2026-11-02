# Local module with explicit provider mapping and several input shapes.
module "child" {
  source = "./modules/child"

  providers = {
    aws         = aws.audit
    aws.replica = aws.replica
  }

  name        = var.string_default
  count_value = var.number_default
  enabled     = var.bool_true
  tags        = { Team = "platform", Pipe = "a|b" }
  objects     = var.map_of_objects
}

# Registry module with count.
module "registry_count" {
  source  = "terraform-aws-modules/s3-bucket/aws"
  version = "4.1.2"
  count   = var.bool_false ? 1 : 0

  bucket = "complete-registry"
}

# Registry module with for_each and depends_on.
module "registry_for_each" {
  source   = "cloudposse/label/null"
  version  = ">= 0.25.0, < 1.0.0"
  for_each = toset(var.list_of_strings)

  name = each.key

  depends_on = [module.child, random_pet.names]
}

# Git source pinned to a ref, no version argument.
module "git_source" {
  source = "git::https://github.com/example/terraform-modules.git//network?ref=v1.2.3"
}

# GitHub shorthand and an HTTP archive.
module "github_shorthand" {
  source = "github.com/example/terraform-module-foo"
}

module "http_archive" {
  source = "https://example.com/modules/vpc-module.zip"
}

# OpenTofu early evaluation: non-literal source and version.
module "dynamic_source" {
  source  = "${local.literal_string}/modules/child"
  version = var.string_default
}

# Module call with no inputs at all.
module "no_inputs" {
  source = "./modules/child"
}
