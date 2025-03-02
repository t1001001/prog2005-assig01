package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

func TestStatusHandler(t *testing.T) {
	// defining test response
	expectedResponse := constants.Status{
		CountriesNowApi:  "200 OK",
		RestCountriesApi: "200 OK",
		Version:          "V1",
		Uptime:           int(time.Since(startTime).Seconds()),
	}

	// creating test http request
	req, err := http.NewRequest("GET", "/status", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// recording response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)

	// serving the request
	handler.ServeHTTP(rr, req)

	// checking if response is OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, status)
	}

	// checking for correct content type
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Expected Content-Type 'application/json', but got %s", ct)
	}

	// checking if response body is correct
	var response constants.Status
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error unmarshalling response body: %v", err)
	}

	// comparing responses
	if response.CountriesNowApi != expectedResponse.CountriesNowApi {
		t.Errorf("Expected CountriesNowApi status %s, but got %s", expectedResponse.CountriesNowApi, response.CountriesNowApi)
	}
	if response.RestCountriesApi != expectedResponse.RestCountriesApi {
		t.Errorf("Expected RestCountriesApi status %s, but got %s", expectedResponse.RestCountriesApi, response.RestCountriesApi)
	}
	if response.Version != expectedResponse.Version {
		t.Errorf("Expected version %s, but got %s", expectedResponse.Version, response.Version)
	}
	if response.Uptime < 0 {
		t.Errorf("Expected uptime to be a positive value, but got %d", response.Uptime)
	}
}
