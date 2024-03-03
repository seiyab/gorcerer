package gost

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/seiyab/gorcerer/process"
)

func Run(dir string) (Report, error) {
	var buf bytes.Buffer
	gost := process.Command("gost", "-json", "./...")
	gost.Stdout = &buf
	gost.Stderr = os.Stderr
	gost.Dir = dir
	if err := gost.Run(); err != nil {
		return nil, err
	}
	var report Report
	if err := json.Unmarshal(buf.Bytes(), &report); err != nil {
		return nil, err
	}
	return report, nil
}
