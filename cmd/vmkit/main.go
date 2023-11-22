package main

import (
	"os"
)

func main() {
	if err := newCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
