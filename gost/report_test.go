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
		{
			name: "multiple",
			input: `{
				"github.com/cli/cli/v2/pkg/cmd/extension": {
					"openfileflag": [
						{
							"posn": "path/to/repo/pkg/cmd/extension/manager.go:379:12",
							"message": "O_TRUNC / O_APPEND / O_EXCL flags are not specified"
						}
					]
				},
				"github.com/cli/cli/v2/pkg/cmd/extension [github.com/cli/cli/v2/pkg/cmd/extension.test]": {
					"openfileflag": [
						{
							"posn": "path/to/repo/pkg/cmd/extension/manager.go:379:12",
							"message": "O_TRUNC / O_APPEND / O_EXCL flags are not specified"
						},
						{
							"posn": "path/to/repo/pkg/cmd/extension/manager_test.go:1161:11",
							"message": "O_TRUNC / O_APPEND / O_EXCL flags are not specified"
						},
						{
							"posn": "path/to/repo/pkg/cmd/extension/manager_test.go:1185:12",
							"message": "O_TRUNC / O_APPEND / O_EXCL flags are not specified"
						}
					]
				}
			}`,
			expected: gost.Report{
				"github.com/cli/cli/v2/pkg/cmd/extension": {
					"openfileflag": []gost.Diagnostic{
						{
							File:    "path/to/repo/pkg/cmd/extension/manager.go",
							Line:    379,
							Column:  12,
							Message: "O_TRUNC / O_APPEND / O_EXCL flags are not specified",
						},
					},
				},
				"github.com/cli/cli/v2/pkg/cmd/extension [github.com/cli/cli/v2/pkg/cmd/extension.test]": {
					"openfileflag": []gost.Diagnostic{
						{
							File:    "path/to/repo/pkg/cmd/extension/manager.go",
							Line:    379,
							Column:  12,
							Message: "O_TRUNC / O_APPEND / O_EXCL flags are not specified",
						},
						{
							File:    "path/to/repo/pkg/cmd/extension/manager_test.go",
							Line:    1161,
							Column:  11,
							Message: "O_TRUNC / O_APPEND / O_EXCL flags are not specified",
						},
						{
							File:    "path/to/repo/pkg/cmd/extension/manager_test.go",
							Line:    1185,
							Column:  12,
							Message: "O_TRUNC / O_APPEND / O_EXCL flags are not specified",
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
