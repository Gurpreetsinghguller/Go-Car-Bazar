package Controller

import (
	"main/Utilities"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "signup.html", nil)
}
