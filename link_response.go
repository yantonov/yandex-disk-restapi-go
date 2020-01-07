package yandexdiskapi

// https://tech.yandex.com/disk/api/reference/response-objects-docpage/#link
type LinkResponse struct {
	Href      string `json:"href"`
	Method    string `json:"method"`
	Templated bool   `json:"templated"`
}
