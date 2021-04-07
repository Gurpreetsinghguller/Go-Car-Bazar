package Controller

import (
	"main/Utilities"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "contact.html", nil)
}
