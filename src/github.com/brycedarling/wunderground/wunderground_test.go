package wunderground

import (
	"fmt"
	"os"
	"testing"
)

func TestGetApiKeyFromEnvironment_InvalidApiKey(t *testing.T) {
	os.Setenv(WUNDERGROUND_API_KEY, "")

	_, err := GetApiKeyFromEnvironment()

	expectedError := fmt.Sprintf("%s environment variable must be set", WUNDERGROUND_API_KEY)

	if err == nil || err.Error() != expectedError {
		t.Errorf("GetApiKeyFromEnvironment() should return the environment variable must be set")
	}
}

func TestGetApiKeyFromEnvironment_ReturnsKey(t *testing.T) {
	fakeApiKey := "foobar"

	os.Setenv(WUNDERGROUND_API_KEY, fakeApiKey)

	apiKey, _ := GetApiKeyFromEnvironment()

	if apiKey != fakeApiKey {
		t.Errorf("GetApiKeyFromEnvironment() should return the api key from the environment")
	}
}

func TestGetBaseApiUrl_MissingEnvironmentVariable(t *testing.T) {
	os.Setenv(WUNDERGROUND_API_KEY, "")

	_, err := GetBaseApiUrl()

	expectedError := fmt.Sprintf("%s environment variable must be set", WUNDERGROUND_API_KEY)

	if err == nil || err.Error() != expectedError {
		t.Errorf("GetBaseApiUrl() should return the environment variable must be set")
	}
}

func TestGetBaseApiUrl_ReturnsBaseApiUrl(t *testing.T) {
	fakeApiKey := "foobar"

	os.Setenv(WUNDERGROUND_API_KEY, fakeApiKey)

	baseApiUrl, _ := GetBaseApiUrl()

	expectedBaseApiUrl := fmt.Sprintf("http://api.wunderground.com/api/%s", fakeApiKey)

	if baseApiUrl != expectedBaseApiUrl {
		t.Errorf("GetBaseApiUrl() should return the base api url containing the api key")
	}
}

func TestParseCityState_EmptyString(t *testing.T) {
	_, err := ParseCityState("")

	if err == nil || err.Error() != "Both city and state are required" {
		t.Errorf("ParseCityState() should return an error that it requires both city and state")
	}
}

func TestParseCityState_StringWithoutComma(t *testing.T) {
	_, err := ParseCityState("Not a city state")

	if err == nil || err.Error() != "Both city and state are required" {
		t.Errorf("ParseCityState() should return an error that it requires both city and state")
	}
}

func TestParseCityState_BlankCity(t *testing.T) {
	_, err := ParseCityState(", Blank City")

	if err == nil || err.Error() != "City must be a non-empty string" {
		t.Errorf("ParseCityState() should return an error that city must be a non-empty string")
	}
}

func TestParseCityState_InvalidStateLength(t *testing.T) {
	_, err := ParseCityState("City, InvalidState")

	if err == nil || err.Error() != "State must be a two character string" {
		t.Errorf("ParseCityState() should return an error that state must be two characters")
	}
}

func TestParseCityState_ReturnsCityAndState(t *testing.T) {
	location, _ := ParseCityState("Asheville, NC")

	if location.City != "Asheville" {
		t.Errorf("ParseCityState() should return the city")
	}

	if location.State != "NC" {
		t.Errorf("ParseCityState() should return the state")
	}
}
