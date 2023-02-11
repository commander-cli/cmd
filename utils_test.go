package cmd

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertEqualWithLineBreak(t *testing.T, expected string, actual string) {
	if runtime.GOOS == "windows" {
		expected = expected + "\r\n"
	} else {
		expected = expected + "\n"
	}

	assert.Equal(t, expected, actual)
}
