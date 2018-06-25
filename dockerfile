FROM golang:alpine

#RUN go get -d -v ./...
#RUN go install -v

WORKDIR /go/src/app
COPY . .

CMD ["app"]
