package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	token      string
	basePath   string
	HttpClient *http.Client
}

func NewClient(token string, client ...*http.Client) *Client {
	return newClientInternal(
		token,
		"https://cloud-api.yandex.net/v1/disk",
		client...)
}

func newClientInternal(token string, basePath string, client ...*http.Client) *Client {
	c := &Client{
		token:    token,
		basePath: basePath,
	}
	if len(client) != 0 {
		c.HttpClient = client[0]
	} else {
		c.HttpClient = http.DefaultClient
	}
	return c
}

type ErrorHandler func(*http.Response) error

var defaultErrorHandler ErrorHandler = func(resp *http.Response) error {
	if resp.StatusCode/100 == 5 {
		return errors.New("server error")
	}

	if resp.StatusCode/100 == 4 {
		var response DiskClientError
		contents, _ := ioutil.ReadAll(resp.Body)
		json.Unmarshal(contents, &response)

		return response
	}

	if resp.StatusCode/100 == 3 {
		return errors.New("redirect error")
	}
	return nil
}

func (httpRequest *httpRequest) run(client *Client) ([]byte, error) {
	var err error

	values := make(url.Values)
	if httpRequest.parameters != nil {
		for k, v := range httpRequest.parameters {
			values.Set(k, fmt.Sprintf("%v", v))
		}
	}

	var req *http.Request
	if httpRequest.method == "POST" {
		// TODO json serialize
		req, err = http.NewRequest(
			"POST",
			client.basePath+httpRequest.path,
			strings.NewReader(values.Encode()))
		if err != nil {
			return nil, err
		}
		// TODO
		// req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(
			httpRequest.method,
			client.basePath+httpRequest.path+"?"+values.Encode(),
			nil)
		if err != nil {
			return nil, err
		}
	}

	for headerName := range httpRequest.headers {
		var headerValues = httpRequest.headers[headerName]
		for _, headerValue := range headerValues {
			req.Header.Set(headerName, headerValue)
		}
	}
	return runRequest(client, req)
}

func runRequest(client *Client, req *http.Request) ([]byte, error) {
	return runRequestWithErrorHandler(client, req, defaultErrorHandler)
}

func runRequestWithErrorHandler(client *Client, req *http.Request, errorHandler ErrorHandler) ([]byte, error) {
	resp, err := client.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return checkResponseForErrorsWithErrorHandler(resp, errorHandler)
}

func checkResponseForErrorsWithErrorHandler(resp *http.Response, errorHandler ErrorHandler) ([]byte, error) {
	if resp.StatusCode/100 > 2 {
		return nil, errorHandler(resp)
	} else {
		return ioutil.ReadAll(resp.Body)
	}
}
