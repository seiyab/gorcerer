package gost_test

import (
	"encoding/json"
	"testing"

	"github.com/seiyab/gorcerer/gost"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalReport(t *testing.T) {
	d := `
{
	"go.uber.org/zap/zapcore [go.uber.org/zap/zapcore.test]": {
		"handleerror": [
			{
				"posn": "/path/to/repo/zapcore/json_encoder_impl_test.go:618:2",
				"message": "unhandled error"
			}
		]
	}
}
	`
	var actual gost.Report
	if err := json.Unmarshal([]byte(d), &actual); err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, gost.Report{
		"go.uber.org/zap/zapcore [go.uber.org/zap/zapcore.test]": {
			"handleerror": []gost.Diagnostic{
				{
					File:    "/path/to/repo/zapcore/json_encoder_impl_test.go",
					Line:    618,
					Column:  2,
					Message: "unhandled error",
				},
			},
		},
	}, actual)
}
