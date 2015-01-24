package main

import (
	//	"APIs/flickr"
	//	"APIs/instagram"
	//	"APIs/twitter"
	"APIs/youtube"
	"fmt"
	//	"strconv"
	"time"
)

func main() {
	//	tweets, err := twitter.SearchTweets("37.7624499", "-122.4602593", time.Now().Add(time.Duration(-86000)).Format("2006-01-02"), 1000, true)
	//	if err != nil {
	//		fmt.Println(err.Message())
	//	}
	//	fmt.Println(len(tweets))
	//	for _, im := range tweets {
	//		fmt.Println(im.User.ScreenName, im.ID, im.RetweetCount)
	//	}

	videos, err := youtube.SearchVideos("37.7624499", "-122.4602593", time.Now().Add(time.Duration((2 * 86000))).Format(time.RFC3339), 1000, true, "")
	if err != nil {
		fmt.Println(err.Message())
	}
	fmt.Println(len(videos))
	for _, im := range videos {
		fmt.Println(im.ID, im.Title, im.Thumb)
	}
}
