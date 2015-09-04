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

const basePath = "https://cloud-api.yandex.net/v1/disk"

type Client struct {
    token      string
    HttpClient *http.Client
}

func NewClient(token string, client ...*http.Client) *Client {
    c := &Client{token: token}
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

func (client *Client) run(method, path string, params map[string]interface{}) ([]byte, error) {
    var err error

    values := make(url.Values)
    for k, v := range params {
        values.Set(k, fmt.Sprintf("%v", v))
    }

    var req *http.Request
    if method == "POST" {
        // TODO json serialize
        req, err = http.NewRequest("POST", basePath+path, strings.NewReader(values.Encode()))
        if err != nil {
            return nil, err
        }
        req.Header.Set("Content-Type", "application/json")
    } else {
        req, err = http.NewRequest(method, basePath+path+"?"+values.Encode(), nil)
        if err != nil {
            return nil, err
        }
    }

    return client.runRequest(req)
}

func (client *Client) runRequest(req *http.Request) ([]byte, error) {
    return client.runRequestWithErrorHandler(req, defaultErrorHandler)
}

func (client *Client) runRequestWithErrorHandler(req *http.Request, errorHandler ErrorHandler) ([]byte, error) {
    req.Header.Set("Authorization", "OAuth "+client.token)
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
