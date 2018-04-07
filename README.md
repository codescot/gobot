# slackbot
Slack Bot written in Golang

# docker set up
Pull down the latest version using docker

> docker pull gurparit/slackbot`

Run docker image with

> docker run -d -v /full/path/to/slackbot.conf:/go/src/app/slackbot.conf --restart=always --name slackbot gurparit/slackbot`

You may need to run the docker commands with sudo.
