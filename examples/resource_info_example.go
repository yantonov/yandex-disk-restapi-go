// disk_info_example.go provides a simple example to fetch a disk info
//
// usage:
//   > go get github.com/yantonov/yandex-disk-restapi-go
//   > cd $GOPATH/github.com/yantonov/yandex-disk-restapi-go/examples
//   > go run resource_info_example.go -token=access_token -path=path_to_resource
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

	fmt.Printf("Fetching resource info...\n")
	info, err := client.NewResourceInfoRequest(path).Exec()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\tPublic key: %s\n", info.Public_key)
	fmt.Printf("\tName: %s\n", info.Name)
	fmt.Printf("\tCreated: %s\n", info.Created)

	custom_properties, err := json.Marshal(info.Custom_properties)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	embedded, err := json.Marshal(info.Embedded)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\tPublic url: %s\n", info.Public_url)
	fmt.Printf("\tEmbedded: %s\n", embedded)
	fmt.Printf("\tOrigin path: %s\n", info.Origin_path)
	fmt.Printf("\tModified: %s\n", info.Modified)
	fmt.Printf("\tCustom_properties: %s\n", custom_properties)
	fmt.Printf("\tPath: %s\n", info.Path)
	fmt.Printf("\tMd5: %s\n", info.Md5)
	fmt.Printf("\tType: %s\n", info.Resource_type)
	fmt.Printf("\tMime type: %s\n", info.Mime_type)
	fmt.Printf("\tSize: %d\n", info.Size)
}
