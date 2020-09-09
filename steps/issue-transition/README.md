# jira-step-issue-transition

The Jira issue transition step container can update (transition) a jira issue to a new state.


| Setting      | Child setting | Child setting | Child setting | Data type        | Description                                                      | Default | Required |
|--------------|---------------|---------------|---------------|------------------|------------------------------------------------------------------|---------|----------|
| `connection` |               |               |               | Relay Connection | Connection to Jira requiring URL and authentication              | None    | True     |
| `issue`      |               |               |               | mapping          | A mapping of the issue values                                    | None    | True     |
|              | `key`         |               |               | string           | The issue key, such as `COVID-19`                                | None    | True     |
|              | `fields`      |               |               | mapping          | A mapping containing the fields for the transition               | None    | True     |
|              |               | `status`      |               | mapping          | A mapping containing the name of the new issue status            | None    | True     |
|              |               |               | `name`        | string           | Status name, such as `In Progress` or `Closed`                   | None    | True     |
|              |               | `resolution`  |               | mapping          | A mapping containing the name of the new issue status resolution | None    | False    |
|              |               |               | `name`        | string           | Resolution name, such as `Fixed` or `Won't Fix`                  | None    | False    |

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
