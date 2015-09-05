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

func Test_ResourceRequest_WithSortMode(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{
		Sort_mode: (&diskclient.SortMode{}).BySize(),
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["sort"] != "size" {
		t.Errorf("invalid sort mode, actual : %s", request.Parameters["sort"])
	}
}

func Test_ResourceRequest_WithReverseSortMode(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{
		Sort_mode: (&diskclient.SortMode{}).BySize().Reverse(),
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["sort"] != "-size" {
		t.Errorf("invalid sort mode, actual : %s", request.Parameters["sort"])
	}

	options = diskclient.ResourceInfoRequestOptions{
		Sort_mode: (&diskclient.SortMode{}).BySize().Reverse().Reverse(),
	}
	request = client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["sort"] != "size" {
		t.Errorf("invalid sort mode, actual : %s", request.Parameters["sort"])
	}
}

func Test_ResourceRequest_WithoutSortMode(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["sort"] != nil {
		t.Errorf("invalid sort mode, actual : %s", request.Parameters["sort"])
	}
}

func Test_ResourceRequest_Limit(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var limit uint32 = 123456
	var options = diskclient.ResourceInfoRequestOptions{
		Limit: &limit,
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["limit"] == nil {
		t.Errorf("limit is undefined")
	}
	request_limit := (request.Parameters["limit"]).(*uint32)
	if *request_limit != 123456 {
		t.Errorf("invalid limit, actual : %d", *request_limit)
	}
}

func Test_ResourceRequest_NoLimit(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["limit"] != nil {
		t.Errorf("limit must be undefined")
	}
}

func Test_ResourceRequest_Offset(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var offset uint32 = 123456
	var options = diskclient.ResourceInfoRequestOptions{
		Offset: &offset,
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["offset"] == nil {
		t.Errorf("offset is undefined")
	}
	request_offset := (request.Parameters["offset"]).(*uint32)
	if *request_offset != 123456 {
		t.Errorf("invalid offset, actual : %d", *request_offset)
	}
}

func Test_ResourceRequest_NoOffset(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["offset"] != nil {
		t.Errorf("offset must be undefined")
	}
}

func Test_ResourceRequest_PreviewSize(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{
		Preview_size: (&diskclient.PreviewSize{}).PredefinedSizeM(),
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	size := request.Parameters["preview_size"]
	if size != "M" {
		t.Errorf("invalid preview_size, actual : %d", size)
	}
}

func Test_ResourceRequest_NoPreviewSize(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{}
	request := client.NewResourceInfoRequest("/path", options).Request()

	size := request.Parameters["preview_size"]
	if size != nil {
		t.Errorf("preview size must be undefined")
	}
}

func Test_ResourceRequest_PreviewCrop(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	crop := true
	var options = diskclient.ResourceInfoRequestOptions{
		Preview_crop: &crop,
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	extracted_crop := request.Parameters["preview_crop"].(*bool)
	if *extracted_crop != true {
		t.Errorf("invalid preview_crop, actual : %v", *extracted_crop)
	}
}

func Test_ResourceRequest_NoPreviewCrop(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{}
	request := client.NewResourceInfoRequest("/path", options).Request()

	if request.Parameters["preview_crop"] != nil {
		t.Errorf("preview_crop must be undefined")
	}
}
