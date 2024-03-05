package gost_test

import (
	"encoding/json"
	"testing"

	"github.com/seiyab/gorcerer/gost"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalReport(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected gost.Report
	}{
		{
			name: "simple",
			input: `{
				"go.uber.org/zap/zapcore [go.uber.org/zap/zapcore.test]": {
					"handleerror": [
						{
							"posn": "/path/to/repo/zapcore/json_encoder_impl_test.go:618:2",
							"message": "unhandled error"
						}
					]
				}
			}`,
			expected: gost.Report{
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
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var actual gost.Report
			if err := json.Unmarshal([]byte(tc.input), &actual); err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tc.expected, actual)
		})
	}
}
