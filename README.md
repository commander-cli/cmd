[![Build Status](https://travis-ci.org/SimonBaeumer/cmd.svg?branch=master)](https://travis-ci.org/SimonBaeumer/cmd)
[![GoDoc](https://godoc.org/github.com/SimonBaeumer/cmd?status.svg)](https://godoc.org/github.com/SimonBaeumer/cmd)
[![Test Coverage](https://api.codeclimate.com/v1/badges/af3487439a313d580619/test_coverage)](https://codeclimate.com/github/SimonBaeumer/cmd/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/af3487439a313d580619/maintainability)](https://codeclimate.com/github/SimonBaeumer/cmd/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/SimonBaeumer/cmd)](https://goreportcard.com/report/github.com/SimonBaeumer/cmd)

# cmd package

A simple package to execute shell commands on linux, darwin and windows.

## Installation

`$ go get -u github.com/SimonBaeumer/cmd@v1.0.0`

## Usage

```go
cmd := NewCommand("echo hello")

err := cmd.Execute()
if err != nil {
    panic(err.Error())    
}

fmt.Println(cmd.Stdout())
fmt.Println(cmd.Stderr())
```

### Stream output to stderr and stdout

```go
cmd := NewCommand("echo hello", cmd.WithStandardStreams)
cmd.Execute()
```

### Set custom options

```go
func SetTimeout(c *Command) {
    c.Timeout = 1 * time.Hour
}

func SetWorkingDir(c *Command) {
    c.WorkingDir = "/tmp/test"
}

cmd := NewCommand("pwd", SetTimeout, SetWorkingDir)
cmd.Execute()
```

## Development

### Running tests

```
make test
```

### ToDo

 - os.Stdout and os.Stderr output access after execution