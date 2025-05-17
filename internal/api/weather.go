package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

const weatherAPIURL = "https://api.weather.gov/points/"

type WeatherResponse struct {
    Properties struct {
        Forecast string `json:"forecast"`
    } `json:"properties"`
}

func GetWeather(latitude, longitude float64) (string, error) {
    url := fmt.Sprintf("%s%f,%f", weatherAPIURL, latitude, longitude)
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch weather data: %s", resp.Status)
    }

    var weatherResponse WeatherResponse
    if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
        return "", err
    }

    return weatherResponse.Properties.Forecast, nil
}