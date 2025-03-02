package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

func TestInfoHandler(t *testing.T) {
	// defining test response
	expectedResponse := constants.Country{
		Name: struct {
			Common string `json:"common"`
		}{
			Common: "Germany",
		},
		Population: 83240525,
		Cities: []string{
			"Aach", "Aachen", "Aalen",
		},
		Borders: []string{
			"AUT", "BEL", "CZE", "DNK", "FRA", "LUX", "NLD", "POL", "CHE",
		},
		Flag: struct {
			Flag string `json:"png"`
		}{
			Flag: "https://flagcdn.com/w320/de.png",
		},
	}

	// creating test http request
	req, err := http.NewRequest("GET", "/info/de?limit=3", nil)
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// recording response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(InfoHandler)

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
	var response constants.Country
	if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
		t.Fatalf("Error unmarshalling response body: %v", err)
	}

	// comparing responses
	if response.Name.Common != expectedResponse.Name.Common {
		t.Errorf("Expected country name '%s', but got '%s'", expectedResponse.Name.Common, response.Name.Common)
	}
	if response.Population != expectedResponse.Population {
		t.Errorf("Expected population %d, but got %d", expectedResponse.Population, response.Population)
	}
	if len(response.Cities) != len(expectedResponse.Cities) {
		t.Errorf("Expected %d cities, but got %d", len(expectedResponse.Cities), len(response.Cities))
	}
	for i, city := range expectedResponse.Cities {
		if response.Cities[i] != city {
			t.Errorf("Expected city '%s', but got '%s'", city, response.Cities[i])
		}
	}
	if len(response.Borders) != len(expectedResponse.Borders) {
		t.Errorf("Expected %d borders, but got %d", len(expectedResponse.Borders), len(response.Borders))
	}
	for i, border := range expectedResponse.Borders {
		if response.Borders[i] != border {
			t.Errorf("Expected border '%s', but got '%s'", border, response.Borders[i])
		}
	}
	if response.Flag.Flag != expectedResponse.Flag.Flag {
		t.Errorf("Expected flag URL '%s', but got '%s'", expectedResponse.Flag.Flag, response.Flag.Flag)
	}
}
