# Contributing, reporting, and obtaining open-docs

open-docs is open source under the Mozilla Public License 2.0. This page
explains how to obtain it, report problems or suggest improvements, and
contribute code.

## Obtain

- **Release binaries** for Linux, macOS, and Windows:
  [latest release](https://github.com/RemoteRabbit/open-docs/releases/latest).
- **Go install**: `go install github.com/remoterabbit/open-docs@latest`
- **Build from source**: clone the repository and run `make build`.

See [Getting started](getting-started.md) for the full walkthrough.

## Provide feedback

Help is welcome in three forms:

- **Bug reports** — use the
  [bug report](https://github.com/RemoteRabbit/open-docs/issues/new?template=bug_report.yml)
  template. Include the open-docs version, the command you ran, and the output
  or error so it can be reproduced.
- **Enhancements / feature ideas** — use the
  [feature request](https://github.com/RemoteRabbit/open-docs/issues/new?template=feature_request.yml)
  template.
- **Questions** — start a
  [discussion](https://github.com/RemoteRabbit/open-docs/discussions) rather
  than an issue.

Report security problems privately as described in
[SECURITY.md](https://github.com/RemoteRabbit/open-docs/blob/main/.github/SECURITY.md).

## Contribute

See [.github/CONTRIBUTING.md](https://github.com/RemoteRabbit/open-docs/blob/main/.github/CONTRIBUTING.md)
for setup, the project layout, and how pull requests and releases work. In
short: branch from `main`, add tests for new or changed behavior, run
`make verify` and `golangci-lint run`, and open a pull request with a
[Conventional Commits](https://www.conventionalcommits.org/) title.
