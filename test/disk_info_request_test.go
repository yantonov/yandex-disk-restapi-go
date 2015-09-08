package test

import (
	"fmt"
	diskclient "github.com/yantonov/yandex-disk-restapi-go/src"
	"net/http"
	"reflect"
	"testing"
)

func Test_DiskInfo_Simple(t *testing.T) {
	var client = NewStubResponseClient("{}", http.StatusOK)
	response, err := client.NewDiskInfoRequest().Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.DiskInfoResponse{}
	expected.System_folders = make(map[string]string)

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}

func Test_NonEmptySystemFolders(t *testing.T) {
	var client = NewStubResponseClient(`{"system_folders": {"applications": "disk:/Applications", "downloads": "disk:/Downloads/"}}`, http.StatusOK)
	response, err := client.NewDiskInfoRequest().Exec()
	if err != nil {
		t.Error(fmt.Sprintf("unexpected error %s", err.Error()))
	}
	var expected = &diskclient.DiskInfoResponse{}
	expected.System_folders = make(map[string]string)
	expected.System_folders["applications"] = "disk:/Applications"
	expected.System_folders["downloads"] = "disk:/Downloads/"

	if !reflect.DeepEqual(response, expected) {
		t.Errorf("should match\nactual   = %v\nexpected = %v", response, expected)
	}
}
