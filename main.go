package main

import (
	"fmt"
	"os"

	"github.com/adnsio/vmkit/cmd"
)

func main() {
	cmd, err := cmd.New()
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
