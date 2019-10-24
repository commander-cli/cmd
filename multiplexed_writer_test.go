package cmd

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMultiplexedWriter(t *testing.T) {
	writer01 := bytes.Buffer{}
	writer02 := bytes.Buffer{}
	// Test another io.Writer interface type
	r, w, _ := os.Pipe()

	writer := NewMultiplexedWriter(&writer01, &writer02, w)
	n, err := writer.Write([]byte(`test`))

	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "test", writer01.String())
	assert.Equal(t, "test", writer02.String())

	data := make([]byte, 4)
	n, err = r.Read(data)
	assert.Nil(t, err)
	assert.Equal(t, 4, n)
	assert.Equal(t, "test", string(data))
}

func TestMultiplexedWriter_SingleWirter(t *testing.T) {
	writer01 := bytes.Buffer{}

	writer := NewMultiplexedWriter(&writer01)

	n, _ := writer.Write([]byte(`another`))

	assert.Equal(t, 7, n)
	assert.Equal(t, "another", writer01.String())
}

func TestMultiplexedWriter_Fail(t *testing.T) {
	writer := NewMultiplexedWriter(InvalidWriter{})

	n, err := writer.Write([]byte(`another`))

	assert.Equal(t, 0, n)
	assert.Equal(t, "Error in writer: failed", err.Error())
}

type InvalidWriter struct {
}

func (w InvalidWriter) Write(p []byte) (n int, err error) {
	return 0, fmt.Errorf("failed")
}