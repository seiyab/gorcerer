package inspect

import (
	"errors"
	"os"
	"os/exec"
	"path"

	"github.com/google/go-github/v60/github"
	"github.com/seiyab/gorcerer/utils"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "inspect",
		RunE: func(cmd *cobra.Command, args []string) error {
			return utils.TempDir("gorcerer-inspect", func(tempDir string) error {
				if err := gitClone(
					"git@github.com:seiyab/gorcerer.git",
					path.Join(tempDir, "gorcerer"),
				); err != nil {
					return err
				}

				_ = github.NewClient(nil)

				return errors.New("not implemented yet")
			})
		},
	}
	return cmd
}

func gitClone(url string, dir string) error {
	cmd := exec.Command("git", "clone",
		"--depth", "1",
		"--", url, dir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
