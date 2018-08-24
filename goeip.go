package goeip

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

//Result holds data returned by the API
type Result struct {
	Error      int
	Details    string `json:"Detail"`
	IPAddr     string
	Hostname   string
	City       string
	State      string
	PostalCode string
	Country    Country
	Timezone   string
	Location   Location
	ASN        ASN
}

//Country holds data about a country
type Country struct {
	Name string `json:"Name"`
	Code string
}

//Location holds data about a location
type Location struct {
	Latitude  float64
	Longitude float64
}

//ASN holds data about an ASN
type ASN struct {
	Number       int
	Organization string
}

//Lookup performs a GET request to the API and returns data about the specified hostname/ipaddr
func Lookup(host string) (*Result, error) {
	escapedHost := url.QueryEscape(host)

	url := fmt.Sprintf("https://kealper.com/util/geoip/api.cgi?type=json&ip=%s", escapedHost)

	data := &Result{}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	err = unmarshal(res, data)

	if err == nil {
		if data.Error > 0 {
			return data, fmt.Errorf("goeip: %s", data.Detail)
		}
	}

	return data, err
}

func unmarshal(body *http.Response, target interface{}) error {
	defer body.Body.Close()
	return json.NewDecoder(body.Body).Decode(target)
}
