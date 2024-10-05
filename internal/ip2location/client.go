package ip2location

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	APIKey  string
	baseURL string
	Client  *http.Client
}

type LocationResponse struct {
	IP          string  `json:"ip"`
	CountryCode string  `json:"country_code"`
	Country     string  `json:"country_name"`
	Region      string  `json:"region_name"`
	City        string  `json:"city_name"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	ZipCode     string  `json:"zip_code"`
	TimeZone    string  `json:"time_zone"`
}

func NewClient(apiKey string, opts ...func(*Client)) *Client {
	client := &Client{
		APIKey:  apiKey,
		baseURL: "https://api.ip2location.io",
		Client:  &http.Client{},
	}

	// apply the variadic options
	for _, opt := range opts {
		opt(client)
	}

	return client
}

func (c *Client) Get() (LocationResponse, error) {
	// Construct ip2Location url
	requestUrl := fmt.Sprintf("%s/?token=%s&format=json", c.baseURL, c.APIKey)

	// use http.Client to fetch data
	res, err := c.Client.Get(requestUrl)

	if err != nil {
		return LocationResponse{}, fmt.Errorf("error fetching location: %w", err)
	}

	defer res.Body.Close()

	// read res which is a *http.Response type
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error reading reponse body: %w", err)
	}

	var locationResponse LocationResponse
	err = json.Unmarshal(body, &locationResponse)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error parsing response body: %w", err)
	}

	return locationResponse, nil
}
