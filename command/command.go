package command

import (
	"compress/gzip"
	"io/ioutil"
	"net/http"

	"github.com/gurparit/marbles/util"
)

// Command the basic command interface
type Command interface {
	Execute(respond func(string), message string)
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
