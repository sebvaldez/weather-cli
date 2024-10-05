package openweather

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `{
					"lat": 45.5122,
					"lon": -122.6587,
					"timezone": "America/Los_Angeles",
					"timezone_offset": -25200,
					"current": {
						"dt": 1728108509,
						"sunrise": 1728051194,
						"sunset": 1728092690,
						"temp": 285.1,
						"feels_like": 284.78,
						"pressure": 1021,
						"humidity": 93,
						"dew_point": 284,
						"uvi": 0,
						"clouds": 75,
						"visibility": 10000,
						"wind_speed": 1.54,
						"wind_deg": 280,
						"weather": [
							{
								"id": 803,
								"main": "Clouds",
								"description": "broken clouds",
								"icon": "04n"
							}
						]
					},
					"minutely": [
						{
							"dt": 1728112080,
							"precipitation": 0
						}
					],
					"hourly": [
						{
							"dt": 1728108000,
							"temp": 285.1,
							"feels_like": 284.78,
							"pressure": 1021,
							"humidity": 93,
							"dew_point": 284,
							"uvi": 0,
							"clouds": 75,
							"visibility": 10000,
							"wind_speed": 0.52,
							"wind_deg": 273,
							"wind_gust": 0.45,
							"weather": [
								{
									"id": 803,
									"main": "Clouds",
									"description": "broken clouds",
									"icon": "04n"
								}
							],
							"pop": 0
						}
					],
					"daily": [
						{
							"dt": 1728068400,
							"sunrise": 1728051194,
							"sunset": 1728092690,
							"moonrise": 1728058500,
							"moonset": 1728094740,
							"moon_phase": 0.06,
							"summary": "Expect a day of partly cloudy with rain",
							"temp": {
								"day": 284.25,
								"min": 283.26,
								"max": 288.2,
								"night": 285.1,
								"eve": 287.84,
								"morn": 285.23
							},
							"feels_like": {
								"day": 283.87,
								"night": 284.78,
								"eve": 287.48,
								"morn": 283.98
							},
							"pressure": 1015,
							"humidity": 94,
							"dew_point": 282.92,
							"wind_speed": 2.6,
							"wind_deg": 219,
							"wind_gust": 6.71,
							"weather": [
								{
									"id": 502,
									"main": "Rain",
									"description": "heavy intensity rain",
									"icon": "10d"
								}
							],
							"clouds": 100,
							"pop": 1,
							"rain": 12.59,
							"uvi": 1.15
						}
					]}`)
	}))

	defer ts.Close()

	client := NewClient("test-api-key")
	client.Client = ts.Client()
	client.baseURL = ts.URL // Modify the unexported baseURL directly

	weather, err := client.Get()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if weather.Current.Temp != 285.1 {
		t.Errorf("Expected temperature 285.1, got %v", weather.Current.Temp)
	}

	if weather.Daily[0].Summary != "Expect a day of partly cloudy with rain" {
		t.Errorf("Expected summary 'Expect a day of partly cloudy with rain', got %s", weather.Daily[0].Summary)
	}
}
