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
		Name          string
		File          string
		ExpectedError error
		Reopen        bool
	}{
		{
			Name:          "Jira Server Issue Transition: In Progress",
			File:          "fixtures/jira-server-issue-transition-in-progress.yaml",
			ExpectedError: nil,
			Reopen:        true,
		},
		{
			Name:          "Jira Server Issue Transition: Close, Won't Do",
			File:          "fixtures/jira-server-issue-transition-close-won't-do.yaml",
			ExpectedError: nil,
			Reopen:        true,
		},
		{
			Name:          "Jira Server Issue Transition: Invalid Status",
			File:          "fixtures/jira-server-issue-transition-invalid-status.yaml",
			ExpectedError: fmt.Errorf("transition %s is not applicable for issue %s", "Close", "RELAY-45"),
			Reopen:        true,
		},
		{
			Name:          "Jira Server Issue Transition: Fields Undefined",
			File:          "fixtures/jira-server-issue-transition-fields-undefined.yaml",
			ExpectedError: issue.ErrNoIssueFieldsAreDefined,
		},
		{
			Name:          "Jira Server Issue Transition: Issue Key Undefined",
			File:          "fixtures/jira-server-issue-transition-issue-key-undefined.yaml",
			ExpectedError: issue.ErrNoIssueKeyIsDefined,
		},
		{
			Name:          "Jira Server Issue Transition: Issue Status Field Undefined",
			File:          "fixtures/jira-server-issue-transition-issue-status-field-undefined.yaml",
			ExpectedError: issue.ErrNoIssueStatusFieldIsDefined,
		},
		{
			Name:          "Jira Server Invalid Auth",
			File:          "fixtures/jira-server-invalid-auth.yaml",
			ExpectedError: issue.ErrInvalidAuth,
		},
	}
	for _, test := range tcs {
		t.Run(fmt.Sprintf("%s", test.Name), func(t *testing.T) {
			content := getTestFixture(test.File)

			var spec issue.Spec
			if err := yaml.Unmarshal(content, &spec); err != nil {
				panic(err)
			}

			if test.Reopen {
				reopen(spec)
			}

			err := issue.TransitionIssue(spec)
			require.Equal(t, test.ExpectedError, err)
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

func reopen(spec issue.Spec) {
	reopen := issue.Spec{
		Connection: spec.Connection,
		Issue: &issue.IssueSpec{
			Key: spec.Issue.Key,
			Fields: &issue.IssueFieldsSpec{
				Status: &issue.IssueStatusSpec{
					Name: "Reopen Issue",
				},
			},
		},
	}

	issue.TransitionIssue(reopen)
}
