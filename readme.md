# Colar

Colar helps makes the error messages from go a bit more readable by adding color highlighting

## Installation
``` bash
go get github.com/CorgiMan/colar
cd $GOPATH/src/github.com/CorgiMan/colar
go install
```

You might want to add the go binaries folder to your path with
``` bash
export PATH=$PATH:$GOPATH/bin`
```

## Usage:
Just run your program like normal and pipe the output into colar
``` bash
go run main.go | colar
```

During compilation errors and panics go files and line numbers turn op red. During a panic colar appends a summary to the end of the output. The summary informs the user about the number of goroutines that were running at the time of the panic and the file the panic occured in.
