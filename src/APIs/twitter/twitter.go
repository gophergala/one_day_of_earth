package twitter

import (
	"config"
	"encoding/json"
	"fmt"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"io/ioutil"
	"lib"
	"net/http"
	"net/url"
	"strconv"
)

type TweetUser struct {
	Name       string
	ScreenName string
	ID         string
	Photo      string
}

type Tweet struct {
	User         TweetUser
	Text         string
	ID           string
	RetweetCount int
	HashTags     []string
}

func api_call(api_url string, query url.Values) (resp *twittergo.APIResponse, err error) {
	var (
		client *twittergo.Client
		req    *http.Request
	)
	auth_config := &oauth1a.ClientConfig{
		ConsumerKey:    config.TWEETER_CONSUMER_KEY,
		ConsumerSecret: config.TWEETER_CONSUMER_SECRET,
	}
	user := oauth1a.NewAuthorizedConfig(config.TWEETER_ACCESS_TOKEN, config.TWEETER_ACCESS_TOKEN_SECRET)
	client = twittergo.NewClient(auth_config, user)

	url := fmt.Sprintf("%s?%v", api_url, query.Encode())
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	resp, err = client.SendRequest(req)
	return
}

func ParseTweets(resp *twittergo.APIResponse) (tweets []Tweet, cerr *lib.CError) {
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	var req_data interface{}
	err = json.Unmarshal(content, &req_data)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}

	st := req_data.(map[string]interface{})
	for _, t := range st["statuses"].([]interface{}) {
		tv := t.(map[string]interface{})
		tweet := Tweet{
			ID:           tv["id_str"].(string),
			Text:         tv["text"].(string),
			RetweetCount: int(tv["retweet_count"].(float64)),
		}
		if tv["hashtags"] != nil {
			for _, h := range tv["hashtags"].([]interface{}) {
				ht := h.(map[string]interface{})
				tweet.HashTags = append(tweet.HashTags, ht["text"].(string))
			}
		}

		u := tv["user"].(map[string]interface{})
		tweet.User = TweetUser{
			Name:       u["name"].(string),
			ID:         u["id_str"].(string),
			ScreenName: u["screen_name"].(string),
			Photo:      u["profile_background_image_url_https"].(string),
		}
		tweets = append(tweets, tweet)
	}
	return
}

func SearchTweets(lat, lng, MinDate string, distance int) (tweets []Tweet, cerr *lib.CError) {
	cerr = nil
	var (
		err  error
		resp *twittergo.APIResponse
	)
	query := url.Values{}
	query.Set("q", "")
	query.Set("geocode", fmt.Sprintf("%s,%s,%skm", lat, lng, strconv.Itoa(distance)))
	query.Set("since", MinDate)
	query.Set("count", "100")
	resp, err = api_call("/1.1/search/tweets.json", query)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	tweets, cerr = ParseTweets(resp)
	return
}
