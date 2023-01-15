[![CI](https://github.com/commander-cli/cmd/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/commander-cli/cmd/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/commander-cli/cmd?status.svg)](https://godoc.org/github.com/commander-cli/cmd)
[![Test Coverage](https://api.codeclimate.com/v1/badges/31911138f62cea099c31/test_coverage)](https://codeclimate.com/github/commander-cli/cmd/test_coverage)
[![Maintainability](https://api.codeclimate.com/v1/badges/af3487439a313d580619/maintainability)](https://codeclimate.com/github/commander-cli/cmd/maintainability)
[![Go Report Card](https://goreportcard.com/badge/github.com/commander-cli/cmd)](https://goreportcard.com/report/github.com/commander-cli/cmd)

# cmd package

A simple package to execute shell commands on linux, darwin and windows.

## Installation

Install the latest version with:

```bash
$ go get -u github.com/commander-cli/cmd
```

or an exact version:

```bash
$ go get -u github.com/commander-cli/cmd@v1.0.0
```

## Usage

```go
c := cmd.NewCommand("echo hello")

err := c.Execute()
if err != nil {
    panic(err.Error())    
}

fmt.Println(c.Stdout())
fmt.Println(c.Stderr())
```

### Configure the command

To configure the command an option function can be passed which receives the
command object as an argument passed by reference.

Default option functions:

```
cmd.WithCustomBaseCommand(*exec.Cmd)
cmd.WithStandardStreams
cmd.WithCustomStdout(...io.Writers)
cmd.WithCustomStderr(...io.Writers)
cmd.WithTimeout(time.Duration)
cmd.WithoutTimeout
cmd.WithWorkingDir(string)
cmd.WithEnvironmentVariables(cmd.EnvVars)
cmd.WithInheritedEnvironment(cmd.EnvVars)
```

See [godocs for details][].

#### Example

```go
c := cmd.NewCommand("echo hello", cmd.WithStandardStreams)
c.Execute()
```

#### Set custom options

```go
setWorkingDir := func (c *Command) {
    c.WorkingDir = "/tmp/test"
}

c := cmd.NewCommand("pwd", setWorkingDir)
c.Execute()
```

## Contributing

If you would like to contribute please submit a pull request.
For bug fixes/minor changes a simple pull request will
suffice. If the change is large or you would like to have a feature
discussion before implementation feel free to open an issue.

If you have a feature request or bug report please open an issue.

### Development

Please fork the project and do your development there. Please use a
meaningful branch name as well as adhere to [commitlint rules][].

If you would like the precommit hooks run:

```
make init
```

To run the test suite:

```
make test
```

*Reminder:* The goal of this project is to ensure we abstract the OS specific
command execution as mush as possible. Ensure your change is compatible
with linux, windows and osx. If unable to test on every operating system
(help needed for windows (:) the CI will take care of that for you.

[commitlint rules]: https://www.conventionalcommits.org/en/v1.0.0/
[godocs for details]: https://godoc.org/github.com/commander-cli/cmd

