# Getting started

## Install

Download a binary for Linux, macOS, or Windows from the
[latest release](https://github.com/RemoteRabbit/open-docs/releases/latest),
or install with Go (the version in `go.mod` or newer):

```bash
go install github.com/remoterabbit/open-docs@latest
```

To build from source, clone the repository and run `make build`; the binary is
written to `bin/open-docs`.

## Generate docs

Point open-docs at a module directory:

```bash
# print to stdout
open-docs ./examples/vpc

# write to a file (replace mode)
open-docs -o ./examples/vpc/README.md ./examples/vpc
```

## Keep a README up to date

Rather than overwriting a whole file, open-docs can update a region of a
hand-written README. Add the markers where the generated docs should go:

```markdown
# My Module

Hand-written intro.

<!-- BEGIN_OPEN_DOC -->
<!-- END_OPEN_DOC -->
```

Then set `mode = "inject"` in a `.open-docs.hcl` next to the module:

```hcl
main {
  mode = "inject"
}
```

Running `open-docs ./my-module` now replaces only the text between the markers
and leaves the rest of the file untouched. See
[Configuration](configuration.md) for every option.

## Add provider schema findings

With `-schema`, open-docs asks `tofu` or `terraform` for the provider schemas
and reports deprecated, unknown, and missing-required attributes:

```bash
open-docs -schema ./examples/vpc
```

!!! note

    `-schema` needs `tofu` or `terraform` on your `PATH` and an initialized
    module (`tofu init`).
