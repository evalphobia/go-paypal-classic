package request

import (
	"net/url"
	"strings"

	"github.com/evalphobia/go-log-wrapper/log"
	"github.com/evalphobia/goreq-wrapper/request"
	"github.com/mitchellh/mapstructure"
)

var _ = log.Nothing

// CallGET sends GET request
func CallGET(url string, params interface{}, result interface{}) error {
	return call("GET", url, params, result)
}

// call sends http request to `url` with `params` and set reqponse to `result`
func call(method, url string, params interface{}, result interface{}) error {
	dsn := request.DSN{
		Method: method,
		Uri:    url,
	}
	b, err := dsn.Param(params).Call()
	if err != nil {
		return err
	}
	body, err := b.ToString()
	if err != nil {
		return err
	}
	m := parseToMap(body)
	return assignFromMap(m, result, "url")
}

// parseToMap converts response string data to map
func parseToMap(str string) map[string]interface{} {
	m := make(map[string]interface{})
	values := strings.Split(str, "&")
	for _, value := range values {
		v := strings.Split(value, "=")
		if len(v) != 2 {
			continue
		}
		m[v[0]] = urlUnescape(v[1])
	}
	return m
}

// assignFromMap set data from map to struct
func assignFromMap(m interface{}, result interface{}, tagName string) error {
	config := &mapstructure.DecoderConfig{
		Metadata: nil,
		Result:   result,
		TagName:  tagName,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(m)
}

func urlUnescape(v string) string {
	unescaped, _ := url.QueryUnescape(v)
	return unescaped
}
