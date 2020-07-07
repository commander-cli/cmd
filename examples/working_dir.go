package examples

import "github.com/commander-cli/cmd"

// CreateWithWorkingDir sets the current working directory
func CreateWithWorkingDir() {
	setWorkingDir := func(c *cmd.Command) {
		c.WorkingDir = "/tmp"
	}
	c := cmd.NewCommand("pwd", cmd.WithStandardStreams, setWorkingDir)
	c.Execute()
}
