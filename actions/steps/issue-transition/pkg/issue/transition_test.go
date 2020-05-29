package issue_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/relay-integrations/relay-jira-server/actions/steps/issue-transition/pkg/issue"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestTransitionIssueFromSpec(t *testing.T) {
	tcs := []struct {
		Name string
		File string
	}{
		{
			Name: "Jira Server Issue Transition: Close, Won't Do",
			File: "fixtures/jira-server-issue-transition-close-won't-do.yaml",
		},
	}
	for _, test := range tcs {
		t.Run(fmt.Sprintf("%s", test.Name), func(t *testing.T) {
			content := getTestFixture(test.File)

			var spec issue.Spec
			if err := yaml.Unmarshal(content, &spec); err != nil {
				panic(err)
			}

			err := issue.TransitionIssue(spec)
			require.NoError(t, err)
		})
	}
}

func getTestFixture(p string) []byte {
	_, file, _, _ := runtime.Caller(0)
	base := filepath.Dir(file)
	content, err := ioutil.ReadFile(path.Join(base, p))
	if err != nil {
		panic(err)
	}

	return []byte(os.ExpandEnv(string(content)))
}
