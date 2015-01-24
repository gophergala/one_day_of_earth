package main

import (
	"APIs/instagram"
	"fmt"
	"time"
)

func main() {
	images, err := instagram.SearchImages("37.7624499", "-122.4602593", int(time.Now().Unix()-86000), 5000)
	if err != nil {
		fmt.Println(err.Message())
	}
	fmt.Println(len(images))
	for _, im := range images {
		fmt.Println(im.Standard.Url)
	}
}
