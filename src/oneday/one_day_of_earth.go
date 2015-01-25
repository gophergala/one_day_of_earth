package main

import (
	//	"APIs/flickr"
	//	"APIs/instagram"
	//	"APIs/twitter"
	//	"APIs/youtube"
	//	"config"
	"cronTask"
	//	"fmt"
	//	"lib"
	//	"strconv"
	//	"time"
	//	"mongodatabase"
)

func main() {

	cronTask.StartCron(10)

	//	videos, err := youtube.SearchVideos("37.7624499", "-122.4602593", lib.YesterdayTime().Format(time.RFC3339), 1000, true, "")
	//	if err != nil {
	//		fmt.Println(err.Message())
	//	}
	//	fmt.Println(len(videos))
	//	for _, im := range videos {
	//		fmt.Println(im.ID, im.Title, im.Thumb)
	//	}
}
