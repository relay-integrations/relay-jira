package issue_test

import (
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/relay-integrations/relay-jira-server/actions/steps/issue-create/pkg/issue"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestCreateIssueFromSpec(t *testing.T) {
	t.SkipNow()

	tt := []struct {
		File string
	}{
		{
			File: "fixtures/jira-server-issue-create-bug.yaml",
		},
		{
			File: "fixtures/jira-server-issue-create-epic.yaml",
		},
	}
	for _, test := range tt {
		content := getTestFixture(test.File)

		var spec issue.Spec
		if err := yaml.Unmarshal(content, &spec); err != nil {
			panic(err)
		}

		issue, err := issue.CreateIssue(spec)
		require.NoError(t, err)
		require.NotNil(t, issue)
		require.NotEmpty(t, issue.Key)

		if issue != nil {
			t.Log(fmt.Sprintf("Created issue %v", issue.Key))
		}
	}
}

func getTestFixture(p string) []byte {
	_, file, _, _ := runtime.Caller(0)
	base := filepath.Dir(file)
	content, err := ioutil.ReadFile(path.Join(base, p))
	if err != nil {
		panic(err)
	}

	return content
}
