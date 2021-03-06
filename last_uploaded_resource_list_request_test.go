package yandexdiskapi

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func Test_LastUploadedResourceList_Simple(t *testing.T) {
	var client = NewStubResponseClient(`{
  "items": [
      {
        "name": "photo2.png",
        "preview": "https://downloader.disk.yandex.ru/preview/...",
        "created": "2014-04-22T14:57:13+04:00",
        "modified": "2014-04-22T14:57:14+04:00",
        "path": "disk:/foo/photo2.png",
        "md5": "53f4dc6379c8f95ddf11b9508cfea271",
        "type": "file",
        "mime_type": "image/png",
        "size": 54321
      },
      {
        "name": "photo1.png",
        "preview": "https://downloader.disk.yandex.ru/preview/...",
        "created": "2014-04-21T14:57:13+04:00",
        "modified": "2014-04-21T14:57:14+04:00",
        "path": "disk:/foo/photo1.png",
        "md5": "4334dc6379c8f95ddf11b9508cfea271",
        "type": "file",
        "mime_type": "image/png",
        "size": 34567
      }
    ],
    "limit": 20,
    "offset": 0
  }`, http.StatusOK)
	response, err := client.NewLastUploadedResourceListRequest().Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &LastUploadedResourceListResponse{}
	var resource1 = ResourceInfoResponse{
		Name:          "photo2.png",
		Preview:       "https://downloader.disk.yandex.ru/preview/...",
		Created:       "2014-04-22T14:57:13+04:00",
		Modified:      "2014-04-22T14:57:14+04:00",
		Path:          "disk:/foo/photo2.png",
		Md5:           "53f4dc6379c8f95ddf11b9508cfea271",
		Resource_type: "file",
		Mime_type:     "image/png",
		Size:          54321,
	}
	var resource2 = ResourceInfoResponse{
		Name:          "photo1.png",
		Preview:       "https://downloader.disk.yandex.ru/preview/...",
		Created:       "2014-04-21T14:57:13+04:00",
		Modified:      "2014-04-21T14:57:14+04:00",
		Path:          "disk:/foo/photo1.png",
		Md5:           "4334dc6379c8f95ddf11b9508cfea271",
		Resource_type: "file",
		Mime_type:     "image/png",
		Size:          34567,
	}
	expected.Items = []ResourceInfoResponse{resource1, resource2}
	var limit uint64 = 20
	expected.Limit = &limit

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_LastUploadedResourceList_NoItemsInResponse(t *testing.T) {
	var client = NewStubResponseClient(`{
    "limit": 20,
    "offset": 0
  }`, http.StatusOK)
	response, err := client.NewLastUploadedResourceListRequest().Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &LastUploadedResourceListResponse{}
	expected.Items = []ResourceInfoResponse{}
	var limit uint64 = 20
	expected.Limit = &limit

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_LastUploadedResourceListRequest_Limit(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var limit uint32 = 123456
	var options = LastUploadedResourceListRequestOptions{
		Limit: &limit,
	}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	if request_limit, ok := request.Parameters["limit"].(uint32); ok {
		if request_limit != 123456 {
			t.Errorf("invalid limit, actual : %d", request_limit)
		}
	} else {
		t.Errorf("limit is undefined")
	}
}

func Test_LastUploadedResourceListRequest_NoLimit(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	if request.Parameters["limit"] != nil {
		t.Errorf("limit must be undefined")
	}
}

func Test_LastUploadedResourceListRequest_PreviewSize(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{
		Preview_size: (&PreviewSize{}).PredefinedSizeM(),
	}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	size := request.Parameters["preview_size"]
	if size != "M" {
		t.Errorf("invalid preview_size, actual : %d", size)
	}
}

func Test_LastUploadedResourceListRequest_NoPreviewSize(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	size := request.Parameters["preview_size"]
	if size != nil {
		t.Errorf("preview size must be undefined")
	}
}

func Test_LastUploadedResourceListRequest_PreviewCrop(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	crop := true
	var options = LastUploadedResourceListRequestOptions{
		Preview_crop: &crop,
	}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	if extracted_crop, ok := request.Parameters["preview_crop"].(bool); ok {
		if extracted_crop != true {
			t.Errorf("invalid preview_crop, actual : %v", extracted_crop)
		}
	} else {
		t.Errorf("preview_crop is undefined")
	}
}

func Test_LastUploadedResourceListRequest_NoPreviewCrop(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	if request.Parameters["preview_crop"] != nil {
		t.Errorf("preview_crop must be undefined")
	}
}

func Test_LastUploadedResourceListRequest_FieldsList(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var fields = []string{"a", "b"}
	var options = LastUploadedResourceListRequestOptions{
		Fields: fields,
	}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	extracted_param := request.Parameters["fields"]
	if extracted_param != "a,b" {
		t.Errorf("invalid fields param, actual : %v", extracted_param)
	}
}

func Test_LastUploadedResourceListRequest_EmptyFieldsList(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	extracted_param := request.Parameters["fields"]
	if extracted_param != nil {
		t.Errorf("fields param must be undefined")
	}
}

func Test_LastUploadedResourceListRequest_MediaTypes(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{}
	options.Media_type = []MediaType{
		*(&MediaType{}).Audio(),
		*(&MediaType{}).Backup(),
	}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	extracted_param := request.Parameters["media_type"]
	if extracted_param != "audio,backup" {
		t.Errorf(fmt.Sprintf("invalid media type, actual = %v", request.Parameters["media_type"]))
	}
}

func Test_LastUploadedResourceListRequest_NoMediaTypes(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = LastUploadedResourceListRequestOptions{}
	request := client.NewLastUploadedResourceListRequest(options).Request()

	extracted_param := request.Parameters["media_type"]
	if extracted_param != nil {
		t.Errorf("media type is not expected")
	}
}
