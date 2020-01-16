package main

import (
	"fmt"
	"os"
)

func main() {

	if runApp, cleanup, err := InitializeApplicationRunFunc(); err != nil {
		fmt.Fprintf(os.Stderr, "Error during dependency injection: %v", err)
		os.Exit(1)
	} else if err := runApp(); err != nil {
		fmt.Fprintf(os.Stderr, "fatal error: %v", err)
		cleanup()
		os.Exit(1)
	} else {
		cleanup()
	}
}
