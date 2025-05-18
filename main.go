package main

import (
	"Weather/models"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	latitude  = flag.Float64("lat", 0.0, "Latitude")
	longitude = flag.Float64("lon", 0.0, "Longitude")
)

func main() {
	flag.Parse()

	if *latitude == 0.0 || *longitude == 0.0 {
		fmt.Println("Usage: Weather -lat=<latitude> -lon=<longitude>")
		os.Exit(1)
	}

	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&daily=temperature_2m_max,temperature_2m_min&current=temperature_2m,relative_humidity_2m,apparent_temperature,rain,wind_speed_10m&timezone=America/Los_Angeles&forecast_days=1&wind_speed_unit=mph&temperature_unit=fahrenheit&precipitation_unit=inch",
		*latitude, *longitude,
	)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("API returned status:", resp.Status)
		os.Exit(1)
	}

	var weather models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		fmt.Println("Error decoding response:", err)
		os.Exit(1)
	}

	fmt.Printf("Current temperature: %.1f°F\n", weather.Current.Temperature)
	fmt.Printf("Relative humidity: %.1f%%\n", weather.Current.RelativeHumidity)
	fmt.Printf("Apparent temperature: %.1f°F\n", weather.Current.ApparentTemperature)
	fmt.Printf("Rain: %.1f inches\n", weather.Current.Rain)
	fmt.Printf("Wind speed: %.1f mph\n", weather.Current.WindSpeed)
}
