package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

var startTime = time.Now()

// send GET request to the endpoint to see whether its available
func checkAPIStatus(baseURL string) string {
	var statusURL string
	if strings.Contains(baseURL, "v3.1") {
		statusURL = baseURL + "alpha/de?fields=capital" //checks for germanys capital
		log.Printf("Checking the RestCountries API: %s", statusURL)
	} else if strings.Contains(baseURL, "api/v0.1") {
		statusURL = baseURL + "countries/info?returns=name" // checks for existing countries
		log.Printf("Checking the CountriesNow API: %s", statusURL)
	} else {
		statusURL = baseURL
	}

	// checking whether the APIs are available
	resp, err := http.Get(statusURL)
	if err != nil {
		log.Printf("Error fetching %s: %v", statusURL, err)
		return "Unavailable"
	}
	defer resp.Body.Close()

	return resp.Status
}

// service status
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// ensuring that only GET methods are allowed
	if r.Method != http.MethodGet {
		http.Error(w, "Only the GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// status response
	status := constants.Status{
		CountriesNowApi:  checkAPIStatus(constants.COUNTRIESNOW_API_URL),
		RestCountriesApi: checkAPIStatus(constants.RESTCOUNTRIES_API_URL),
		Version:          constants.VERSION,
		Uptime:           int(time.Since(startTime).Seconds()),
	}

	// setting header
	w.Header().Set("Content-Type", "application/json")

	// encoding
	if err := json.NewEncoder(w).Encode(status); err != nil {
		log.Printf("Error encoding JSON: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
