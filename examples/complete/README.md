# Module `complete`

<!-- VERSION-START -->
This documentation was generated with [open-docs dev](https://github.com/RemoteRabbit/open-docs)
<!-- VERSION-END -->

- **Required core:** `>= 1.10.0, != 1.10.1`

## Required Providers

| Name | Source | Version | Position |
|------|--------|---------|----------|
| `aws` | `hashicorp/aws` | `~> 5.70` | [`./versions.tf`](./versions.tf#L9-L13) |
| `http` | `hashicorp/http` | - | [`./versions.tf`](./versions.tf#L26-L28) |
| `local` | `registry.opentofu.org/hashicorp/local` | `2.5.2` | [`./versions.tf`](./versions.tf#L29-L32) |
| `null` | `hashicorp/null` | `~> 3.2` | [`./versions.tf`](./versions.tf#L18-L21) |
| `random` | `hashicorp/random` | `>= 3.6.0, < 4.0.0` | [`./versions.tf`](./versions.tf#L14-L17) |
| `time` | - | - | [`./versions.tf`](./versions.tf#L34) |
| `tls` | `hashicorp/tls` | `~> 4.0` | [`./versions.tf`](./versions.tf#L22-L25) |

## Providers

| Name | Alias | Position |
|------|-------|----------|
| `aws` | - | [`./providers.tf`](./providers.tf#L1) |
| `aws` | `by_region` | [`./providers.tf`](./providers.tf#L6) |
| `aws` | `audit` | [`./providers.tf`](./providers.tf#L12) |
| `random` | - | [`./providers.tf`](./providers.tf#L21) |
| `http` | - | [`./providers.tf`](./providers.tf#L23) |

## Inputs

| Name | Type | Default | Description | Comment | Sensitive | Nullable | Ephemeral | Validation Count | Position |
|------|------|---------|-------------|---------|-----------|----------|-----------|------------------|----------|
| `json_variable` | `string` | `"from-json"` | A variable declared in a .tf.json file. | - | - | - | - | 0 | [`./extra.tf.json`](./extra.tf.json#L4) |
| `json_object` | `object({name = string, size = optional(number, 1)})` | `{name = "json"}` | An object-typed variable declared in JSON. | - | - | - | - | 0 | [`./extra.tf.json`](./extra.tf.json#L9) |
| `string_required` | `string` | - | A required string with no default. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L3) |
| `string_default` | `string` | `"hello"` | A string with a default. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L8) |
| `string_empty` | `string` | `""` | A string whose default is the empty string. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L14) |
| `number_default` | `number` | `42` | A number. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L20) |
| `number_float` | `number` | `3.14159` | A floating point number. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L26) |
| `number_negative` | `number` | `-1` | A negative number. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L32) |
| `bool_true` | `bool` | `true` | A bool defaulting to true. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L38) |
| `bool_false` | `bool` | `false` | A bool defaulting to false. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L44) |
| `null_default` | `string` | `null` | Optional because the default is null. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L50) |
| `untyped` | - | - | No type constraint at all. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L56) |
| `any_type` | `any` | `{anything = ["goes", 1, true]}` | Type `any` with a mixed default. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L60) |
| `list_of_strings` | `list(string)` | `["a", "b", "c"]` | A list of strings. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L68) |
| `list_empty` | `list(number)` | `[]` | An empty list. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L74) |
| `set_of_strings` | `set(string)` | `["x", "y"]` | A set of strings. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L80) |
| `map_of_numbers` | `map(number)` | `{large = 100, medium = 10, small = 1}` | A map of numbers. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L86) |
| `map_empty` | `map(string)` | `{}` | An empty map. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L96) |
| `tuple_type` | `tuple([string, number, bool])` | `["one", 2, true]` | A tuple with mixed element types. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L102) |
| `list_of_lists` | `list(list(string))` | `[["a", "b"], ["c"]]` | Nested collections. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L108) |
| `map_of_lists` | `map(list(string))` | `{admins = ["alice", "bob"], readers = []}` | A map whose values are lists. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L114) |
| `regions` | `set(string)` | `["us-east-1", "us-west-2", "eu-west-1"]` | Regions for the multi-instance `aws.by_region` provider. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L123) |
| `simple_object` | `object({enabled = bool, name = string, port = number})` | `{enabled = true, name = "svc", port = 8080}` | An object with only required attributes. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L131) |
| `optional_object` | `object({image = optional(string), labels = optional(map(string), {}), name = string, ports = optional(list(number), [80, 443]), replicas = optional(number, 1)})` | - | An object with optional attributes and defaults. Required, so no default. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L145) |
| `nested_object` | `object({database = optional(object({engine = optional(string, "postgres"), storage = optional(object({encrypted = optional(bool, true), iops = optional(number), size_gb = optional(number, 20)}), {encrypted = null, iops = null, size_gb = null}), version = optional(string, "16")}), {engine = null, storage = null, version = null}), network = object({cidr = string, subnets = list(object({cidr = string, name = string, public = optional(bool, false)}))})})` | `{network = {cidr = "10.0.0.0/16", subnets = [{cidr = "10.0.1.0/24", name = "a", public = true}, {cidr = "10.0.2.0/24", name = "b"}]}}` | Deeply nested objects with optional attributes at several levels. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L156) |
| `list_of_objects` | `list(object({name = string, value = string}))` | `[{name = "first", value = "1"}, {name = "second", value = "2"}]` | A list of objects. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L188) |
| `map_of_objects` | `map(object({instance_type = string, max_size = optional(number, 3), min_size = optional(number, 1)}))` | `{web = {instance_type = "t3.small"}, worker = {instance_type = "c6i.large", max_size = 10}}` | A map of objects, the classic node-group shape. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L200) |
| `sensitive_string` | `string` | `"hunter2"` | A sensitive variable with a default. | - | ✓ | - | - | 0 | [`./variables.tf`](./variables.tf#L215) |
| `state_passphrase` | `string` | - | Passphrase for OpenTofu state encryption. | - | ✓ | - | - | 0 | [`./variables.tf`](./variables.tf#L222) |
| `nullable_false` | `string` | `"never-null"` | Explicitly not nullable. | - | - | ✗ | - | 0 | [`./variables.tf`](./variables.tf#L228) |
| `nullable_true` | `string` | `null` | Explicitly nullable. | - | - | ✓ | - | 0 | [`./variables.tf`](./variables.tf#L235) |
| `ephemeral_token` | `string` | - | An ephemeral, sensitive input that is never persisted. | - | ✓ | - | ✓ | 0 | [`./variables.tf`](./variables.tf#L242) |
| `many_validations` | `string` | `"abc-123"` | Three validation blocks. | - | - | - | - | 3 | [`./variables.tf`](./variables.tf#L249) |
| `pipe_in_description` | `string` | `"a\|b"` | Contains a pipe \| character that would break a Markdown table \| if unescaped. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L272) |
| `markdown_in_description` | `string` | `"*"` | Has **bold**, _italic_, `code`, a [link](https://opentofu.org), and <b>raw HTML</b>. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L278) |
| `multiline_description` | `string` | `"x"` | A heredoc description spanning several lines. - It has a bullet list - and a blank line above And a trailing paragraph. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L284) |
| `unicode_description` | `string` | `"héllo wörld ✓"` | Unicode: 日本語, emoji 🚀, arrows → ←, and accents café. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L297) |
| `long_description` | `string` | `"long"` | This description is intentionally very long so that table cells wrap or overflow in renderers. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L303) |
| `long_default` | `list(string)` | `["10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24", "10.0.7.0/24", "10.0.8.0/24", "10.0.9.0/24"]` | A default value long enough to stress column widths. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L309) |
| `quotes_in_default` | `string` | `"she said \"hi\" and left `quickly`"` | Escaped quotes and backticks inside a default. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L315) |
| `heredoc_default` | `string` | `"#!/bin/bash\necho \"user data\"\n"` | A multi-line heredoc default. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L321) |
| `interpolated_default_like` | `string` | `"${not_interpolated}"` | A literal that looks like interpolation. | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L330) |
| `comment_only` | `string` | `"c"` | - | Leading doc-comment, no description attribute. The renderer should surface this in the Comment column. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L338) |
| `comment_and_description` | `string` | `"both"` | Has both a description and a leading block comment. | Block-style doc-comment above a variable that also has a description. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L346) |
| `slash_comment` | `number` | `0` | - | Double-slash comment style. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L353) |
| `no_description` | `string` | `"undocumented"` | - | - | - | - | - | 0 | [`./variables.tf`](./variables.tf#L358) |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|-----------|-----------|------------|----------|
| json_output | `"${null_resource.from_json.id}"` | `null_resource.from_json` | An output declared in a .tf.json file. | - | - | - | - | [`./extra.tf.json`](./extra.tf.json#L28) |
| literal | `"constant"` | - | A literal string output. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L1) |
| simple_ref | `null_resource.empty.id` | `null_resource.empty` | A single attribute reference. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L6) |
| splat | `random_pet.names[*].id` | `random_pet.names` | A splat expression over a counted resource. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L11) |
| for_expression | `{ for k, v in random_string.per_key : k => v.result }` | `random_string.per_key` | A for expression over a for_each resource. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L16) |
| many_references | `{ vars = [var.string_default, var.number_default, var.bool_true] locals = [local.literal_string, local.interpolated] data = data.aws_caller_identity.current.account_id module = module.child.passthrough managed = aws_s3_bucket.replica.arn }` | `var.string_default, var.number_default, var.bool_true, local.literal_string, local.interpolated, data.aws_caller_identity.current, module.child.passthrough, aws_s3_bucket.replica` | References every kind of address: var, local, data, module, managed resource. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L21) |
| conditional | `var.bool_true ? module.child.passthrough : null` | `var.bool_true, module.child.passthrough` | A conditional referencing a module output. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L32) |
| sensitive_value | `random_password.db.result` | `random_password.db` | A sensitive output. | - | ✓ | - | - | [`./outputs.tf`](./outputs.tf#L37) |
| ephemeral_value | `ephemeral.random_password.session.result` | `ephemeral.random_password.session` | An ephemeral output (only valid in a child module). | - | ✓ | ✓ | - | [`./outputs.tf`](./outputs.tf#L43) |
| with_depends_on | `"done"` | - | An output with explicit depends_on. | - | - | - | `time_sleep.wait, module.registry_for_each` | [`./outputs.tf`](./outputs.tf#L50) |
| with_precondition | `local.primary` | `local.primary` | An output with a precondition block. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L56) |
| pipe_in_value | `"a \| b \| c"` | - | Pipe characters in the value and description \| here. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L66) |
| heredoc | `<<-EOT multi line EOT` | - | A heredoc value. | - | - | - | - | [`./outputs.tf`](./outputs.tf#L71) |
| comment_only | `local.literal_number` | `local.literal_number` | - | Comment-only output with no description. | - | - | - | [`./outputs.tf`](./outputs.tf#L80) |
| no_description | `local.literal_list` | `local.literal_list` | - | - | - | - | - | [`./outputs.tf`](./outputs.tf#L84) |

## Locals

| Name | Value | References | Position |
|------|-------|------------|----------|
| conditional | `var.bool_true ? "yes" : "no"` | `var.bool_true` | [`./locals.tf`](./locals.tf#L11) |
| for_list | `[for s in var.list_of_strings : upper(s)]` | `var.list_of_strings` | [`./locals.tf`](./locals.tf#L12) |
| for_map | `{ for k, v in var.map_of_numbers : k => v * 2 if v > 1 }` | `var.map_of_numbers` | [`./locals.tf`](./locals.tf#L13) |
| from_var | `var.string_default` | `var.string_default` | [`./locals.tf`](./locals.tf#L9) |
| function_chain | `join(",", sort(distinct(concat(var.list_of_strings, ["z"]))))` | `var.list_of_strings` | [`./locals.tf`](./locals.tf#L15) |
| heredoc_value | `<<-EOT line one line two EOT` | - | [`./locals.tf`](./locals.tf#L18-L21) |
| interpolated | `"${var.string_default}-${local.literal_number}"` | `var.string_default, local.literal_number` | [`./locals.tf`](./locals.tf#L10) |
| json_local | `"${var.json_variable}-local"` | `var.json_variable` | [`./extra.tf.json`](./extra.tf.json#L16) |
| literal_bool | `true` | - | [`./locals.tf`](./locals.tf#L4) |
| literal_list | `[1, 2, 3]` | - | [`./locals.tf`](./locals.tf#L6) |
| literal_map | `{ a = 1, b = "two" }` | - | [`./locals.tf`](./locals.tf#L7) |
| literal_null | `null` | - | [`./locals.tf`](./locals.tf#L5) |
| literal_number | `7` | - | [`./locals.tf`](./locals.tf#L3) |
| literal_string | `"plain"` | - | [`./locals.tf`](./locals.tf#L2) |
| many_refs | `merge(var.map_of_objects, { extra = { instance_type = local.literal_string } }, module.child.passthrough)` | `var.map_of_objects, local.literal_string, module.child.passthrough` | [`./locals.tf`](./locals.tf#L16) |
| pipe_value | `"left \| right"` | - | [`./locals.tf`](./locals.tf#L17) |
| primary | `local.regions_list[0]` | `local.regions_list` | [`./locals.tf`](./locals.tf#L27) |
| regions_list | `tolist(var.regions)` | `var.regions` | [`./locals.tf`](./locals.tf#L26) |
| splat | `random_pet.names[*].id` | `random_pet.names` | [`./locals.tf`](./locals.tf#L14) |

## Managed Resources

| Name | Mode | Type | Count | For Each | Depends On | Nested Body | LifeCycle | Comment | Position |
|------|------|------|-------|----------|------------|-------------|-----------|---------|----------|
| from_json | managed | `null_resource` | - | - |  | - | - | - | [`./extra.tf.json`](./extra.tf.json#L20) |
| empty | managed | `null_resource` | - | - |  | - | - | Minimal resource: no attributes, no meta-arguments. | [`./resources.tf`](./resources.tf#L4) |
| names | managed | `random_pet` | `3` | - |  | `length, prefix, separator` | - | - | [`./resources.tf`](./resources.tf#L6) |
| per_key | managed | `random_string` | - | `var.map_of_numbers` |  | `length, special` | - | - | [`./resources.tf`](./resources.tf#L14) |
| db | managed | `random_password` | - | - |  | `length, override_special, special` | - | - | [`./resources.tf`](./resources.tf#L21) |
| kitchen_sink | managed | `aws_s3_bucket` | - | `{ for name, cfg in var.map_of_objects : name => cfg if cfg.max_size > 1 }` | random_pet.names, null_resource.empty, module.child | `bucket, force_destroy, tags` | `create_before_destroy=true; prevent_destroy=true; ignore_changes=[tags,force_destroy]; replace_triggered_by=[random_password.db.id,null_resource.empty]; preconditions=2; postconditions=1` | Resource with every meta-argument and a full lifecycle block. | [`./resources.tf`](./resources.tf#L30) |
| replica | managed | `aws_s3_bucket` | - | - |  | `bucket` | - | - | [`./resources.tf`](./resources.tf#L71) |
| ignore_all | managed | `aws_iam_role` | - | - |  | `assume_role_policy, name` | `ignore_changes=[all]` | - | [`./resources.tf`](./resources.tf#L76) |
| nested | managed | `aws_security_group` | - | - |  | `description, name_prefix` | - | Nested blocks, dynamic blocks, and provisioners. | [`./resources.tf`](./resources.tf#L95) |
| ca | managed | `tls_private_key` | - | - |  | `algorithm, ecdsa_curve` | - | - | [`./resources.tf`](./resources.tf#L138) |
| rendered | managed | `local_file` | - | - |  | `content, file_permission, filename` | - | - | [`./resources.tf`](./resources.tf#L143) |
| wait | managed | `time_sleep` | - | - | aws_s3_bucket.kitchen_sink | `create_duration` | - | - | [`./resources.tf`](./resources.tf#L149) |
| replacement | managed | `terraform_data` | - | - |  | `input, triggers_replace` | - | - | [`./resources.tf`](./resources.tf#L154) |

## Data Resources

| Name | Mode | Type | Count | For Each | Depends On | Nested Body | LifeCycle | Comment | Position |
|------|------|------|-------|----------|------------|-------------|-----------|---------|----------|
| current | data | `aws_caller_identity` | - | - |  | - | - | - | [`./resources.tf`](./resources.tf#L161) |
| per_region | data | `aws_iam_policy_document` | - | `var.regions` |  | - | - | - | [`./resources.tf`](./resources.tf#L163) |
| health | data | `http` | `var.bool_true ? 1 : 0` | - |  | `request_headers, url` | `postconditions=1` | - | [`./resources.tf`](./resources.tf#L173) |
| readme | data | `local_file` | - | - | local_file.rendered | `filename` | - | - | [`./resources.tf`](./resources.tf#L190) |

## Ephemeral Resources

| Name | Mode | Type | Count | For Each | Depends On | Nested Body | LifeCycle | Comment | Position |
|------|------|------|-------|----------|------------|-------------|-----------|---------|----------|
| session | ephemeral | `random_password` | - | - |  | `length, special` | - | - | [`./resources.tf`](./resources.tf#L197) |
| db | ephemeral | `aws_secretsmanager_secret_version` | `var.bool_true ? 1 : 0` | - | aws_iam_role.ignore_all | `secret_id` | - | - | [`./resources.tf`](./resources.tf#L202) |
| per_region | ephemeral | `tls_private_key` | - | `var.regions` |  | `algorithm` | - | - | [`./resources.tf`](./resources.tf#L211) |

## Module Calls

| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |
|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|
| child | ./modules/child |  | - | - | - | `aws=aws.audit, aws.replica=aws.replica` | `count_value, enabled, name, objects, tags` | Local module with explicit provider mapping and several input shapes. | [`./modules.tf`](./modules.tf#L2) |
| registry_count | terraform-aws-modules/s3-bucket/aws | 4.1.2 | `var.bool_false ? 1 : 0` | - | - | - | `bucket` | Registry module with count. | [`./modules.tf`](./modules.tf#L18) |
| registry_for_each | cloudposse/label/null | >= 0.25.0, < 1.0.0 | - | `toset(var.list_of_strings)` | `module.child, random_pet.names` | - | `name` | Registry module with for_each and depends_on. | [`./modules.tf`](./modules.tf#L27) |
| git_source | git::https://github.com/example/terraform-modules.git//network?ref=v1.2.3 |  | - | - | - | - | - | Git source pinned to a ref, no version argument. | [`./modules.tf`](./modules.tf#L38) |
| github_shorthand | github.com/example/terraform-module-foo |  | - | - | - | - | - | GitHub shorthand and an HTTP archive. | [`./modules.tf`](./modules.tf#L43) |
| http_archive | https://example.com/modules/vpc-module.zip |  | - | - | - | - | - | - | [`./modules.tf`](./modules.tf#L47) |
| dynamic_source |  |  | - | - | - | - | - | OpenTofu early evaluation: non-literal source and version. | [`./modules.tf`](./modules.tf#L52) |
| no_inputs | ./modules/child |  | - | - | - | - | - | Module call with no inputs at all. | [`./modules.tf`](./modules.tf#L58) |

## Moved

| From | To | Position |
|------|----|----------|
| null_resource.old_name | null_resource.empty | [`./blocks.tf`](./blocks.tf#L3) |
| random_pet.single | random_pet.names[0] | [`./blocks.tf`](./blocks.tf#L8) |
| module.old_child | module.child | [`./blocks.tf`](./blocks.tf#L13) |
| module.child.aws_s3_bucket.legacy | module.child.aws_s3_bucket.this | [`./blocks.tf`](./blocks.tf#L18) |

## Imports

| To | ID | Provider | Position |
|----|----|----------|----------|
| `null_resource.empty` | `"static-id-123"` | - | [`./blocks.tf`](./blocks.tf#L25) |
| `aws_iam_role.ignore_all` | `"complete-audit"` | `aws.audit` | [`./blocks.tf`](./blocks.tf#L30) |
| `aws_s3_bucket.replica` | `"${local.literal_string}-replica"` | - | [`./blocks.tf`](./blocks.tf#L36) |
| `aws_s3_bucket.kitchen_sink[each.key]` | `"complete-${each.key}"` | - | [`./blocks.tf`](./blocks.tf#L42) |

## Removed

| From | Destroy on Drop | Position |
|------|-----------------|----------|
| `null_resource.gone` | - | [`./blocks.tf`](./blocks.tf#L50) |
| `aws_s3_bucket.retired` | ✗ | [`./blocks.tf`](./blocks.tf#L54) |
| `module.deprecated_child` | ✓ | [`./blocks.tf`](./blocks.tf#L62) |

## Checks

| Name | Data Source | Assertions | Position |
|------|-------------|------------|----------|
| `health_endpoint` | `http.probe` | 2 | [`./blocks.tf`](./blocks.tf#L72) |
| `inputs_sane` | - | 1 | [`./blocks.tf`](./blocks.tf#L89) |
| `three_assertions` | - | 3 | [`./blocks.tf`](./blocks.tf#L96) |
