package issue

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/andygrunwald/go-jira"
)

var (
	ErrInvalidAuth                 = errors.New("Invalid authentication")
	ErrNoIssueFieldsAreDefined     = errors.New("No issue fields are defined")
	ErrNoIssueKeyIsDefined         = errors.New("No issue key is defined")
	ErrNoIssueStatusFieldIsDefined = errors.New("No issue status field is defined")
)

type ConnectionSpec struct {
	URL      string `spec:"url"`
	Username string `spec:"username"`
	Password string `spec:"password"`
}

type IssueResolutionSpec struct {
	Name string `spec:"name`
}

type IssueStatusSpec struct {
	Name string `spec:"name`
}

type IssueFieldsSpec struct {
	Status     *IssueStatusSpec     `spec:"status"`
	Resolution *IssueResolutionSpec `spec:"resolution"`
}

type IssueSpec struct {
	Key    string           `spec:"key"`
	Fields *IssueFieldsSpec `spec:"fields"`
}

type Spec struct {
	Connection *ConnectionSpec `spec:"connection"`
	Issue      *IssueSpec      `spec:"issue"`
}

func TransitionIssue(spec Spec) error {
	if spec.Issue.Key == "" {
		return ErrNoIssueKeyIsDefined
	}
	if spec.Issue.Fields == nil {
		return ErrNoIssueFieldsAreDefined
	}
	if spec.Issue.Fields.Status == nil {
		return ErrNoIssueStatusFieldIsDefined
	}

	tp := jira.BasicAuthTransport{
		Username: spec.Connection.Username,
		Password: spec.Connection.Password,
	}

	jiraClient, err := jira.NewClient(tp.Client(), spec.Connection.URL)
	if err != nil {
		return err
	}

	transitions, response, err := jiraClient.Issue.GetTransitions(spec.Issue.Key)
	if err != nil {
		return handleResponseError(response.Response, err)
	}

	for _, transition := range transitions {
		if transition.Name == spec.Issue.Fields.Status.Name {
			payload := jira.CreateTransitionPayload{
				Transition: jira.TransitionPayload{
					ID: transition.ID,
				},
			}

			if spec.Issue.Fields.Resolution != nil {
				payload.Fields = jira.TransitionPayloadFields{
					Resolution: &jira.Resolution{
						Name: spec.Issue.Fields.Resolution.Name,
					},
				}
			}

			response, err := jiraClient.Issue.DoTransitionWithPayload(spec.Issue.Key, payload)
			if err != nil {
				return handleResponseError(response.Response, err)
			}

			return nil
		}
	}

	return fmt.Errorf("transition %s is not applicable for issue %s", spec.Issue.Fields.Status.Name, spec.Issue.Key)
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
