package twitter

import (
	"config"
	"fmt"
	"github.com/kurrik/oauth1a"
	"github.com/kurrik/twittergo"
	"lib"
	"net/http"
	"net/url"
	"strconv"
)

type TweetUser struct {
	Name       string
	ScreenName string
	ID         string
}

type Tweet struct {
	User TweetUser
	Text string
	ID   string
}

func SearchTweets(lat, lng, MinDate string, distance int) (tweets []Tweet, cerr *lib.CError) {
	cerr = nil
	var (
		err     error
		client  *twittergo.Client
		req     *http.Request
		resp    *twittergo.APIResponse
		results *twittergo.SearchResults
	)
	auth_config := &oauth1a.ClientConfig{
		ConsumerKey:    config.TWEETER_CONSUMER_KEY,
		ConsumerSecret: config.TWEETER_CONSUMER_SECRET,
	}
	user := oauth1a.NewAuthorizedConfig(config.TWEETER_ACCESS_TOKEN, config.TWEETER_ACCESS_TOKEN_SECRET)
	client = twittergo.NewClient(auth_config, user)

	query := url.Values{}
	query.Set("q", "")
	query.Set("geocode", fmt.Sprintf("%s,%s,%skm", lat, lng, strconv.Itoa(distance)))
	query.Set("since", MinDate)
	query.Set("count", "100")
	url := fmt.Sprintf("/1.1/search/tweets.json?%v", query.Encode())
	req, err = http.NewRequest("GET", url, nil)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	resp, err = client.SendRequest(req)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	results = &twittergo.SearchResults{}
	err = resp.Parse(results)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	for _, tweet := range results.Statuses() {
		user := tweet.User()
		t := Tweet{}
		t.ID = strconv.Itoa(int(tweet.Id()))
		t.Text = tweet.Text()
		t.User = TweetUser{
			Name:       user.Name(),
			ScreenName: user.ScreenName(),
			ID:         user.IdStr(),
		}
		tweets = append(tweets, t)
	}
	return
}
