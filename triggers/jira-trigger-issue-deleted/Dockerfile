FROM relaysh/core:latest-python
WORKDIR /trigger
COPY handler.py .
CMD python handler.py

LABEL "org.opencontainers.image.title"="Jira issue deleted"
LABEL "org.opencontainers.image.description"="This trigger fires when a Jira issue is deleted."
LABEL "com.puppet.nebula.sdk.version"="v1"
