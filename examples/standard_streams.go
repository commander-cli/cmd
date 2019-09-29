package examples

import "github.com/SimonBaeumer/cmd"

func CreateNewCommandWithStandardStream() {
	c := cmd.NewCommand("echo hello; sleep 1; echo another;", cmd.WithStandardStreams)
	c.Execute()
}
