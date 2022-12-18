## Motivation

As a Go developer, I actively use [`golangci-lint`][golangci-lint] as a static linter,
but it doesn't support the [`looppointer`][looppointer] linter. I need to extend it
to include the linter, which you can enable through the configuration:

```yaml
linters:

  enable:
    - looppointer
```

or passing the argument

```bash
$ golangci-lint run --enable looppointer ./...
```

[golangci-lint]:  https://golangci-lint.run
[looppointer]:    https://github.com/kyoh86/looppointer

## What's changed

- The `golangci-lint` supports the `looppointer` linter.
- The tool is available by `brew install kamilsk/tap/golangci-lint`.
- The installation is also possible by
  ```bash
  $ curl -sSfL https://install.octolab.org/golangci-lint | sh -s -- -b /usr/local/bin
  ```

## Related issues

- https://github.com/golangci/golangci-lint/issues/1404
- https://github.com/golangci/golangci-lint/pull/1924
- https://github.com/golangci/golangci-lint/compare/master...kamilsk:extended

<p align="right">made with ❤️ for everyone by <a href="https://www.octolab.org/">OctoLab</a></p>
