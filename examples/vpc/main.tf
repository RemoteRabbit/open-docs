terraform {
  required_version = ">= 1.5"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">= 5.0"
    }
  }
}

provider "aws" {
  region = var.region
}

variable "region" {
  type        = string
  description = "AWS region to deploy into."
}

variable "cidr_block" {
  type        = string
  default     = "10.0.0.0/16"
  description = "CIDR block for the VPC."
}

variable "tags" {
  type        = map(string)
  default     = {}
  description = "Tags applied to all resources."
}

resource "aws_vpc" "this" {
  cidr_block = var.cidr_block
  tags       = var.tags
}

data "aws_availability_zones" "available" {
  state = "available"
}

module "subnets" {
  source  = "terraform-aws-modules/subnets/aws"
  version = "1.2.0"
}

output "vpc_id" {
  value       = aws_vpc.this.id
  description = "The ID of the created VPC."
}

output "secret" {
  value       = "shh"
  description = "A sensitive output."
  sensitive   = true
}
