# Deliberately broken module for testing how diagnostics render. Each block
# below triggers a different loader diagnostic while the rest of the module
# still loads.

terraform {
  required_providers {
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}

variable "description_source" {
  type    = string
  default = "from a variable"
}

variable "flag_source" {
  type    = bool
  default = false
}

# Validation block missing both condition and error_message.
variable "malformed_validation" {
  type = string

  validation {
  }
}

# Unknown type keyword.
variable "bad_type" {
  type    = strng
  default = "typo"
}

# description must be a literal string.
output "non_literal_description" {
  value       = "hello"
  description = "from ${var.description_source}"
}

# sensitive must be a literal bool.
output "non_literal_sensitive" {
  value     = "hello"
  sensitive = var.flag_source
}

# Lifecycle precondition missing its required attributes.
resource "null_resource" "malformed_precondition" {
  lifecycle {
    precondition {
    }
  }
}

# A perfectly valid resource, so the healthy parts still render.
resource "null_resource" "valid" {
  triggers = {
    note = "I load fine."
  }
}
