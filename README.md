# open-doc

A documentation generator for Terraform/OpenTofu modules, built on
[open-inspector](https://github.com/RemoteRabbit/open-inspector).

The goal: a modern alternative to `terraform-docs` that matches its day-to-day
workflow, then goes further using data the inspector captures that
`terraform-docs` discards (source ranges, schema findings, reference graphs,
structured types).

> **Status: early / proof-of-concept.** Today open-doc emits a verbose,
> debug-oriented dump of everything the inspector parses. The curated,
> reader-facing output and the workflow features below are planned. See the
> [feature checklist](#features) for what works now vs. what is coming, and
> [`ROADMAP.md`](./ROADMAP.md) for the phased plan.

## Why open-doc

`terraform-docs` renders the fields it understands and throws the rest away.
open-inspector keeps the parts that make docs actionable, and open-doc turns
them into output:

- **Source permalinks.** Every block carries `filename` + start/end
  `line:column:byte`, so docs can deep-link to the exact source lines on
  GitHub/GitLab/Bitbucket.
- **Health, not just a field dump.** Schema findings (deprecated, unknown,
  missing-required attributes) and missing descriptions become an actionable
  checklist with severities.
- **Diagrams.** The inspector's module graph (tree/DOT/Mermaid) becomes inline
  dependency and resource diagrams; GitHub renders Mermaid natively.
- **Richer data.** Leading `# ...` doc-comments, structured variable types
  (`object({ ... })` with `optional`), decoded default values, and
  expression-level references (`var.x`, `module.y.z`, `aws_s3_bucket.b`).
- **Fuller block coverage** than terraform-docs: `moved`, `import`, `removed`,
  `check`, `ephemeral`, and OpenTofu `encryption`.

## Features

Legend: `[x]` available now &nbsp;·&nbsp; `[~]` partial / debug-only
&nbsp;·&nbsp; `[ ]` planned.

### Core rendering

- [x] Inspect a module directory and render Markdown for every block type
      (providers, inputs, outputs, locals, managed/data/ephemeral resources,
      module calls, moved/import/removed, checks, schema findings, diagnostics).
- [x] Full source position per row (`` `file.tf` [L:C:B -> L:C:B] ``).
- [x] Doc-comments, structured types, decoded defaults, and expression
      references surfaced in the table view.
- [~] Curated, reader-facing default layout. *Today the output is the verbose
      all-columns debug view; a clean default with `--verbose` as the opt-in
      debug mode is planned.*
- [ ] Section toggles (show/hide) and sort order (by name or source position).
- [ ] Deterministic, diff-stable output across runs and platforms.

### Workflow / terraform-docs parity (table stakes)

- [x] Write to stdout or a file (`-o`).
- [x] **README injection** between `<!-- BEGIN_OPEN_DOC -->` /
      `<!-- END_OPEN_DOC -->` markers instead of overwriting the file
      (`output { mode = "inject" }`).
- [~] **Config file** (`.open-doc.hcl`): `output { file, mode }` works today;
      formatter, sections, sort, and header/footer source are planned.
- [x] **prek hooks** for file hygiene, gofmt, `go vet`, and `go test`.
- [ ] **GitHub Action** wrapper.
- [ ] **CI gate**: `--check` (fail if on-disk docs are stale) and
      `--fail-on warning|error`.

### Differentiators (the wedge)

- [x] Schema-aware inspection via `-schema` (shells out to `tofu`/`terraform`).
- [~] **Health section**: render schema findings + missing descriptions as a
      severity checklist. *Findings render as a table today; the curated
      health view is planned.*
- [ ] **Source permalinks**: ranges -> host links at the right ref/line, with
      git remote/ref detection and config override.
- [ ] **Inline Mermaid graph**: module dependency and resource-relationship
      diagrams.
- [ ] **Cross-references** ("used by"): drive from expression references
      (e.g. which outputs/resources consume `var.x`).
- [ ] **Lint/quality**: unused variables, unreferenced outputs, undocumented
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

## Usage

```bash
go build -o open-doc .

# print to stdout
./open-doc ./examples/vpc

# write to a file (replace mode)
./open-doc -o ./examples/vpc/README.md ./examples/vpc

# enrich with provider schema (deprecated / missing-required findings)
# requires `tofu` or `terraform` available and the module initialized
./open-doc -schema ./examples/vpc
```

### Configuration (`.open-doc.hcl`)

Drop a `.open-doc.hcl` in the module directory (or pass `-config <file>`) to
control output. open-doc looks for it in the module directory, then the current
working directory.

```hcl
output {
  file = "README.md"  # relative to the module directory
  mode = "inject"     # "stdout" (default) | "replace" | "inject"
}
```

In `inject` mode, open-doc replaces the text between the markers and leaves the
rest of the file untouched (creating the region if it is missing):

```markdown
# My Module

Hand-written intro.

<!-- BEGIN_OPEN_DOC -->
<!-- END_OPEN_DOC -->
```

The `-o <file>` flag is a shortcut for `replace` mode to that path and takes
precedence over the config file.

## Example

See [examples/vpc](./examples/vpc) for a sample module and its generated
[README.md](./examples/vpc/README.md).

## Development

The project uses [devenv](https://devenv.sh/) for a pinned Go, OpenTofu, and
prek toolchain. Enter it directly or allow direnv once:

```bash
devenv shell
# or
direnv allow
```

The shell installs the prek Git shim. Common commands are:

```bash
build                 # compile ./open-doc
format                # format all Go files
verify                # format check, build, vet, test, and smoke test
prek run --all-files  # run every repository hook
devenv test           # run the complete devenv verification
```

## How it works

open-doc calls `inspector.Inspect(dir)` to parse the module into a
`*model.Module`, then walks that struct in [markdown.go](./markdown.go) to
render Markdown tables. All HCL parsing lives in open-inspector; open-doc owns
presentation only.

## Project layout

- [`main.go`](./main.go): CLI flag parsing, calls `inspector.Inspect`, writes output.
- [`markdown.go`](./markdown.go): renders a `*model.Module` to Markdown.
- [`examples/`](./examples): sample modules and their generated `README.md`.
- [`ROADMAP.md`](./ROADMAP.md): phased delivery plan and design decisions.
- [`INSPECTOR_WISHLIST.md`](./INSPECTOR_WISHLIST.md): capabilities open-doc
  wants from open-inspector (and what has shipped).

## Related

- [open-inspector](https://github.com/RemoteRabbit/open-inspector): the HCL
  inspection engine open-doc is built on.
- [terraform-docs](https://github.com/terraform-docs/terraform-docs): the
  incumbent open-doc aims to match, then surpass.
</content>

</invoke>
