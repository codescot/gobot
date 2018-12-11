# gobot
Twitch/IRC Bot written in Golang

# docker build image

> docker build -t repo/twitchbot .

# docker set up
Pull down the latest version using docker

> docker pull gurparit/twitchbot

Run docker image with

> docker run -d -env-file /path/to/env.list --restart=always --name gobot gurparit/twitchbot

You may need to run the docker commands with sudo.
