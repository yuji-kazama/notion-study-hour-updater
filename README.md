# Notion Study Hour Updater

## Requirement
* Go 1.16
* direnv

## Config
Create ".env" file, and set the following as env valuables.
* NOTION_INTEGRATION_TOKEN
* NOTION_DATABASE_ID
* INDIFY_URL

## Deploy in Cloud Functions
gcloud functions deploy updateStudyHour --entry-point Function --trigger-http --runtime go116 --region asia-northeast1 --memory=2048 --set-env-vars NOTION_INTEGRATION_TOKEN=[ANY_VALUE],NOTION_DATABASE_ID=[ANY_VALUE],INDIFY_URL=[ANY_VALUE]


