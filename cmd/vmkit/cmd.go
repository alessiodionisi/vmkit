package main

import (
	"log/slog"
	"os"

	"github.com/alessiodionisi/vmkit/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type globalOptions struct {
	logLevel string
	daemon   string
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "vmkit",
		Long: "VMKit is a tool to manage virtual machines",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			programLevel := new(slog.LevelVar)
			slogHandler := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
				Level: programLevel,
			})

			slog.SetDefault(slog.New(slogHandler))

			if err := programLevel.UnmarshalText([]byte(globalOptions.logLevel)); err != nil {
				return err
			}

			return nil
		},
	}

	cmd.AddCommand(newDiskCommand())
	cmd.AddCommand(newImageCommand())
	cmd.AddCommand(newVirtualMachineCommand())

	cmd.PersistentFlags().String("log-level", "info", "Logging level, one of debug, info, warn or error")
	cmd.PersistentFlags().String("daemon", "unix://~/.vmkit/vmkit.sock", "Daemon socket to connect to")

	return cmd
}

func newGlobalOptions(cmd *cobra.Command) (*globalOptions, error) {
	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return nil, err
	}

	daemon, err := cmd.Flags().GetString("daemon")
	if err != nil {
		return nil, err
	}

	return &globalOptions{
		logLevel: logLevel,
		daemon:   daemon,
	}, nil
}

func newClient(opts *globalOptions) (client proto.VMKitClient, conn *grpc.ClientConn, err error) {
	conn, err = grpc.Dial(opts.daemon, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	client = proto.NewVMKitClient(conn)

	return client, conn, nil
}
