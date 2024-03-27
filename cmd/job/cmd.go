package job

import (
	"os"

	"github.com/google/go-github/v60/github"
	"github.com/seiyab/gorcerer/output"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var owner string
	var repository string
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
				target := output.IssueCommentTarget{
					Owner:      "seiyab",
					Repository: "gorcerer",
					Issue:      issue,
				}
				if owner != "" {
					target.Owner = owner
				}
				if repository != "" {
					target.Repository = repository
				}
				out = output.NewIssueComment(c, target)
			}
			if err := out(m); err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().IntVar(&issue, "issue", 0, "issue number")
	cmd.Flags().StringVar(&owner, "owner", "", "owner")
	cmd.Flags().StringVar(&repository, "repository", "", "repository")
	return cmd
}
