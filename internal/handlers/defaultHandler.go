package handlers

import (
	_ "embed"
	"html/template"
	"log"
	"net/http"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

//go:embed templates/index.html
var indexHTML string

// default page of the service
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// debugging
	log.Println("Using embedded template")

	// parsing the template
	tmpl, err := template.New("index").Parse(indexHTML)
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	// template data
	data := struct {
		InfoPath       string
		PopulationPath string
		StatusPath     string
	}{
		InfoPath:       constants.INFO_PATH,
		PopulationPath: constants.POPULATION_PATH,
		StatusPath:     constants.STATUS_PATH,
	}

	// executing the template with the data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
