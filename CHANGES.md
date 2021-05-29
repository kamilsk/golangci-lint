The fork has solution for the [issue][] ([diff][]).

You can enable it in your configuration:

```yaml
linters:
  ...

  enable:
    - looppointer
```

## How to install

### Homebrew

```bash
$ brew install kamilsk/tap/golangci-lint
```

### Binary

```bash
$ curl -sSfL https://bit.ly/install-golangci | sh -s -- -b $(go env GOPATH)/bin
```

Happy hacking!

<!-- references -->

[diff]:     https://github.com/golangci/golangci-lint/compare/master...kamilsk:looppointer
[issue]:    https://github.com/golangci/golangci-lint/issues/1404
[pr]:       https://github.com/golangci/golangci-lint/pull/1924
