package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newImagesCommand() *cobra.Command {
	cmd := &cobra.Command{
		Short: "List images",
		Use:   "images",
		RunE: func(cmd *cobra.Command, args []string) error {
			globalOptions, err := newGlobalOptions(cmd)
			if err != nil {
				return err
			}

			if err := runImages(globalOptions); err != nil {
				fmt.Printf("Error: %s\n", err)
				os.Exit(1)
			}

			return nil
		},
	}

	return cmd
}

func runImages(opts *globalOptions) error {
	eng, err := newEngine(opts)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	imgs := eng.ListImages()

	tableRows := make([][]string, 0, len(imgs))
	for _, image := range imgs {
		pulled := "No"

		imgPulled, err := image.Pulled()
		if err != nil {
			return err
		}

		if imgPulled {
			pulled = "Yes"
		}

		tableRows = append(tableRows, []string{
			image.Name,
			image.Version,
			image.Description,
			pulled,
		})
	}

	writeTable(&writeTableOptions{
		writer: os.Stdout,
		header: []string{
			"Name",
			"Version",
			"Description",
			"Pulled",
		},
		rows: tableRows,
	})

	return nil
}
