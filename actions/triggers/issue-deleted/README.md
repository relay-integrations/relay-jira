# jira-trigger-issue-deleted

This [Atlassian Jira](https://www.atlassian.com/software/jira) trigger fires when a new issue is deleted. 

## Data Emitted 

| Name | Data type | Description | 
|------|-----------|-------------|
| `key` | string | issue key | 
| `type` | string | issue type | 
| `project` | string | issue project | 
| `priority` | string | issue priority | 
| `summary` | string | issue summary |
| `description` | string | description of the issue | 
| `reporter` | string | reporter of the issue |
| `issue` | mapping | raw issue output | 

## Example Trigger

```
parameters:
  key:
    default: ""
  type:
    default: ""
  project:
    default: ""
  priority:
    default: ""  
  summary:
    default: ""
  description:
    default: ""
triggers:
- name: jira-issue-deleted
  source:
    type: webhook
    image: relaysh/jira-trigger-issue-deleted:latest
  binding:
    parameters:
      key: !Data key
      type: !Data type
      project: !Data project
      priority: !Data priority
      summary: !Data summary
      description: !Data description
```

## Example Raw Data 

```
{
  "timestamp": 1591420829690,
  "webhookEvent": "jira:issue_deleted",
  "user": {
    "self": "https://relaysh.atlassian.net/rest/api/2/user?accountId=557058%3A182fed4d-0782-4cf4-a2cd-3083b91d0452",
    "accountId": "557058:182fed4d-0782-4cf4-a2cd-3083b91d0452",
    "avatarUrls": {
      "48x48": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
      "24x24": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
      "16x16": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
      "32x32": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png"
    },
    "displayName": "Kenaz Kwa",
    "active": true,
    "timeZone": "Etc/GMT",
    "accountType": "atlassian"
  },
  "issue": {
    "id": "10000",
    "self": "https://relaysh.atlassian.net/rest/api/2/10000",
    "key": "REL-1",
    "fields": {
      "statuscategorychangedate": "2020-06-05T16:09:54.262+0000",
      "issuetype": {
        "self": "https://relaysh.atlassian.net/rest/api/2/issuetype/10001",
        "id": "10001",
        "description": "Functionality or a feature expressed as a user goal.",
        "iconUrl": "https://relaysh.atlassian.net/secure/viewavatar?size=medium&avatarId=10315&avatarType=issuetype",
        "name": "Story",
        "subtask": false,
        "avatarId": 10315
      },
      "timespent": null,
      "project": {
        "self": "https://relaysh.atlassian.net/rest/api/2/project/10000",
        "id": "10000",
        "key": "REL",
        "name": "RELAY",
        "projectTypeKey": "software",
        "simplified": false,
        "avatarUrls": {
          "48x48": "https://relaysh.atlassian.net/secure/projectavatar?pid=10000&avatarId=10408",
          "24x24": "https://relaysh.atlassian.net/secure/projectavatar?size=small&s=small&pid=10000&avatarId=10408",
          "16x16": "https://relaysh.atlassian.net/secure/projectavatar?size=xsmall&s=xsmall&pid=10000&avatarId=10408",
          "32x32": "https://relaysh.atlassian.net/secure/projectavatar?size=medium&s=medium&pid=10000&avatarId=10408"
        }
      },
      "fixVersions": [],
      "aggregatetimespent": null,
      "resolution": null,
      "customfield_10028": null,
      "resolutiondate": null,
      "workratio": -1,
      "watches": {
        "self": "https://relaysh.atlassian.net/rest/api/2/issue/REL-1/watchers",
        "watchCount": 1,
        "isWatching": true
      },
      "lastViewed": "2020-06-06T05:20:28.182+0000",
      "created": "2020-06-05T16:09:54.158+0000",
      "customfield_10020": null,
      "customfield_10021": null,
      "customfield_10022": null,
      "customfield_10023": null,
      "priority": {
        "self": "https://relaysh.atlassian.net/rest/api/2/priority/3",
        "iconUrl": "https://relaysh.atlassian.net/images/icons/priorities/medium.svg",
        "name": "Medium",
        "id": "3"
      },
      "customfield_10024": null,
      "customfield_10025": null,
      "labels": [],
      "customfield_10016": null,
      "customfield_10017": null,
      "customfield_10018": {
        "hasEpicLinkFieldDependency": false,
        "showField": false,
        "nonEditableReason": {
          "reason": "PLUGIN_LICENSE_ERROR",
          "message": "The Parent Link is only available to Jira Premium users."
        }
      },
      "customfield_10019": "0|hzzzzz:",
      "timeestimate": null,
      "aggregatetimeoriginalestimate": null,
      "versions": [],
      "issuelinks": [],
      "assignee": null,
      "updated": "2020-06-05T16:09:54.158+0000",
      "status": {
        "self": "https://relaysh.atlassian.net/rest/api/2/status/10000",
        "description": "",
        "iconUrl": "https://relaysh.atlassian.net/",
        "name": "Backlog",
        "id": "10000",
        "statusCategory": {
          "self": "https://relaysh.atlassian.net/rest/api/2/statuscategory/2",
          "id": 2,
          "key": "new",
          "colorName": "blue-gray",
          "name": "To Do"
        }
      },
      "components": [],
      "timeoriginalestimate": null,
      "description": null,
      "customfield_10010": null,
      "customfield_10014": null,
      "customfield_10015": null,
      "timetracking": {},
      "customfield_10005": null,
      "customfield_10006": null,
      "customfield_10007": null,
      "security": null,
      "customfield_10008": null,
      "attachment": [],
      "aggregatetimeestimate": null,
      "customfield_10009": null,
      "summary": "Test Issue #1",
      "creator": {
        "self": "https://relaysh.atlassian.net/rest/api/2/user?accountId=557058%3A182fed4d-0782-4cf4-a2cd-3083b91d0452",
        "accountId": "557058:182fed4d-0782-4cf4-a2cd-3083b91d0452",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
          "24x24": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
          "16x16": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
          "32x32": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png"
        },
        "displayName": "Kenaz Kwa",
        "active": true,
        "timeZone": "Etc/GMT",
        "accountType": "atlassian"
      },
      "subtasks": [],
      "reporter": {
        "self": "https://relaysh.atlassian.net/rest/api/2/user?accountId=557058%3A182fed4d-0782-4cf4-a2cd-3083b91d0452",
        "accountId": "557058:182fed4d-0782-4cf4-a2cd-3083b91d0452",
        "avatarUrls": {
          "48x48": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
          "24x24": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
          "16x16": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png",
          "32x32": "https://secure.gravatar.com/avatar/8dab52cfb60225bd67ca0bb5b821e460?d=https%3A%2F%2Favatar-management--avatars.us-west-2.prod.public.atl-paas.net%2Finitials%2FKK-1.png"
        },
        "displayName": "Kenaz Kwa",
        "active": true,
        "timeZone": "Etc/GMT",
        "accountType": "atlassian"
      },
      "customfield_10000": "{}",
      "aggregateprogress": {
        "progress": 0,
        "total": 0
      },
      "customfield_10001": null,
      "customfield_10002": null,
      "customfield_10003": null,
      "customfield_10004": null,
      "environment": null,
      "duedate": null,
      "progress": {
        "progress": 0,
        "total": 0
      },
      "votes": {
        "self": "https://relaysh.atlassian.net/rest/api/2/issue/REL-1/votes",
        "votes": 0,
        "hasVoted": false
      },
      "comment": {
        "comments": [],
        "maxResults": 0,
        "total": 0,
        "startAt": 0
      },
      "worklog": {
        "startAt": 0,
        "maxResults": 20,
        "total": 0,
        "worklogs": []
      }
    }
  }
}
```