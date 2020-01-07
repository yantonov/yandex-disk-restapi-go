package yandexdiskapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test_Download_Simple(t *testing.T) {
	var client = NewStubResponseClient(`{
  "href": "https://cloud-api.yandex.net/v1/disk/resources?path=disk%3A%2Ffoo%2Fphoto.png",
  "method": "GET",
  "templated": true
}`, http.StatusOK)
	response, err := client.NewDownloadRequest("/some_dir/photo.png").Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &LinkResponse{}
	expected.Href = "https://cloud-api.yandex.net/v1/disk/resources?path=disk%3A%2Ffoo%2Fphoto.png"
	expected.Method = "GET"
	expected.Templated = true

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}
