# Module `vpc`

<!-- VERSION-START -->
This documentation was generated with [open-docs dev](https://github.com/RemoteRabbit/open-docs)
<!-- VERSION-END -->

- **Required core:** `>= 1.5, < 2.0`

## Required Providers

| Name | Source | Version | Position |
|------|--------|---------|----------|
| `aws` | `hashicorp/aws` | `>= 5.0, < 7.0` | [`./versions.tf`](./versions.tf#L5-L8) |
| `random` | `hashicorp/random` | `~> 3.6` | [`./versions.tf`](./versions.tf#L9-L12) |

## Providers

| Name | Alias | Position |
|------|-------|----------|
| `aws` | - | [`./providers.tf`](./providers.tf#L3) |
| `aws` | `dr` | [`./providers.tf`](./providers.tf#L12) |

## Inputs

| Name | Type | Default | Description | Comment | Sensitive | Nullable | Ephemeral | Validation Count | Position |
|------|------|---------|-------------|---------|-----------|----------|-----------|------------------|----------|
| `name` | `string` | - | Name prefix applied to every resource created by this module. | - | - | - | - | 1 | [`./variables.tf`](./variables.tf#L1) |
| `environment` | `string` | `"dev"` | Deployment environment. Drives tagging and a few cost-related defaults. | - | - | - | - | 1 | [`./variables.tf`](./variables.tf#L11) |
| `region` | `string` | - | AWS region to deploy into. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L22) |
| `dr_region` | `string` | `"us-west-2"` | Secondary region used for disaster-recovery copies of flow logs. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L27) |
| `cidr_block` | `string` | `"10.0.0.0/16"` | IPv4 CIDR block for the VPC. | - | - | - | - | 2 | [`./variables.tf`](./variables.tf#L33) |
| `secondary_cidr_blocks` | `list(string)` | `[]` | Additional IPv4 CIDR blocks to associate with the VPC. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L49) |
| `azs` | `list(string)` | `null` | Availability zones to spread subnets across. Defaults to the first three available in the region. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L55) |
| `public_subnets` | `list(string)` | `["10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24"]` | CIDR blocks for public subnets, one per availability zone. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L61) |
| `private_subnets` | `list(string)` | `["10.0.10.0/24", "10.0.11.0/24", "10.0.12.0/24"]` | CIDR blocks for private subnets, one per availability zone. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L67) |
| `database_subnets` | `list(string)` | `[]` | CIDR blocks for isolated database subnets. Leave empty to skip. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L73) |
| `enable_dns_hostnames` | `bool` | `true` | Enable DNS hostnames in the VPC. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L79) |
| `enable_dns_support` | `bool` | `true` | Enable DNS resolution in the VPC. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L85) |
| `enable_nat_gateway` | `bool` | `true` | Provision NAT gateways so private subnets can reach the internet. | NAT gateways are the single largest cost in this module. In non-prod environments prefer `single_nat_gateway = true`. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L93) |
| `single_nat_gateway` | `bool` | `false` | Share one NAT gateway across all private subnets instead of one per AZ. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L99) |
| `nat_gateway_config` | `object({connectivity_type = optional(string, "public"), eip_allocation_ids = optional(list(string), []), reuse_eips = optional(bool, false), tags = optional(map(string), {})})` | `{}` | Advanced NAT gateway settings. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L105) |
| `flow_logs` | `object({destination = optional(string, "cloud-watch-logs"), enabled = bool, replicate_to_dr = optional(bool, false), retention_in_days = optional(number, 30), traffic_type = optional(string, "ALL")})` | `{enabled = true}` | VPC flow log configuration. | - | - | - | - | 2 | [`./variables.tf`](./variables.tf#L116) |
| `vpc_endpoints` | `set(string)` | `["s3", "dynamodb"]` | Gateway or interface endpoint services to create, e.g. `s3`, `ecr.api`. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L140) |
| `import_existing_vpc_id` | `string` | `""` | Adopt an existing VPC into state instead of creating a new one. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L146) |
| `vpn_preshared_key` | `string` | `null` | Pre-shared key for the optional site-to-site VPN connection. | - | ✓ | - | - | 0 | [`./variables.tf`](./variables.tf#L152) |
| `tags` | `map(string)` | `{}` | Tags applied to all resources. | - | - | ✗ | - | 0 | [`./variables.tf`](./variables.tf#L159) |
| `subnet_tags` | `object({database = optional(map(string), {}), private = optional(map(string), {}), public = optional(map(string), {})})` | `{}` | Extra tags per subnet tier, merged over `tags`. | - | - | ✗ | - | 0 | [`./variables.tf`](./variables.tf#L166) |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|-----------|-----------|------------|----------|
| vpc_id | `aws_vpc.this.id` | `aws_vpc.this` | The ID of the created VPC. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L1) |
| vpc_arn | `aws_vpc.this.arn` | `aws_vpc.this` | The ARN of the VPC. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L6) |
| vpc_cidr_block | `aws_vpc.this.cidr_block` | `aws_vpc.this` | The primary IPv4 CIDR block of the VPC. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L11) |
| all_cidr_blocks | `concat([aws_vpc.this.cidr_block], [for a in aws_vpc_ipv4_cidr_block_association.secondary : a.cidr_block])` | `aws_vpc.this, aws_vpc_ipv4_cidr_block_association.secondary` | Primary and secondary CIDR blocks associated with the VPC. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L16) |
| azs | `local.azs` | `local.azs` | Availability zones the subnets were spread across. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L21) |
| public_subnet_ids | `{ for az, s in aws_subnet.public : az => s.id }` | `aws_subnet.public` | Map of availability zone to public subnet ID. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L26) |
| private_subnet_ids | `{ for az, s in aws_subnet.private : az => s.id }` | `aws_subnet.private` | Map of availability zone to private subnet ID. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L31) |
| database_subnet_ids | `aws_subnet.database[*].id` | `aws_subnet.database` | List of database subnet IDs. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L36) |
| database_subnet_group_name | `try(aws_db_subnet_group.database[0].name, null)` | `aws_db_subnet_group.database` | Name of the database subnet group, or null when no database subnets were requested. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L41) |
| nat_gateway_ids | `aws_nat_gateway.this[*].id` | `aws_nat_gateway.this` | NAT gateway IDs, one per AZ (or one total with `single_nat_gateway`). | - | - | - | - | [`./outputs.tf`](./outputs.tf#L46) |
| nat_public_ips | `aws_nat_gateway.this[*].public_ip` | `aws_nat_gateway.this` | Elastic IPs attached to the NAT gateways. | - | - | - | `aws_route_table.private` | [`./outputs.tf`](./outputs.tf#L51) |
| route_table_ids | `{ public = aws_route_table.public.id private = { for az, rt in aws_route_table.private : az => rt.id } }` | `aws_route_table.public, aws_route_table.private` | Route table IDs grouped by tier. | Callers wiring additional routes need these. | - | - | - | [`./outputs.tf`](./outputs.tf#L58) |
| flow_log_id | `one(aws_flow_log.this[*].id)` | `aws_flow_log.this` | ID of the VPC flow log, if enabled. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L66) |
| endpoint_ids | `module.endpoints.endpoints` | `module.endpoints.endpoints` | VPC endpoints created by the endpoints submodule. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L71) |
| region | `data.aws_region.current.name` | `data.aws_region.current` | Region the VPC was deployed into. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L76) |
| vpn_preshared_key | `var.vpn_preshared_key` | `var.vpn_preshared_key` | Pre-shared key passed through for the VPN connection. | - | ✓ | - | - | [`./outputs.tf`](./outputs.tf#L81) |

## Locals

| Name | Value | References | Position |
|------|-------|------------|----------|
| azs | `coalesce(var.azs, slice(data.aws_availability_zones.available.names, 0, 3))` | `var.azs, data.aws_availability_zones.available` | [`./locals.tf`](./locals.tf#L4) |
| create_database_subnets | `length(var.database_subnets) > 0` | `var.database_subnets` | [`./locals.tf`](./locals.tf#L15) |
| create_vpn | `var.vpn_preshared_key != null` | `var.vpn_preshared_key` | [`./locals.tf`](./locals.tf#L17) |
| flow_log_destination_arn | `var.flow_logs.destination == "s3" ? module.flow_log_bucket[0].bucket_arn : aws_cloudwatch_log_group.flow_logs[0].arn` | `var.flow_logs, module.flow_log_bucket, aws_cloudwatch_log_group.flow_logs` | [`./locals.tf`](./locals.tf#L19) |
| name_prefix | `"${var.name}-${var.environment}"` | `var.name, var.environment` | [`./locals.tf`](./locals.tf#L2) |
| nat_gateway_count | `var.enable_nat_gateway ? (var.single_nat_gateway ? 1 : length(local.azs)) : 0` | `var.enable_nat_gateway, var.single_nat_gateway, local.azs` | [`./locals.tf`](./locals.tf#L16) |
| private_subnets | `{ for i, cidr in var.private_subnets : local.azs[i] => cidr }` | `var.private_subnets, local.azs` | [`./locals.tf`](./locals.tf#L13) |
| public_subnets | `{ for i, cidr in var.public_subnets : local.azs[i] => cidr }` | `var.public_subnets, local.azs` | [`./locals.tf`](./locals.tf#L12) |
| tags | `merge(var.tags, { Module = "vpc" Environment = var.environment ManagedBy = "terraform" })` | `var.tags, var.environment` | [`./locals.tf`](./locals.tf#L6-L10) |

## Managed Resources

| Name | Mode | Type | Count | For Each | Depends On | Nested Body | LifeCycle | Comment | Position |
|------|------|------|-------|----------|------------|-------------|-----------|---------|----------|
| this | managed | `aws_vpc` | - | - |  | `cidr_block, enable_dns_hostnames, enable_dns_support, tags` | `prevent_destroy=true; ignore_changes=[tags["CreatedAt"]]` | The VPC itself. Protected from accidental destruction in every environment because subnets, ENIs, and peering all hang off it. | [`./main.tf`](./main.tf#L3) |
| secondary | managed | `aws_vpc_ipv4_cidr_block_association` | - | `toset(var.secondary_cidr_blocks)` |  | `cidr_block, vpc_id` | - | - | [`./main.tf`](./main.tf#L16) |
| this | managed | `aws_internet_gateway` | `length(var.public_subnets) > 0 ? 1 : 0` | - |  | `tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L23) |
| public | managed | `aws_subnet` | - | `local.public_subnets` |  | `availability_zone, cidr_block, map_public_ip_on_launch, tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L32) |
| private | managed | `aws_subnet` | - | `local.private_subnets` |  | `availability_zone, cidr_block, tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L46) |
| database | managed | `aws_subnet` | `length(var.database_subnets)` | - |  | `availability_zone, cidr_block, tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L59) |
| database | managed | `aws_db_subnet_group` | `local.create_database_subnets ? 1 : 0` | - |  | `name, subnet_ids` | - | - | [`./main.tf`](./main.tf#L72) |
| nat | managed | `aws_eip` | `var.nat_gateway_config.reuse_eips ? 0 : local.nat_gateway_count` | - | aws_internet_gateway.this | `domain, tags` | - | - | [`./main.tf`](./main.tf#L81) |
| this | managed | `aws_nat_gateway` | `local.nat_gateway_count` | - | aws_internet_gateway.this, aws_eip.nat | `allocation_id, connectivity_type, subnet_id, tags` | `create_before_destroy=true; preconditions=1` | - | [`./main.tf`](./main.tf#L90) |
| public | managed | `aws_route_table` | - | - |  | `tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L113) |
| public_internet | managed | `aws_route` | - | - |  | `destination_cidr_block, gateway_id, route_table_id` | - | - | [`./main.tf`](./main.tf#L118) |
| private | managed | `aws_route_table` | - | `local.private_subnets` |  | `tags, vpc_id` | `ignore_changes=[propagating_vgws]` | - | [`./main.tf`](./main.tf#L128) |
| public | managed | `aws_route_table_association` | - | `aws_subnet.public` |  | `route_table_id, subnet_id` | - | - | [`./main.tf`](./main.tf#L147) |
| private | managed | `aws_route_table_association` | - | `aws_subnet.private` |  | `route_table_id, subnet_id` | - | - | [`./main.tf`](./main.tf#L154) |
| flow_logs | managed | `aws_cloudwatch_log_group` | `var.flow_logs.enabled && var.flow_logs.destination == "cloud-watch-logs" ? 1 : 0` | - |  | `name, retention_in_days` | `prevent_destroy=false` | - | [`./main.tf`](./main.tf#L163) |
| flow_logs | managed | `aws_iam_role` | `var.flow_logs.enabled ? 1 : 0` | - |  | `assume_role_policy, name_prefix` | - | - | [`./main.tf`](./main.tf#L174) |
| flow_logs | managed | `aws_iam_role_policy` | `var.flow_logs.enabled ? 1 : 0` | - |  | `policy, role` | - | - | [`./main.tf`](./main.tf#L181) |
| this | managed | `aws_flow_log` | `var.flow_logs.enabled ? 1 : 0` | - | aws_iam_role_policy.flow_logs | `iam_role_arn, log_destination, log_destination_type, traffic_type, vpc_id` | `replace_triggered_by=[aws_iam_role.flow_logs[0].id]; postconditions=1` | - | [`./main.tf`](./main.tf#L188) |
| this | managed | `aws_vpn_gateway` | `local.create_vpn ? 1 : 0` | - |  | `vpc_id` | - | - | [`./main.tf`](./main.tf#L211) |
| vpn_suffix | managed | `random_id` | `local.create_vpn ? 1 : 0` | - |  | `byte_length, keepers` | - | - | [`./main.tf`](./main.tf#L217) |

## Data Resources

| Name | Mode | Type | Count | For Each | Depends On | Nested Body | LifeCycle | Comment | Position |
|------|------|------|-------|----------|------------|-------------|-----------|---------|----------|
| available | data | `aws_availability_zones` | - | - |  | `state` | - | - | [`./data.tf`](./data.tf#L1) |
| current | data | `aws_region` | - | - |  | - | - | - | [`./data.tf`](./data.tf#L10) |
| current | data | `aws_caller_identity` | - | - |  | - | - | - | [`./data.tf`](./data.tf#L12) |
| flow_logs_assume | data | `aws_iam_policy_document` | - | - |  | - | - | Trust policy that lets the VPC flow-log service assume the delivery role. | [`./data.tf`](./data.tf#L15) |
| flow_logs_publish | data | `aws_iam_policy_document` | - | - |  | - | - | - | [`./data.tf`](./data.tf#L33) |

## Module Calls

| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |
|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|
| endpoints | terraform-aws-modules/vpc/aws//modules/vpc-endpoints | ~> 5.8 | - | - | - | - | `endpoints, security_group_ids, subnet_ids, tags, vpc_id` | Gateway and interface endpoints so private subnets reach AWS APIs without traversing the NAT gateway. | [`./modules.tf`](./modules.tf#L3) |
| flow_log_bucket | ./modules/log-bucket |  | `var.flow_logs.enabled && var.flow_logs.destination == "s3" ? 1 : 0` | - | - | `aws=aws, aws.replica=aws.dr` | `account_id, bucket_prefix, expiration_days, replicate` | - | [`./modules.tf`](./modules.tf#L22) |
| subnet_labels | cloudposse/label/null | 0.25.0 | - | `toset(local.azs)` | `aws_subnet.public, aws_subnet.private` | - | `attributes, environment, namespace` | - | [`./modules.tf`](./modules.tf#L37) |

## Moved

| From | To | Position |
|------|----|----------|
| `aws_subnet.public[0]` | `aws_subnet.public["us-east-1a"]` | [`./refactor.tf`](./refactor.tf#L2) |
| `aws_subnet.public[1]` | `aws_subnet.public["us-east-1b"]` | [`./refactor.tf`](./refactor.tf#L7) |
| `aws_nat_gateway.main` | `aws_nat_gateway.this[0]` | [`./refactor.tf`](./refactor.tf#L12) |
| `module.endpoints_legacy` | `module.endpoints` | [`./refactor.tf`](./refactor.tf#L17) |

## Imports

| To | ID | Provider | Position |
|----|----|----------|----------|
| `aws_vpc.this` | `var.import_existing_vpc_id` | - | [`./refactor.tf`](./refactor.tf#L22) |
| `aws_internet_gateway.this[0]` | `"igw-0123456789abcdef0"` | - | [`./refactor.tf`](./refactor.tf#L27) |

## Removed

| From | Destroy on Drop | Position |
|------|-----------------|----------|
| `aws_default_security_group.this` | ✗ | [`./refactor.tf`](./refactor.tf#L32) |
| `aws_vpn_connection.legacy` | - | [`./refactor.tf`](./refactor.tf#L40) |
