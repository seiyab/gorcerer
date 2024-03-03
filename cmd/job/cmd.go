package job

import (
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "job",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			j, err := Parse(args[0])
			if err != nil {
				return err
			}

			m, err := j()
			if err != nil {
				return err
			}
			_, err = fmt.Println(m)
			if err != nil {
				return err
			}
			_ = github.NewClient(nil)
			return nil
		},
	}
	return cmd
}
