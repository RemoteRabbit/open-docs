# Module `invalid`

<!-- VERSION-START -->
This documentation was generated with [open-docs dev](https://github.com/RemoteRabbit/open-docs)
<!-- VERSION-END -->


## Required Providers

| Name | Source | Version | Position |
|------|--------|---------|----------|
| `null` | `hashicorp/null` | `~> 3.2` | [`./main.tf`](./main.tf#L7-L10) |

## Inputs

| Name | Type | Default | Description | Comment | Sensitive | Nullable | Ephemeral | Validation Count | Position |
|------|------|---------|-------------|---------|-----------|----------|-----------|------------------|----------|
| `description_source` | `string` | `"from a variable"` | - | - | - | - | - | 0 | [`./main.tf`](./main.tf#L14) |
| `flag_source` | `bool` | `false` | - | - | - | - | - | 0 | [`./main.tf`](./main.tf#L19) |
| `malformed_validation` | `string` | - | - | Validation block missing both condition and error_message. | - | - | - | 0 | [`./main.tf`](./main.tf#L25) |
| `bad_type` | `strng` | `"typo"` | - | Unknown type keyword. | - | - | - | 0 | [`./main.tf`](./main.tf#L33) |

## Outputs

| Name | Value | References | Description | Comment | Sensitive | Ephemeral | Depends On | Position |
|------|-------|------------|-------------|---------|-----------|-----------|------------|----------|
| non_literal_description | `"hello"` | - | - | description must be a literal string. | - | - | - | [`./main.tf`](./main.tf#L39) |
| non_literal_sensitive | `"hello"` | - | - | sensitive must be a literal bool. | - | - | - | [`./main.tf`](./main.tf#L45) |

## Managed Resources

| Name | Mode | Type | Count | For Each | Depends On | Nested Body | LifeCycle | Comment | Position |
|------|------|------|-------|----------|------------|-------------|-----------|---------|----------|
| malformed_precondition | managed | `null_resource` | - | - |  | - | - | Lifecycle precondition missing its required attributes. | [`./main.tf`](./main.tf#L51) |
| valid | managed | `null_resource` | - | - |  | `triggers` | - | A perfectly valid resource, so the healthy parts still render. | [`./main.tf`](./main.tf#L59) |
| broken | managed | `null_resource` | - | - |  | `triggers` | - | Missing closing brace: a hard HCL syntax error with a source range. | [`./syntax.tf`](./syntax.tf#L2) |

## Diagnostics

| Summary | Detail | Subject | Context |
|---------|--------|---------|---------|
| Unclosed configuration block | There is no closing brace for this block before the end of the file. This may be caused by incorrect brace nesting elsewhere in this file. | [`./syntax.tf`](./syntax.tf#L2) | - |
| Missing required argument | The argument "condition" is required, but no definition was found. | [`./main.tf`](./main.tf#L28) | - |
| Missing required argument | The argument "error_message" is required, but no definition was found. | [`./main.tf`](./main.tf#L28) | - |
| Invalid type specification | The keyword "strng" is not a valid type specification. | [`./main.tf`](./main.tf#L34) | - |
| Variables not allowed | Variables may not be used here. | [`./main.tf`](./main.tf#L41) | - |
| Variables not allowed | Variables may not be used here. | [`./main.tf`](./main.tf#L47) | - |
| Missing required argument | The argument "condition" is required, but no definition was found. | [`./main.tf`](./main.tf#L53) | - |
| Missing required argument | The argument "error_message" is required, but no definition was found. | [`./main.tf`](./main.tf#L53) | - |
