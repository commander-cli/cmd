# cmd package

A simple package to execute shell commands on windows, darwin and windows.

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
cmd := NewCommand("echo hello", WithStandardStreams)
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

 - Reports
    - Coverage reports
    - Go report
    - Codeclimate
 - Travis-Pipeline
 - Documentation
 - os.Stdout and os.Stderr output access after execution