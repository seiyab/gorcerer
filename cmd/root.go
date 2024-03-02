package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/seiyab/gorcerer/cmd/inspect"
	"github.com/spf13/cobra"
)

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "gorcerer",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.Help(); err != nil {
				return err
			}
			return errors.New("no command specified")
		},
	}
	return cmd
}

func Execute() {
	cmd := rootCmd()
	cmd.AddCommand(inspect.Cmd())
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
