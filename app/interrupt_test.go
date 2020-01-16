package app

import (
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProvideInterruptChannel(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Interrupt test not supported on windows")
	}

	// Setup
	myProcess, err := os.FindProcess(os.Getpid())
	if err != nil {
		assert.Fail(t, "failed to find my process: %v", err)
	}

	// Tested Code
	interrupt := ProvideInterruptChannel()
	if err := myProcess.Signal(os.Interrupt); err != nil {
		assert.Fail(t, "failed to send interrupt signal: %v", err)
	}

	// Asserts
	assert.Equal(t, os.Interrupt, <-interrupt)
}
