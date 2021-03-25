FROM relaysh/core:latest-python
WORKDIR /trigger
COPY handler.py .
CMD python handler.py

LABEL "org.opencontainers.image.title"="Jira issue created"
LABEL "org.opencontainers.image.description"="This trigger fires when a Jira issue is created."
LABEL "com.puppet.nebula.sdk.version"="v1"
