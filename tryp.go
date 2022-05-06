package tryp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// clientInterface denotes functions of the client
type clientInterface interface {
	encodeUri(Request) string
	Get(Request) Response
}

// Client stores global config
type Client struct {
	clientInterface
	key string
}

// NewClient from key
func NewClient(key string) (*Client, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("no key")
	}

	return &Client{key: key}, nil
}

// encodeUri from Client Request
func (c *Client) encodeUri(r Request) string {
	uri := url.Values{}

	uri.Set("key", c.key)

	// send | delimited origins and destinations
	uri.Set("origins", strings.Join(r.Origins, "|"))
	uri.Set("destinations", strings.Join(r.Destinations, "|"))

	// only set departure_time if non-zero value
	if !r.DepartureTime.IsZero() {
		time := fmt.Sprintf("%d", r.DepartureTime.Unix())
		uri.Set("departure_time", time)
	}

	// only set arrival_time if non-zero value
	if !r.ArrivalTime.IsZero() {
		time := fmt.Sprintf("%d", r.ArrivalTime.Unix())
		uri.Set("arrival_time", time)
	}

	// set strings
	if r.Avoid != "" {
		uri.Set("avoid", r.Avoid)
	}
	if r.Units != "" {
		uri.Set("units", r.Units)
	}
	if r.Language != "" {
		uri.Set("language", r.Language)
	}
	if r.Mode != "" {
		uri.Set("mode", r.Mode)
	}
	if r.Region != "" {
		uri.Set("region", r.Region)
	}
	if r.TrafficModel != "" {
		uri.Set("traffic_model", r.TrafficModel)
	}
	if r.TransitMode != "" {
		uri.Set("transit_mode", r.TransitMode)
	}
	if r.TransitRoutingPreference != "" {
		uri.Set("transit_routing_preference", r.TransitRoutingPreference)
	}

	return uri.Encode()
}

// Get distance matrix via HTTP request
func (c *Client) Get(r Request) (*Response, error) {
	// encode and build url
	query := c.encodeUri(r)
	url := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?%s", query)
	method := "GET"

	// build and make request
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// read body and convert json to Response
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
