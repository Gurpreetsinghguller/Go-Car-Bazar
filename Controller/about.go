package Controller

import (
	"main/Utilities"
	"net/http"
)

func About(w http.ResponseWriter, r *http.Request) {
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "about-us.html", nil)
}
