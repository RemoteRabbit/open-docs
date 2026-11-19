# Configuration

open-docs reads an optional `.open-docs.hcl`. Put it in the module directory or
the repository root, or pass `-config <file>`. open-docs uses the nearest one,
searching from the module directory up to the repository root.

Your file is decoded on top of the [built-in defaults](#defaults), so set only
what you want to change.

## `main`

```hcl
main {
  docsfile = "README.md" # relative to the module directory
  mode     = "inject"    # "stdout" (default) | "replace" | "inject"
  header   = false       # omit the module title line (default: true)
  verbose  = true        # add line:column:byte ranges to positions (default: false)
  version  = false       # omit the "generated with open-docs" line (default: true)
}
```

| Mode      | Behavior                                                                  |
|-----------|---------------------------------------------------------------------------|
| `stdout`  | Print the docs. The default.                                              |
| `replace` | Overwrite `docsfile` with the docs.                                       |
| `inject`  | Replace the text between `<!-- BEGIN_OPEN_DOC -->` and `<!-- END_OPEN_DOC -->` in `docsfile`, adding the region if it is missing. |

The `-o <file>` flag is a shortcut for `replace` mode to that path and takes
precedence over the config file.

## Sections and columns

Every section has a block with a `render` switch and one switch per column:

```hcl
# Hide two columns of the inputs table.
variables {
  comment  = false
  position = false
}

# Drop the outputs section entirely.
outputs {
  render = false
}
```

## Defaults

These are the built-in defaults, included straight from
`internal/config/defaults.hcl`, so they always match the released code.

```hcl
--8<-- "internal/config/defaults.hcl"
```
