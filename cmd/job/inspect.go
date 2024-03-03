package job

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/seiyab/gorcerer/gost"
	"github.com/seiyab/gorcerer/process"
	"github.com/seiyab/gorcerer/utils"
)

func inspect(repository string) (string, error) {
	var result gost.Report
	if err := utils.TempDir("gorcerer-inspect", func(tempDir string) error {
		repoDir := path.Join(tempDir, "repo")
		if err := gitClone(
			repository,
			repoDir,
		); err != nil {
			return err
		}

		var err error
		if result, err = gost.Run(repoDir); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return "", err
	}
	return formatReport(result), nil
}

func gitClone(url string, dir string) error {
	cmd := process.Command("git", "clone",
		"--depth", "1",
		"--", url, dir)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func formatReport(report gost.Report) string {
	var lines []string
	for _, as := range report {
		for _, a := range as {
			for _, d := range a {
				lines = append(lines,
					fmt.Sprintf(
						"%s:%d:%d %s",
						d.File,
						d.Line,
						d.Column,
						d.Message,
					),
				)
			}
		}
	}
	return strings.Join(lines, "\n")
}
