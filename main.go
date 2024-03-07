package main

import (
	"os"

	"github.com/alessiodionisi/vmkit/cmd"
)

func main() {
	if err := cmd.New().Execute(); err != nil {
		os.Exit(1)
	}
}
