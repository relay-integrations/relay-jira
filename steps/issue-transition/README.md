# jira-step-issue-transition

The Jira issue transition step container can update (transition) a jira issue to a new state.

## Examples

```YAML
steps:

# ...
- name: jira-issue-resolve
  image: relaysh/jira-step-issue-transition
  spec:
    connection: !Connection { type: jira, name: jira-bot }
    issue:
      key: !Parameter jiraIssueKey
      fields:
        status:
          name: Resolved
        resolution:
          name: Fixed
```
