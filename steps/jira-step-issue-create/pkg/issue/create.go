package issue

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/andygrunwald/go-jira"
	"github.com/mitchellh/mapstructure"
	"github.com/trivago/tgo/tcontainer"
)

var (
	ErrInvalidAuth             = errors.New("Invalid authentication")
	ErrNoIssueFieldsAreDefined = errors.New("No issue fields are defined")
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

type AssigneeSpec struct {
	Name string `spec:"name"`
}

type IssueFieldsSpec struct {
	Summary     string         `spec:"summary"`
	Description string         `spec:"description"`
	Type        *IssueTypeSpec `spec:"type"`
	Project     *ProjectSpec   `spec:"project"`
	Assignee    *AssigneeSpec  `spec:"assignee"`
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
	if spec.Issue == nil || spec.Issue.Fields == nil {
		return nil, ErrNoIssueFieldsAreDefined
	}

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
		return nil, handleResponseError(response.Response, err)
	}

	return issue, nil
}

func handleResponseError(response *http.Response, err error) error {
	if err != nil {
		if response.StatusCode == http.StatusUnauthorized {
			return ErrInvalidAuth
		}
		if response != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			return errors.New(string(body))
		}
		return err
	}

	return nil
}
