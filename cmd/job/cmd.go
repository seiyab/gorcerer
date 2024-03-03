package job

import (
	"os"

	"github.com/google/go-github/v60/github"
	"github.com/seiyab/gorcerer/output"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var issue int
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

			out := output.Println
			if issue != 0 {
				c := github.NewClient(nil).
					WithAuthToken(os.Getenv("GITHUB_TOKEN"))
				out = output.NewIssueComment(c, issue)
			}
			if err := out(m); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().IntVar(&issue, "issue", 0, "issue number")
	return cmd
}
