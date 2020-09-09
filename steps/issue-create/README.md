# jira-step-issue-create

The Jira issue create step container creates a jira ticket.

| Setting      | Child setting  | Child setting | Child setting | Data type        | Description                                              | Default | Required |
|--------------|----------------|---------------|---------------|------------------|----------------------------------------------------------|---------|----------|
| `connection` |                |               |               | Relay Connection | Connection to Jira requiring URL and authentication      | None    | True     |
| `issue`      |                |               |               | mapping          | A mapping of the issue values                            | None    | True     |
|              | `fields`       |               |               | mapping          | A mapping of the issue fields                            | None    | True     |
|              |                | `summary`     |               | string           | A summary of the issue                                   | None    | True     |
|              |                | `description` |               | string           | A description of the issue                               | None    | False    |
|              |                | `type`        |               | mapping          | A mapping containing the name of the issue type          | None    | True     |
|              |                |               | `name`        | string           | The name of an issue type, such as `Story` or `Bug`      | None    | True     |
|              |                | `assignee`    |               | mapping          | A mapping containing the name of the assignee            | None    | True     |
|              |                |               | `name`        | string           | The name of an assignee within jira                      | None    | False    |
|              |                | `project`     |               | mapping          | A mapping containing the issue project key               | None    | True     |
|              |                |               | `key`         | string           | The project key for the new issue, such as `OPS`         | None    | True     |
|              | `customFields` |               |               | mapping          | An arbitrary mapping of any custom field keys and values | None    | False    |

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
