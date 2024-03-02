package utils

import (
	"os"

	"go.uber.org/zap"
)

func TempDir(pattern string, proc func(path string) error) error {
	tempDir, err := os.MkdirTemp(os.TempDir(), pattern)
	if err != nil {
		return err
	}
	zap.L().Debug("created temp dir", zap.String("dir", tempDir))
	defer os.RemoveAll(tempDir)

	return proc(tempDir)
}
