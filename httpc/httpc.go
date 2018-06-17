package httpc

import (
	"net/http"
	"compress/gzip"
	"io/ioutil"
	"encoding/json"
	"io"
	"net/url"
	"strings"
	"fmt"
	"errors"
)

type URL func()string

type HTTP struct {
	TargetURL string
	Method    string
	Headers   map[string]string
	Form      map[string]string
}

// Result get a json formatted http response.
func (h HTTP) httpResult() ([]byte, int, error) {
	var err error
	var body []byte
	var data io.Reader = nil

	if h.Form != nil {
		values := url.Values{}
		for key, value := range h.Form {
			values.Add(key, value)
		}

		data = strings.NewReader(values.Encode())
	}

	request, err := http.NewRequest(h.Method, h.TargetURL, data)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	request.Header.Add("Accept-Encoding", "gzip")

	if h.Headers != nil {
		for key, value := range h.Headers {
			request.Header.Add(key, value)
		}
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, response.StatusCode, err
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

	return body, response.StatusCode, err
}

// JSON makes a json request and returns the result
func (h HTTP) JSON(result interface{}) error {
	body, status, err := h.httpResult()
	if err != nil {
		return err
	}

	switch status {
	case http.StatusOK:
		return json.Unmarshal(body, result)
	default:
		response := string(body)
		return errors.New(response)
	}
}

func (h HTTP) String() (string, error) {
	body, status, err := h.httpResult()
	if err != nil {
		return "", err
	}

	switch status {
	case http.StatusOK:
		return string(body), nil
	default:
		response := string(body)
		return "", errors.New(response)
	}
}
