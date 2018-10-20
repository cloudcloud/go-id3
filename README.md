# go-id3

Go lib and binary for working with [id3v1](http://id3.org/ID3v1),
[id3v2](http://id3.org/id3v2-00), [id3v2.3](http://id3.org/id3v2.3.0),
and [id3v2.4](http://id3.org/id3v2.4.0-structure) tags.

[![GoDoc](https://godoc.org/github.com/cloudcloud/go-id3?status.svg)](https://godoc.org/github.com/cloudcloud/go-id3)
[![Circle CI](https://circleci.com/gh/cloudcloud/go-id3.svg?style=svg)](https://circleci.com/gh/cloudcloud/go-id3)
[![Go Report Card](https://goreportcard.com/badge/github.com/cloudcloud/go-id3)](https://goreportcard.com/report/github.com/cloudcloud/go-id3)
[![codecov](https://codecov.io/gh/cloudcloud/go-id3/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudcloud/go-id3)
[![Maintainability](https://api.codeclimate.com/v1/badges/843a328c85524bf0ff66/maintainability)](https://codeclimate.com/github/cloudcloud/go-id3/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/843a328c85524bf0ff66/test_coverage)](https://codeclimate.com/github/cloudcloud/go-id3/test_coverage)

## explanation

Whilst putting together a simple single-streaming music application, there was a need to have a simple and comprehensive
tag processor. To be more cross-platform and readily available, this side-project was born to suit the purpose. A
library is provided for other project usage, along with binary for specific command-line usage when required.

## requirements

* ``go``

## installation

Standard installation is simply using ``go get github.com/cloudcloud/go-id3/...`` to import the libraries and compile
the binaries. Your go environment will need to be correctly configured and available for this to function correctly though.

