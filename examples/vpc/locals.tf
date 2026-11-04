locals {
  name_prefix = "${var.name}-${var.environment}"

  azs = coalesce(var.azs, slice(data.aws_availability_zones.available.names, 0, 3))

  tags = merge(var.tags, {
    Module      = "vpc"
    Environment = var.environment
    ManagedBy   = "terraform"
  })

  public_subnets  = { for i, cidr in var.public_subnets : local.azs[i] => cidr }
  private_subnets = { for i, cidr in var.private_subnets : local.azs[i] => cidr }

  create_database_subnets = length(var.database_subnets) > 0
  nat_gateway_count       = var.enable_nat_gateway ? (var.single_nat_gateway ? 1 : length(local.azs)) : 0
  create_vpn              = var.vpn_preshared_key != null

  flow_log_destination_arn = var.flow_logs.destination == "s3" ? module.flow_log_bucket[0].bucket_arn : aws_cloudwatch_log_group.flow_logs[0].arn
}
