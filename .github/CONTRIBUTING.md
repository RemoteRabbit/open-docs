# Contributing to open-docs

Thanks for helping. Bug reports, feature ideas, and pull requests are all
welcome. For questions, use
[Discussions](https://github.com/RemoteRabbit/open-docs/discussions). Report
security problems privately as described in [SECURITY.md](./SECURITY.md).

By participating you agree to the [Code of Conduct](./CODE_OF_CONDUCT.md).

## Before you start

For anything larger than a small fix, open an issue or discussion first so we
can agree on the approach. open-docs is early and its output format is still
changing.

## Setup

You need:

- **Go**: the version in [`go.mod`](../go.mod).
- **[prek](https://prek.j178.dev/)**: runs the git hooks in
  [`prek.toml`](../prek.toml) (file hygiene, gofmt, `go vet`, `go test`).
  Install the hooks once with `prek install`.
- **[golangci-lint](https://golangci-lint.run/) v2** (optional locally; CI
  runs it).

Run `make` to list the available targets. The ones you will use most:

| Command           | What it does                                               |
|-------------------|------------------------------------------------------------|
| `make verify`     | gofmt check, `go vet`, tests with `-race`, smoke test.     |
| `make readmes`    | Regenerates `examples/*/README.md` with the current code.  |
| `make fmt`        | Formats Go sources with gofmt and golines.                 |
| `make build`      | Builds `bin/open-docs`.                                    |
| `make docs-serve` | Previews the docs site ([Zensical](https://zensical.org)). |
| `make zizmor`     | Audits workflows with [zizmor](https://docs.zizmor.sh).    |

## Making a change

1. Branch from `main`.
2. Make the change, with tests for new behavior or bug fixes.
3. If the change affects rendered output, run `make readmes` and commit the
   regenerated example READMEs. The [examples](../examples) are the visual
   fixtures for the output, and CI fails when they are stale.
4. If you add a configuration option, add its default to
   [`internal/config/defaults.hcl`](../internal/config/defaults.hcl) and
   document it in the README.
5. Run `make verify` and `golangci-lint run` before pushing.

## Pull requests

PRs are squash-merged, and the **PR title becomes the commit message**.
[release-please](https://github.com/googleapis/release-please) reads it to
choose the next version and write the changelog, so titles must follow
[Conventional Commits](https://www.conventionalcommits.org/):

```text
<type>(<optional scope>): <summary>
```

| Type                                                  | Effect                    |
|-------------------------------------------------------|---------------------------|
| `feat`                                                | Minor bump, changelog     |
| `fix`, `perf`, `deps`, `revert`, `docs`               | Patch bump, changelog     |
| `refactor`, `test`, `build`, `ci`, `chore`, `style`   | No release on their own   |

Add `!` after the type (for example `feat!: ...`) for a breaking change; before
1.0 that bumps the minor version. A CI check validates the title.

CI must pass before merge: prek hooks, golangci-lint, tests on Linux, macOS,
and Windows, the generated-files check, a GoReleaser snapshot build, and the
security scans.

## Project layout

- [`main.go`](../main.go): CLI flags, config loading, and output dispatch.
- [`internal/config`](../internal/config): `.open-docs.hcl` discovery,
  decoding, and built-in defaults.
- [`internal/render`](../internal/render): turns the inspected module into
  Markdown, one file per section.
- [`internal/output`](../internal/output): writes to stdout, replaces a file,
  or injects between markers.
- [`examples/`](../examples): sample modules and their generated READMEs.

HCL parsing belongs in
[open-inspector](https://github.com/RemoteRabbit/open-inspector); open-docs
owns presentation only. If you need data the inspector does not expose, open
an issue there.
