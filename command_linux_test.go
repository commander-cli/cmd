package cmd

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func TestCommand_ExecuteStderr(t *testing.T) {
	cmd := NewCommand(">&2 echo hello")

	err := cmd.Execute()

	assert.Nil(t, err)
	assert.Equal(t, "hello\n", cmd.Stderr())
}

func TestCommand_WithTimeout(t *testing.T) {
	cmd := NewCommand("sleep 0.1;", WithTimeout(1*time.Millisecond))

	err := cmd.Execute()

	assert.NotNil(t, err)
	// Sadly a process can not be killed every time :(
	containsMsg := strings.Contains(err.Error(), "Timeout occurred and can not kill process with pid") || strings.Contains(err.Error(), "Command timed out after 1ms")
	assert.True(t, containsMsg)
}

func TestCommand_WithValidTimeout(t *testing.T) {
	cmd := NewCommand("sleep 0.01;", WithTimeout(500*time.Millisecond))

	err := cmd.Execute()

	assert.Nil(t, err)
}

func TestCommand_WithWorkingDir(t *testing.T) {
	setWorkingDir := func(c *Command) {
		c.WorkingDir = "/tmp"
	}

	cmd := NewCommand("pwd", setWorkingDir)
	cmd.Execute()

	assert.Equal(t, "/tmp\n", cmd.Stdout())
}

func TestCommand_WithStandardStreams(t *testing.T) {
	tmpFile, _ := ioutil.TempFile("/tmp", "stdout_")
	originalStdout := os.Stdout
	os.Stdout = tmpFile

	// Reset os.Stdout to its original value
	defer func() {
		os.Stdout = originalStdout
	}()

	cmd := NewCommand("echo hey", WithStandardStreams)
	cmd.Execute()

	r, err := ioutil.ReadFile(tmpFile.Name())
	assert.Nil(t, err)
	assert.Equal(t, "hey\n", string(r))
}

func TestCommand_WithoutTimeout(t *testing.T) {
	cmd := NewCommand("sleep 0.001; echo hello", WithoutTimeout)

	err := cmd.Execute()

	assert.Nil(t, err)
	assert.Equal(t, "hello\n", cmd.Stdout())
}

func TestCommand_WithInvalidDir(t *testing.T) {
	cmd := NewCommand("echo hello", WithWorkingDir("/invalid"))
	err := cmd.Execute()
	assert.NotNil(t, err)
	assert.Equal(t, "chdir /invalid: no such file or directory", err.Error())
}

func TestWithInheritedEnvironment(t *testing.T) {
	os.Setenv("FROM_OS", "is on os")
	os.Setenv("OVERWRITE", "is on os but should be overwritten")
	defer func() {
		os.Unsetenv("FROM_OS")
		os.Unsetenv("OVERWRITE")
	}()

	c := NewCommand(
		"echo $FROM_OS $OVERWRITE",
		WithInheritedEnvironment(map[string]string{"OVERWRITE": "overwritten"}))
	c.Execute()

	assertEqualWithLineBreak(t, "is on os overwritten", c.Stdout())
}

func TestWithCustomStderr(t *testing.T) {
	writer := bytes.Buffer{}
	c := NewCommand(">&2 echo stderr; sleep 0.01; echo stdout;", WithCustomStderr(&writer))
	c.Execute()

	assertEqualWithLineBreak(t, "stderr", writer.String())
	assertEqualWithLineBreak(t, "stdout", c.Stdout())
	assertEqualWithLineBreak(t, "stderr", c.Stderr())
	assertEqualWithLineBreak(t, "stderr\nstdout", c.Combined())
}

func TestWithCustomStdout(t *testing.T) {
	writer := bytes.Buffer{}
	c := NewCommand(">&2 echo stderr; sleep 0.01; echo stdout;", WithCustomStdout(&writer))
	c.Execute()

	assertEqualWithLineBreak(t, "stdout", writer.String())
	assertEqualWithLineBreak(t, "stdout", c.Stdout())
	assertEqualWithLineBreak(t, "stderr", c.Stderr())
	assertEqualWithLineBreak(t, "stderr\nstdout", c.Combined())
}

func TestWithStandardStreams(t *testing.T) {
	out, err := CaptureStandardOutput(func() interface{} {
		c := NewCommand(">&2 echo stderr; sleep 0.01; echo stdout;", WithStandardStreams)
		err := c.Execute()
		return err
	})

	assertEqualWithLineBreak(t, "stderr\nstdout", out)
	assert.Nil(t, err)
}
