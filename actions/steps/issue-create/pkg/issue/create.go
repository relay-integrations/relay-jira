package issue

import (
	"errors"
	"io/ioutil"

	"github.com/andygrunwald/go-jira"
	"github.com/trivago/tgo/tcontainer"
)

// ConnectionSpec represents the connection/authentication data
type ConnectionSpec struct {
	URL      string `spec:"url"`
	Username string `spec:"username"`
	Password string `spec:"password"`
}

// IssueSpec represents the issue data
type IssueSpec struct {
	Fields       *jira.IssueFields
	CustomFields *jira.CustomFields
}

// Spec represents the encompassing specification structure
type Spec struct {
	Connection *ConnectionSpec
	Issue      *IssueSpec
}

// CreateIssue creates a new Jira issue
func CreateIssue(spec Spec) (*jira.Issue, error) {
	tp := jira.BasicAuthTransport{
		Username: spec.Connection.Username,
		Password: spec.Connection.Password,
	}

	jiraClient, err := jira.NewClient(tp.Client(), spec.Connection.URL)
	if err != nil {
		return nil, err
	}

	cf := tcontainer.NewMarshalMap()
	if spec.Issue.CustomFields != nil {
		for name, value := range *spec.Issue.CustomFields {
			cf[name] = value
		}
	}

	spec.Issue.Fields.Unknowns = cf

	i := &jira.Issue{
		Fields: spec.Issue.Fields,
	}

	issue, response, err := jiraClient.Issue.Create(i)
	if err != nil {
		if response != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return nil, err
			}
			return nil, errors.New(string(body))
		}
		return nil, err
	}

	return issue, nil
}
