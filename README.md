# gobot
IRC Bot written in Golang

# docker build image

> docker build -t repo/gobot .

# docker set up
Pull down the latest version using docker

> docker pull codescot/gobot

Run docker image with

> docker run -d -env-file /path/to/env.list --restart=always --name gobot codescot/gobot

You may need to run the docker commands with sudo.
