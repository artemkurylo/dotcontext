// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"os"

	"github.com/artemkurylo/dotcontext/internal/check"
	initcmd "github.com/artemkurylo/dotcontext/internal/init"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dotcontext",
		Short: "Manage project context alongside your code",
		Long:  "dotcontext is a CLI for the dotcontext standard — colocate project documentation with code in a .context/ directory.",
	}

	// init command
	var forceInit bool
	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a .context/ directory",
		Long:  "Create the .context/ directory structure with a starter CONTEXT.md template.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not determine working directory: %w", err)
			}
			return initcmd.Run(dir, forceInit)
		},
	}
	initCmd.Flags().BoolVar(&forceInit, "force", false, "Reinitialize even if .context/ already exists")
	rootCmd.AddCommand(initCmd)

	// check command
	checkCmd := &cobra.Command{
		Use:   "check",
		Short: "Validate the .context/ directory",
		Long:  "Check that the .context/ directory follows the dotcontext standard.",
		RunE: func(cmd *cobra.Command, args []string) error {
			dir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("could not determine working directory: %w", err)
			}
			results, err := check.Run(dir)
			if err != nil {
				return err
			}
			if !check.PrintResults(results) {
				os.Exit(1)
			}
			return nil
		},
	}
	rootCmd.AddCommand(checkCmd)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
