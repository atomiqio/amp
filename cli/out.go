// NOTICE
// This file is copyrighted by Docker under the Apache 2.0 license
// The original can be found here:
// https://github.com/appcelerator/amp/docker/docker/blob/master/cli/command/out.go

package cli

import (
	"io"
	"log"
	"os"

	"github.com/appcelerator/amp/docker/docker/pkg/term"
)

// OutStream is an output stream used by the Cli to write normal program
// output.
type OutStream struct {
	out        io.Writer
	fd         uintptr
	isTerminal bool
	state      *term.State
}

func (o *OutStream) Write(p []byte) (int, error) {
	return o.out.Write(p)
}

// FD returns the file descriptor number for this stream
func (o *OutStream) FD() uintptr {
	return o.fd
}

// IsTerminal returns true if this stream is connected to a terminal
func (o *OutStream) IsTerminal() bool {
	return o.isTerminal
}

// SetRawTerminal sets raw mode on the output terminal
func (o *OutStream) SetRawTerminal() (err error) {
	if os.Getenv("NORAW") != "" || !o.isTerminal {
		return nil
	}
	o.state, err = term.SetRawTerminalOutput(o.fd)
	return err
}

// RestoreTerminal restores normal mode to the terminal
func (o *OutStream) RestoreTerminal() {
	if o.state != nil {
		term.RestoreTerminal(o.fd, o.state) // nolint
	}
}

// GetTtySize returns the height and width in characters of the tty
func (o *OutStream) GetTtySize() (uint, uint) {
	if !o.isTerminal {
		return 0, 0
	}
	ws, err := term.GetWinsize(o.fd)
	if err != nil {
		log.Printf("Error getting size: %s\n", err)
		if ws == nil {
			return 0, 0
		}
	}
	return uint(ws.Height), uint(ws.Width)
}

// NewOutStream returns a new OutStream object from a Writer
func NewOutStream(out io.Writer) *OutStream {
	fd, isTerminal := term.GetFdInfo(out)
	return &OutStream{out: out, fd: fd, isTerminal: isTerminal}
}
