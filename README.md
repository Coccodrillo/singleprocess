# Single process checker

singleprocess is a library for Go that implements checking if a process with
the same name already exists. It is useful for checking cron jobs you don't
want to overlap.

It uses Mitchel Hashimoto's go-ps. I have only tested it on "darwin" and "linux" systems.

## Installation

Install using standard `go get`:

```
$ go get github.com/coccodrillo/singleprocess
```

Usage:
```
import (
	...
	"github.com/coccodrillo/singleprocess"
	...
)

func main() {
	if singleprocess.IsAnotherInstanceRunning() {
		// Do something, for instance exit
		os.Exit(1)
	}
}
```

## TODO

Adapt for other plaforms, namely Windows, FreeBSD...