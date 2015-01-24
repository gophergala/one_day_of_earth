package main

import (
	//	"APIs/flickr"
	//	"APIs/instagram"
	"APIs/twitter"
	"fmt"
	//	"strconv"
	"time"
)

func main() {
	tweets, err := twitter.SearchTweets("37.7624499", "-122.4602593", time.Now().Add(time.Duration(86000)).Format("2006-01-02"), 1000)
	if err != nil {
		fmt.Println(err.Message())
	}
	fmt.Println(len(tweets))
	for _, im := range tweets {
		fmt.Println(im.User.ScreenName, im.ID, im.RetweetCount)
	}
}
