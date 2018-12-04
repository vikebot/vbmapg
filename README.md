# vbmapg

vbmapg is a procedural generation library for 2D-maps in Go. Based on perlin-noise.

[![Build Status](https://travis-ci.org/vikebot/vbmapg.svg?branch=master)](https://travis-ci.org/vikebot/vbmapg)
[![codecov](https://codecov.io/gh/vikebot/vbmapg/branch/master/graph/badge.svg)](https://codecov.io/gh/vikebot/vbmapg)
[![Go Report Card](https://goreportcard.com/badge/github.com/vikebot/vbmapg)](https://goreportcard.com/report/github.com/vikebot/vbmapg)
[![GoDoc](https://godoc.org/github.com/vikebot/vbmapg?status.svg)](https://godoc.org/github.com/vikebot/vbmapg)

## Usage

```shell
$ gen <width> <height> [Optional: <filename.jpeg>]
```

### Create an example map encoded in JPG

```shell
$ go run -mod=vendor cmd/gen/gen.go 100 100 example.jpg
```

## Examples

|  |  |
| ----- | ----- |
| ![Sample1](./example/sample1.jpg) | ![Sample2](./example/sample2.jpg) |
| ![Sample3](./example/sample3.jpg) | ![Sample4](./example/sample4.jpg) |
| ![Sample5](./example/sample5.jpg) |  |
