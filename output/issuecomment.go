package output

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
)

type Output func(a ...any) error

func Println(a ...any) error {
	_, err := fmt.Println(a...)
	return err
}

var _ Output = Println

func NewIssueComment(client *github.Client, issue int) Output {
	return func(a ...any) error {
		s := fmt.Sprintln(a...)
		_, _, err := client.Issues.CreateComment(
			context.Background(),
			"seiyab", "gorcerer",
			issue,
			&github.IssueComment{Body: &s},
		)
		if err != nil {
			return err
		}
		return nil
	}
}
