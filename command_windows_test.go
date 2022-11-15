package cmd

import (
	"strings"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestCommand_ExecuteStderr(t *testing.T) {
	cmd := NewCommand("echo hello 1>&2")

	err := cmd.Execute()

	assert.Nil(t, err)
	assertEqualWithLineBreak(t, "hello ", cmd.Stderr())
}

func TestCommand_WithTimeout(t *testing.T) {
	cmd := NewCommand("timeout 0.005;", WithTimeout(5*time.Millisecond))

	err := cmd.Execute()

	assert.NotNil(t, err)
	// This is needed because windows sometimes can not kill the process :(
	containsMsg := strings.Contains(err.Error(), "Timeout occurred and can not kill process with pid") || strings.Contains(err.Error(), "Command timed out after 5ms")
	assert.True(t, containsMsg)
}

func TestCommand_WithValidTimeout(t *testing.T) {
	cmd := NewCommand("timeout 0.01;", WithTimeout(1000*time.Millisecond))

	err := cmd.Execute()

	assert.Nil(t, err)
}

func TestCommand_WithUser(t *testing.T) {
	onehundred := 100
	token := syscall.Token(uintptr(unsafe.Pointer(&onehundred)))
	cmd := NewCommand("echo hello", WithUser(token))
	err := cmd.Execute()
	assert.Error(t, err)
}
