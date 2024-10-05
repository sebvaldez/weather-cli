package openweather

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

// Create a new openweather client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:  apiKey,
		baseURL: "https://api.openweathermap.org/data/3.0/onecall",
		Client:  &http.Client{},
	}
}

// Return weather data
func (c *Client) Get() (WeatherResponse, error) {
	// Construct openweather url with hard coded lon and lat of   "lat": 45.5122, "lon": -122.6587,
	requestUrl := fmt.Sprintf("%s?lat=45.5122&lon=-122.6587&appid=%s", c.baseURL, c.APIKey)

	res, err := c.Client.Get(requestUrl)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("error fetching weather: %w", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("error reading response body: %w", err)
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(body, &weatherResponse)
	if err != nil {
		return WeatherResponse{}, fmt.Errorf("error parsing response body: %w", err)
	}

	return weatherResponse, nil
}
