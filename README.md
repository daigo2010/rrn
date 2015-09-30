# rrn

Recursive Rename Command.

## Description

This command can change the file name to in the sub directory simply.

## Usage

```bash
rrn -n '/.txt$' '.xml' . # dry-run
rrn '/.txt$' '.xml' .
```

## Install

To install, use `go get`:

```bash
$ go get -d github.com/daigo2010/rrn
```

## Contribution

1. Fork ([https://github.com/daigo2010/rrn/fork](https://github.com/daigo2010/rrn/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[daigo2010](https://github.com/daigo2010)
