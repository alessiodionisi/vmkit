package main

import (
	"errors"
	"net"
	"os"
	"os/signal"

	"github.com/adnsio/vmkit/pkg/driver"
	"github.com/adnsio/vmkit/pkg/engine"
	"github.com/adnsio/vmkit/pkg/rpc"
	"github.com/adnsio/vmkit/pkg/util"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/spf13/cobra"
)

type rootOptions struct {
	address string
	driver  string
	qemu    string
	avfvm   string
}

func newRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vmkitd",
		Short: "VMKit Daemon",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			_, err := cmd.Flags().GetString("log-level")
			if err != nil {
				return err
			}

			zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
			zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

			zerolog.SetGlobalLevel(zerolog.DebugLevel)

			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			address, err := cmd.Flags().GetString("address")
			if err != nil {
				return err
			}

			driver, err := cmd.Flags().GetString("driver")
			if err != nil {
				return err
			}

			avfvm, err := cmd.Flags().GetString("avfvm")
			if err != nil {
				return err
			}

			qemu, err := cmd.Flags().GetString("qemu")
			if err != nil {
				return err
			}

			opts := &rootOptions{
				address: address,
				driver:  driver,
				avfvm:   avfvm,
				qemu:    qemu,
			}

			if err := runRoot(opts); err != nil {
				log.Error().Err(err).Stack().Send()
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringP("log-level", "l", "info", "sets the log level (debug, info, error)")
	cmd.PersistentFlags().StringP("address", "a", "unix:///var/run/vmkitd.sock", "sets the listening socket or port (unix, tcp)")
	cmd.PersistentFlags().String("avfvm", driver.AVFVMExecutableName, "sets the avfvm binary")
	cmd.PersistentFlags().String("qemu", driver.QEMUExecutableName, "sets the qemu binary")
	cmd.PersistentFlags().String("driver", string(driver.Preferred), "sets the driver (avfvm, qemu)")

	return cmd
}

func runRoot(opts *rootOptions) error {
	var drv driver.Driver

	switch driver.DriverType(opts.driver) {
	case driver.DriverTypeAVFVM:
		var err error

		drv, err = driver.NewAVFVM(opts.avfvm)
		if err != nil {
			return err
		}
	case driver.DriverTypeQEMU:
		var err error

		drv, err = driver.NewQEMU(opts.qemu)
		if err != nil {
			return err
		}
	default:
		return driver.ErrNotSupported
	}

	eng, err := engine.NewEngine(drv)
	if err != nil {
		return err
	}

	if err := eng.ReloadVirtualMachines(); err != nil {
		return err
	}

	rpcSrv, err := rpc.NewServer(&rpc.NewServerOptions{
		Engine: eng,
	})
	if err != nil {
		return err
	}

	parsedAddress, err := util.NewNetworkAddress(opts.address)
	if err != nil {
		return err
	}

	log.Debug().Msgf("starting listener on %s %s...", parsedAddress.Network, parsedAddress.Address)
	l, err := net.Listen(parsedAddress.Network, parsedAddress.Address)
	if err != nil {
		return err
	}
	defer func() {
		log.Debug().Msg("gracefully closing listener...")

		if err := l.Close(); err != nil {
			log.Error().Err(err).Stack().Send()
			os.Exit(1)
		}
	}()

	go func() {
		for {
			conn, err := l.Accept()
			if err != nil {
				if errors.Is(err, net.ErrClosed) {
					continue
				}

				log.Error().Err(err).Stack().Send()
				os.Exit(1)
			}

			log.Debug().Msgf("accepted connection from %s", conn.RemoteAddr())

			go rpcSrv.ServeConn(conn)
		}
	}()

	interruptCh := make(chan os.Signal, 1)
	signal.Notify(interruptCh, os.Interrupt)

	<-interruptCh
	log.Debug().Msg("got an interrupt signal")

	return nil
}
