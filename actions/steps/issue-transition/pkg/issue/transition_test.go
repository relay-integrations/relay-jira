package issue_test

import (
	"io/ioutil"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/relay-integrations/relay-jira-server/actions/steps/issue-transition/pkg/issue"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestTransitionIssueFromSpec(t *testing.T) {
	t.SkipNow()

	tt := []struct {
		File string
	}{
		{
			File: "fixtures/jira-server-issue-transition-close-won't-do.yaml",
		},
	}
	for _, test := range tt {
		content := getTestFixture(test.File)

		var spec issue.Spec
		if err := yaml.Unmarshal(content, &spec); err != nil {
			panic(err)
		}

		err := issue.TransitionIssue(spec)
		require.NoError(t, err)
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
