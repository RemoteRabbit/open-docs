output "literal" {
  value       = "constant"
  description = "A literal string output."
}

output "simple_ref" {
  value       = null_resource.empty.id
  description = "A single attribute reference."
}

output "splat" {
  value       = random_pet.names[*].id
  description = "A splat expression over a counted resource."
}

output "for_expression" {
  value       = { for k, v in random_string.per_key : k => v.result }
  description = "A for expression over a for_each resource."
}

output "many_references" {
  value = {
    vars    = [var.string_default, var.number_default, var.bool_true]
    locals  = [local.literal_string, local.interpolated]
    data    = data.aws_caller_identity.current.account_id
    module  = module.child.passthrough
    managed = aws_s3_bucket.replica.arn
  }
  description = "References every kind of address: var, local, data, module, managed resource."
}

output "conditional" {
  value       = var.bool_true ? module.child.passthrough : null
  description = "A conditional referencing a module output."
}

output "sensitive_value" {
  value       = random_password.db.result
  description = "A sensitive output."
  sensitive   = true
}

output "ephemeral_value" {
  value       = ephemeral.random_password.session.result
  description = "An ephemeral output (only valid in a child module)."
  ephemeral   = true
  sensitive   = true
}

output "with_depends_on" {
  value       = "done"
  description = "An output with explicit depends_on."
  depends_on  = [time_sleep.wait, module.registry_for_each]
}

output "with_precondition" {
  value       = local.primary
  description = "An output with a precondition block."

  precondition {
    condition     = length(var.regions) > 0
    error_message = "At least one region is required."
  }
}

output "pipe_in_value" {
  value       = "a | b | c"
  description = "Pipe characters in the value and description | here."
}

output "heredoc" {
  value       = <<-EOT
    multi
    line
  EOT
  description = "A heredoc value."
}

# Comment-only output with no description.
output "comment_only" {
  value = local.literal_number
}

output "no_description" {
  value = local.literal_list
}
