package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

// default page of the service
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	// defining the path
	templatePath := filepath.Join("internal", "templates", "index.html")

	// debugging
	absPath, _ := filepath.Abs(templatePath)
	log.Println(absPath)

	// parsing the template
	template, err := template.ParseFiles(templatePath)
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
	err = template.Execute(w, data)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}
