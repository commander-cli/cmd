package main

import "github.com/SimonBaeumer/cmd"

func main() {
	c := cmd.NewCommand("echo hello; sleep 1; echo another;", cmd.WithStandardStreams)
	c.Execute()
}
