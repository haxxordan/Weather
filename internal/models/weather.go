package models

type Weather struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Conditions  string  `json:"conditions"`
}
type WeatherResponse struct {
	Weather []Weather `json:"weather"` // List of weather conditions
	Main    Main      `json:"main"`    // Main weather data
}

// Main represents main weather data such as temperature and humidity.
type Main struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}
