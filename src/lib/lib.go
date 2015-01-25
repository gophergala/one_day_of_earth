package lib

import (
	"bytes"
	"net/url"
	"strings"
	"time"
)

type CError struct {
	message string
}

func (e *CError) Message() string {
	return e.message
}

func (e *CError) SetMessage(s string) {
	e.message = s
}

func GenerateURL(api_url string, params map[string]string) string {
	var buffer bytes.Buffer //Using Buffer because it should be faster
	buffer.WriteString(api_url)
	i := strings.Index(api_url, "?")
	if i == -1 {
		buffer.WriteString("?")
	} else if i < (len(api_url) - 1) {
		buffer.WriteString("?")
	}
	buffer.WriteString(URLEncodeParams(params))
	return buffer.String()
}

func URLEncodeParams(params map[string]string) string {
	val := UrlValues(params)
	return val.Encode()
}

func UrlValues(params map[string]string) url.Values {
	val := url.Values{}
	for k, v := range params {
		val.Add(k, v)
	}
	return val
}

func YesterdayTime() time.Time {
	now := time.Now().UTC()
	return time.Date(now.Year(), now.Month(), (now.Day() - 1), 0, 0, 0, 0, time.UTC)
}
