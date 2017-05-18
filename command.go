package main

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"

	irc "github.com/thoj/go-ircevent"
)

// Command the basic command interface
type Command interface {
	execute(ircobj *irc.Connection, event *irc.Event)
}

// HTTPCommand http command helper methods
type HTTPCommand struct{}

func (httpCommand HTTPCommand) getJSONResult(targetURL string) ([]byte, error) {
	var err error
	var body []byte

	request, err := http.NewRequest("GET", targetURL, nil)
	request.Header.Add("Accept-Encoding", "gzip")

	client := &http.Client{}
	response, err := client.Do(request) //http.Get(targetURL)
	defer response.Body.Close()

	switch request.Header.Get("Content-Encoding") {
	case "gzip":
		gzipReader, _ := gzip.NewReader(response.Body)
		body, err = ioutil.ReadAll(gzipReader)
	default:
		body, err = ioutil.ReadAll(response.Body)
	}

	return body, err
}
