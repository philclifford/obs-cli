//go:build man
// +build man

package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var (
	manCmd = &cobra.Command{
		Use:    "man [directory]",
		Short:  "Generates manpages",
		Long:   `Generates manpages in the given directory`,
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("man needs a target directory")
			}
			return generateManPages(args[0])
		},
	}
)

func init() {
	rootCmd.AddCommand(manCmd)
}

func generateManPages(targetDirectory string) error {
	header := &doc.GenManHeader{
		Section: "1",
		Source:  "Auto generated by muesli/obs-cli",
	}

	return doc.GenManTree(rootCmd, header, targetDirectory)
}