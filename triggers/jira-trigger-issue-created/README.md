# jira-trigger-issue-created

This [Atlassian Jira](https://www.atlassian.com/software/jira) trigger fires when a new issue is created. 

## Setup Instructions 

**NOTE: Configuring webhooks requires Jira administrator privileges**

### Step 1: Create new webhook
- Navigate to your System Settings in your Jira console 
- Select "Webhooks" from the Advanced Section. 
- Click "Create a Webhook" 

![Setting up new webhook in Jira](../../media/jira-webhook.gif)

### Step 2: Configure the webhook 
- Name your trigger (e.g. "relay")
- Copy the Webhook URL from your Relay workflow and paste it in the URL field.
- Toggle the box for "Issue: Created" to configure the webhook to trigger when issues are created.

![Configuring new webhook in Jira](../../media/configure-trigger.gif)