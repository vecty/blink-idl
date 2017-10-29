# Blink IDL files

This repository contains IDL files from the Blink browser engine. Effectively, these files describe the public JavaScript API exposed by Google Chrome and other Chromium browsers.

## Why

The actual blink repository is around 5 GB of data, and can take multiple hours to clone. This repository only contains the IDL files, which means cloning is substantially quicker.

The IDL files are used by [vectapi](https://github.com/vecty/vectapi) to generate the Vecty API.

## Updating

To update the IDL files in this repository, run `go run update.go`.
