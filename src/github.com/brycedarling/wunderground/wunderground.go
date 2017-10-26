package wunderground

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

type Location struct {
	City  string
	State string
}

const (
	WUNDERGROUND_API_KEY = "WUNDERGROUND_API_KEY"
)

func GetApiKeyFromEnvironment() (string, error) {
	apiKey := os.Getenv(WUNDERGROUND_API_KEY)

	if len(apiKey) == 0 {
		return "", errors.New(fmt.Sprintf("%s environment variable must be set", WUNDERGROUND_API_KEY))
	}

	return apiKey, nil
}

func GetBaseApiUrl() (string, error) {
	apiKey, err := GetApiKeyFromEnvironment()

	if err != nil {
		return "", err
	}

	baseApiUrl := fmt.Sprintf("http://api.wunderground.com/api/%s", apiKey)

	return baseApiUrl, nil
}

func ParseCityState(cityAndState string) (Location, error) {
	var location Location

	parts := strings.Split(cityAndState, ", ")

	if len(parts) != 2 {
		return location, errors.New("Both city and state are required")
	}

	location.City = parts[0]
	location.State = parts[1]

	if len(location.City) == 0 {
		return location, errors.New("City must be a non-empty string")
	}

	if len(location.State) != 2 {
		return location, errors.New("State must be a two character string")
	}

	return location, nil
}
