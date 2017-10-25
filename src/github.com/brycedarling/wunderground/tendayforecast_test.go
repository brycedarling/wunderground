package wunderground

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGetTenDayForecastUrlBadCityState(t *testing.T) {
	_, err := GetTenDayForecastUrl("")

	if err == nil {
		t.Errorf("GetTenDayForecastUrl() should have returned an error related to bad city state value")
	}
}

func TestGetTenDayForecastUrlMissingEnvironmentVariable(t *testing.T) {
	os.Setenv(WUNDERGROUND_API_KEY, "")

	_, err := GetTenDayForecastUrl("Asheville, NC")

	if err == nil {
		t.Errorf("GetTenDayForecastUrl() should have returned an error about missing environment variable")
	}
}

func TestGetTenDayForecastUrlReturnsUrl(t *testing.T) {
	fakeApiKey := "foobar"
	city := "Asheville"
	state := "NC"

	os.Setenv(WUNDERGROUND_API_KEY, fakeApiKey)

	url, _ := GetTenDayForecastUrl(fmt.Sprintf("%s, %s", city, state))

	expectedUrl := fmt.Sprintf("http://api.wunderground.com/api/%s/forecast10day/q/%s/%s.json", fakeApiKey, state, city)

	if url != expectedUrl {
		t.Errorf("GetTenDayForecastUrl() should have returned the expected url")
	}
}

func TestGetTenDayForecastDataHandlesErrors(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	_, err := GetTenDayForecastData(ts.URL)

	if err == nil {
		t.Errorf("GetTenDayForecastData() should return the response error")
	}
}

func TestGetTenDayForecastDataReturnsAstronomyData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake": data}`)
	}))
	defer ts.Close()

	forecast, _ := GetTenDayForecastData(ts.URL)

	if reflect.TypeOf(forecast) != reflect.TypeOf((*TenDayForecast)(nil)).Elem() {
		t.Errorf("GetTenDayForecastData() should return a TenDayForecast struct when there is no error")
	}
}
