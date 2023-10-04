# GO Template

![Golang](https://github.com/zaphiro-technologies/go-template/actions/workflows/golang.yaml/badge.svg)

A golang basic template for starting new golang projects at Zaphiro.

## Default libraries

The following libraries are used in this template:

- [logging](https://github.com/zaphiro-technologies/logging) a wrapper
  preconfigured for [zerolog](https://pkg.go.dev/github.com/rs/zerolog).
- [testify](https://github.com/stretchr/testify) for test asserting.
- [godotenv](https://github.com/joho/godotenv) for loading .env content.
- [go-envconfig](https://github.com/sethvargo/go-envconfig) for parsing and
  casting .env.

## Workflows

The repository configures the default workflows available in
[github-workflows](https://github.com/zaphiro-technologies/github-workflows/):

- add-to-project workflow;
- check-pr;
- clean-up-storage;
- markdown;
- release-notes.

Additionally, it includes a
[golang dedicated workflow](.github/workflows/golang.yaml) that:

- The _lint_ job lints files using [`golangci-lint`](https://golangci-lint.run/)
  and [`go-lines`](https://pkg.go.dev/github.com/segmentio/golines).
  Configuration for `golangci-lint` is defined in
  [`.golangci.yaml`](.golangci.yaml).
- The _test_ job tests go and measure coverage using
  [barecheck](https://www.barecheck.com/).
- The _benchmark_ job checks benchmark regression using
  [continuous-benchmark](https://github.com/marketplace/actions/continuous-benchmark).
