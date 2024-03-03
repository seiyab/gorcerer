package inspect

import (
	"errors"
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "inspect",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := inspect(args[0])
			if err != nil {
				return err
			}
			_ = github.NewClient(nil)
			fmt.Println(m)
			return errors.New("not implemented yet")
		},
	}
	return cmd
}
