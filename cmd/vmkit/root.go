// Spin up Linux VMs with QEMU and Apple virtualization framework
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

	"github.com/adnsio/vmkit/pkg/driver/qemu"
	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/spf13/cobra"
)

type globalOptions struct {
	driver               string
	driverExecutableName string
	configPath           string
}

func newRootCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Short: "Spin up Linux VMs with QEMU and Apple virtualization framework",
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

	var defaultDriver engine.Driver
	var defaultDriverExecutableName string

	switch runtime.GOOS {
	case "darwin":
		defaultDriver = engine.DriverQEMU

	case "linux":
		defaultDriver = engine.DriverQEMU

	default:
		return nil, ErrUnsupportedOperatingSystem
	}

	if defaultDriver == engine.DriverAVFVM {
		defaultDriverExecutableName = "avfvm"
	} else {
		switch runtime.GOARCH {
		case "arm64":
			defaultDriverExecutableName = qemu.Aarch64ExecutableName

		case "amd64":
			defaultDriverExecutableName = qemu.X86_64ExecutableName

		default:
			return nil, ErrUnsupportedArchitecture
		}
	}

	configPathEnv := os.Getenv("VMKIT_CONFIG_PATH")
	if configPathEnv != "" {
		defaultConfigPath = configPathEnv
	}

	driverEnv := os.Getenv("VMKIT_DRIVER")
	if driverEnv != "" {
		defaultDriver = engine.Driver(driverEnv)
	}

	driverExecutableNameEnv := os.Getenv("VMKIT_DRIVER_EXECUTABLE_NAME")
	if driverExecutableNameEnv != "" {
		defaultDriverExecutableName = driverExecutableNameEnv
	}

	cmd.PersistentFlags().String("config-path", defaultConfigPath, "configuration path (env VMKIT_CONFIG_PATH)")
	cmd.PersistentFlags().String("driver-executable-name", defaultDriverExecutableName, "driver executable name (env VMKIT_DRIVER_EXECUTABLE_NAME)")
	cmd.PersistentFlags().String("driver", string(defaultDriver), "driver to use (env VMKIT_DRIVER)")

	return cmd, nil
}

func newGlobalOptions(cmd *cobra.Command) (*globalOptions, error) {
	configPath, err := cmd.Flags().GetString("config-path")
	if err != nil {
		return nil, err
	}

	driver, err := cmd.Flags().GetString("driver")
	if err != nil {
		return nil, err
	}

	driverExecutableName, err := cmd.Flags().GetString("driver-executable-name")
	if err != nil {
		return nil, err
	}

	return &globalOptions{
		driver:               driver,
		driverExecutableName: driverExecutableName,
		configPath:           configPath,
	}, nil
}

func newEngine(opts *globalOptions) (*engine.Engine, error) {
	return engine.New(&engine.NewOptions{
		Driver:               engine.Driver(opts.driver),
		DriverExecutableName: opts.driverExecutableName,
		Path:                 opts.configPath,
		Writer:               os.Stderr,
	})
}
