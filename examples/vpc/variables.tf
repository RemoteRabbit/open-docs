variable "name" {
  type        = string
  description = "Name prefix applied to every resource created by this module."

  validation {
    condition     = can(regex("^[a-z][a-z0-9-]{1,30}[a-z0-9]$", var.name))
    error_message = "name must be 3-32 lowercase alphanumeric characters or hyphens, starting with a letter."
  }
}

variable "environment" {
  type        = string
  default     = "dev"
  description = "Deployment environment. Drives tagging and a few cost-related defaults."

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "environment must be one of: dev, staging, prod."
  }
}

variable "region" {
  type        = string
  description = "AWS region to deploy into."
}

variable "dr_region" {
  type        = string
  default     = "us-west-2"
  description = "Secondary region used for disaster-recovery copies of flow logs."
}

variable "cidr_block" {
  type        = string
  default     = "10.0.0.0/16"
  description = "IPv4 CIDR block for the VPC."

  validation {
    condition     = can(cidrhost(var.cidr_block, 0))
    error_message = "cidr_block must be a valid IPv4 CIDR."
  }

  validation {
    condition     = tonumber(split("/", var.cidr_block)[1]) <= 24
    error_message = "cidr_block must be /24 or larger."
  }
}

variable "secondary_cidr_blocks" {
  type        = list(string)
  default     = []
  description = "Additional IPv4 CIDR blocks to associate with the VPC."
}

variable "azs" {
  type        = list(string)
  default     = null
  description = "Availability zones to spread subnets across. Defaults to the first three available in the region."
}

variable "public_subnets" {
  type        = list(string)
  default     = ["10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24"]
  description = "CIDR blocks for public subnets, one per availability zone."
}

variable "private_subnets" {
  type        = list(string)
  default     = ["10.0.10.0/24", "10.0.11.0/24", "10.0.12.0/24"]
  description = "CIDR blocks for private subnets, one per availability zone."
}

variable "database_subnets" {
  type        = list(string)
  default     = []
  description = "CIDR blocks for isolated database subnets. Leave empty to skip."
}

variable "enable_dns_hostnames" {
  type        = bool
  default     = true
  description = "Enable DNS hostnames in the VPC."
}

variable "enable_dns_support" {
  type        = bool
  default     = true
  description = "Enable DNS resolution in the VPC."
}

# NAT gateways are the single largest cost in this module. In non-prod
# environments prefer `single_nat_gateway = true`.
variable "enable_nat_gateway" {
  type        = bool
  default     = true
  description = "Provision NAT gateways so private subnets can reach the internet."
}

variable "single_nat_gateway" {
  type        = bool
  default     = false
  description = "Share one NAT gateway across all private subnets instead of one per AZ."
}

variable "nat_gateway_config" {
  type = object({
    connectivity_type  = optional(string, "public")
    reuse_eips         = optional(bool, false)
    eip_allocation_ids = optional(list(string), [])
    tags               = optional(map(string), {})
  })
  default     = {}
  description = "Advanced NAT gateway settings."
}

variable "flow_logs" {
  type = object({
    enabled           = bool
    traffic_type      = optional(string, "ALL")
    retention_in_days = optional(number, 30)
    destination       = optional(string, "cloud-watch-logs")
    replicate_to_dr   = optional(bool, false)
  })
  default = {
    enabled = true
  }
  description = "VPC flow log configuration."

  validation {
    condition     = contains(["ACCEPT", "REJECT", "ALL"], var.flow_logs.traffic_type)
    error_message = "flow_logs.traffic_type must be ACCEPT, REJECT, or ALL."
  }

  validation {
    condition     = contains([1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 3653], var.flow_logs.retention_in_days)
    error_message = "flow_logs.retention_in_days must be a CloudWatch-supported retention period."
  }
}

variable "vpc_endpoints" {
  type        = set(string)
  default     = ["s3", "dynamodb"]
  description = "Gateway or interface endpoint services to create, e.g. `s3`, `ecr.api`."
}

variable "import_existing_vpc_id" {
  type        = string
  default     = ""
  description = "Adopt an existing VPC into state instead of creating a new one."
}

variable "vpn_preshared_key" {
  type        = string
  default     = null
  sensitive   = true
  description = "Pre-shared key for the optional site-to-site VPN connection."
}

variable "tags" {
  type        = map(string)
  default     = {}
  nullable    = false
  description = "Tags applied to all resources."
}

variable "subnet_tags" {
  type = object({
    public   = optional(map(string), {})
    private  = optional(map(string), {})
    database = optional(map(string), {})
  })
  default     = {}
  nullable    = false
  description = "Extra tags per subnet tier, merged over `tags`."
}
