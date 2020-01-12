package app

import (
	"os"
	"os/signal"
	"syscall"
)

// InterruptChannel is a channel containing interrupts to stop the application
type InterruptChannel <-chan os.Signal

// ProvideInterruptChannel builds an InterruptChannel from SIGTERM signals
func ProvideInterruptChannel() InterruptChannel {
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	return interrupt
}
