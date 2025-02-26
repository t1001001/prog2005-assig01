package constants

// general country information
type Country struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       struct {
		Flag string `json:"png"`
	} `json:"flags"`
	Capital []string `json:"capital"`
	Cities  []string `json:"cities"`
}

// population of a country
type Population struct {
	Country string `json:"country"`
	Mean    int    `json:"mean"`
	Values  []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	} `json:"values"`
}

// api status
type Status struct {
	CountriesNowApi  string `json:"countriesNowApi"`
	RestCountriesApi string `json:"restCountriesApi"`
	Version          string `json:"version"`
	Uptime           int    `json:"uptime"`
}
