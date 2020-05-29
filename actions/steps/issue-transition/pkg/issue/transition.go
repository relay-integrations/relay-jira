package issue

import (
	"errors"
	"io/ioutil"

	"github.com/andygrunwald/go-jira"
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
		if response != nil {
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				return err
			}
			return errors.New(string(body))
		}
		return err
	}

	for _, transition := range transitions {
		if transition.Name == spec.Issue.Fields.Status.Name {
			payload := jira.CreateTransitionPayload{
				Transition: jira.TransitionPayload{
					ID: transition.ID,
				},
				Fields: jira.TransitionPayloadFields{
					Resolution: &jira.Resolution{
						Name: spec.Issue.Fields.Resolution.Name,
					},
				},
			}

			response, err := jiraClient.Issue.DoTransitionWithPayload(spec.Issue.Key, payload)
			if err != nil {
				if response != nil {
					body, err := ioutil.ReadAll(response.Body)
					if err != nil {
						return err
					}
					return errors.New(string(body))
				}
				return err
			}

			break
		}
	}

	return nil
}
