package cmd

import (
	"os/exec"
	"syscall"
)

func createBaseCommand(c *Command) *exec.Cmd {
	cmd := exec.Command("/bin/sh", "-c", c.Command)
	return cmd
}

func WithUser(credential syscall.Credential) func(c *Command) {
	return func(c *Command) {
		c.baseCommand.SysProcAttr = &syscall.SysProcAttr{
			Credential: &credential,
		}
	}
}
