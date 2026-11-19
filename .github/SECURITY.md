# Security Policy

## Supported versions

open-docs is pre-1.0. Security fixes land on `main` and ship in the next
release; only the latest release is supported.

## Reporting a vulnerability

Please do not open a public issue, discussion, or pull request for a security
problem.

Report it privately through GitHub instead:
[**Report a vulnerability**](https://github.com/RemoteRabbit/open-docs/security/advisories/new).
If you cannot use GitHub, email trjahnke@protonmail.com.

Include what you can of:

- the open-docs version or commit,
- the command and configuration you ran,
- a minimal module or input that triggers the problem,
- the impact you expect (for example, file writes outside the target path,
  command execution, or leaking secrets into generated docs).

You should get an acknowledgement within a week. Once a fix is ready, it is
released and the advisory is published with credit to you, unless you prefer
otherwise.

## Scope

open-docs reads Terraform/OpenTofu source and writes Markdown. Areas most
likely to matter:

- **File writes**: `-o`, `docsfile`, and inject mode writing outside the
  intended file.
- **`-schema`**: this flag runs `tofu` or `terraform` in the module directory,
  so only use it on modules you trust.
- **Generated output**: content from the inspected module rendered in a way
  that injects unexpected HTML or Markdown.
