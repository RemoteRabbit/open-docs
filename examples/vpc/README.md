# Module `vpc`

- **Path:** `/home/remoterabbit/Projects/open-docs/examples/vpc`
- **Required core:** `>= 1.5`

## Required Providers

| Name | Source | Version Constraints | Configuration Aliases | Position |
|------|--------|---------------------|-----------------------|----------|
| `aws` | `hashicorp/aws` | `>= 5.0` | - | [`./main.tf`](./main.tf#L4-L7) [4:5:69 -> 7:6:139] |

## Providers

| Name | Alias | For Each | Position |
|------|-------|----------|----------|
| `aws` | - | - | [`./main.tf`](./main.tf#L11) [11:1:147 -> 11:15:161] |

## Inputs

| Name | Type | Default | Required | Description | Comment | Sensitive | Nullable | Ephemeral | Validations | Position |
|------|------|---------|:--------:|-------------|---------|:---------:|:--------:|:---------:|:-----------:|----------|
| `region` | `string` | - | yes | AWS region to deploy into. | - | - | - | - | 0 | [`./main.tf`](./main.tf#L15) [15:1:189 -> 15:18:206] |
| `cidr_block` | `string` | `"10.0.0.0/16"` | no | CIDR block for the VPC. | - | - | - | - | 0 | [`./main.tf`](./main.tf#L20) [20:1:280 -> 20:22:301] |
| `tags` | `map(string)` | `{}` | no | Tags applied to all resources. | - | - | - | - | 0 | [`./main.tf`](./main.tf#L26) [26:1:402 -> 26:16:417] |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|:---------:|:---------:|------------|----------|
| `vpc_id` | `aws_vpc.this.id` | `aws_vpc.this` | The ID of the created VPC. | - | - | - | - | [`./main.tf`](./main.tf#L46) [46:1:763 -> 46:16:778] |
| `secret` | `"shh"` | - | A sensitive output. | - | ✓ | - | - | [`./main.tf`](./main.tf#L51) [51:1:861 -> 51:16:876] |

## Managed Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `managed` | `aws_vpc` | `this` | - | - | - | - | `cidr_block, tags` | - | - | [`./main.tf`](./main.tf#L32) [32:1:519 -> 32:26:544] |

## Data Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `data` | `aws_availability_zones` | `available` | - | - | - | - | `state` | - | - | [`./main.tf`](./main.tf#L37) [37:1:604 -> 37:42:645] |

## Module Calls

| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |
|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|
| `subnets` | `terraform-aws-modules/subnets/aws` | `1.2.0` | - | - | - | - | - | - | [`./main.tf`](./main.tf#L41) [41:1:673 -> 41:17:689] |
