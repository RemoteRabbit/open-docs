terraform {
  required_providers {
    aws = {
      source                = "hashicorp/aws"
      configuration_aliases = [aws.replica]
    }
  }
}

variable "name" {
  type        = string
  default     = "child"
  description = "Name passed down from the parent."
}

variable "count_value" {
  type    = number
  default = 1
}

variable "enabled" {
  type    = bool
  default = true
}

variable "tags" {
  type    = map(string)
  default = {}
}

variable "objects" {
  type    = any
  default = {}
}

resource "aws_s3_bucket" "this" {
  count  = var.enabled ? var.count_value : 0
  bucket = "${var.name}-${count.index}"
  tags   = var.tags
}

output "passthrough" {
  value       = var.objects
  description = "Echoes the objects input back to the caller."
}
