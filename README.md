# Blink IDL files [![GoDoc](https://godoc.org/github.com/vecty/blink-idl?status.svg)](https://godoc.org/github.com/vecty/blink-idl)

This repository contains IDL files from the Blink browser engine. Effectively, these files describe the public JavaScript API exposed by Google Chrome and other Chromium browsers.

## Why

The actual blink repository is around 5 GB of data, and can take multiple hours to clone. This repository only contains the IDL files, which means cloning is substantially quicker.

The IDL files are used by [vectapi](https://github.com/vecty/vectapi) to generate the Vecty API.

## Updating

Install `vfsgendev`:

```Bash
go get -u github.com/shurcooL/vfsgen/cmd/vfsgendev
```

To update the IDL files in this repository, run `go run update.go`.

## Importable

All of the files are also importable as a Go VFS asset package:

```Go
import "github.com/vecty/blink-idl"
```
