// disk_info_example.go provides a simple example to fetch a disk info
//
// usage:
//   > go get github.com/yantonov/yandex-disk-restapi-go
//   > cd $GOPATH/github.com/yantonov/yandex-disk-restapi-go/examples
//   > go run resource_info_example.go -token=access_token
//
//   You can find an access_token for your app at https://oauth.yandex.ru
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/yantonov/yandex-disk-restapi-go/src"
	"os"
)

func main() {
	var accessToken string
	var path string

	flag.StringVar(&accessToken, "token", "", "Access Token")
	flag.StringVar(&path, "path", "/", "Path to resource")

	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://oauth.yandex.ru")

		flag.PrintDefaults()
		os.Exit(1)
	}

	client := src.NewClient(accessToken)

	fmt.Printf("Fetching flat file list ...\n")
	info, err := client.NewFlatFileListRequest().Exec()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	items, err := json.Marshal(info.Items)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\tItems: %s\n", items)
	if info.Limit != nil {
		fmt.Printf("\tLimit: %d\n", info.Limit)
	}
	if info.Offset != nil {
		fmt.Printf("\tOffset: %d\n", *info.Offset)
	}
}
