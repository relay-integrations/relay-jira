from nebula_sdk import Interface, WebhookServer
from quart import Quart, request

import logging
import json

relay = Interface()
app = Quart('issue-deleted')

logging.getLogger().setLevel(logging.INFO)

@app.route('/', methods=['POST'])
async def handler():
    if request.headers.get('x-atlassian-webhook-identifier') is None:
        return {'message': 'not a valid Jira event'}, 400, {}

    data = await request.get_json()
    logging.info("Received the following webhook payload: \n%s", json.dumps(data, indent=4))

    if data['webhookEvent'] != 'jira:issue_deleted':
        return {'message': 'not a valid Jira issue deleted trigger'}, 400, {}
    else:
        issue = data['issue']
        relay.events.emit({
            'issue': issue, 
            'key': issue['key'],
            'type': issue['fields']['issuetype']['name'],
            'project': issue['fields']['project']['name'],
            'priority': issue['fields']['priority']['name'],
            'summary': issue['fields']['summary'],
            'description': issue['fields']['description'],
            'reporter': issue['fields']['reporter']['displayName']
        })
        return {'success': True}

if __name__ == '__main__':
    WebhookServer(app).serve_forever()

