package cmd

import (
	"fmt"
	"io"
)

// NewMultiplexedWriter returns a new multiplexer
func NewMultiplexedWriter(writers ...io.Writer) *MultiplexedWriter {
	return &MultiplexedWriter{Writers: writers}
}

// MultiplexedWriter writes to multiple writers at once
type MultiplexedWriter struct {
	Writers []io.Writer
}

// Write writes the given bytes. If one write fails it returns the error
// and bytes of the failed write operation
func (w MultiplexedWriter) Write(p []byte) (n int, err error) {
	for _, o := range w.Writers {
		n, err = o.Write(p)
		if err != nil {
			return 0, fmt.Errorf("Error in writer: %s", err.Error())
		}
	}

	return n, nil
}
