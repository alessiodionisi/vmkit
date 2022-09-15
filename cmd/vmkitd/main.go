package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func handleCmdError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	rootCmd := newRootCmd()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
