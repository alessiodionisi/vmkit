package cmd

import "errors"

var (
	ErrUnsupportedArchitecture    = errors.New("vmkit: unsupported architecture")
	ErrUnsupportedOperatingSystem = errors.New("vmkit: unsupported operating system")
)
