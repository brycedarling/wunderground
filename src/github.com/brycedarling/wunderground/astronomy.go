package wunderground

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Astronomy struct {
	MoonPhase MoonPhase `json:"moon_phase"`
	SunPhase  SunPhase  `json:"sun_phase"`
}

type MoonPhase struct {
	PercentIlluminated string `json:"percentIlluminated"`
	AgeOfMoon          string `json:"ageOfMoon"`
	PhaseOfMoon        string `json:"phaseofMoon"`
	Hemisphere         string `json:"hemisphere"`
	CurrentTime        Time   `json:"current_time"`
	Sunrise            Time   `json:"sunrise"`
	Sunset             Time   `json:"sunset"`
	Moonrise           Time   `json:"moonrise"`
	Moonset            Time   `json:"moonset"`
}

type SunPhase struct {
	Sunrise Time `json:"sunrise"`
	Sunset  Time `json:"sunset"`
}

type Time struct {
	Hour   string `json:"hour"`
	Minute string `json:"minute"`
}

func GetAstronomyUrl(cityAndState string) (string, error) {
	baseApiUrl, err := GetBaseApiUrl()

	if err != nil {
		return "", err
	}

	location, err := ParseCityState(cityAndState)

	if err != nil {
		return "", err
	}

	city := strings.Replace(location.City, " ", "_", -1)

	url := fmt.Sprintf("%s/astronomy/q/%s/%s.json", baseApiUrl, location.State, city)

	return url, nil
}

func GetAstronomyData(url string) (Astronomy, error) {
	var astronomy Astronomy

	response, err := http.Get(url)

	if err != nil {
		return astronomy, err
	}

	err = json.NewDecoder(response.Body).Decode(&astronomy)

	return astronomy, err
}
