package Controller

import (
	"main/Utilities"
	"net/http"
)

func SalesDashboard(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Role") != "Admin" {
		w.Write([]byte("Unauthorized accesss"))
		return
	}
	tmpl := Utilities.GetTemplate()
	tmpl.ExecuteTemplate(w, "salesdashboard.html", nil)
}
