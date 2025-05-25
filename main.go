package main

import (
	"Weather/models"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	latitude  = flag.Float64("lat", 0.0, "Latitude")
	longitude = flag.Float64("lon", 0.0, "Longitude")
)

// weatherIcons maps weather status codes to their corresponding icons
var weatherIcons = map[string]string{
	"sunnyDay":         "󰖙",
	"clearNight":       "󰖔",
	"cloudyFoggyDay":   "",
	"cloudyFoggyNight": "",
	"rainyDay":         "",
	"rainyNight":       "",
	"snowyIcyDay":      "",
	"snowyIcyNight":    "",
	"severe":           "",
	"default":          "",
}

var wmoToIconKey = map[int]string{
	0:  "sunnyDay",       // Clear sky
	1:  "sunnyDay",       // Mainly clear
	2:  "cloudyFoggyDay", // Partly cloudy
	3:  "cloudyFoggyDay", // Overcast
	45: "cloudyFoggyDay", // Fog
	48: "cloudyFoggyDay", // Depositing rime fog
	51: "rainyDay",       // Drizzle: Light
	53: "rainyDay",       // Drizzle: Moderate
	55: "rainyDay",       // Drizzle: Dense
	56: "rainyDay",       // Freezing Drizzle: Light
	57: "rainyDay",       // Freezing Drizzle: Dense
	61: "rainyDay",       // Rain: Slight
	63: "rainyDay",       // Rain: Moderate
	65: "rainyDay",       // Rain: Heavy
	66: "rainyDay",       // Freezing Rain: Light
	67: "rainyDay",       // Freezing Rain: Heavy
	71: "snowyIcyDay",    // Snow fall: Slight
	73: "snowyIcyDay",    // Snow fall: Moderate
	75: "snowyIcyDay",    // Snow fall: Heavy
	77: "snowyIcyDay",    // Snow grains
	80: "rainyDay",       // Rain showers: Slight
	81: "rainyDay",       // Rain showers: Moderate
	82: "rainyDay",       // Rain showers: Violent
	85: "snowyIcyDay",    // Snow showers: Slight
	86: "snowyIcyDay",    // Snow showers: Heavy
	95: "severe",         // Thunderstorm: Slight/Moderate
	96: "severe",         // Thunderstorm with slight hail
	99: "severe",         // Thunderstorm with heavy hail
}

func parseFloat(value string) float64 {
	result, _ := strconv.ParseFloat(value, 64)
	return result
}

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

	// Get WMO code from API response
	wmoCode := weather.Current.WeatherCode // Adjust field name as per your models.WeatherResponse

	// Map WMO code to icon key
	iconKey := wmoToIconKey[wmoCode]
	if iconKey == "" {
		iconKey = "default"
	}
	icon := weatherIcons[iconKey]

	// Use icon and status as needed
	fmt.Printf("Current temperature: %.1f°F %s\n", weather.Current.Temperature, icon)

	temp := weather.Current.Temperature

	// Current status phrase
	// status := doc.Find("div[data-testid='wxPhrase']").Text()
	// if len(status) > 17 {
	// 	status = status[:16] + ".."
	// }

	// // Status code
	// statusCode := strings.Split(doc.Find("#regionHeader").AttrOr("class", ""), " ")[2]
	// statusCode = strings.Split(statusCode, "-")[2]

	// // Status icon
	// icon = weatherIcons[statusCode]
	// if icon == "" {
	// 	icon = weatherIcons["default"]
	// }

	// Temperature feels like
	tempFeelFloat := weather.Current.ApparentTemperature
	tempFeel := strconv.FormatFloat(tempFeelFloat, 'f', 0, 64)
	tempFeelText := "Feels like " + tempFeel

	// Min-max temperature
	tempMin := weather.Daily.TempMin[0]
	tempMax := weather.Daily.TempMax[0]
	tempMinMax := fmt.Sprintf("  %s\t\t  %s", tempMin, tempMax)

	// Wind speed
	windSpeedFloat := weather.Current.WindSpeed
	windSpeed := strconv.FormatFloat(windSpeedFloat, 'f', 0, 64)
	windText := "  " + windSpeed

	// Humidity
	humidityFloat := weather.Current.RelativeHumidity
	humidity := strconv.FormatFloat(humidityFloat, 'f', 0, 64)
	humidityText := "  " + humidity

	// Hourly rain prediction
	prediction := doc.Find("section[aria-label='Hourly Forecast'] div[data-testid='SegmentPrecipPercentage'] > span").Text()
	prediction = strings.Replace(prediction, "Chance of Rain", "", -1)
	if len(prediction) > 0 {
		prediction = "\n\n (hourly) " + prediction
	}

	// Tooltip text
	tooltipText := fmt.Sprintf("\t\t%s\t\t\n%s\n%s\n%s\n\n%s\n%s\n%s\tAQI %s\n<i> %s</i>",
		"<span size=\"xx-large\">"+temp+"</span>",
		"<big> "+icon+"</big>",
		"<b>"+status+"</b>",
		"<small>"+tempFeelText+"</small>",
		"<b>"+tempMinMax+"</b>",
		windText+"\t"+humidityText,
		visibilityText,
		airQualityIndex,
		prediction,
	)

	// Print waybar module data
	outData := map[string]string{
		"text":    icon + "  " + temp,
		"alt":     status,
		"tooltip": tooltipText,
		"class":   statusCode,
	}
	outDataJSON, _ := json.Marshal(outData)
	fmt.Println(string(outDataJSON))

	simpleWeather := fmt.Sprintf("%s  %s\n  %s (%s)\n%s \n%s \n%s AQI%s\n",
		icon, status, temp, tempFeelText, windText, humidityText, visibilityText, airQualityIndex)

	// Write to cache
	err = os.WriteFile(os.ExpandEnv("$HOME/.cache/.weather_cache"), []byte(simpleWeather), 0644)
	if err != nil {
		fmt.Println("Error writing to cache:", err)
	}
}
