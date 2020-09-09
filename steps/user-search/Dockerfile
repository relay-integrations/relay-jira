FROM relaysh/core:latest-python
RUN pip install datadog
COPY "./datadog-event-post.py" "/entrypoint.py"
ENTRYPOINT []
CMD ["python3", "/entrypoint.py"]

LABEL "org.opencontainers.image.title"="Datadog Event Creation"
LABEL "org.opencontainers.image.description"="This step posts an event to the Datadog event stream"
LABEL "com.puppet.nebula.sdk.version"="v1"
