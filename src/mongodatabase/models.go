package mongodatabase

import (
	"APIs"
	"APIs/twitter"
	"APIs/youtube"
)

type TwitterCollection struct {
	Index        string //MD5 for combination lat, lng, date, distance
	LocationHash string //MD5 for lat, lng
	DateStr      string //Date for compiring on 1 day period
	Tweets       []twitter.Tweet
}

type YoutubeCollection struct {
	Index        string //MD5 for combination lat, lng, date, distance
	LocationHash string //MD5 for lat, lng
	DateStr      string //Date for compiring on 1 day period
	Videos       []youtube.Video
}

type InstagramCollection struct {
	Index        string //MD5 for combination lat, lng, date, distance
	LocationHash string //MD5 for lat, lng
	DateStr      string //Date for compiring on 1 day period
	Images       []APIs.ApiImage
}

type FlickrCollection struct {
	Index        string //MD5 for combination lat, lng, date, distance
	LocationHash string //MD5 for lat, lng
	DateStr      string //Date for compiring on 1 day period
	Images       []APIs.ApiImage
}
