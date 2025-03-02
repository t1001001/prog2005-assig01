package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

// send GET request for population data of a specific country
func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	// extracting the iso2 from the url
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[2] == "" {
		http.Error(w, "Invalid or missing ISO2 country code", http.StatusBadRequest)
		return
	}
	isoCode := strings.ToUpper(parts[2])
	log.Printf("Requested ISO2 Code: %s", isoCode)

	// fetching country name based on iso2 code
	countryName, err := getCountryNameByISO(isoCode)
	if err != nil {
		http.Error(w, "Error fetching country name, ISO2 may not exist", http.StatusInternalServerError)
		return
	}

	// extracting year range
	yearRange := r.URL.Query().Get("limit")
	var startYear, endYear int
	if yearRange != "" {
		years := strings.Split(yearRange, "-")
		if len(years) != 2 {
			http.Error(w, "Invalid limit format, expected YYYY-YYYY", http.StatusBadRequest)
			return
		}
		startYear, err = strconv.Atoi(years[0])
		if err != nil {
			http.Error(w, "Invalid start year", http.StatusBadRequest)
			return
		}
		endYear, err = strconv.Atoi(years[1])
		if err != nil {
			http.Error(w, "Invalid end year", http.StatusBadRequest)
			return
		}
	}

	// fetching data
	populationData, err := getPopulationData()
	if err != nil {
		http.Error(w, "Error fetching population data", http.StatusInternalServerError)
		return
	}

	// finding data based on country name
	var countryPopulationData []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	}

	for _, country := range populationData {
		if strings.EqualFold(country.Country, countryName) {
			countryPopulationData = append(countryPopulationData, country.PopulationCounts...)
			break
		}
	}

	// error handling when theres no data to be found
	if len(countryPopulationData) == 0 {
		http.Error(w, fmt.Sprintf("Population data not found for ISO2 code: %s", isoCode), http.StatusNotFound)
		return
	}

	// filtering data based on year range
	var filteredValues []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	}
	var totalPopulation int
	var count int

	for _, entry := range countryPopulationData {
		if (startYear == 0 || entry.Year >= startYear) && (endYear == 0 || entry.Year <= endYear) {
			filteredValues = append(filteredValues, entry)
			totalPopulation += entry.Value
			count++
		}
	}

	// calculating mean
	var meanPopulation int
	if count > 0 {
		meanPopulation = totalPopulation / count
	}

	// constructing response
	response := constants.Population{
		Country: countryName,
		Mean:    meanPopulation,
		Values:  filteredValues,
	}

	// json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Error encoding response: %s", err)
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// helper function to fetch country name based on the iso2 code
func getCountryNameByISO(isoCode string) (string, error) {
	infoURL := fmt.Sprintf("%scountries/info?returns=iso2", constants.COUNTRIESNOW_API_URL)
	log.Printf("Fetching data from CountriesNow API: %s", infoURL)

	// error handling for fetching
	resp, err := http.Get(infoURL)
	if err != nil {
		log.Printf("Error fetching country info: %s", err)
		return "", err
	}
	defer resp.Body.Close()

	// checking if request was successful
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching country info, status code: %d", resp.StatusCode)
		return "", fmt.Errorf("failed to fetch country info")
	}

	// constructing response
	var apiResponse struct {
		Error bool `json:"error"`
		Data  []struct {
			ISO2 string `json:"iso2"`
			Name string `json:"name"`
		} `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		log.Printf("Error parsing response: %s", err)
		return "", err
	}

	// searching for the country based on the iso2 code
	for _, country := range apiResponse.Data {
		if strings.EqualFold(country.ISO2, isoCode) {
			return country.Name, nil
		}
	}
	return "", fmt.Errorf("country not found for ISO2 code: %s", isoCode)
}

// helper function to get the population data
func getPopulationData() ([]struct {
	Country          string `json:"country"`
	PopulationCounts []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	} `json:"populationCounts"`
}, error) {
	populationURL := fmt.Sprintf("%scountries/population", constants.COUNTRIESNOW_API_URL)
	log.Printf("Fetching data from CountriesNow API: %s", populationURL)
	resp, err := http.Get(populationURL)
	if err != nil {
		log.Printf("Error fetching population data: %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error fetching population data, status code: %d", resp.StatusCode)
		return nil, fmt.Errorf("failed to fetch population data")
	}

	// reading and parsing the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %s", err)
		return nil, err
	}

	var apiResponse struct {
		Error bool `json:"error"`
		Data  []struct {
			Country          string `json:"country"`
			ISO2             string `json:"iso2"`
			PopulationCounts []struct {
				Year  int `json:"year"`
				Value int `json:"value"`
			} `json:"populationCounts"`
		} `json:"data"`
	}

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		log.Printf("Error parsing response: %s", err)
		return nil, err
	}

	// creating a slice that only contains the necessary data
	var populationData []struct {
		Country          string `json:"country"`
		PopulationCounts []struct {
			Year  int `json:"year"`
			Value int `json:"value"`
		} `json:"populationCounts"`
	}

	for _, country := range apiResponse.Data {
		populationData = append(populationData, struct {
			Country          string `json:"country"`
			PopulationCounts []struct {
				Year  int `json:"year"`
				Value int `json:"value"`
			} `json:"populationCounts"`
		}{
			Country:          country.Country,
			PopulationCounts: country.PopulationCounts,
		})
	}

	// returning the modigied data
	return populationData, nil
}
