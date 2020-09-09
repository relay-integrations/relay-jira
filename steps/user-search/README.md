# jira-step-user-search

Search for a users ID given their email address

This step translates a human user id to the corresponding internal ID, which can be used by subsequent steps like assigning issues

It requires a `jira` connection type. Workflows should provide the following parameters in the spec:

* `user_email`: (string) the user's email address to search for


Example usage:

```yaml
parameters:
  user_email:
    description: the email address of the user to search for

steps:
- name: search-user-id
  image: relaysh/jira-step-user-search
  spec:
    user_email: !Parameter user_email
    connection: !Connection { type: jira, name: my-jira-login }
- name: echo-output
  image: relaysh/core
  spec:
    user_email: !Parameter user_email
    user_id: !Output [search-user-id, user_id]
  input:
    - echo "Looked up $(ni get -p {.user_email}) as id $(ni get -p {.user_id})"
```
