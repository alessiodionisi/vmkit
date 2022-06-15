package main

import (
	"fmt"
	"os"
)

func handleCmdError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
