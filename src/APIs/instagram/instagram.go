package instagram

import (
	"config"
	"encoding/json"
	"io/ioutil"
	"lib"
	"net/http"
	"strconv"
)

//Private function only for this package, just for making REST API calls
func api_call(method int, api_url string, params map[string]string) (ret_data interface{}, cerr *lib.CError) {
	var (
		resp *http.Response
		err  error
	)
	cerr = nil
	params["access_token"] = config.INSTAGRAM_ACCESS_TOKEN
	switch method {
	case config.GET:
		{
			resp, err = http.Get(lib.GenerateURL(api_url, params))
		}
	case config.POST:
		{
			resp, err = http.PostForm(api_url, lib.UrlValues(params))
		}
	default:
		cerr = &lib.CError{}
		cerr.SetMessage("Unknown Request type")
		return
	}

	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		cerr = &lib.CError{}
		cerr.SetMessage(err.Error())
		return
	}
	json.Unmarshal(body, &ret_data)
	return
}

type Image struct {
	Width  int
	Height int
	Url    string
}

type InstaUser struct {
	UserName     string
	FullName     string
	ID           string
	ProfilePhoto string //URl to profile photo
}

type InstaImage struct {
	User      InstaUser
	Id        string
	Instalink string //Link to Instagram for that image
	Standard  Image
	Thumb     Image
	Small     Image
}

func SearchImages(lat, lng string, minTimeStamp, distance int) (images []InstaImage, cerr *lib.CError) {
	cerr = nil
	data, cerr := api_call(config.GET, "https://api.instagram.com/v1/media/search", map[string]string{
		"lat":           lat,
		"lng":           lng,
		"min_timestamp": strconv.Itoa(minTimeStamp),
		"distance":      strconv.Itoa(distance),
	})
	if cerr != nil {
		return
	}
	i := data.(map[string]interface{})
	//Getting Status code for request
	s := i["meta"].(map[string]interface{})
	if int(s["code"].(float64)) != 200 {
		cerr = &lib.CError{}
		cerr.SetMessage(s["error_message"].(string))
		return
	}

	for _, img_interface := range i["data"].([]interface{}) {
		insta_image := InstaImage{}
		img := img_interface.(map[string]interface{})
		insta_image.Id = img["id"].(string)
		insta_image.Instalink = img["link"].(string)
		imgs := img["images"].(map[string]interface{})
		im_standard := imgs["standard_resolution"].(map[string]interface{})
		insta_image.Standard = Image{
			Width:  int(im_standard["width"].(float64)),
			Height: int(im_standard["height"].(float64)),
			Url:    im_standard["url"].(string),
		}
		im_thumbnail := imgs["thumbnail"].(map[string]interface{})
		insta_image.Standard = Image{
			Width:  int(im_thumbnail["width"].(float64)),
			Height: int(im_thumbnail["height"].(float64)),
			Url:    im_thumbnail["url"].(string),
		}
		im_small := imgs["low_resolution"].(map[string]interface{})
		insta_image.Standard = Image{
			Width:  int(im_small["width"].(float64)),
			Height: int(im_small["height"].(float64)),
			Url:    im_small["url"].(string),
		}

		user_data := img["user"].(map[string]interface{})
		insta_image.User = InstaUser{
			UserName:     user_data["username"].(string),
			FullName:     user_data["full_name"].(string),
			ProfilePhoto: user_data["profile_picture"].(string),
			ID:           user_data["id"].(string),
		}
		images = append(images, insta_image)
	}

	return
}
