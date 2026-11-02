# Module `invalid`

<!-- VERSION-START -->
This documentation was generated with [open-docs v1](https://github.com/RemoteRabbit/open-docs)
<!-- VERSION-END -->


## Required Providers

| Name | Source | Version | Position |
|------|--------|---------|----------|
| `null` | `hashicorp/null` | `~> 3.2` | [`./main.tf`](./main.tf#L7-L10) [7:5:207 -> 10:6:279] |

## Inputs

| Name | Type | Default | Description | Comment | Sensitive | Nullable | Ephemeral | Validation Count | Position |
|------|------|---------|-------------|---------|-----------|----------|-----------|------------------|----------|
| `description_source` | `string` | `"from a variable"` |  |  | - | - | - | 0 | [`./main.tf`](./main.tf#L14) [14:1:287 -> 14:30:316] |
| `flag_source` | `bool` | `false` |  |  | - | - | - | 0 | [`./main.tf`](./main.tf#L19) [19:1:371 -> 19:23:393] |
| `malformed_validation` | `string` | - |  | Validation block missing both condition and error_message. | - | - | - | 0 | [`./main.tf`](./main.tf#L25) [25:1:495 -> 25:32:526] |
| `bad_type` | `strng` | `"typo"` |  | Unknown type keyword. | - | - | - | 0 | [`./main.tf`](./main.tf#L33) [33:1:592 -> 33:20:611] |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|-----------|-----------|------------|----------|
| non_literal_description | `"hello"` | - |  | description must be a literal string. | - | - | - | [`./main.tf`](./main.tf#L39) [39:1:694 -> 39:33:726] |
| non_literal_sensitive | `"hello"` | - |  | sensitive must be a literal bool. | - | - | - | [`./main.tf`](./main.tf#L45) [45:1:841 -> 45:31:871] |

## Managed Resources

| Mode | Type | Name | Provider | Count | For Each | Depends On | Attributes | Lifecycle | Comment | Position |
|------|------|------|----------|-------|----------|------------|------------|-----------|---------|----------|
| `managed` | `null_resource` | `malformed_precondition` | - | - | - | - | - | - | Lifecycle precondition missing its required attributes. | [`./main.tf`](./main.tf#L51) [51:1:987 -> 51:50:1036] |
| `managed` | `null_resource` | `valid` | - | - | - | - | `triggers` | - | A perfectly valid resource, so the healthy parts still render. | [`./main.tf`](./main.tf#L59) [59:1:1150 -> 59:33:1182] |
| `managed` | `null_resource` | `broken` | - | - | - | - | `triggers` | - | Missing closing brace: a hard HCL syntax error with a source range. | [`./syntax.tf`](./syntax.tf#L2) [2:1:70 -> 2:34:103] |

## Diagnostics

| Severity | Summary | Detail | Subject | Context |
|----------|---------|--------|---------|---------|
| error | Unclosed configuration block | There is no closing brace for this block before the end of the file. This may be caused by incorrect brace nesting elsewhere in this file. | [`./syntax.tf`](./syntax.tf#L2) [2:35:104 -> 2:36:105] | - |
| error | Missing required argument | The argument "condition" is required, but no definition was found. | [`./main.tf`](./main.tf#L28) [28:14:559 -> 28:14:559] | - |
| error | Missing required argument | The argument "error_message" is required, but no definition was found. | [`./main.tf`](./main.tf#L28) [28:14:559 -> 28:14:559] | - |
| error | Invalid type specification | The keyword "strng" is not a valid type specification. | [`./main.tf`](./main.tf#L34) [34:13:626 -> 34:18:631] | - |
| error | Variables not allowed | Variables may not be used here. | [`./main.tf`](./main.tf#L41) [41:25:777 -> 41:28:780] | - |
| error | Variables not allowed | Variables may not be used here. | [`./main.tf`](./main.tf#L47) [47:15:910 -> 47:18:913] | - |
| error | Missing required argument | The argument "condition" is required, but no definition was found. | [`./main.tf`](./main.tf#L53) [53:18:1070 -> 53:18:1070] | - |
| error | Missing required argument | The argument "error_message" is required, but no definition was found. | [`./main.tf`](./main.tf#L53) [53:18:1070 -> 53:18:1070] | - |
