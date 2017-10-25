package wunderground

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
)

func TestGetAstronomyUrlBadCityState(t *testing.T) {
	_, err := GetAstronomyUrl("")

	if err == nil {
		t.Errorf("GetAstronomyUrl() should have returned an error related to bad city state value")
	}
}

func TestGetAstronomyUrlMissingEnvironmentVariable(t *testing.T) {
	os.Setenv(WUNDERGROUND_API_KEY, "")

	_, err := GetAstronomyUrl("Asheville, NC")

	if err == nil {
		t.Errorf("GetAstronomyUrl() should have returned an error about missing environment variable")
	}
}

func TestGetAstronomyUrlReturnsUrl(t *testing.T) {
	fakeApiKey := "foobar"
	city := "Asheville"
	state := "NC"

	os.Setenv(WUNDERGROUND_API_KEY, fakeApiKey)

	url, _ := GetAstronomyUrl(fmt.Sprintf("%s, %s", city, state))

	expectedUrl := fmt.Sprintf("http://api.wunderground.com/api/%s/astronomy/q/%s/%s.json", fakeApiKey, state, city)

	if url != expectedUrl {
		t.Errorf("GetAstronomyUrl() should have returned the expected url")
	}
}

func TestGetAstronomyDataHandlesErrors(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	_, err := GetAstronomyData(ts.URL)

	if err == nil {
		t.Errorf("GetAstronomyData() should return the response error")
	}
}

func TestGetAstronomyDataReturnsAstronomyData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, `{"fake": data}`)
	}))
	defer ts.Close()

	astronomy, _ := GetAstronomyData(ts.URL)

	if reflect.TypeOf(astronomy) != reflect.TypeOf((*Astronomy)(nil)).Elem() {
		t.Errorf("GetAstronomyData() should return an Astronomy struct when there is no error")
	}
}
