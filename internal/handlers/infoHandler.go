package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

// send GET request for basic information of a specific country
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// extracting the iso2 from the url
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || len(parts[2]) != 2 {
		http.Error(w, "Invalid or missing ISO2 country code", http.StatusBadRequest)
		return
	}
	isoCode := strings.ToUpper(parts[2])
	log.Printf("Requested ISO2 Code: %s", isoCode)

	// extracting city limit
	limitParam := r.URL.Query().Get("limit")
	limit := -1 // default is no limit
	if limitParam != "" {
		if parsedLimit, err := strconv.Atoi(limitParam); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		} else {
			http.Error(w, "Invalid limit value", http.StatusBadRequest)
			return
		}
	}

	// fetching data from RestCountries
	restCountriesURL := fmt.Sprintf("%salpha/%s", constants.RESTCOUNTRIES_API_URL, isoCode)
	log.Printf("Fetching data from REST Countries API: %s", restCountriesURL)

	// error handling for fetching
	restCountriesResp, err := http.Get(restCountriesURL)
	if err != nil {
		log.Printf("Error fetching REST Countries API: %v", err)
		http.Error(w, "Error fetching country data", http.StatusInternalServerError)
		return
	}
	defer restCountriesResp.Body.Close()

	// decoding response from RestCountries
	var restCountries []constants.Country
	if err := json.NewDecoder(restCountriesResp.Body).Decode(&restCountries); err != nil {
		log.Printf("Error decoding REST Countries API response: %v", err)
		http.Error(w, "Error processing country data", http.StatusInternalServerError)
		return
	}

	if len(restCountries) == 0 {
		http.Error(w, "No country data available", http.StatusNotFound)
		return
	}

	// fetching data from CountriesNow
	countriesNowURL := fmt.Sprintf("%scountries/info?returns=iso2,cities", constants.COUNTRIESNOW_API_URL)
	log.Printf("Fetching data from CountriesNow API: %s", countriesNowURL)

	// error handling for fetching
	countriesNowResp, err := http.Get(countriesNowURL)
	if err != nil {
		log.Printf("Error fetching CountriesNow API: %v", err)
		http.Error(w, "Error fetching cities data", http.StatusInternalServerError)
		return
	}
	defer countriesNowResp.Body.Close()

	// decoding response from CountriesNow
	var cnResponse struct {
		Data []struct {
			ISO2   string   `json:"iso2"`
			Cities []string `json:"cities"`
		} `json:"data"`
	}

	// decoding response from CountriesNow
	if err := json.NewDecoder(countriesNowResp.Body).Decode(&cnResponse); err != nil {
		log.Printf("Error decoding CountriesNow API response: %v", err)
		http.Error(w, "Error processing cities data", http.StatusInternalServerError)
		return
	}

	// finding matching country
	var cities []string
	for _, country := range cnResponse.Data {
		if strings.ToUpper(country.ISO2) == isoCode {
			cities = country.Cities
			break
		}
	}

	// applying limit
	if limit > 0 && len(cities) > limit {
		cities = cities[:limit]
	}

	// combining data
	country := restCountries[0]
	country.Cities = cities

	// json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(country); err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
