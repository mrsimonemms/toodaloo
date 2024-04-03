# Pre-commit hook

A supported [pre-commit](https://pre-commit.com) hook is provided to scan repos.

```yaml
repos:
  - repo: https://github.com/mrsimonemms/toodaloo
    rev: "" # Use the ref you want to point at
    hooks:
      - id: scan
```

This makes some changes to the default options passed in to the `scan` command.
1. The `--git-files` flag is always applied. This ensures that it scans the files
   in the `HEAD` of the repo.
2. It defaults to output the data in [markdown](./commands#output-format) and
   names the file `toodaloo.md`. This can be overridden if desired.

See the [scan command](./commands#scan) for more details on the available flag.
