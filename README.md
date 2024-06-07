# Go custom logging by Neat Path Networks

Copyright ©️ 2024 Neat Path Networks GmbH

Authors(s):
  - Dragos Vingarzan - dragos@neatpath.net

## About

Super basic for now and only meant for console and mostly for humans to read.

Custom because:
- adds filename and line
- pads log prefix to align the text of the log

See [log.go](log.go) for how to use - basically copy-paste the commented-out code, set your package name and done.

## ToDos

- consider [slog](https://pkg.go.dev/golang.org/x/exp/slog) for output to other tools.
- remove from prefix year, nano-seconds, etc to make it more compact
