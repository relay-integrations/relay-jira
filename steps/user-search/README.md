# jira-step-user-search

Search for a users ID given their email address

This step translates a human user id to the corresponding internal ID, which can be used by subsequent steps like assigning issues

It requires a `jira` connection type. Workflows should provide the following parameters in the spec:

* `userEmail`: (string) the user's email address to search for

