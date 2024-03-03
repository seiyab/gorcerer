package output

import (
	"bytes"
	"context"
	"io"

	"github.com/google/go-github/v60/github"
	"github.com/seiyab/gorcerer/utils"
)

type IssueComment struct {
	gh    *github.Client
	issue *github.Issue
	buf   bytes.Buffer
}

var _ io.WriteCloser = &IssueComment{}

func NewIssueComment(client *github.Client, issue *github.Issue) IssueComment {
	return IssueComment{
		gh:    client,
		issue: issue,
		buf:   bytes.Buffer{},
	}
}

func (ic *IssueComment) Write(p []byte) (n int, err error) {
	return ic.buf.Write(p)
}

func (ic *IssueComment) Close() error {
	repo := ic.issue.Repository
	_, _, err := ic.gh.Issues.CreateComment(context.Background(),
		repo.Owner.GetName(), repo.GetName(), *ic.issue.Number,
		&github.IssueComment{Body: utils.Ref(ic.buf.String())},
	)
	return err
}
