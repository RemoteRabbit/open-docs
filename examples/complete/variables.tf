# --- Primitive types and defaults ------------------------------------------

variable "string_required" {
  type        = string
  description = "A required string with no default."
}

variable "string_default" {
  type        = string
  default     = "hello"
  description = "A string with a default."
}

variable "string_empty" {
  type        = string
  default     = ""
  description = "A string whose default is the empty string."
}

variable "number_default" {
  type        = number
  default     = 42
  description = "A number."
}

variable "number_float" {
  type        = number
  default     = 3.14159
  description = "A floating point number."
}

variable "number_negative" {
  type        = number
  default     = -1
  description = "A negative number."
}

variable "bool_true" {
  type        = bool
  default     = true
  description = "A bool defaulting to true."
}

variable "bool_false" {
  type        = bool
  default     = false
  description = "A bool defaulting to false."
}

variable "null_default" {
  type        = string
  default     = null
  description = "Optional because the default is null."
}

variable "untyped" {
  description = "No type constraint at all."
}

variable "any_type" {
  type        = any
  default     = { anything = ["goes", 1, true] }
  description = "Type `any` with a mixed default."
}

# --- Collection types -------------------------------------------------------

variable "list_of_strings" {
  type        = list(string)
  default     = ["a", "b", "c"]
  description = "A list of strings."
}

variable "list_empty" {
  type        = list(number)
  default     = []
  description = "An empty list."
}

variable "set_of_strings" {
  type        = set(string)
  default     = ["x", "y"]
  description = "A set of strings."
}

variable "map_of_numbers" {
  type = map(number)
  default = {
    small  = 1
    medium = 10
    large  = 100
  }
  description = "A map of numbers."
}

variable "map_empty" {
  type        = map(string)
  default     = {}
  description = "An empty map."
}

variable "tuple_type" {
  type        = tuple([string, number, bool])
  default     = ["one", 2, true]
  description = "A tuple with mixed element types."
}

variable "list_of_lists" {
  type        = list(list(string))
  default     = [["a", "b"], ["c"]]
  description = "Nested collections."
}

variable "map_of_lists" {
  type = map(list(string))
  default = {
    admins  = ["alice", "bob"]
    readers = []
  }
  description = "A map whose values are lists."
}

variable "regions" {
  type        = set(string)
  default     = ["us-east-1", "us-west-2", "eu-west-1"]
  description = "Regions for the multi-instance `aws.by_region` provider."
}

# --- Structural types -------------------------------------------------------

variable "simple_object" {
  type = object({
    name    = string
    enabled = bool
    port    = number
  })
  default = {
    name    = "svc"
    enabled = true
    port    = 8080
  }
  description = "An object with only required attributes."
}

variable "optional_object" {
  type = object({
    name     = string
    replicas = optional(number, 1)
    image    = optional(string)
    labels   = optional(map(string), {})
    ports    = optional(list(number), [80, 443])
  })
  description = "An object with optional attributes and defaults. Required, so no default."
}

variable "nested_object" {
  type = object({
    network = object({
      cidr = string
      subnets = list(object({
        name   = string
        cidr   = string
        public = optional(bool, false)
      }))
    })
    database = optional(object({
      engine  = optional(string, "postgres")
      version = optional(string, "16")
      storage = optional(object({
        size_gb   = optional(number, 20)
        encrypted = optional(bool, true)
        iops      = optional(number)
      }), {})
    }), {})
  })
  default = {
    network = {
      cidr = "10.0.0.0/16"
      subnets = [
        { name = "a", cidr = "10.0.1.0/24", public = true },
        { name = "b", cidr = "10.0.2.0/24" },
      ]
    }
  }
  description = "Deeply nested objects with optional attributes at several levels."
}

variable "list_of_objects" {
  type = list(object({
    name  = string
    value = string
  }))
  default = [
    { name = "first", value = "1" },
    { name = "second", value = "2" },
  ]
  description = "A list of objects."
}

variable "map_of_objects" {
  type = map(object({
    instance_type = string
    min_size      = optional(number, 1)
    max_size      = optional(number, 3)
  }))
  default = {
    web    = { instance_type = "t3.small" }
    worker = { instance_type = "c6i.large", max_size = 10 }
  }
  description = "A map of objects, the classic node-group shape."
}

# --- Meta-arguments ---------------------------------------------------------

variable "sensitive_string" {
  type        = string
  default     = "hunter2"
  sensitive   = true
  description = "A sensitive variable with a default."
}

variable "state_passphrase" {
  type        = string
  sensitive   = true
  description = "Passphrase for OpenTofu state encryption."
}

variable "nullable_false" {
  type        = string
  default     = "never-null"
  nullable    = false
  description = "Explicitly not nullable."
}

variable "nullable_true" {
  type        = string
  default     = null
  nullable    = true
  description = "Explicitly nullable."
}

variable "ephemeral_token" {
  type        = string
  ephemeral   = true
  sensitive   = true
  description = "An ephemeral, sensitive input that is never persisted."
}

variable "many_validations" {
  type        = string
  default     = "abc-123"
  description = "Three validation blocks."

  validation {
    condition     = length(var.many_validations) >= 3
    error_message = "Must be at least 3 characters."
  }

  validation {
    condition     = length(var.many_validations) <= 63
    error_message = "Must be at most 63 characters."
  }

  validation {
    condition     = can(regex("^[a-z0-9-]+$", var.many_validations))
    error_message = "Only lowercase letters, digits, and hyphens: got ${var.many_validations}."
  }
}

# --- Description and comment edge cases ------------------------------------

variable "pipe_in_description" {
  type        = string
  default     = "a|b"
  description = "Contains a pipe | character that would break a Markdown table | if unescaped."
}

variable "markdown_in_description" {
  type        = string
  default     = "*"
  description = "Has **bold**, _italic_, `code`, a [link](https://opentofu.org), and <b>raw HTML</b>."
}

variable "multiline_description" {
  type        = string
  default     = "x"
  description = <<-EOT
    A heredoc description spanning several lines.

    - It has a bullet list
    - and a blank line above

    And a trailing paragraph.
  EOT
}

variable "unicode_description" {
  type        = string
  default     = "héllo wörld ✓"
  description = "Unicode: 日本語, emoji 🚀, arrows → ←, and accents café."
}

variable "long_description" {
  type        = string
  default     = "long"
  description = "This description is intentionally very long so that table cells wrap or overflow in renderers. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."
}

variable "long_default" {
  type        = list(string)
  default     = ["10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24", "10.0.7.0/24", "10.0.8.0/24", "10.0.9.0/24"]
  description = "A default value long enough to stress column widths."
}

variable "quotes_in_default" {
  type        = string
  default     = "she said \"hi\" and left `quickly`"
  description = "Escaped quotes and backticks inside a default."
}

variable "heredoc_default" {
  type        = string
  default     = <<-EOT
    #!/bin/bash
    echo "user data"
  EOT
  description = "A multi-line heredoc default."
}

variable "interpolated_default_like" {
  type        = string
  default     = "$${not_interpolated}"
  description = "A literal that looks like interpolation."
}

# Leading doc-comment, no description attribute.
# The renderer should surface this in the Comment column.
variable "comment_only" {
  type    = string
  default = "c"
}

/*
 * Block-style doc-comment above a variable that also has a description.
 */
variable "comment_and_description" {
  type        = string
  default     = "both"
  description = "Has both a description and a leading block comment."
}

// Double-slash comment style.
variable "slash_comment" {
  type    = number
  default = 0
}

variable "no_description" {
  type    = string
  default = "undocumented"
}
