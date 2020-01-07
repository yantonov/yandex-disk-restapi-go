package yandexdiskapi

type apiRequest interface {
	Request() *httpRequest
}
