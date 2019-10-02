package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"runtime"
	"testing"
	"time"
)

func TestCommand_NewCommand(t *testing.T) {
	cmd := NewCommand("")
	assert.False(t, cmd.Executed())
}

func TestCommand_Execute(t *testing.T) {
	cmd := NewCommand("echo hello")

	err := cmd.Execute()

	assert.Nil(t, err)
	assert.True(t, cmd.Executed())
	assertEqualWithLineBreak(t, "hello", cmd.Stdout())
}

func TestCommand_ExitCode(t *testing.T) {
	cmd := NewCommand("exit 120")

	err := cmd.Execute()

	assert.Nil(t, err)
	assert.Equal(t, 120, cmd.ExitCode())
}

func TestCommand_WithEnvVariables(t *testing.T) {
	envVar := "$TEST"
	if runtime.GOOS == "windows" {
		envVar = "%TEST%"
	}
	cmd := NewCommand(fmt.Sprintf("echo %s", envVar))
	cmd.Env = []string{"TEST=hey"}

	_ = cmd.Execute()

	assertEqualWithLineBreak(t, "hey", cmd.Stdout())
}

func TestCommand_Executed(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			assert.Contains(t, r, "Can not read Stdout if command was not executed")
		}
		assert.NotNil(t, r)
	}()

	c := NewCommand("echo will not be executed")
	_ = c.Stdout()
}

func TestCommand_AddEnv(t *testing.T) {
	c := NewCommand("echo test")
	c.AddEnv("key", "value")
	assert.Equal(t, []string{"key=value"}, c.Env)
}

func TestCommand_AddEnvWithShellVariable(t *testing.T) {
	const TestEnvKey = "COMMANDER_TEST_SOME_KEY"
	os.Setenv(TestEnvKey, "test from shell")
	defer os.Unsetenv(TestEnvKey)

	c := NewCommand(getCommand())
	c.AddEnv("SOME_KEY", fmt.Sprintf("${%s}", TestEnvKey))

	err := c.Execute()

	assert.Nil(t, err)
	assertEqualWithLineBreak(t, "test from shell", c.Stdout())
}

func TestCommand_AddMultipleEnvWithShellVariable(t *testing.T) {
	const TestEnvKeyPlanet = "CMD_TEST_PLANET"
	const TestEnvKeyName = "CMD_TEST_NAME"
	os.Setenv(TestEnvKeyPlanet, "world")
	os.Setenv(TestEnvKeyName, "Simon")
	defer func() {
		os.Unsetenv(TestEnvKeyPlanet)
		os.Unsetenv(TestEnvKeyName)
	}()

	c := NewCommand(getCommand())
	envValue := fmt.Sprintf("Hello ${%s}, I am ${%s}", TestEnvKeyPlanet, TestEnvKeyName)
	c.AddEnv("SOME_KEY", envValue)

	err := c.Execute()

	assert.Nil(t, err)
	assertEqualWithLineBreak(t, "Hello world, I am Simon", c.Stdout())
}

func getCommand() string {
	command := "echo $SOME_KEY"
	if runtime.GOOS == "windows" {
		command = "echo %SOME_KEY%"
	}
	return command
}

func TestCommand_SetOptions(t *testing.T) {
	writer := &bytes.Buffer{}

	setWriter := func(c *Command) {
		c.StdoutWriter = writer
	}

	setTimeout := func(c *Command) {
		c.Timeout = 1 * time.Second
	}

	c := NewCommand("echo test", setTimeout, setWriter)
	err := c.Execute()

	assert.Nil(t, err)
	assert.Equal(t, time.Duration(1000000000), c.Timeout)
	assertEqualWithLineBreak(t, "test", writer.String())
}

func assertEqualWithLineBreak(t *testing.T, expected string, actual string) {
	if runtime.GOOS == "windows" {
		expected = expected + "\r\n"
	} else {
		expected = expected + "\n"
	}

	assert.Equal(t, expected, actual)
}
