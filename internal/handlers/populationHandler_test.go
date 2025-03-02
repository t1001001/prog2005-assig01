package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

func TestPopulationHandler(t *testing.T) {
	// defining test response
	expectedResponse := constants.Population{
		Country: "Germany",
		Mean:    80825912,
		Values: []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		}{
			{Year: 2010, Value: 81776930},
			{Year: 2011, Value: 80274983},
			{Year: 2012, Value: 80425823},
		},
	}

	// creating test http request
	req, err := http.NewRequest("GET", "/population/de?limit=2010-2012", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// recording response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PopulationHandler)

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
	var response constants.Population
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error unmarshalling response body: %v", err)
	}

	// comparing responses
	if response.Country != expectedResponse.Country {
		t.Errorf("Expected country %s, but got %s", expectedResponse.Country, response.Country)
	}
	if response.Mean != expectedResponse.Mean {
		t.Errorf("Expected mean population %d, but got %d", expectedResponse.Mean, response.Mean)
	}
	if len(response.Values) != len(expectedResponse.Values) {
		t.Errorf("Expected %d population values, but got %d", len(expectedResponse.Values), len(response.Values))
	}
	for i, expected := range expectedResponse.Values {
		if response.Values[i].Year != expected.Year {
			t.Errorf("Expected year %d, but got %d", expected.Year, response.Values[i].Year)
		}
		if response.Values[i].Value != expected.Value {
			t.Errorf("Expected value %d, but got %d", expected.Value, response.Values[i].Value)
		}
	}
}
