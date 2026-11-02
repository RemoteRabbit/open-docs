# Module `complete`

<!-- VERSION-START -->
This documentation was generated with [open-docs v1](https://github.com/RemoteRabbit/open-docs)
<!-- VERSION-END -->

- **Required core:** `>= 1.10.0, != 1.10.1`

## Required Providers

| Name | Source | Version | Position |
|------|--------|---------|----------|
| `aws` | `hashicorp/aws` | `~> 5.70` | [`./versions.tf`](./versions.tf#L9-L13) [9:5:250 -> 13:6:414] |
| `http` | `hashicorp/http` | - | [`./versions.tf`](./versions.tf#L26-L28) [26:5:663 -> 28:6:709] |
| `local` | `registry.opentofu.org/hashicorp/local` | `2.5.2` | [`./versions.tf`](./versions.tf#L29-L32) [29:5:714 -> 32:6:809] |
| `null` | `hashicorp/null` | `~> 3.2` | [`./versions.tf`](./versions.tf#L18-L21) [18:5:511 -> 21:6:583] |
| `random` | `hashicorp/random` | `>= 3.6.0, < 4.0.0` | [`./versions.tf`](./versions.tf#L14-L17) [14:5:419 -> 17:6:506] |
| `time` | - | - | [`./versions.tf`](./versions.tf#L34) [34:5:857 -> 34:14:866] |
| `tls` | `hashicorp/tls` | `~> 4.0` | [`./versions.tf`](./versions.tf#L22-L25) [22:5:588 -> 25:6:658] |

## Providers

| Name | Position |
|------|----------|
| `aws` | [`./providers.tf`](./providers.tf#L1) [1:1:0 -> 1:15:14] |
| `aws` | [`./providers.tf`](./providers.tf#L6) [6:1:84 -> 6:15:98] |
| `aws` | [`./providers.tf`](./providers.tf#L12) [12:1:176 -> 12:15:190] |
| `random` | [`./providers.tf`](./providers.tf#L21) [21:1:313 -> 21:18:330] |
| `http` | [`./providers.tf`](./providers.tf#L23) [23:1:335 -> 23:16:350] |

## Inputs

| Name | Type | Default | Description | Comment | Sensitive | Nullable | Ephemeral | Validation Count | Position |
|------|------|---------|-------------|---------|-----------|----------|-----------|------------------|----------|
| `json_variable` | `string` | `"from-json"` | A variable declared in a .tf.json file. |  | - | - | - | 0 | [`./extra.tf.json`](./extra.tf.json#L4) [4:22:118 -> 4:23:119] |
| `json_object` | `object({name = string, size = optional(number, 1)})` | `{name = "json"}` | An object-typed variable declared in JSON. |  | - | - | - | 0 | [`./extra.tf.json`](./extra.tf.json#L9) [9:20:263 -> 9:21:264] |
| `string_required` | `string` | - | A required string with no default. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L3) [3:1:79 -> 3:27:105] |
| `string_default` | `string` | `"hello"` | A string with a default. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L8) [8:1:187 -> 8:26:212] |
| `string_empty` | `string` | `""` | A string whose default is the empty string. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L14) [14:1:308 -> 14:24:331] |
| `number_default` | `number` | `42` | A number. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L20) [20:1:441 -> 20:26:466] |
| `number_float` | `number` | `3.14159` | A floating point number. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L26) [26:1:542 -> 26:24:565] |
| `number_negative` | `number` | `-1` | A negative number. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L32) [32:1:661 -> 32:27:687] |
| `bool_true` | `bool` | `true` | A bool defaulting to true. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L38) [38:1:772 -> 38:21:792] |
| `bool_false` | `bool` | `false` | A bool defaulting to false. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L44) [44:1:885 -> 44:22:906] |
| `null_default` | `string` | `null` | Optional because the default is null. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L50) [50:1:1001 -> 50:24:1024] |
| `untyped` | - | - | No type constraint at all. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L56) [56:1:1130 -> 56:19:1148] |
| `any_type` | `any` | `{anything = ["goes", 1, true]}` | Type `any` with a mixed default. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L60) [60:1:1199 -> 60:20:1218] |
| `list_of_strings` | `list(string)` | `["a", "b", "c"]` | A list of strings. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L68) [68:1:1424 -> 68:27:1450] |
| `list_empty` | `list(number)` | `[]` | An empty list. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L74) [74:1:1554 -> 74:22:1575] |
| `set_of_strings` | `set(string)` | `["x", "y"]` | A set of strings. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L80) [80:1:1662 -> 80:26:1687] |
| `map_of_numbers` | `map(number)` | `{large = 100, medium = 10, small = 1}` | A map of numbers. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L86) [86:1:1784 -> 86:26:1809] |
| `map_empty` | `map(string)` | `{}` | An empty map. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L96) [96:1:1938 -> 96:21:1958] |
| `tuple_type` | `tuple([string, number, bool])` | `["one", 2, true]` | A tuple with mixed element types. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L102) [102:1:2043 -> 102:22:2064] |
| `list_of_lists` | `list(list(string))` | `[["a", "b"], ["c"]]` | Nested collections. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L108) [108:1:2201 -> 108:25:2225] |
| `map_of_lists` | `map(list(string))` | `{admins = ["alice", "bob"], readers = []}` | A map whose values are lists. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L114) [114:1:2340 -> 114:24:2363] |
| `regions` | `set(string)` | `["us-east-1", "us-west-2", "eu-west-1"]` | Regions for the multi-instance `aws.by_region` provider. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L123) [123:1:2510 -> 123:19:2528] |
| `simple_object` | `object({enabled = bool, name = string, port = number})` | `{enabled = true, name = "svc", port = 8080}` | An object with only required attributes. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L131) [131:1:2773 -> 131:25:2797] |
| `optional_object` | `object({image = optional(string), labels = optional(map(string), {}), name = string, ports = optional(list(number), [80, 443]), replicas = optional(number, 1)})` | - | An object with optional attributes and defaults. Required, so no default. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L145) [145:1:3022 -> 145:27:3048] |
| `nested_object` | `object({database = optional(object({engine = optional(string, "postgres"), storage = optional(object({encrypted = optional(bool, true), iops = optional(number), size_gb = optional(number, 20)}), {encrypted = null, iops = null, size_gb = null}), version = optional(string, "16")}), {engine = null, storage = null, version = null}), network = object({cidr = string, subnets = list(object({cidr = string, name = string, public = optional(bool, false)}))})})` | `{network = {cidr = "10.0.0.0/16", subnets = [{cidr = "10.0.1.0/24", name = "a", public = true}, {cidr = "10.0.2.0/24", name = "b"}]}}` | Deeply nested objects with optional attributes at several levels. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L156) [156:1:3348 -> 156:25:3372] |
| `list_of_objects` | `list(object({name = string, value = string}))` | `[{name = "first", value = "1"}, {name = "second", value = "2"}]` | A list of objects. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L188) [188:1:4158 -> 188:27:4184] |
| `map_of_objects` | `map(object({instance_type = string, max_size = optional(number, 3), min_size = optional(number, 1)}))` | `{web = {instance_type = "t3.small"}, worker = {instance_type = "c6i.large", max_size = 10}}` | A map of objects, the classic node-group shape. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L200) [200:1:4387 -> 200:26:4412] |
| `sensitive_string` | `string` | `"hunter2"` | A sensitive variable with a default. |  | ✓ | - | - | 0 | [`./variables.tf`](./variables.tf#L215) [215:1:4821 -> 215:28:4848] |
| `state_passphrase` | `string` | - | Passphrase for OpenTofu state encryption. |  | ✓ | - | - | 0 | [`./variables.tf`](./variables.tf#L222) [222:1:4979 -> 222:28:5006] |
| `nullable_false` | `string` | `"never-null"` | Explicitly not nullable. |  | - | ✗ | - | 0 | [`./variables.tf`](./variables.tf#L228) [228:1:5116 -> 228:26:5141] |
| `nullable_true` | `string` | `null` | Explicitly nullable. |  | - | ✓ | - | 0 | [`./variables.tf`](./variables.tf#L235) [235:1:5264 -> 235:25:5288] |
| `ephemeral_token` | `string` | - | An ephemeral, sensitive input that is never persisted. |  | ✓ | - | ✓ | 0 | [`./variables.tf`](./variables.tf#L242) [242:1:5398 -> 242:27:5424] |
| `many_validations` | `string` | `"abc-123"` | Three validation blocks. |  | - | - | - | 3 | [`./variables.tf`](./variables.tf#L249) [249:1:5568 -> 249:28:5595] |
| `pipe_in_description` | `string` | `"a\|b"` | Contains a pipe | character that would break a Markdown table | if unescaped. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L272) [272:1:6212 -> 272:31:6242] |
| `markdown_in_description` | `string` | `"*"` | Has **bold**, _italic_, `code`, a [link](https://opentofu.org), and <b>raw HTML</b>. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L278) [278:1:6389 -> 278:35:6423] |
| `multiline_description` | `string` | `"x"` | A heredoc description spanning several lines.

- It has a bullet list
- and a blank line above

And a trailing paragraph.
 |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L284) [284:1:6575 -> 284:33:6607] |
| `unicode_description` | `string` | `"héllo wörld ✓"` | Unicode: 日本語, emoji 🚀, arrows → ←, and accents café. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L297) [297:1:6823 -> 297:31:6853] |
| `long_description` | `string` | `"long"` | This description is intentionally very long so that table cells wrap or overflow in renderers. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L303) [303:1:7003 -> 303:28:7030] |
| `long_default` | `list(string)` | `["10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24", "10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24", "10.0.7.0/24", "10.0.8.0/24", "10.0.9.0/24"]` | A default value long enough to stress column widths. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L309) [309:1:7427 -> 309:24:7450] |
| `quotes_in_default` | `string` | `"she said \"hi\" and left `quickly`"` | Escaped quotes and backticks inside a default. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L315) [315:1:7723 -> 315:29:7751] |
| `heredoc_default` | `string` | `"#!/bin/bash\necho \"user data\"\n"` | A multi-line heredoc default. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L321) [321:1:7898 -> 321:27:7924] |
| `interpolated_default_like` | `string` | `"${not_interpolated}"` | A literal that looks like interpolation. |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L330) [330:1:8067 -> 330:37:8103] |
| `comment_only` | `string` | `"c"` |  | Leading doc-comment, no description attribute.
The renderer should surface this in the Comment column. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L338) [338:1:8337 -> 338:24:8360] |
| `comment_and_description` | `string` | `"both"` | Has both a description and a leading block comment. | Block-style doc-comment above a variable that also has a description. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L346) [346:1:8481 -> 346:35:8515] |
| `slash_comment` | `number` | `0` |  | Double-slash comment style. | - | - | - | 0 | [`./variables.tf`](./variables.tf#L353) [353:1:8668 -> 353:25:8692] |
| `no_description` | `string` | `"undocumented"` |  |  | - | - | - | 0 | [`./variables.tf`](./variables.tf#L358) [358:1:8731 -> 358:26:8756] |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|-----------|-----------|------------|----------|
| json_output | `"${null_resource.from_json.id}"` | `null_resource.from_json` | An output declared in a .tf.json file. |  | - | - | - | [`./extra.tf.json`](./extra.tf.json#L28) [28:20:701 -> 28:21:702] |
| literal | `"constant"` | - | A literal string output. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L1) [1:1:0 -> 1:17:16] |
| simple_ref | `null_resource.empty.id` | `null_resource.empty` | A single attribute reference. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L6) [6:1:92 -> 6:20:111] |
| splat | `random_pet.names[*].id` | `random_pet.names` | A splat expression over a counted resource. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L11) [11:1:204 -> 11:15:218] |
| for_expression | `{ for k, v in random_string.per_key : k => v.result }` | `random_string.per_key` | A for expression over a for_each resource. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L16) [16:1:325 -> 16:24:348] |
| many_references | `{ vars = [var.string_default, var.number_default, var.bool_true] locals = [local.literal_string, local.interpolated] data = data.aws_caller_identity.current.account_id module = module.child.passthrough managed = aws_s3_bucket.replica.arn }` | `var.string_default, var.number_default, var.bool_true, local.literal_string, local.interpolated, data.aws_caller_identity.current, module.child.passthrough, aws_s3_bucket.replica` | References every kind of address: var, local, data, module, managed resource. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L21) [21:1:485 -> 21:25:509] |
| conditional | `var.bool_true ? module.child.passthrough : null` | `var.bool_true, module.child.passthrough` | A conditional referencing a module output. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L32) [32:1:891 -> 32:21:911] |
| sensitive_value | `random_password.db.result` | `random_password.db` | A sensitive output. |  | ✓ | - | - | [`./outputs.tf`](./outputs.tf#L37) [37:1:1042 -> 37:25:1066] |
| ephemeral_value | `ephemeral.random_password.session.result` | `ephemeral.random_password.session` | An ephemeral output (only valid in a child module). |  | ✓ | ✓ | - | [`./outputs.tf`](./outputs.tf#L43) [43:1:1173 -> 43:25:1197] |
| with_depends_on | `"done"` | - | An output with explicit depends_on. |  | - | - | `time_sleep.wait, module.registry_for_each` | [`./outputs.tf`](./outputs.tf#L50) [50:1:1372 -> 50:25:1396] |
| with_precondition | `local.primary` | `local.primary` | An output with a precondition block. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L56) [56:1:1539 -> 56:27:1565] |
| pipe_in_value | `"a \| b \| c"` | - | Pipe characters in the value and description | here. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L66) [66:1:1777 -> 66:23:1799] |
| heredoc | `<<-EOT multi line EOT` | - | A heredoc value. |  | - | - | - | [`./outputs.tf`](./outputs.tf#L71) [71:1:1904 -> 71:17:1920] |
| comment_only | `local.literal_number` | `local.literal_number` |  | Comment-only output with no description. | - | - | - | [`./outputs.tf`](./outputs.tf#L80) [80:1:2052 -> 80:22:2073] |
| no_description | `local.literal_list` | `local.literal_list` |  |  | - | - | - | [`./outputs.tf`](./outputs.tf#L84) [84:1:2110 -> 84:24:2133] |

## Locals

| Name |
|------|
| conditional |
| for_list |
| for_map |
| from_var |
| function_chain |
| heredoc_value |
| interpolated |
| json_local |
| literal_bool |
| literal_list |
| literal_map |
| literal_null |
| literal_number |
| literal_string |
| many_refs |
| pipe_value |
| primary |
| regions_list |
| splat |

## Managed Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `managed` | `null_resource` | `from_json` | - | - | - | - | - | - | - | [`./extra.tf.json`](./extra.tf.json#L20) [20:20:574 -> 20:21:575] |
| `managed` | `null_resource` | `empty` | - | - | - | - | - | - | Minimal resource: no attributes, no meta-arguments. | [`./resources.tf`](./resources.tf#L4) [4:1:134 -> 4:33:166] |
| `managed` | `random_pet` | `names` | - | `3` | - | - | `length, prefix, separator` | - | - | [`./resources.tf`](./resources.tf#L6) [6:1:171 -> 6:30:200] |
| `managed` | `random_string` | `per_key` | - | - | `var.map_of_numbers` | - | `length, special` | - | - | [`./resources.tf`](./resources.tf#L14) [14:1:286 -> 14:35:320] |
| `managed` | `random_password` | `db` | - | - | - | - | `length, override_special, special` | - | - | [`./resources.tf`](./resources.tf#L21) [21:1:404 -> 21:32:435] |
| `managed` | `aws_s3_bucket` | `kitchen_sink` | `aws.by_region` | - | `{ for name, cfg in var.map_of_objects : name => cfg if cfg.max_size > 1 }` | `random_pet.names, null_resource.empty, module.child` | `bucket, force_destroy, tags` | `create_before_destroy=true; prevent_destroy=true; ignore_changes=[tags,force_destroy]; replace_triggered_by=[random_password.db.id,null_resource.empty]; preconditions=2; postconditions=1` | Resource with every meta-argument and a full lifecycle block. | [`./resources.tf`](./resources.tf#L30) [30:1:607 -> 30:40:646] |
| `managed` | `aws_s3_bucket` | `replica` | `aws.replica` | - | - | - | `bucket` | - | - | [`./resources.tf`](./resources.tf#L71) [71:1:1621 -> 71:35:1655] |
| `managed` | `aws_iam_role` | `ignore_all` | `aws.audit` | - | - | - | `assume_role_policy, name` | `ignore_changes=[all]` | - | [`./resources.tf`](./resources.tf#L76) [76:1:1718 -> 76:37:1754] |
| `managed` | `aws_security_group` | `nested` | - | - | - | - | `description, name_prefix` | - | Nested blocks, dynamic blocks, and provisioners. | [`./resources.tf`](./resources.tf#L95) [95:1:2116 -> 95:39:2154] |
| `managed` | `tls_private_key` | `ca` | - | - | - | - | `algorithm, ecdsa_curve` | - | - | [`./resources.tf`](./resources.tf#L138) [138:1:2927 -> 138:32:2958] |
| `managed` | `local_file` | `rendered` | - | - | - | - | `content, file_permission, filename` | - | - | [`./resources.tf`](./resources.tf#L143) [143:1:3011 -> 143:33:3043] |
| `managed` | `time_sleep` | `wait` | - | - | - | `aws_s3_bucket.kitchen_sink` | `create_duration` | - | - | [`./resources.tf`](./resources.tf#L149) [149:1:3239 -> 149:29:3267] |
| `managed` | `terraform_data` | `replacement` | - | - | - | - | `input, triggers_replace` | - | - | [`./resources.tf`](./resources.tf#L154) [154:1:3348 -> 154:40:3387] |

## Data Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `data` | `aws_caller_identity` | `current` | - | - | - | - | - | - | - | [`./resources.tf`](./resources.tf#L161) [161:1:3570 -> 161:37:3606] |
| `data` | `aws_iam_policy_document` | `per_region` | - | - | `var.regions` | - | - | - | - | [`./resources.tf`](./resources.tf#L163) [163:1:3611 -> 163:44:3654] |
| `data` | `http` | `health` | - | `var.bool_true ? 1 : 0` | - | - | `request_headers, url` | `postconditions=1` | - | [`./resources.tf`](./resources.tf#L173) [173:1:3830 -> 173:21:3850] |
| `data` | `local_file` | `readme` | - | - | - | `local_file.rendered` | `filename` | - | - | [`./resources.tf`](./resources.tf#L190) [190:1:4146 -> 190:27:4172] |

## Ephemeral Resources

| Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `random_password` | `session` | - | - | - | - | `length, special` | - | - | [`./resources.tf`](./resources.tf#L197) [197:1:4337 -> 197:38:4374] |
| `aws_secretsmanager_secret_version` | `db` | `aws.audit` | `var.bool_true ? 1 : 0` | - | `aws_iam_role.ignore_all` | `secret_id` | - | - | [`./resources.tf`](./resources.tf#L202) [202:1:4412 -> 202:51:4462] |
| `tls_private_key` | `per_region` | - | - | `var.regions` | - | `algorithm` | - | - | [`./resources.tf`](./resources.tf#L211) [211:1:4597 -> 211:41:4637] |

## Module Calls

| Name | Source | Version | Count | For Each | Depends On | Providers | Inputs | Comment | Position |
|------|--------|---------|-------|----------|------------|-----------|--------|---------|----------|
| `child` | `./modules/child` | - | - | - | - | `aws=aws.audit` | `count_value, enabled, name, objects, tags` | Local module with explicit provider mapping and several input shapes. | [`./modules.tf`](./modules.tf#L2) [2:1:72 -> 2:15:86] |
| `registry_count` | `terraform-aws-modules/s3-bucket/aws` | `4.1.2` | `var.bool_false ? 1 : 0` | - | - | - | `bucket` | Registry module with count. | [`./modules.tf`](./modules.tf#L18) [18:1:418 -> 18:24:441] |
| `registry_for_each` | `cloudposse/label/null` | `>= 0.25.0, < 1.0.0` | - | `toset(var.list_of_strings)` | `module.child, random_pet.names` | - | `name` | Registry module with for_each and depends_on. | [`./modules.tf`](./modules.tf#L27) [27:1:632 -> 27:27:658] |
| `git_source` | `git::https://github.com/example/terraform-modules.git//network?ref=v1.2.3` | - | - | - | - | - | - | Git source pinned to a ref, no version argument. | [`./modules.tf`](./modules.tf#L38) [38:1:894 -> 38:20:913] |
| `github_shorthand` | `github.com/example/terraform-module-foo` | - | - | - | - | - | - | GitHub shorthand and an HTTP archive. | [`./modules.tf`](./modules.tf#L43) [43:1:1046 -> 43:26:1071] |
| `http_archive` | `https://example.com/modules/vpc-module.zip` | - | - | - | - | - | - | - | [`./modules.tf`](./modules.tf#L47) [47:1:1130 -> 47:22:1151] |
| `dynamic_source` | - | - | - | - | - | - | - | OpenTofu early evaluation: non-literal source and version. | [`./modules.tf`](./modules.tf#L52) [52:1:1274 -> 52:24:1297] |
| `no_inputs` | `./modules/child` | - | - | - | - | - | - | Module call with no inputs at all. | [`./modules.tf`](./modules.tf#L58) [58:1:1423 -> 58:19:1441] |

## Moved

| From | To | Position |
|------|----|----------|
| `null_resource.old_name` | `null_resource.empty` | [`./blocks.tf`](./blocks.tf#L3) [3:1:80 -> 3:6:85] |
| `random_pet.single` | `random_pet.names` | [`./blocks.tf`](./blocks.tf#L8) [8:1:152 -> 8:6:157] |
| `module.old_child` | `module.child` | [`./blocks.tf`](./blocks.tf#L13) [13:1:219 -> 13:6:224] |
| `module.child.aws_s3_bucket.legacy` | `module.child.aws_s3_bucket.this` | [`./blocks.tf`](./blocks.tf#L18) [18:1:278 -> 18:6:283] |

## Imports

| To | ID | Provider | Position |
|----|----|----------|----------|
| `null_resource.empty` | `"static-id-123"` | - | [`./blocks.tf`](./blocks.tf#L25) [25:1:453 -> 25:7:459] |
| `aws_iam_role.ignore_all` | `"complete-audit"` | `aws.audit` | [`./blocks.tf`](./blocks.tf#L30) [30:1:515 -> 30:7:521] |
| `aws_s3_bucket.replica` | `"${local.literal_string}-replica"` | - | [`./blocks.tf`](./blocks.tf#L36) [36:1:617 -> 36:7:623] |
| - | `"complete-${each.key}"` | - | [`./blocks.tf`](./blocks.tf#L42) [42:1:749 -> 42:7:755] |

## Removed

| From | Destroy On Drop | Position |
|------|-----------------|----------|
| `null_resource.gone` | - | [`./blocks.tf`](./blocks.tf#L50) [50:1:959 -> 50:8:966] |
| `aws_s3_bucket.retired` | ✗ | [`./blocks.tf`](./blocks.tf#L54) [54:1:1000 -> 54:8:1007] |
| `module.deprecated_child` | ✓ | [`./blocks.tf`](./blocks.tf#L62) [62:1:1083 -> 62:8:1090] |

## Checks

| Name | Data Source | Assertions | Position |
|------|-------------|:----------:|----------|
| `health_endpoint` | `http.probe` | 2 | [`./blocks.tf`](./blocks.tf#L72) [72:1:1247 -> 72:24:1270] |
| `inputs_sane` | - | 1 | [`./blocks.tf`](./blocks.tf#L89) [89:1:1685 -> 89:20:1704] |
| `three_assertions` | - | 3 | [`./blocks.tf`](./blocks.tf#L96) [96:1:1823 -> 96:25:1847] |

## Diagnostics

| Severity | Summary | Detail | Subject | Context |
|----------|---------|--------|---------|---------|
| error | Invalid expression | A single static variable reference is required: only attribute access and indexing with constant keys. No calculations, function calls, template expressions, etc are allowed here. | [`./blocks.tf`](./blocks.tf#L44) [44:14:803 -> 44:50:839] | - |
| error | Ambiguous attribute key | If this expression is intended to be a reference, wrap it in parentheses. If it's instead intended as a literal name containing periods, wrap it in quotes to create a string literal. | [`./modules.tf`](./modules.tf#L7) [7:5:167 -> 7:16:178] | - |
| error | Invalid expression | A single static variable reference is required: only attribute access and indexing with constant keys. No calculations, function calls, template expressions, etc are allowed here. | [`./resources.tf`](./resources.tf#L165) [165:14:3695 -> 165:37:3718] | - |
