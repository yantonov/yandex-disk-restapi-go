// disk_info_example.go provides a simple example to fetch a disk info
//
// usage:
//   > go get github.com/yantonov/yandex-disk-restapi-go
//   > cd $GOPATH/github.com/yantonov/yandex-disk-restapi-go/examples
//   > go run disk_info_example.go -token=access_token
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

	flag.StringVar(&accessToken, "token", "", "Access Token")

	flag.Parse()

	if accessToken == "" {
		fmt.Println("\nPlease provide an access_token, one can be found at https://oauth.yandex.ru")

		flag.PrintDefaults()
		os.Exit(1)
	}

	client := src.NewClient(accessToken)

	fmt.Printf("Fetching disk info...\n")
	info, err := client.NewDiskInfoRequest().Exec()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\tTrash size: %d\n", info.Trash_size)
	fmt.Printf("\tTotal space: %d\n", info.Total_space)
	fmt.Printf("\tUsed space: %d\n", info.Used_space)
	sys_folders, err := json.Marshal(info.System_folders)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("\tSystem folders: %s\n", sys_folders)
	// TODO format map info.System_folders
}
