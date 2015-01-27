# Colar

Colar helps makes the error messages from go a bit more readable by adding color highlighting

## Installation
`go get github.com/CorgiMan/colar`
`cd $GOPATH/src/github.com/CorgiMan/colar`
`go install`

You might want to add the go binaries folder to your path with
`export PATH=$PATH:$GOPATH/bin`

## Usage:
Just run your program like normal and pipe the output into colar
`go run main.go | colar`

During compilation errors and panics go files and line numbers turn op red
During panic a summary at the end of the output is generated that informs the user about the number of goroutines that were running and the file the panic occured in.
 