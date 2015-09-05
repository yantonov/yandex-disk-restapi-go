package src

type apiRequest interface {
	request() *httpRequest
}
