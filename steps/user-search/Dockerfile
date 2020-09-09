FROM relaysh/core:latest-python
COPY "./jira-user-search.py" "/entrypoint.py"
ENTRYPOINT []
CMD ["python3", "/entrypoint.py"]

LABEL "org.opencontainers.image.title"="Jira User Search"
LABEL "org.opencontainers.image.description"="This step returns the internal user id for a JIRA user, given an email"
LABEL "com.puppet.nebula.sdk.version"="v1"
