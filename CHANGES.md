## Motivation

`golangci-lint` doesn't support the [`looppointer`][looppointer] linter.
You can enable it in your configuration:

```yaml
linters:
  ...

  enable:
    - looppointer
```

## Related issues

- https://github.com/golangci/golangci-lint/issues/1404
- https://github.com/golangci/golangci-lint/pull/1924
- https://github.com/golangci/golangci-lint/compare/master...kamilsk:looppointer

## What's changed

- The `golangci-lint` supports the `looppointer` linter.
- The package is available by `brew install kamilsk/tap/golangci-lint`.
- The installation is also possible by
  `curl -sSfL https://raw.githubusercontent.com/kamilsk/golangci-lint/extended/bin/install | sh -s -- -b /usr/local/bin`

<!-- references -->

[looppointer]: https://github.com/kyoh86/looppointer
