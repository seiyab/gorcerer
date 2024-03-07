package job

import (
	"strings"
	"testing"

	"github.com/seiyab/gorcerer/gost"
	"github.com/stretchr/testify/assert"
)

func TestFormatReport(t *testing.T) {
	report := gost.Report{
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
	}

	actual := formatReport(report)
	expected := strings.Join([]string{
		"path/to/repo/pkg/cmd/extension/manager.go:379:12 O_TRUNC / O_APPEND / O_EXCL flags are not specified",
		"path/to/repo/pkg/cmd/extension/manager_test.go:1161:11 O_TRUNC / O_APPEND / O_EXCL flags are not specified",
		"path/to/repo/pkg/cmd/extension/manager_test.go:1185:12 O_TRUNC / O_APPEND / O_EXCL flags are not specified",
	}, "\n")

	assert.Equal(t, expected, actual)
}
