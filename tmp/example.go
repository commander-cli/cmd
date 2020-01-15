package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	command := `
echo stdout.
echo stdout..

>&2 echo stderr.
>&2 echo stderr..

echo stdout.
echo stdout..

>&2 echo stderr.
>&2 echo stderr..

echo stdout.
echo stdout..

>&2 echo stderr.
>&2 echo stderr..

sleep 1
`

	// Create the writer which will be used in stderr and stdout
	combinedWriter := &bytes.Buffer{}

	// Create multiplexed writer
	stdout := MultiplexedWriter{[]io.Writer{combinedWriter}}
	stderr := MultiplexedWriter{[]io.Writer{combinedWriter}}

	// Execute the command
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf(combinedWriter.String())
}

// MultiplexedWriter is used to write into multiple writers,
type MultiplexedWriter struct {
	writers []io.Writer
}

func (w MultiplexedWriter) Write(p []byte) (n int, err error) {
	for _, o := range w.writers {
		n, err = o.Write(p)
		if err != nil {
			return 0, fmt.Errorf("Error occured %s", err.Error())
		}
	}

	return n, nil
}
