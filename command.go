package main

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"

	irc "github.com/thoj/go-ircevent"
)

// Command the basic command interface
type Command interface {
	Execute(ircobj *irc.Connection, event *irc.Event)
}

// HTTPCommand http command helper methods
type HTTPCommand struct {
	Headers map[string]string
}

// JSONResult get a json http request.
func (httpCommand HTTPCommand) JSONResult(targetURL string) ([]byte, error) {
	var err error
	var body []byte

	request, err := http.NewRequest("GET", targetURL, nil)
	request.Header.Add("Accept-Encoding", "gzip")

	for key, value := range httpCommand.Headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	defer response.Body.Close()

	contentEncoding := response.Header.Get("Content-Encoding")

	switch contentEncoding {
	case "gzip":
		gzipReader, _ := gzip.NewReader(response.Body)
		body, err = ioutil.ReadAll(gzipReader)
		break
	default:
		body, err = ioutil.ReadAll(response.Body)
		break
	}

	return body, err
}
