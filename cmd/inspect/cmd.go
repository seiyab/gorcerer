package inspect

import (
	"errors"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/google/go-github/v60/github"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "inspect",
		RunE: func(cmd *cobra.Command, args []string) error {
			tempDir, err := os.MkdirTemp(os.TempDir(), "gorcerer-inspect")
			if err != nil {
				return err
			}
			defer os.RemoveAll(tempDir)
			log.Default().Printf("tempDir: %s", tempDir)

			if err := gitClone(
				"git@github.com:seiyab/gorcerer.git",
				path.Join(tempDir, "gorcerer"),
			); err != nil {
				return err
			}

			_ = github.NewClient(nil)

			return errors.New("not implemented yet")
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
