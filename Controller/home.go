package Controller

import (
	"main/Utilities"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {

	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "index.html", nil)
}
