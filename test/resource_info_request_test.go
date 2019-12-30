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
"_embedded": {
    "sort": "name",
    "path": "disk:/foo",
    "items": [
      {
        "path": "disk:/foo/bar",
        "type": "dir",
        "name": "bar",
        "modified": "2014-04-22T10:32:49+04:00",
        "created": "2014-04-22T10:32:49+04:00"
      },
      {
        "name": "photo.png",
        "preview": "https://downloader.disk.yandex.ru/preview/...",
        "created": "2014-04-21T14:57:13+04:00",
        "modified": "2014-04-21T14:57:14+04:00",
        "path": "disk:/foo/photo.png",
        "md5": "4334dc6379c8f95ddf11b9508cfea271",
        "type": "file",
        "mime_type": "image/png",
        "size": 34567
      }
    ],
    "limit": 20,
    "offset": 0
  },
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
	expected.Custom_properties = custom_properties

	expected.Public_url = "https://yadi.sk/d/2rEgCiNTZGiYX"
	expected.Origin_path = "disk:/foo/photo.png"
	expected.Modified = "2014-04-21T14:57:14+04:00"
	expected.Path = "disk:/foo/photo.png"
	expected.Md5 = "4334dc6379c8f95ddf11b8508cfea271"
	expected.Resource_type = "file"
	expected.Mime_type = "application/x-www-form-urlencoded"
	expected.Size = 34567

	var limit uint64 = 20
	var offset uint64 = 0
	var embedded diskclient.ResourceListResponse = diskclient.ResourceListResponse{
		Sort:   (&diskclient.SortMode{}).ByName(),
		Path:   "disk:/foo",
		Limit:  &limit,
		Offset: &offset,
		Items: []diskclient.ResourceInfoResponse{
			diskclient.ResourceInfoResponse{
				Path:          "disk:/foo/bar",
				Resource_type: "dir",
				Name:          "bar",
				Modified:      "2014-04-22T10:32:49+04:00",
				Created:       "2014-04-22T10:32:49+04:00"},
			diskclient.ResourceInfoResponse{
				Name:          "photo.png",
				Preview:       "https://downloader.disk.yandex.ru/preview/...",
				Created:       "2014-04-21T14:57:13+04:00",
				Modified:      "2014-04-21T14:57:14+04:00",
				Path:          "disk:/foo/photo.png",
				Md5:           "4334dc6379c8f95ddf11b9508cfea271",
				Resource_type: "file",
				Mime_type:     "image/png",
				Size:          34567,
			},
		},
	}
	expected.Embedded = &embedded

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_ResourceInfo_EmptyEmbeddedSortMode(t *testing.T) {
	var client = NewStubResponseClient(`{
"_embedded": {
    "sort": ""
},
  "name": "photo.png",
  "size": 34567
}`, http.StatusOK)
	response, err := client.NewResourceInfoRequest("/some_dir").Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.ResourceInfoResponse{}
	expected.Name = "photo.png"
	var custom_properties = make(map[string]interface{})
	expected.Custom_properties = custom_properties
	expected.Size = 34567

	var embedded diskclient.ResourceListResponse = diskclient.ResourceListResponse{
		Sort:  (&diskclient.SortMode{}).Default(),
		Items: []diskclient.ResourceInfoResponse{},
	}
	expected.Embedded = &embedded
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

	if request_limit, ok := (request.Parameters["limit"]).(uint32); ok {
		if request_limit != 123456 {
			t.Errorf("invalid limit, actual : %d", request_limit)
		}
	} else {
		t.Errorf("limit is undefined")
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

	if request_offset, ok := (request.Parameters["offset"]).(uint32); ok {
		if request_offset != 123456 {
			t.Errorf("invalid offset, actual : %d", request_offset)
		}
	} else {
		t.Errorf("offset is undefined")
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

	if extracted_crop, ok := request.Parameters["preview_crop"].(bool); ok {
		if extracted_crop != true {
			t.Errorf("invalid preview_crop, actual : %v", extracted_crop)
		}
	} else {
		t.Errorf("preview_crop is undefined")
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

func Test_ResourceRequest_FieldsList(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var fields = []string{"a", "b"}
	var options = diskclient.ResourceInfoRequestOptions{
		Fields: fields,
	}
	request := client.NewResourceInfoRequest("/path", options).Request()

	extracted_param := request.Parameters["fields"]
	if extracted_param != "a,b" {
		t.Errorf("invalid fields param, actual : %v", extracted_param)
	}
}

func Test_ResourceRequest_EmptyFieldsList(t *testing.T) {
	var client = NewStubResponseClient(`{}`, http.StatusOK)
	var options = diskclient.ResourceInfoRequestOptions{}
	request := client.NewResourceInfoRequest("/path", options).Request()

	extracted_param := request.Parameters["fields"]
	if extracted_param != nil {
		t.Errorf("fields param must be undefined")
	}
}
