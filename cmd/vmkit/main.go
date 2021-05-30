package main

import "os"

func main() {
	rootCommand := newRootCommand()

	if err := rootCommand.Execute(); err != nil {
		os.Exit(1)
	}
}
