package examples

import "github.com/SimonBaeumer/cmd"

func CreateWithWorkingDir() {
	setWorkingDir := func(c *cmd.Command) {
		c.WorkingDir = "/tmp"
	}
	c := cmd.NewCommand("pwd", cmd.WithStandardStreams, setWorkingDir)
	c.Execute()
}
