# open-docs

A documentation generator for Terraform/OpenTofu modules, built on
[open-inspector](https://github.com/RemoteRabbit/open-inspector).

The goal: a modern alternative to `terraform-docs` that matches its day-to-day
workflow, then goes further using data the inspector captures that
`terraform-docs` discards (source ranges, schema findings, reference graphs,
structured types).

**Documentation and example gallery:** <https://remoterabbit.github.io/open-docs/>

[![CI](https://img.shields.io/badge/CI-555555)](https://github.com/RemoteRabbit/open-docs/actions/workflows/ci.yml)
[![Latest release](https://img.shields.io/github/v/release/RemoteRabbit/open-docs)](https://github.com/RemoteRabbit/open-docs/releases/latest)
[![Go Version](https://img.shields.io/github/go-mod/go-version/RemoteRabbit/open-docs/main)](./go.mod)
[![Docs](https://img.shields.io/badge/docs-visit-0d619e?link=https://remoterabbit.github.io/open-docs/)](https://remoterabbit.github.io/open-docs/)
[![License: MPL 2.0](https://img.shields.io/github/license/RemoteRabbit/open-docs)](./LICENSE)

<!--
[![Coverage](https://codecov.io/gh/RemoteRabbit/open-docs/branch/main/graph/badge.svg)](https://codecov.io/gh/RemoteRabbit/open-docs)
[![Downloads](https://img.shields.io/github/downloads/RemoteRabbit/open-docs/total)](https://github.com/RemoteRabbit/open-docs/releases/latest)
[![Sponsor](https://img.shields.io/github/sponsors/RemoteRabbit)](https://github.com/sponsors/RemoteRabbit)
-->

> **Status: early / proof-of-concept.** Today open-docs emits a verbose,
> debug-oriented dump of everything the inspector parses. The curated,
> reader-facing output and the workflow features below are planned. See the
> [feature checklist](#features) for what works now vs. what is coming.

# Table of Contents

- [Why open-docs](#why-open-docs)
- [Features](#features)
  - [Core rendering](#core-rendering)
  - [Workflow / terraform-docs parity](#workflow-/-terraform-docs-parity)
  - [Differentiators](#differentiators)
  - [Power features (later)](#power-features-later)
- [Installation](#installation)
  - [Binary](#binary)
  - [Homebrew](#homebrew)
    - [All in one](#all-in-one)
    - [Tap then install](#tap-then-install)
    - [Brewfile](#brewfile)
  - [With Go](#with-go)
  - [Source](#source)
- [Usage](#usage)
  - [Configuration (`.open-docs.hcl`)](#configuration-(`.open-docs.hcl`))
- [Example](#example)
- [How it works](#how-it-works)
- [Contributing](#contributing)
- [Related](#related)
- [License](#license)

## Why open-docs

`terraform-docs` renders the fields it understands and throws the rest away.
open-inspector keeps the parts that make docs actionable, and open-docs turns
them into output:

- **Source permalinks.** Every block carries `filename` + start/end
  `line:column:byte`, so docs can deep-link to the exact source lines on
  GitHub/GitLab/Bitbucket.
- **Health, not just a field dump.** Schema findings (deprecated, unknown,
  missing-required attributes) and missing descriptions become an actionable
  checklist with severity.
- **Diagrams.** The inspector's module graph (tree/DOT/Mermaid) becomes inline
  dependency and resource diagrams; GitHub renders Mermaid natively.
- **Richer data.** Leading `# ...` doc-comments, structured variable types
  (`object({ ... })` with `optional`), decoded default values, and
  expression-level references (`var.x`, `module.y.z`, `aws_s3_bucket.b`).
- **Fuller block coverage** than terraform-docs: `moved`, `import`, `removed`,
  `check`, `ephemeral`, and OpenTofu `encryption`.

## Features

### Core rendering

- [x] Inspect a module directory and render Markdown for every block type
      (providers, inputs, outputs, locals, managed/data/ephemeral resources,
      module calls, moved/import/removed, checks, schema findings, diagnostics).
- [x] Source link per row (`` [`./file.tf`](./file.tf#L4-L7) ``);
      `main { verbose = true }` adds full `line:column:byte` ranges.
- [x] Doc-comments, structured types, decoded defaults, and expression
      references surfaced in the table view.
- [X] Curated, reader-facing default layout. *Today most columns are on by
      default and `main { verbose = true }` is the opt-in debug detail.*
- [X] Section and column toggles: each section has a config block with
      `render` and per-column switches.
- [ ] Sort order (by name or source position).
- [~] Deterministic, diff-stable output across runs and platforms.

### Workflow / terraform-docs parity

- [x] Write to stdout or a file (`-o`).
- [x] **README injection** between `<!-- BEGIN_OPEN_DOC -->` /
      `<!-- END_OPEN_DOC -->` markers instead of overwriting the file
      (`main { mode = "inject" }`).
- [~] **Config file** (`.open-docs.hcl`): `main { docsfile, mode, header,
      verbose, version }` and per-section toggles work today; formatter, sort,
      and header/footer source are planned.
- [x] **prek hooks** for file hygiene, gofmt, `go vet`, and `go test`.
- [X] **Homebrew** installation; right now only via a cask, part of core in the future.
- [ ] **GitHub Action** wrapper.
- [ ] **GitLab Component** wrapper.
- [ ] **CI gate**: `--check` (fail if on-disk docs are stale) and
      `--fail-on warning|error`.

### Differentiators

- [x] Schema-aware inspection via `-schema` (shells out to `tofu`/`terraform`).
- [~] **Health section**: render schema findings + missing descriptions as a
      severity checklist. *Findings render as a table today; the curated
      health view is planned.*
- [~] **Source permalinks**: rows link to the source file and line range,
      relative to the output file. *Host permalinks (GitHub/GitLab/Bitbucket
      at a git ref) with remote detection and a config override are planned.*
- [ ] **Inline Mermaid graph**: module dependency and resource-relationship
      diagrams.
- [ ] **Cross-references** ("used by"): drive from expression references
      (e.g. which outputs/resources consume `var.x`).
- [~] **Lint/quality**: unused variables, unreferenced outputs, undocumented
      inputs/outputs, doc-completeness score.

### Power features (later)

- [ ] **Dynamic config** (HCL eval context): build section lists with
      `for`/conditionals that react to the inspected module.
- [ ] **Templates / custom content**: arrange generated sections and interleave
      prose or included files.
- [ ] **More formats**: Markdown document, `json`, `mdx`, `asciidoc`.
- [ ] **Recursive / multi-module docs** via the module graph.
- [ ] **Examples extraction**: include `examples/*/main.tf` snippets.
- [ ] Optional static HTML/MDX site output.

## Installation

### Binary

Download a binary for Linux, macOS, or Windows from the
[latest release](https://github.com/RemoteRabbit/open-docs/releases/latest)

### Homebrew

A cask can be found [here](https://github.com/remoterabbit/homebrew-open-docs).

#### All in one

```shell
brew install remoterabbit/homebrew-open-docs/open-docs
```

#### Tap then install

```shell
brew tap remoterabbit/homebrew-open-docs
brew install remoterabbit/open-docs
```

#### Brewfile

```Brewfile
tap "remoterabbit/open-docs"
cask "open-docs"
```

> [!NOTE]
> Eventually this will be a formula in the homebrew-core but it is not ready yet.

### With Go

Install with Go (the version in [`go.mod`](./go.mod) or newer):

```bash
go install github.com/remoterabbit/open-docs@latest
```

### Source

To build from source, clone the repository and run `make build`; the binary is
written to `bin/open-docs`.

## Usage

```bash
# print to stdout
open-docs ./examples/vpc

# write to a file (replace mode)
open-docs -o ./examples/vpc/README.md ./examples/vpc

# update the marked region of a file (inject mode)
open-docs -m inject -o ./examples/vpc/README.md ./examples/vpc

# skip sections for one run (block names from .open-docs.hcl)
open-docs -hide locals,resources ./examples/vpc

# in CI: fail if the README is out of date, without writing it
open-docs -check -o ./examples/vpc/README.md ./examples/vpc

# enrich with provider schema (deprecated / missing-required findings)
# requires `tofu` or `terraform` available and the module initialized
open-docs -schema ./examples/vpc
```

### Configuration (`.open-docs.hcl`)

> [!note]
> You can find the defaults [here](./internal/config/defaults.hcl).

Drop a `.open-docs.hcl` in the module directory or the repository root (or pass
`-config <file>`) to control output. open-docs uses the nearest one, searching
from the module directory up to the repository root.

```hcl
main {
  docsfile = "README.md" # relative to the module directory
  mode     = "inject"    # "stdout" (default) | "replace" | "inject"
  header   = false       # omit the module title line (default: true)
  verbose  = true        # add line:column:byte ranges to positions (default: false)
  version  = false       # omit the "generated with open-docs" line (default: true)
}

# Every section has a block with `render` and one switch per column.
variables {
  comment  = false
  position = false
}

outputs {
  render = false
}
```

In `inject` mode, open-docs replaces the text between the markers and leaves the
rest of the file untouched (creating the region if it is missing):

```markdown
# My Module

Hand-written intro.

<!-- BEGIN_OPEN_DOC -->
<!-- END_OPEN_DOC -->
```

The `-m <mode>` flag sets the mode from the command line, and `-o <file>` sets
the output file. Both take precedence over the config file. `-o` alone is a
shortcut for `replace` mode to that path; combine it with `-m inject` to update
the markers in that file instead.

`-show` and `-hide` take comma-separated section block names, such as
`-hide locals,resources`, and override each section's `render` switch for one
run; `-hide` wins when both name a section. `-check` renders the docs and exits
non-zero if writing them would change the docs file, without writing it.

## Example

Each directory under [examples/](./examples) is a module with its generated
`README.md`, used to check how output renders:

- [vpc](./examples/vpc): a realistic multi-file AWS VPC module with a local
  child module.
- [complete](./examples/complete): every block type open-inspector understands
  (including ephemeral, check, OpenTofu encryption and provider `for_each`,
  and a `.tf.json` file), plus table edge cases such as pipes, heredocs,
  Unicode, and long values.
- [invalid](./examples/invalid): deliberately broken configuration for the
  diagnostics table.

## How it works

open-docs calls `inspector.Inspect(dir)` to parse the module into a
`*model.Module`, then walks that struct in
[`internal/render`](./internal/render) to render Markdown tables. All HCL
parsing lives in open-inspector; open-docs owns presentation only.

## Contributing

See [CONTRIBUTING.md](./.github/CONTRIBUTING.md) for setup, the project layout,
and how pull requests and releases work. Please report security issues
privately as described in [SECURITY.md](./.github/SECURITY.md).

## Related

- [open-inspector](https://github.com/RemoteRabbit/open-inspector): the HCL
  inspection engine open-docs is built on.
- [terraform-docs](https://github.com/terraform-docs/terraform-docs): the
  incumbent open-docs aims to match, then surpass.

## License

[Mozilla Public License 2.0](./LICENSE).
