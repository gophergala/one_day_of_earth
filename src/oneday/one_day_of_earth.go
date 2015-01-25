package main

import (
	//	"APIs/flickr"
	//	"APIs/instagram"
	//	"APIs/twitter"
	//	"APIs/youtube"
	"config"
	"cronTask"
	"fmt"
	"lib"
	"strconv"
	"time"
	//	"mongodatabase"
)

func main() {
	var err *lib.CError
	err = cronTask.RunCronTask(config.TWITTER_CRON, "37.7624499", "-122.4602593", lib.YesterdayTime().Format("2006-01-02"), "1000")

	if err != nil {
		fmt.Println(err.Message())
	} else {
		fmt.Println("Twitter Done")
	}

	err = cronTask.RunCronTask(config.YOUTUBE_CRON, "37.7624499", "-122.4602593", lib.YesterdayTime().Format(time.RFC3339), "1000")

	if err != nil {
		fmt.Println(err.Message())
	} else {
		fmt.Println("Youtube Done")
	}

	err = cronTask.RunCronTask(config.INSTAGRAM_CRON, "37.7624499", "-122.4602593", strconv.Itoa(lib.YesterdayTime().Second()), "5000")

	if err != nil {
		fmt.Println(err.Message())
	} else {
		fmt.Println("Instagram Done")
	}

	err = cronTask.RunCronTask(config.FLICKR_CRON, "37.7624499", "-122.4602593", lib.YesterdayTime().Format("2006-01-02"), "20")

	if err != nil {
		fmt.Println(err.Message())
	} else {
		fmt.Println("Flickr Done")
	}

	//	videos, err := youtube.SearchVideos("37.7624499", "-122.4602593", lib.YesterdayTime().Format(time.RFC3339), 1000, true, "")
	//	if err != nil {
	//		fmt.Println(err.Message())
	//	}
	//	fmt.Println(len(videos))
	//	for _, im := range videos {
	//		fmt.Println(im.ID, im.Title, im.Thumb)
	//	}
}
