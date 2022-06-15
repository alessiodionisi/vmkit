package main

import (
	"os"
	"os/signal"

	"github.com/adnsio/vmkit/server"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type rootCmdOptions struct {
}

func newRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "vmkitd",
		RunE: func(_ *cobra.Command, _ []string) error {
			if err := runRootCmd(rootCmdOptions{}); err != nil {
				handleCmdError(err)
				return nil
			}

			return nil
		},
	}

	return rootCmd
}

func runRootCmd(_ rootCmdOptions) error {
	srv := server.NewServer()

	log.Info().Msg("starting server")
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal().Err(err)
		}
	}()
	defer srv.Shutdown()

	interruptSig := make(chan os.Signal, 1)
	signal.Notify(interruptSig, os.Interrupt)

	<-interruptSig
	log.Info().Msg("interrupt signal received, gracefully shutting down")

	return nil
}
