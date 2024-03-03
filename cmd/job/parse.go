package job

import (
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Job func() (string, error)

func Parse(instruction string) (Job, error) {
	zap.L().Debug("parsing instruction", zap.String("instruction", instruction))
	segments := strings.Split(strings.Trim(instruction, " \t\n"), " ")
	if len(segments) == 0 {
		return nil, errors.New("empty instruction")
	}
	command := segments[0]
	zap.L().Debug("parsed command", zap.String("command", command))
	switch command {
	case "ping":
		if len(segments) != 1 {
			return nil, errors.Errorf(
				"ping does not take any argument but got %d arguments",
				len(segments),
			)
		}
		return func() (string, error) {
			return "pong", nil
		}, nil
	case "inspect":
		if len(segments) != 2 {
			return nil, errors.Errorf(
				"inspect requires 1 argument but got %d arguments",
				len(segments),
			)
		}
		return func() (string, error) {
			return inspect(segments[1])
		}, nil
	default:
		return nil, errors.Errorf("unknown command %q", command)
	}
}
