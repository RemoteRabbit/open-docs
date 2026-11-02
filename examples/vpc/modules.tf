# Gateway and interface endpoints so private subnets reach AWS APIs without
# traversing the NAT gateway.
module "endpoints" {
  source  = "terraform-aws-modules/vpc/aws//modules/vpc-endpoints"
  version = "~> 5.8"

  vpc_id             = aws_vpc.this.id
  subnet_ids         = [for s in aws_subnet.private : s.id]
  security_group_ids = []

  endpoints = {
    for svc in var.vpc_endpoints : svc => {
      service         = svc
      route_table_ids = [for rt in aws_route_table.private : rt.id]
      tags            = { Name = "${local.name_prefix}-${svc}" }
    }
  }

  tags = local.tags
}

module "flow_log_bucket" {
  source = "./modules/log-bucket"
  count  = var.flow_logs.enabled && var.flow_logs.destination == "s3" ? 1 : 0

  providers = {
    aws         = aws
    aws.replica = aws.dr
  }

  bucket_prefix   = "${local.name_prefix}-flow-logs"
  replicate       = var.flow_logs.replicate_to_dr
  expiration_days = var.flow_logs.retention_in_days
  account_id      = data.aws_caller_identity.current.account_id
}

module "subnet_labels" {
  source   = "cloudposse/label/null"
  version  = "0.25.0"
  for_each = toset(local.azs)

  namespace   = var.name
  environment = var.environment
  attributes  = [each.key]

  depends_on = [aws_subnet.public, aws_subnet.private]
}
