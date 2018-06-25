FROM golang:alpine

WORKDIR /go/app

ADD gobot /go/app

CMD ["./gobot"]
