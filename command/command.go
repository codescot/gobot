package command

import (
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gurparit/slackbot/util"
)

// Command the basic command interface
type Command interface {
	Execute(func(string), string)
}

// HTTPCommand http command helper methods
type HTTPCommand struct {
	Headers map[string]string
	URL     string
}

// Result get a json formatted http response.
func (httpCommand HTTPCommand) Result() ([]byte, error) {
	var err error
	var body []byte

	request, err := http.NewRequest("GET", httpCommand.URL, nil)
	request.Header.Add("Accept-Encoding", "gzip")

	for key, value := range httpCommand.Headers {
		request.Header.Add(key, value)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if util.IsError(err) {
		return nil, err
	}

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

// JSON makes a json request and returns the result
func JSON(getURL func() string, result interface{}) error {
	var err error

	targetURL := getURL()

	httpCommand := HTTPCommand{URL: targetURL}
	body, err := httpCommand.Result()

	err = json.Unmarshal(body, result)

	return err
}
