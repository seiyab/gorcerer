package job

import (
	"errors"
	"strings"
)

type Job func() (string, error)

func Parse(instruction string) (Job, error) {
	segments := strings.Split(instruction, " ")
	if len(segments) == 0 {
		return nil, errors.New("empty instruction")
	}
	command := segments[0]
	switch command {
	case "ping":
		if len(segments) != 1 {
			return nil, errors.New("ping does not take any argument")
		}
		return func() (string, error) {
			return "pong", nil
		}, nil
	case "inspect":
		if len(segments) != 2 {
			return nil, errors.New("inspect requires 1 argument")
		}
		return func() (string, error) {
			return inspect(segments[1])
		}, nil
	default:
		return nil, errors.New("unknown command")
	}
}
