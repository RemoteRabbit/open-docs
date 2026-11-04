# Module `vpc`

<!-- VERSION-START -->
This documentation was generated with [open-docs v1](https://github.com/RemoteRabbit/open-docs)
<!-- VERSION-END -->

- **Required core:** `>= 1.5, < 2.0`

## Required Providers

| Name | Source | Version | Position |
|------|--------|---------|----------|
| `aws` | `hashicorp/aws` | `>= 5.0, < 7.0` | [`./versions.tf`](./versions.tf#L5-L8) [5:5:77 -> 8:6:154] |
| `random` | `hashicorp/random` | `~> 3.6` | [`./versions.tf`](./versions.tf#L9-L12) [9:5:159 -> 12:6:235] |

## Providers

| Name | Position |
|------|----------|
| `aws` | [`./providers.tf`](./providers.tf#L3) [3:1:96 -> 3:15:110] |
| `aws` | [`./providers.tf`](./providers.tf#L12) [12:1:251 -> 12:15:265] |

## Inputs

| Name | Type | Default | Description | Comment | Sensitive | Nullable | Ephemeral | Validation Count | Position |
|------|------|---------|-------------|---------|-----------|----------|-----------|------------------|----------|
| `name` | `string` | - | Name prefix applied to every resource created by this module. |  | - | - | - | 1 | [`./variables.tf`](./variables.tf#L1) [1:1:0 -> 1:16:15] |
| `environment` | `string` | `"dev"` | Deployment environment. Drives tagging and a few cost-related defaults. |  | - | - | - | 1 | [`./variables.tf`](./variables.tf#L11) [11:1:329 -> 11:23:351] |
| `region` | `string` | - | AWS region to deploy into. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L22) [22:1:656 -> 22:18:673] |
| `dr_region` | `string` | `"us-west-2"` | Secondary region used for disaster-recovery copies of flow logs. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L27) [27:1:747 -> 27:21:767] |
| `cidr_block` | `string` | `"10.0.0.0/16"` | IPv4 CIDR block for the VPC. |  | - | - | - | 2 | [`./variables.tf`](./variables.tf#L33) [33:1:907 -> 33:22:928] |
| `secondary_cidr_blocks` | `list(string)` | `[]` | Additional IPv4 CIDR blocks to associate with the VPC. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L49) [49:1:1309 -> 49:33:1341] |
| `azs` | `list(string)` | `null` | Availability zones to spread subnets across. Defaults to the first three available in the region. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L55) [55:1:1468 -> 55:15:1482] |
| `public_subnets` | `list(string)` | `["10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24"]` | CIDR blocks for public subnets, one per availability zone. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L61) [61:1:1654 -> 61:26:1679] |
| `private_subnets` | `list(string)` | `["10.0.10.0/24", "10.0.11.0/24", "10.0.12.0/24"]` | CIDR blocks for private subnets, one per availability zone. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L67) [67:1:1853 -> 67:27:1879] |
| `database_subnets` | `list(string)` | `[]` | CIDR blocks for isolated database subnets. Leave empty to skip. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L73) [73:1:2057 -> 73:28:2084] |
| `enable_dns_hostnames` | `bool` | `true` | Enable DNS hostnames in the VPC. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L79) [79:1:2220 -> 79:32:2251] |
| `enable_dns_support` | `bool` | `true` | Enable DNS resolution in the VPC. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L85) [85:1:2350 -> 85:30:2379] |
| `enable_nat_gateway` | `bool` | `true` | Provision NAT gateways so private subnets can reach the internet. | NAT gateways are the single largest cost in this module. In non-prod
environments prefer `single_nat_gateway = true`. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L93) [93:1:2601 -> 93:30:2630] |
| `single_nat_gateway` | `bool` | `false` | Share one NAT gateway across all private subnets instead of one per AZ. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L99) [99:1:2762 -> 99:30:2791] |
| `nat_gateway_config` | `object({connectivity_type = optional(string, "public"), eip_allocation_ids = optional(list(string), []), reuse_eips = optional(bool, false), tags = optional(map(string), {})})` | `{}` | Advanced NAT gateway settings. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L105) [105:1:2930 -> 105:30:2959] |
| `flow_logs` | `object({destination = optional(string, "cloud-watch-logs"), enabled = bool, replicate_to_dr = optional(bool, false), retention_in_days = optional(number, 30), traffic_type = optional(string, "ALL")})` | `{enabled = true}` | VPC flow log configuration. |  | - | - | - | 2 | [`./variables.tf`](./variables.tf#L116) [116:1:3258 -> 116:21:3278] |
| `vpc_endpoints` | `set(string)` | `["s3", "dynamodb"]` | Gateway or interface endpoint services to create, e.g. `s3`, `ecr.api`. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L140) [140:1:4060 -> 140:25:4084] |
| `import_existing_vpc_id` | `string` | `""` | Adopt an existing VPC into state instead of creating a new one. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L146) [146:1:4243 -> 146:34:4276] |
| `vpn_preshared_key` | `string` | `null` | Pre-shared key for the optional site-to-site VPN connection. |  | ✓ | - | - | 0 | [`./variables.tf`](./variables.tf#L152) [152:1:4406 -> 152:29:4434] |
| `tags` | `map(string)` | `{}` | Tags applied to all resources. |  | - | ✗ | - | 0 | [`./variables.tf`](./variables.tf#L159) [159:1:4584 -> 159:16:4599] |
| `subnet_tags` | `object({database = optional(map(string), {}), private = optional(map(string), {}), public = optional(map(string), {})})` | `{}` | Extra tags per subnet tier, merged over `tags`. |  | - | ✗ | - | 0 | [`./variables.tf`](./variables.tf#L166) [166:1:4723 -> 166:23:4745] |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|-----------|-----------|------------|----------|
| vpc_id | `aws_vpc.this.id` | `aws_vpc.this` | The ID of the created VPC. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L1) [1:1:0 -> 1:16:15] |
| vpc_arn | `aws_vpc.this.arn` | `aws_vpc.this` | The ARN of the VPC. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L6) [6:1:98 -> 6:17:114] |
| vpc_cidr_block | `aws_vpc.this.cidr_block` | `aws_vpc.this` | The primary IPv4 CIDR block of the VPC. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L11) [11:1:191 -> 11:24:214] |
| all_cidr_blocks | `concat([aws_vpc.this.cidr_block], [for a in aws_vpc_ipv4_cidr_block_association.secondary : a.cidr_block])` | `aws_vpc.this, aws_vpc_ipv4_cidr_block_association.secondary` | Primary and secondary CIDR blocks associated with the VPC. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L16) [16:1:318 -> 16:25:342] |
| azs | `local.azs` | `local.azs` | Availability zones the subnets were spread across. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L21) [21:1:548 -> 21:13:560] |
| public_subnet_ids | `{ for az, s in aws_subnet.public : az => s.id }` | `aws_subnet.public` | Map of availability zone to public subnet ID. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L26) [26:1:661 -> 26:27:687] |
| private_subnet_ids | `{ for az, s in aws_subnet.private : az => s.id }` | `aws_subnet.private` | Map of availability zone to private subnet ID. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L31) [31:1:821 -> 31:28:848] |
| database_subnet_ids | `aws_subnet.database[*].id` | `aws_subnet.database` | List of database subnet IDs. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L36) [36:1:984 -> 36:29:1012] |
| database_subnet_group_name | `try(aws_db_subnet_group.database[0].name, null)` | `aws_db_subnet_group.database` | Name of the database subnet group, or null when no database subnets were requested. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L41) [41:1:1107 -> 41:36:1142] |
| nat_gateway_ids | `aws_nat_gateway.this[*].id` | `aws_nat_gateway.this` | NAT gateway IDs, one per AZ (or one total with `single_nat_gateway`). |  | - | - | - | [`./outputs.tf`](./outputs.tf#L46) [46:1:1314 -> 46:25:1338] |
| nat_public_ips | `aws_nat_gateway.this[*].public_ip` | `aws_nat_gateway.this` | Elastic IPs attached to the NAT gateways. |  | - | - | `aws_route_table.private` | [`./outputs.tf`](./outputs.tf#L51) [51:1:1475 -> 51:24:1498] |
| route_table_ids | `{ public = aws_route_table.public.id private = { for az, rt in aws_route_table.private : az => rt.id } }` | `aws_route_table.public, aws_route_table.private` | Route table IDs grouped by tier. | Callers wiring additional routes need these. | - | - | - | [`./outputs.tf`](./outputs.tf#L58) [58:1:1703 -> 58:25:1727] |
| flow_log_id | `one(aws_flow_log.this[*].id)` | `aws_flow_log.this` | ID of the VPC flow log, if enabled. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L66) [66:1:1910 -> 66:21:1930] |
| endpoint_ids | `module.endpoints.endpoints` | `module.endpoints.endpoints` | VPC endpoints created by the endpoints submodule. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L71) [71:1:2035 -> 71:22:2056] |
| region | `data.aws_region.current.name` | `data.aws_region.current` | Region the VPC was deployed into. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L76) [76:1:2173 -> 76:16:2188] |
| vpn_preshared_key | `var.vpn_preshared_key` | `var.vpn_preshared_key` | Pre-shared key passed through for the VPN connection. |  | ✓ | - | - | [`./outputs.tf`](./outputs.tf#L81) [81:1:2291 -> 81:27:2317] |

## Locals

| Name |
|------|
| azs |
| create_database_subnets |
| create_vpn |
| flow_log_destination_arn |
| name_prefix |
| nat_gateway_count |
| private_subnets |
| public_subnets |
| tags |

## Managed Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `managed` | `aws_vpc` | `this` | - | - | - | - | `cidr_block, enable_dns_hostnames, enable_dns_support, tags` | `prevent_destroy=true; ignore_changes=[tags]` | The VPC itself. Protected from accidental destruction in every environment because subnets, ENIs, and peering all hang off it. | [`./main.tf`](./main.tf#L3) [3:1:131 -> 3:26:156] |
| `managed` | `aws_vpc_ipv4_cidr_block_association` | `secondary` | - | - | `toset(var.secondary_cidr_blocks)` | - | `cidr_block, vpc_id` | - | - | [`./main.tf`](./main.tf#L16) [16:1:446 -> 16:59:504] |
| `managed` | `aws_internet_gateway` | `this` | - | `length(var.public_subnets) > 0 ? 1 : 0` | - | - | `tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L23) [23:1:614 -> 23:39:652] |
| `managed` | `aws_subnet` | `public` | - | - | `local.public_subnets` | - | `availability_zone, cidr_block, map_public_ip_on_launch, tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L32) [32:1:863 -> 32:31:893] |
| `managed` | `aws_subnet` | `private` | - | - | `local.private_subnets` | - | `availability_zone, cidr_block, tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L46) [46:1:1207 -> 46:32:1238] |
| `managed` | `aws_subnet` | `database` | - | `length(var.database_subnets)` | - | - | `availability_zone, cidr_block, tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L59) [59:1:1505 -> 59:33:1537] |
| `managed` | `aws_db_subnet_group` | `database` | - | `local.create_database_subnets ? 1 : 0` | - | - | `name, subnet_ids` | - | - | [`./main.tf`](./main.tf#L72) [72:1:1856 -> 72:42:1897] |
| `managed` | `aws_eip` | `nat` | - | `var.nat_gateway_config.reuse_eips ? 0 : local.nat_gateway_count` | - | `aws_internet_gateway.this` | `domain, tags` | - | - | [`./main.tf`](./main.tf#L81) [81:1:2113 -> 81:25:2137] |
| `managed` | `aws_nat_gateway` | `this` | - | `local.nat_gateway_count` | - | `aws_internet_gateway.this, aws_eip.nat` | `allocation_id, connectivity_type, subnet_id, tags` | `create_before_destroy=true; preconditions=1` | - | [`./main.tf`](./main.tf#L90) [90:1:2343 -> 90:34:2376] |
| `managed` | `aws_route_table` | `public` | - | - | - | - | `tags, vpc_id` | - | - | [`./main.tf`](./main.tf#L113) [113:1:3216 -> 113:36:3251] |
| `managed` | `aws_route` | `public_internet` | - | - | - | - | `destination_cidr_block, gateway_id, route_table_id` | - | - | [`./main.tf`](./main.tf#L118) [118:1:3336 -> 118:39:3374] |
| `managed` | `aws_route_table` | `private` | - | - | `local.private_subnets` | - | `tags, vpc_id` | `ignore_changes=[propagating_vgws]` | - | [`./main.tf`](./main.tf#L128) [128:1:3567 -> 128:37:3603] |
| `managed` | `aws_route_table_association` | `public` | - | - | `aws_subnet.public` | - | `route_table_id, subnet_id` | - | - | [`./main.tf`](./main.tf#L147) [147:1:4034 -> 147:48:4081] |
| `managed` | `aws_route_table_association` | `private` | - | - | `aws_subnet.private` | - | `route_table_id, subnet_id` | - | - | [`./main.tf`](./main.tf#L154) [154:1:4197 -> 154:49:4245] |
| `managed` | `aws_cloudwatch_log_group` | `flow_logs` | - | `var.flow_logs.enabled && var.flow_logs.destination == "cloud-watch-logs" ? 1 : 0` | - | - | `name, retention_in_days` | `prevent_destroy=false` | - | [`./main.tf`](./main.tf#L163) [163:1:4452 -> 163:48:4499] |
| `managed` | `aws_iam_role` | `flow_logs` | - | `var.flow_logs.enabled ? 1 : 0` | - | - | `assume_role_policy, name_prefix` | - | - | [`./main.tf`](./main.tf#L174) [174:1:4762 -> 174:36:4797] |
| `managed` | `aws_iam_role_policy` | `flow_logs` | - | `var.flow_logs.enabled ? 1 : 0` | - | - | `policy, role` | - | - | [`./main.tf`](./main.tf#L181) [181:1:4975 -> 181:43:5017] |
| `managed` | `aws_flow_log` | `this` | - | `var.flow_logs.enabled ? 1 : 0` | - | `aws_iam_role_policy.flow_logs` | `iam_role_arn, log_destination, log_destination_type, traffic_type, vpc_id` | `replace_triggered_by=[aws_iam_role.flow_logs.id]; postconditions=1` | - | [`./main.tf`](./main.tf#L188) [188:1:5167 -> 188:31:5197] |
| `managed` | `aws_vpn_gateway` | `this` | - | `local.create_vpn ? 1 : 0` | - | - | `vpc_id` | - | - | [`./main.tf`](./main.tf#L211) [211:1:5874 -> 211:34:5907] |
| `managed` | `random_id` | `vpn_suffix` | - | `local.create_vpn ? 1 : 0` | - | - | `byte_length, keepers` | - | - | [`./main.tf`](./main.tf#L217) [217:1:5976 -> 217:34:6009] |

## Data Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `data` | `aws_availability_zones` | `available` | - | - | - | - | `state` | - | - | [`./data.tf`](./data.tf#L1) [1:1:0 -> 1:42:41] |
| `data` | `aws_region` | `current` | - | - | - | - | - | - | - | [`./data.tf`](./data.tf#L10) [10:1:151 -> 10:28:178] |
| `data` | `aws_caller_identity` | `current` | - | - | - | - | - | - | - | [`./data.tf`](./data.tf#L12) [12:1:183 -> 12:37:219] |
| `data` | `aws_iam_policy_document` | `flow_logs_assume` | - | - | - | - | - | - | Trust policy that lets the VPC flow-log service assume the delivery role. | [`./data.tf`](./data.tf#L15) [15:1:300 -> 15:50:349] |
| `data` | `aws_iam_policy_document` | `flow_logs_publish` | - | - | - | - | - | - | - | [`./data.tf`](./data.tf#L33) [33:1:689 -> 33:51:739] |

## Module Calls

| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |
|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|
| `endpoints` | `terraform-aws-modules/vpc/aws//modules/vpc-endpoints` | `~> 5.8` | - | - | - | - | `endpoints, security_group_ids, subnet_ids, tags, vpc_id` | Gateway and interface endpoints so private subnets reach AWS APIs without traversing the NAT gateway. | [`./modules.tf`](./modules.tf#L3) [3:1:106 -> 3:19:124] |
| `flow_log_bucket` | `./modules/log-bucket` | - | `var.flow_logs.enabled && var.flow_logs.destination == "s3" ? 1 : 0` | - | - | `aws=aws` | `account_id, bucket_prefix, expiration_days, replicate` | - | [`./modules.tf`](./modules.tf#L22) [22:1:597 -> 22:25:621] |
| `subnet_labels` | `cloudposse/label/null` | `0.25.0` | - | `toset(local.azs)` | `aws_subnet.public, aws_subnet.private` | - | `attributes, environment, namespace` | - | [`./modules.tf`](./modules.tf#L37) [37:1:1027 -> 37:23:1049] |

## Moved

| From | To | Position |
|------|----|----------|
| `aws_subnet.public` | `aws_subnet.public` | [`./refactor.tf`](./refactor.tf#L2) [2:1:65 -> 2:6:70] |
| `aws_subnet.public` | `aws_subnet.public` | [`./refactor.tf`](./refactor.tf#L7) [7:1:147 -> 7:6:152] |
| `aws_nat_gateway.main` | `aws_nat_gateway.this` | [`./refactor.tf`](./refactor.tf#L12) [12:1:229 -> 12:6:234] |
| `module.endpoints_legacy` | `module.endpoints` | [`./refactor.tf`](./refactor.tf#L17) [17:1:303 -> 17:6:308] |

## Imports

| To | ID | Provider | Position |
|----|----|----------|----------|
| `aws_vpc.this` | `var.import_existing_vpc_id` | - | [`./refactor.tf`](./refactor.tf#L22) [22:1:373 -> 22:7:379] |
| `aws_internet_gateway.this` | `"igw-0123456789abcdef0"` | - | [`./refactor.tf`](./refactor.tf#L27) [27:1:439 -> 27:7:445] |

## Removed

| From | Destroy On Drop | Position |
|------|-----------------|----------|
| `aws_default_security_group.this` | ✗ | [`./refactor.tf`](./refactor.tf#L32) [32:1:518 -> 32:8:525] |
| `aws_vpn_connection.legacy` | - | [`./refactor.tf`](./refactor.tf#L40) [40:1:611 -> 40:8:618] |

## Diagnostics

| Severity | Summary | Detail | Subject | Context |
|----------|---------|--------|---------|---------|
| error | Ambiguous attribute key | If this expression is intended to be a reference, wrap it in parentheses. If it's instead intended as a literal name containing periods, wrap it in quotes to create a string literal. | [`./modules.tf`](./modules.tf#L28) [28:5:779 -> 28:16:790] | - |
