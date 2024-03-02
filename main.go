package main

import (
	"os"

	"github.com/seiyab/gorcerer/cmd"
	"go.uber.org/zap"
)

func main() {
	setupLogger()
	defer zap.L().Sync()

	cmd.Execute()
}

func setupLogger() {
	if os.Getenv("APP_ENV") == "development" {
		zap.ReplaceGlobals(zap.Must(zap.NewDevelopment()))
	}
}
