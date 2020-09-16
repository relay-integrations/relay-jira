# jira-step-issue-create

The Jira issue create step container creates a Jira ticket.

## Examples

```YAML
steps:

# ...
- name: jira-issue-create
  image: relaysh/jira-step-issue-create
  spec:
    connection: !Connection { type: jira, name: jira-bot }
    issue:
      fields:
        project:
          key: !Parameter jiraProjectKey
        type:
          name: Task
        summary: !Parameter eventTitle
        description: !Fn.convertMarkdown [jira, !Parameter eventBody]
```
