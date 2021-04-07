package Controller

import (
	"main/Utilities"
	"net/http"
)

func Cars(w http.ResponseWriter, r *http.Request) {
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "cars.html", nil)
}
