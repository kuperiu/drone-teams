# drone-teams

Drone plugin for sending MS teams notifications.

## Drone Configuration
- Create a MS Teams incoming webhook
- Add the webhook to drone secrets
- Add the following snippet at the end of your pipeline (Don't foregt to change the secret)
```yaml
  notify:
    image: kuperiu/drone-teams
    webhook: 
      from_secret: teams_webhook
    when: { status: [ success, failure ] }
```

## Available Drone Environment Variables
| Environment Variables  | 
| ------------- |
| TEAMS_WEBHOOK |
| DRONE_REPO_OWNER |
| DRONE_REPO_NAME |
| DRONE_COMMIT_SHA |
| DRONE_COMMIT_BRANCH |
| DRONE_COMMIT_AUTHOR |
| DRONE_COMMIT_AUTHOR_EMAIL |
| DRONE_COMMIT_AUTHOR_AVATAR |
| DRONE_COMMIT_AUTHOR_NAME |
| DRONE_BUILD_NUMBER |
| DRONE_BUILD_STATUS |
| DRONE_BUILD_LINK |
| DRONE_TAG |

## MS Teams
You can create new cards using the [link](https://docs.microsoft.com/en-us/microsoftteams/platform/concepts/cards/cards-reference#office-365-connector-card)

