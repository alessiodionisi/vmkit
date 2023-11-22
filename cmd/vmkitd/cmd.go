package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"os/signal"
	"path"

	"github.com/alessiodionisi/vmkit/engine"
	"github.com/alessiodionisi/vmkit/server"
	"github.com/spf13/cobra"
)

type options struct {
	logLevel string
	listen   string
}

func newCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "vmkitd",
		Long: "VMKit daemon is the server component of VMKit, which exposes a gRPC API to manage virtual machines and their resources",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newOptions(cmd)
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
		RunE: func(cmd *cobra.Command, args []string) error {
			options, err := newOptions(cmd)
			if err != nil {
				return err
			}

			if err := run(options); err != nil {
				slog.Error(fmt.Sprintf("error running vmkitd: %s", err))
				os.Exit(1)
			}

			return nil
		},
	}

	cmd.PersistentFlags().String("log-level", "info", "Logging level, one of debug, info, warn or error")

	cmd.Flags().String("listen", "unix://~/.vmkit/vmkit.sock", "Set the listening address")

	return cmd
}

func newOptions(cmd *cobra.Command) (*options, error) {
	logLevel, err := cmd.Flags().GetString("log-level")
	if err != nil {
		return nil, err
	}

	listen, err := cmd.Flags().GetString("listen")
	if err != nil {
		return nil, err
	}

	return &options{
		logLevel: logLevel,
		listen:   listen,
	}, nil
}

func run(opts *options) error {
	parsedListen, err := url.Parse(opts.listen)
	if err != nil {
		return fmt.Errorf("error parsing listening address: %w", err)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting user home directory: %w", err)
	}

	vmkitDir := path.Join(homeDir, ".vmkit")
	if _, err := os.Stat(vmkitDir); errors.Is(err, os.ErrNotExist) {
		slog.Debug(fmt.Sprintf("creating vmkit directory: %s", vmkitDir))

		if err := os.Mkdir(vmkitDir, 0755); err != nil {
			return fmt.Errorf("error creating vmkit directory: %w", err)
		}
	}

	slog.Info("initializing engine")
	eng, err := engine.New(
		slog.Default(),
		vmkitDir,
	)
	if err != nil {
		return fmt.Errorf("error initializing engine: %w", err)
	}

	srv := server.New(eng, slog.Default())

	errorChan := make(chan error, 1)
	go func() {
		slog.Info(fmt.Sprintf("server listening on %s", parsedListen.String()))
		if err := srv.ListenAndServe(parsedListen.Scheme, parsedListen.Host); err != nil {
			errorChan <- err
		}
	}()

	interruptSignalChan := make(chan os.Signal, 1)
	go func() {
		signal.Notify(interruptSignalChan, os.Interrupt)
	}()

	select {
	case err := <-errorChan:
		return fmt.Errorf("error starting server: %w", err)
	case <-interruptSignalChan:
		slog.Info("received interrupt signal, shutting down")

		if err := srv.Shutdown(); err != nil {
			return fmt.Errorf("error shutting down server: %w", err)
		}
	}

	return nil
}
