# toodaloo

Say goodbye to your todos

<!-- toc -->

* [Purpose](#purpose)
* [Install toodaloo](#install-toodaloo)
  * [Install from GitHub releases](#install-from-github-releases)
  * [Use in Docker](#use-in-docker)
* [Pre-commit hooks](#pre-commit-hooks)
* [Contributing](#contributing)
  * [Open in Gitpod](#open-in-gitpod)
  * [Open in a container](#open-in-a-container)
* [Stargazers over time](#stargazers-over-time)

<!-- Regenerate with "pre-commit run -a markdown-toc" -->

<!-- tocstop -->

## Purpose

Toodaloo is a simple application to detect todos in your codebase. It generates
a report in the root of your project so you can work to reduce them over time.
It's easy to write a todo in your codebase, but they're not easy to keep track
of unless you happen to see them in the file.

This helps with that.

## Install toodaloo

### Install from GitHub releases

### Use in Docker

```shell
docker pull ghcr.io/mrsimonemms/toodaloo
```

## Pre-commit hooks

```yaml
repos:
  - repo: https://github.com/mrsimonemms/
    rev: "" # Run pre-commit autoupdate for latest version
    hooks:
      - id: toodaloo
```

## Contributing

### Open in Gitpod

* [Open in Gitpod](https://gitpod.io/from-referrer/)

### Open in a container

* [Open in a container](https://code.visualstudio.com/docs/devcontainers/containers)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/mrsimonemms/toodaloo.svg)](https://starchart.cc/mrsimonemms/toodaloo)
