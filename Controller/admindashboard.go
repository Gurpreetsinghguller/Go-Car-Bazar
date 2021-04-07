package Controller

import (
	"encoding/json"
	"main/Models"
	"main/Utilities"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte(Utilities.GetEnv("SESSION_KEY")))

func AdminDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Unauthorized accesss"))
		return
	}
	session, err := store.Get(r, "session-name")
	if err != nil {
		var err Models.Error
		err = Utilities.SetError(err, "Failed to generate token")
		Utilities.SetError(err, "Already in Session")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "dashboard.html", nil)
}
