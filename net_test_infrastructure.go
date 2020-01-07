package yandexdiskapi

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type storeRequestTransport struct {
	http.Transport
	request *http.Request
}

type stubResponseTransport struct {
	http.Transport
	content    string
	statusCode int
}

func NewStubResponseClient(content string, statusCode ...int) *Client {
	c := NewClient("")
	t := &stubResponseTransport{content: content}

	if len(statusCode) != 0 {
		t.statusCode = statusCode[0]
	}

	c.HttpClient = &http.Client{Transport: t}

	return c
}

func (t *stubResponseTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	resp := &http.Response{
		Status:     http.StatusText(t.statusCode),
		StatusCode: t.statusCode,
	}
	resp.Body = ioutil.NopCloser(strings.NewReader(t.content))

	return resp, nil
}
