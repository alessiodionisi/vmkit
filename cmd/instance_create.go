package cmd

import (
	"fmt"
	"os"

	"github.com/alessiodionisi/vmkit/bytesize"
	"github.com/alessiodionisi/vmkit/engine"
	"github.com/spf13/cobra"
)

type instanceCreateOptions struct {
	*globalOptions

	volume   string
	image    string
	memory   string
	name     string
	userData string
	vcpu     uint
}

func newInstanceCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [name]",
		Args:  cobra.ExactArgs(1),
		Short: "Create a new instance",
		Example: `  Create a Fedora 39 instance with 4 vcpus, 4 gibibytes of memory and 20 gigabytes of root volume:
    vmkit instance create fedora-vm --image fedora-39 --vcpu 4 --memory 4GiB --volume 20GB`,
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			options, err := newInstanceCreateOptions(globalOptions, cmd, args)
			if err != nil {
				return err
			}

			if err := instanceCreate(options); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.Flags().String("volume", "10GB", "Size of the root volume")
	cmd.Flags().String("image", "", "Name of the image to use as root volume")
	cmd.Flags().String("memory", "1GiB", "Amount of memory to allocate")
	cmd.Flags().String("user-data", "", "Path to user data file to use for cloud-init")
	cmd.Flags().Uint("vcpu", 1, "Number of vCPU to allocate")

	cmd.MarkFlagRequired("image")

	return cmd
}

func newInstanceCreateOptions(globalOptions *globalOptions, cmd *cobra.Command, args []string) (*instanceCreateOptions, error) {
	volume, err := cmd.Flags().GetString("volume")
	if err != nil {
		return nil, err
	}

	image, err := cmd.Flags().GetString("image")
	if err != nil {
		return nil, err
	}

	memory, err := cmd.Flags().GetString("memory")
	if err != nil {
		return nil, err
	}

	userData, err := cmd.Flags().GetString("user-data")
	if err != nil {
		return nil, err
	}

	vcpu, err := cmd.Flags().GetUint("vcpu")
	if err != nil {
		return nil, err
	}

	return &instanceCreateOptions{
		globalOptions: globalOptions,

		volume:   volume,
		image:    image,
		memory:   memory,
		name:     args[0],
		userData: userData,
		vcpu:     vcpu,
	}, nil
}

func instanceCreate(opts *instanceCreateOptions) error {
	e, err := newEngine(opts.globalOptions)
	if err != nil {
		return err
	}

	memory, err := bytesize.Parse(opts.memory)
	if err != nil {
		return err
	}

	volume, err := bytesize.Parse(opts.volume)
	if err != nil {
		return err
	}

	_, err = e.CreateInstance(&engine.CreateInstanceOptions{
		VCPU:     opts.vcpu,
		Image:    opts.image,
		Memory:   memory,
		Name:     opts.name,
		UserData: opts.userData,
		Volume:   volume,
	})
	if err != nil {
		return err
	}

	return nil
}
