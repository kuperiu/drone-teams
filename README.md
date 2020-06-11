[![image
pulls](https://shields.beevelop.com/docker/pulls/kuperiu/drone-teams.svg)](https://hub.docker.com/r/kuperiu/drone-teams)


# drone-teams

Drone plugin for sending Microsoft Teams notifications.

## Drone Configuration
- Create a MS Teams incoming webhook
- Add the webhook to drone secrets
- Add the following snippet at the end of your pipeline (Don't foregt to change the secret)
```yaml
  notify: 
    image: "kuperiu/drone-teams"
    secrets: 
      - 
        source: "teams_webhook"
        target: "teams_webhook"
    when: 
      status: 
        - "success"
        - "failure"
```

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o drone-teams main.go plugin.go
docker build --rm -t kuperiu/drone-teams .
```

## Usage

Execute from the working directory:

```
docker run --rm \
  -e TEAMS_WEBHOOK=https://... \
  -e DRONE_REPO_OWNER=octocat \
  -e DRONE_REPO_NAME=hello-world \
  -e DRONE_COMMIT_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
  -e DRONE_COMMIT_BRANCH=master \
  -e DRONE_COMMIT_AUTHOR=octocat \
  -e DRONE_COMMIT_AUTHOR_EMAIL=octocat@github.com \
  -e DRONE_COMMIT_AUTHOR_AVATAR="https://avatars0.githubusercontent.com/u/583231?s=460&v=4" \
  -e DRONE_COMMIT_AUTHOR_NAME="The Octocat" \
  -e DRONE_BUILD_NUMBER=1 \
  -e DRONE_BUILD_STATUS=success \
  -e DRONE_BUILD_LINK=http://github.com/octocat/hello-world \
  -e DRONE_TAG=1.0.0 \
  kuperiu/drone-teams
```
