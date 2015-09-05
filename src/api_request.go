package src

type apiRequest interface {
	Request() *httpRequest
}
