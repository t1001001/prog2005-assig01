package handlers

import (
	"fmt"
	"net/http"

	"github.com/t1001001/prog2005-assignment-01/internal/constants"
)

// default handler when path is empty
func DefaultHandler(w http.ResponseWriter, r *http.Request) {

	// client takes this as html
	w.Header().Set("content-type", "text/html")

	// default message
	output := fmt.Sprintf(
		"Greetings traveler, this is the Country Information Service API.<br>"+
			"There is nothing to see here, so move along. <br>"+
			"<br>"+
			"<a href=\"%s\">%s</a> is used for country details,<br>"+
			"<a href=\"%s\">%s</a> is used for historical population data,<br>"+
			"<a href=\"%s\">%s</a> is used for service diagnostics.",
		constants.INFO_PATH, constants.INFO_PATH,
		constants.POPULATION_PATH, constants.POPULATION_PATH,
		constants.STATUS_PATH, constants.STATUS_PATH,
	)

	// client output
	_, err := fmt.Fprintf(w, "%v", output)

	// in case there are some funny errors
	if err != nil {
		http.Error(w, "Cannot return output", http.StatusInternalServerError)
	}
}
