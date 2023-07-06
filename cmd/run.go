package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/alessiodionisi/vmkit/engine"
	"github.com/spf13/cobra"
)

type runOptions struct {
	*globalOptions
	cpu          int
	image        string
	memory       int
	name         string
	diskSize     int
	portForwards []string
}

func newRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Example: `  Create and start a Debian 12 (Bookworm) virtual machine:
    vmkit run vm1 -i debian:bookworm

  Create and start a virtual machine with 4 cpus and 4096 mebibytes of ram:
    vmkit run vm1 -i debian:bookworm -c 4 -m 4096
	
	Create and start a virtual machine with 4 cpus, 4096 mebibytes of ram, 20 GB of disk:
    vmkit run vm1 -i debian:bookworm -c 4 -m 4096 -d 20
	
	Create and start a virtual machine and forward port 80 to 8080 on the host:
    vmkit run vm1 -i debian:bookworm -p 8080-80`,
		Args:  cobra.ExactArgs(1),
		Short: "Create and start a new virtual machine",
		Use:   "run [name]",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			image, err := cmd.Flags().GetString("image")
			if err != nil {
				return err
			}

			cpu, err := cmd.Flags().GetInt("cpu")
			if err != nil {
				return err
			}

			memory, err := cmd.Flags().GetInt("memory")
			if err != nil {
				return err
			}

			diskSize, err := cmd.Flags().GetInt("disk-size")
			if err != nil {
				return err
			}

			portForwards, err := cmd.Flags().GetStringSlice("port-forward")
			if err != nil {
				return err
			}

			opts := runOptions{
				cpu:           cpu,
				globalOptions: globalOptions,
				image:         image,
				memory:        memory,
				name:          args[0],
				diskSize:      diskSize,
				portForwards:  portForwards,
			}

			if err := runRun(opts); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().StringP("image", "i", "", "image to use")
	cmd.Flags().IntP("cpu", "c", 1, "number of cpu(s)")
	cmd.Flags().IntP("memory", "m", 512, "ram in mebibytes (MiB)")
	cmd.Flags().IntP("disk-size", "d", 10, "disk size in gigabytes (GB)")
	cmd.Flags().StringSliceP("port-forward", "p", []string{}, "forward host port to the virtual machine")

	cmd.MarkFlagRequired("image")

	return cmd
}

func runRun(opts runOptions) error {
	eng, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	portForwards := map[string]string{}

	for _, portForward := range opts.portForwards {
		hostPort, vmPort, found := strings.Cut(portForward, "-")
		if !found {
			continue
		}

		portForwards[vmPort] = hostPort
	}

	vm, err := eng.CreateVirtualMachine(engine.CreateVirtualMachineOptions{
		CPU:          opts.cpu,
		Image:        opts.image,
		Memory:       opts.memory,
		DiskSize:     opts.diskSize,
		Name:         opts.name,
		PortForwards: portForwards,
	})
	if err != nil {
		return err
	}

	return vm.Start()
}
