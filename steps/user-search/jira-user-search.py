#!/usr/bin/env python
# extracts an incident's id, given its friendly name

import requests, os, base64
from relay_sdk import Interface, Dynamic as D

relay = Interface()

# build up the request from workflow parameters
url = relay.get(D.connection.url) + '/rest/api/3/user/search'

headers = {
  'Accept': 'application/json'
}

userEmail = relay.get(D.userEmail)
params = {'query': userEmail }

r = requests.get(url, headers=headers, params=params,
  auth=(relay.get(D.connection.username),relay.get(D.connection.password)) )

r.raise_for_status()

print('Sent query to JIRA, got response: ', r.text, "\n\n")


response = r.json()

user_id = 'not_found'

# Caveat: the response will be an array of user objects, but
# if there's >1 we have no idea which is correct. 
userID = response[0]['accountId']

print('Matched ', userEmail, ' to ', userID)

if userID == 'not_found':
  print("Reached end of user list with no match")
  exit(1)
else:
  relay.outputs.set("userID",userID)
