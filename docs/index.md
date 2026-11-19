# open-docs

A documentation generator for Terraform/OpenTofu modules, built on
[open-inspector](https://github.com/RemoteRabbit/open-inspector).

The goal: a modern alternative to `terraform-docs` that matches its day-to-day
workflow, then goes further using data the inspector captures that
`terraform-docs` discards (source ranges, schema findings, reference graphs,
structured types).

!!! warning "Early / proof-of-concept"

    Today open-docs emits a verbose, debug-oriented dump of everything the
    inspector parses. The curated, reader-facing output and several workflow
    features are planned; the
    [feature checklist](https://github.com/RemoteRabbit/open-docs#features)
    tracks what works now and what is coming.

## Why open-docs

`terraform-docs` renders the fields it understands and throws the rest away.
open-inspector keeps the parts that make docs actionable, and open-docs turns
them into output:

- **Source permalinks.** Every block carries `filename` + start/end
  `line:column:byte`, so docs can deep-link to the exact source lines.
- **Health, not just a field dump.** Schema findings (deprecated, unknown,
  missing-required attributes) and missing descriptions become an actionable
  checklist with severity.
- **Diagrams.** The inspector's module graph becomes inline dependency and
  resource diagrams.
- **Richer data.** Leading `# ...` doc-comments, structured variable types
  (`object({ ... })` with `optional`), decoded default values, and
  expression-level references (`var.x`, `module.y.z`, `aws_s3_bucket.b`).
- **Fuller block coverage** than terraform-docs: `moved`, `import`, `removed`,
  `check`, `ephemeral`, and OpenTofu `encryption`.

## See it

The [example gallery](examples/index.md) shows open-docs output for real
modules, rendered exactly as the tool generates it.

## How it works

open-docs calls `inspector.Inspect(dir)` to parse the module into a
`*model.Module`, then walks that struct in `internal/render` to render Markdown
tables. All HCL parsing lives in open-inspector; open-docs owns presentation
only.
