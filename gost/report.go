package gost

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Report map[string]AnalyzerReport

type AnalyzerReport map[string][]Diagnostic

type Diagnostic struct {
	File    string
	Line    int64
	Column  int64
	Message string
}

func (d *Diagnostic) UnmarshalJSON(b []byte) error {
	var raw struct {
		Posn    string `json:"posn"`
		Message string `json:"message"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	sections := strings.Split(raw.Posn, ":")
	if len(sections) != 3 {
		return fmt.Errorf("unexpected posn: %q", raw.Posn)
	}
	d.File = sections[0]
	line, err := strconv.Atoi(sections[1])
	if err != nil {
		return err
	}
	d.Line = int64(line)
	column, err := strconv.Atoi(sections[2])
	if err != nil {
		return err
	}
	d.Column = int64(column)
	d.Message = raw.Message
	return nil
}
