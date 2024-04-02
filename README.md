# Toodaloo

Say goodbye to your todos

<!-- toc -->

* [Purpose](#purpose)
* [Install](#install)
  * [Docker](#docker)
  * [Go](#go)
* [Commands](#commands)
  * [Scan](#scan)
* [Pre-commit hook](#pre-commit-hook)
* [Contributing](#contributing)
  * [Open in a container](#open-in-a-container)

<!-- Regenerate with "pre-commit run -a markdown-toc" -->

<!-- tocstop -->

## Purpose

When writing code, it's really _REALLY_ easy to put a `TODO` in your code as a
reminder to look at something later on. In small projects, you'll see those reminders
every time you open up the relevant files. In big projects with multiple contributors,
these can get unwieldly.

Without a proper strategy to work through them, your todos will only ever expand.
Toodaloo extracts these into a central file which can be tracked.

## Install

### Docker

```shell
docker run \
  -it \
  --rm \
  -v /path/to/dir:/data \
  ghcr.io/mrsimonemms/toodaloo
```

### Go

```shell
go install github.com/mrsimonemms/toodaloo@latest
```

## Commands

### Scan

```shell
Scan a project

Usage:
  toodaloo scan [flags]

Flags:
      --git-files              get files from the git tree
      --glob string            glob pattern - ignored if files provided as arguments (default "**/*")
  -h, --help                   help for scan
      --ignore-paths strings   ignore scanning these files (default [.git/**/*])
  -o, --output string          output type (default "yaml")
  -s, --save-path string       save report to path - use "-" to output to stdout (default ".toodaloo.yaml")
  -t, --tags strings           todo tags (default [fixme,todo,@todo])

Global Flags:
  -d, --directory string   working directory (default "/workspaces/toodaloo2")
  -l, --log-level string   log level: trace, debug, info, warning, error, fatal, panic (default "info")
```

## Pre-commit hook

A supported [pre-commit](https://pre-commit.com) hook is provided to scan repos.

```yaml
repos:
  - repo: https://github.com/mrsimonemms/toodaloo
    rev: "" # Use the ref you want to point at
    hooks:
      - id: scan
```

This will generate a [Markdown](https://www.markdownguide.org) formatted file
at `toodaloo.md`. It also only scans files in the git tree.

## Contributing

### Open in a container

* [Open in a container](https://code.visualstudio.com/docs/devcontainers/containers)
