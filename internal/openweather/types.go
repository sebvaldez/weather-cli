package openweather

// Root of the API response
type WeatherResponse struct {
	Lat            float64         `json:"lat"`
	Lon            float64         `json:"lon"`
	Timezone       string          `json:"timezone"`
	TimezoneOffset int             `json:"timezone_offset"`
	Current        CurrentWeather  `json:"current"`
	Minutely       []Minutely      `json:"minutely"`
	Hourly         []HourlyWeather `json:"hourly"`
	Daily          []DailyWeather  `json:"daily"`
}

// Current weather data
type CurrentWeather struct {
	Dt         int64     `json:"dt"`
	Sunrise    int64     `json:"sunrise"`
	Sunset     int64     `json:"sunset"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	Weather    []Weather `json:"weather"`
}

// Weather represents weather conditions
type Weather struct {
	ID          int    `json:"id"`
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// Minutely forecast
type Minutely struct {
	Dt            int64   `json:"dt"`
	Precipitation float64 `json:"precipitation"`
}

// Hourly forecast
type HourlyWeather struct {
	Dt         int64     `json:"dt"`
	Temp       float64   `json:"temp"`
	FeelsLike  float64   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float64   `json:"dew_point"`
	Uvi        float64   `json:"uvi"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float64   `json:"wind_speed"`
	WindDeg    int       `json:"wind_deg"`
	Weather    []Weather `json:"weather"`
	Pop        float64   `json:"pop"`
}

// Daily forecast
type DailyWeather struct {
	Dt        int64     `json:"dt"`
	Sunrise   int64     `json:"sunrise"`
	Sunset    int64     `json:"sunset"`
	Moonrise  int64     `json:"moonrise"`
	Moonset   int64     `json:"moonset"`
	MoonPhase float64   `json:"moon_phase"`
	Temp      Temp      `json:"temp"`
	FeelsLike FeelsLike `json:"feels_like"`
	Pressure  int       `json:"pressure"`
	Humidity  int       `json:"humidity"`
	DewPoint  float64   `json:"dew_point"`
	WindSpeed float64   `json:"wind_speed"`
	WindDeg   int       `json:"wind_deg"`
	Weather   []Weather `json:"weather"`
	Clouds    int       `json:"clouds"`
	Pop       float64   `json:"pop"`
	Rain      float64   `json:"rain,omitempty"`
	Summary   string    `json:"summary,omitempty"`
}

// Temperature details in daily forecast
type Temp struct {
	Day   float64 `json:"day"`
	Min   float64 `json:"min"`
	Max   float64 `json:"max"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}

// Feels like temperature in daily forecast
type FeelsLike struct {
	Day   float64 `json:"day"`
	Night float64 `json:"night"`
	Eve   float64 `json:"eve"`
	Morn  float64 `json:"morn"`
}
