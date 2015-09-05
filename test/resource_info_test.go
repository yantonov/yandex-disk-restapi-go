package test

import (
	"fmt"
	diskclient "github.com/yantonov/yandex-disk-restapi-go/src"
	"net/http"
	"reflect"
	"testing"
)

func Test_ResourceInfo_Simple(t *testing.T) {
	var client = NewStubResponseClient(`{
  "public_key": "HQsmHLoeyBlJf8Eu1jlmzuU+ZaLkjPkgcvmoktUCIo8=",
  "name": "photo.png",
  "created": "2014-04-21T14:57:13+04:00",
  "custom_properties": {"foo": "1", "bar": "2"},
  "public_url": "https://yadi.sk/d/2rEgCiNTZGiYX",
  "origin_path": "disk:/foo/photo.png",
  "modified": "2014-04-21T14:57:14+04:00",
  "path": "disk:/foo/photo.png",
  "md5": "4334dc6379c8f95ddf11b8508cfea271",
  "type": "file",
  "mime_type": "application/x-www-form-urlencoded",
  "size": 34567
}`, http.StatusOK)
	response, err := client.NewResourceInfoRequest("/some_dir").Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.ResourceInfoResponse{}
	expected.Public_key = "HQsmHLoeyBlJf8Eu1jlmzuU+ZaLkjPkgcvmoktUCIo8="
	expected.Name = "photo.png"
	expected.Created = "2014-04-21T14:57:13+04:00"
	var custom_properties = make(map[string]interface{})
	custom_properties["foo"] = "1"
	custom_properties["bar"] = "2"
	expected.Public_url = "https://yadi.sk/d/2rEgCiNTZGiYX"
	expected.Custom_properties = custom_properties
	expected.Origin_path = "disk:/foo/photo.png"
	expected.Modified = "2014-04-21T14:57:14+04:00"
	expected.Path = "disk:/foo/photo.png"
	expected.Md5 = "4334dc6379c8f95ddf11b8508cfea271"
	expected.Resource_type = "file"
	expected.Mime_type = "application/x-www-form-urlencoded"
	expected.Size = 34567

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_ResourceInfo_EmptyCustomProperties(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	response, err := client.NewResourceInfoRequest("/path").Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.ResourceInfoResponse{}
	expected.Custom_properties = make(map[string]interface{})
	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}
