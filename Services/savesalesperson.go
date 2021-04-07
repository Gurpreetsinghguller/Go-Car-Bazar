package Services

import (
	"fmt"
	"log"
	"main/Models"
	"main/Utilities"
	"net/http"
)

func SaveSalesPerson() {
	if r.Method == "POST" {
		pass := r.FormValue("password")
		fmt.Println(pass)
		pass, err := Utilities.HashPassword(pass)
		if err != nil {
			log.Fatalln("Error in Password Hashing")
		}
		var details Models.SignupAdmin
		details.Name = r.FormValue("name")
		details.Email = r.FormValue("email")
		details.Password = pass
		details.MobileNumber = r.FormValue("mobilenumber")
		details.City = r.FormValue("city")

		http.Redirect(w, r, "/signin", 301)
	}
}
