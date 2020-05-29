package issue

import (
	"errors"
	"io/ioutil"

	"github.com/andygrunwald/go-jira"
	"github.com/mitchellh/mapstructure"
	"github.com/trivago/tgo/tcontainer"
)

type ConnectionSpec struct {
	URL      string `spec:"url"`
	Username string `spec:"username"`
	Password string `spec:"password"`
}

type ProjectSpec struct {
	Key string `spec:"key"`
}

type IssueTypeSpec struct {
	Name string `spec:"name"`
}

type IssueFieldsSpec struct {
	Summary     string         `spec:"summary"`
	Description string         `spec:"description"`
	Type        *IssueTypeSpec `spec:"type"`
	Project     *ProjectSpec   `spec:"project"`
}

type IssueSpec struct {
	Fields       *IssueFieldsSpec  `spec:"fields"`
	CustomFields map[string]string `spec:"customFields"`
}

type Spec struct {
	Connection *ConnectionSpec `spec:"connection"`
	Issue      *IssueSpec      `spec:"issue"`
}

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
		for name, value := range spec.Issue.CustomFields {
			cf[name] = value
		}
	}

	f := &jira.IssueFields{}
	mapstructure.Decode(spec.Issue.Fields, f)

	f.Unknowns = cf
	i := &jira.Issue{
		Fields: f,
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
