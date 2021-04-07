package Controller

import (
	"main/Utilities"
	"net/http"
)

func Reset(w http.ResponseWriter, r *http.Request) {
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "reset.html", nil)
}
