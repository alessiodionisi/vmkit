// Spin up Linux VMs with QEMU
// Copyright (C) 2021 VMKit Authors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"os"
	"path"
	"runtime"

	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/adnsio/vmkit/pkg/qemu"
	"github.com/spf13/cobra"
)

type globalOptions struct {
	// driver             string
	qemuExecutableName string
	configPath         string
}

func newRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Short: "Spin up Linux VMs with QEMU",
		Use:   "vmkit",
	}

	// cmd.AddCommand(newLogsCommand())
	cmd.AddCommand(newCompletionCommand())
	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newExecCommand())
	cmd.AddCommand(newImagesCommand())
	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newPullCommand())
	cmd.AddCommand(newRemoveCommand())
	cmd.AddCommand(newRestartCommand())
	cmd.AddCommand(newSSHCommand())
	cmd.AddCommand(newStartCommand())
	cmd.AddCommand(newStopCommand())
	cmd.AddCommand(newXtermCommand())

	homePath, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	defaultConfigPath := path.Join(homePath, ".vmkit")

	var defaultQEMUExecutableName string
	switch runtime.GOARCH {
	case "arm64":
		defaultQEMUExecutableName = qemu.Aarch64ExecutableName

	case "amd64":
		defaultQEMUExecutableName = qemu.X8664ExecutableName

	default:
		return nil, ErrUnsupportedArchitecture
	}

	configPathEnv := os.Getenv("VMKIT_CONFIG_PATH")
	if configPathEnv != "" {
		defaultConfigPath = configPathEnv
	}

	qemuExecutableNameEnv := os.Getenv("VMKIT_QEMU_EXECUTABLE_NAME")
	if qemuExecutableNameEnv != "" {
		defaultQEMUExecutableName = qemuExecutableNameEnv
	}

	cmd.PersistentFlags().String("config-path", defaultConfigPath, "configuration path (env VMKIT_CONFIG_PATH)")
	cmd.PersistentFlags().String("qemu-executable-name", defaultQEMUExecutableName, "qemu executable name (env VMKIT_QEMU_EXECUTABLE_NAME)")

	return cmd, nil
}

func newGlobalOptions(cmd *cobra.Command) (*globalOptions, error) {
	configPath, err := cmd.Flags().GetString("config-path")
	if err != nil {
		return nil, err
	}

	qemuExecutableName, err := cmd.Flags().GetString("qemu-executable-name")
	if err != nil {
		return nil, err
	}

	return &globalOptions{
		qemuExecutableName: qemuExecutableName,
		configPath:         configPath,
	}, nil
}

func newEngine(opts *globalOptions) (*engine.Engine, error) {
	return engine.New(&engine.NewOptions{
		QEMUExecutableName: opts.qemuExecutableName,
		Path:               opts.configPath,
		Writer:             os.Stderr,
	})
}
