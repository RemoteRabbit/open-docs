terraform {
  required_version = ">= 1.5"

  required_providers {
    aws = {
      source                = "hashicorp/aws"
      version               = ">= 5.0"
      configuration_aliases = [aws.replica]
    }
  }
}

variable "bucket_prefix" {
  type        = string
  description = "Prefix for the generated bucket name."
}

variable "account_id" {
  type        = string
  description = "Account ID allowed to write flow logs into the bucket."
}

variable "expiration_days" {
  type        = number
  default     = 90
  description = "Days before log objects expire."
}

variable "replicate" {
  type        = bool
  default     = false
  description = "Replicate objects to a bucket in the `aws.replica` region."
}

resource "aws_s3_bucket" "this" {
  bucket_prefix = var.bucket_prefix
  force_destroy = false
}

resource "aws_s3_bucket_lifecycle_configuration" "this" {
  bucket = aws_s3_bucket.this.id

  rule {
    id     = "expire"
    status = "Enabled"

    expiration {
      days = var.expiration_days
    }
  }
}

resource "aws_s3_bucket" "replica" {
  count    = var.replicate ? 1 : 0
  provider = aws.replica

  bucket_prefix = "${var.bucket_prefix}-replica"
}

output "bucket_arn" {
  value       = aws_s3_bucket.this.arn
  description = "ARN of the log bucket."
}

output "replica_bucket_arn" {
  value       = try(aws_s3_bucket.replica[0].arn, null)
  description = "ARN of the replica bucket, if replication is enabled."
}
