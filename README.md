# gobot
Slack Bot written in Golang (soon to be irc again)

# docker build image

> docker build -t repo/gobot .

# docker set up
Pull down the latest version using docker

> docker pull gurparit/gobot

Run docker image with

> docker run -d -env-file /path/to/env.list --restart=always --name gobot gurparit/gobot

You may need to run the docker commands with sudo.
