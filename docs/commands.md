# Commands

## Scan

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

### Listing files to scan

There are three ways to list the files to scan. They're subtly different, so it's
important to understand the differences.

#### Glob

This is the default.

The value of the `--glob` argument decides the files that are scanned. This is
useful to combine with `--ignore-paths`.

#### Arguments

The `toodaloo scan` command accepts an infinite number of arguments. These files
will then be scanned.

#### Git files

Passing the argument `--git-files` will find all the files in the `HEAD` of the
Git repo.

It is the functional equivalent of the command `git ls-tree -r HEAD --name-only`.

### Tags

There are many ways of declaring a TODO in your source code. By default, the tags
`fixme`, `todo` and `@todo` are used.

### Output format

There are various different output formats available. This tracks the TODOs in a
way that can be used in your project.

Ultimately, this makes it very obvious when new TODOs are added in your pull
requests.

#### YAML

This is the default output. This is useful if you want to programmitically use
the TODO output.

<!-- markdownlint-disable-next-line MD024 -->
##### Example

```yaml
- author: sje
  file: pkg/output/markdown.go
  lineNumber: 30
  message: implement report
- author: sje
  file: pkg/scanner/scan.go
  lineNumber: 189
  message: get the author from the Git history
- author: sje
  file: pkg/scanner/scan.go
  lineNumber: 198
  message: add the URL to the file/line number
```

#### Markdown

This generates a table in markdown. This will then be displayed.

<!-- markdownlint-disable-next-line MD024 -->
##### Example

```markdown
| File | Line Number | Author | Message |
| --- | --- | --- | --- |
| pkg/output/markdown.go | 30 | sje | implement report |
| pkg/scanner/scan.go | 189 | sje | get the author from the Git history |
| pkg/scanner/scan.go | 198 | sje | add the URL to the file/line number |
```
