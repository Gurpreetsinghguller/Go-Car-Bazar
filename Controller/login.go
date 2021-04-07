package Controller

import (
	"main/Utilities"
	"net/http"
)

func Signin(w http.ResponseWriter, r *http.Request) {
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "signin.html", nil)
}
