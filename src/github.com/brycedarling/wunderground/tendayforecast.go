package wunderground

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type TenDayForecast struct {
	Forecast Forecast `json:"forecast"`
}

type Forecast struct {
	TextForecast   TextForecast   `json:"txt_forecast"`
	SimpleForecast SimpleForecast `json:"simpleforecast"`
}

type TextForecast struct {
	Date         string            `json:"date"`
	ForecastDays []TextForecastDay `json:"forecastday"`
}

type TextForecastDay struct {
	Period             int    `json:"period"`
	Icon               string `json:"icon"`
	IconUrl            string `json:"icon_url"`
	Title              string `json:"title"`
	ForecastText       string `json:"fcttext"`
	ForecastTextMetric string `json:"fcttext_metric"`
	Pop                string `json:"pop"`
}

type SimpleForecast struct {
	ForecastDays []SimpleForecastDay `json:"forecastday"`
}

// TODO: this structure is crazy huge, and implementing the rest is just busy work at this point...
// TODO: skipping for now, implement later if you actually end up needing it
type SimpleForecastDay struct {
	Date Date        `json:"date"`
	High Temperature `json:"high"`
	Low  Temperature `json:"low"`
}

type Date struct {
	Epoch  string `json:"epoch"`
	Pretty string `json:"pretty"`
}

type Temperature struct {
	Fahrenheit string `json:"fahrenheit"`
	Celsius    string `json:"celsius"`
}

func GetTenDayForecastUrl(cityAndState string) (string, error) {
	baseApiUrl, err := GetBaseApiUrl()

	if err != nil {
		return "", err
	}

	location, err := ParseCityState(cityAndState)

	if err != nil {
		return "", err
	}

	city := strings.Replace(location.City, " ", "_", -1)

	url := fmt.Sprintf("%s/forecast10day/q/%s/%s.json", baseApiUrl, location.State, city)

	return url, nil
}

func GetTenDayForecastData(url string) (TenDayForecast, error) {
	var forecast TenDayForecast

	response, err := http.Get(url)

	if err != nil {
		return forecast, err
	}

	err = json.NewDecoder(response.Body).Decode(&forecast)

	return forecast, err
}
