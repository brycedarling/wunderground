package wunderground

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

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

func ParseCityState(cityAndState string) (string, string, error) {
	parts := strings.Split(cityAndState, ", ")

	if len(parts) != 2 {
		return "", "", errors.New("Both city and state are required")
	}

	city, state := parts[0], parts[1]

	if len(city) == 0 {
		return "", "", errors.New("City must be a non-empty string")
	}

	if len(state) != 2 {
		return "", "", errors.New("State must be a two character string")
	}

	return city, state, nil
}
