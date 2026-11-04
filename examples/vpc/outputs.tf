output "vpc_id" {
  value       = aws_vpc.this.id
  description = "The ID of the created VPC."
}

output "vpc_arn" {
  value       = aws_vpc.this.arn
  description = "The ARN of the VPC."
}

output "vpc_cidr_block" {
  value       = aws_vpc.this.cidr_block
  description = "The primary IPv4 CIDR block of the VPC."
}

output "all_cidr_blocks" {
  value       = concat([aws_vpc.this.cidr_block], [for a in aws_vpc_ipv4_cidr_block_association.secondary : a.cidr_block])
  description = "Primary and secondary CIDR blocks associated with the VPC."
}

output "azs" {
  value       = local.azs
  description = "Availability zones the subnets were spread across."
}

output "public_subnet_ids" {
  value       = { for az, s in aws_subnet.public : az => s.id }
  description = "Map of availability zone to public subnet ID."
}

output "private_subnet_ids" {
  value       = { for az, s in aws_subnet.private : az => s.id }
  description = "Map of availability zone to private subnet ID."
}

output "database_subnet_ids" {
  value       = aws_subnet.database[*].id
  description = "List of database subnet IDs."
}

output "database_subnet_group_name" {
  value       = try(aws_db_subnet_group.database[0].name, null)
  description = "Name of the database subnet group, or null when no database subnets were requested."
}

output "nat_gateway_ids" {
  value       = aws_nat_gateway.this[*].id
  description = "NAT gateway IDs, one per AZ (or one total with `single_nat_gateway`)."
}

output "nat_public_ips" {
  value       = aws_nat_gateway.this[*].public_ip
  description = "Elastic IPs attached to the NAT gateways."
  depends_on  = [aws_route_table.private]
}

# Callers wiring additional routes need these.
output "route_table_ids" {
  value = {
    public  = aws_route_table.public.id
    private = { for az, rt in aws_route_table.private : az => rt.id }
  }
  description = "Route table IDs grouped by tier."
}

output "flow_log_id" {
  value       = one(aws_flow_log.this[*].id)
  description = "ID of the VPC flow log, if enabled."
}

output "endpoint_ids" {
  value       = module.endpoints.endpoints
  description = "VPC endpoints created by the endpoints submodule."
}

output "region" {
  value       = data.aws_region.current.name
  description = "Region the VPC was deployed into."
}

output "vpn_preshared_key" {
  value       = var.vpn_preshared_key
  description = "Pre-shared key passed through for the VPN connection."
  sensitive   = true
}
