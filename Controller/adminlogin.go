package Controller

import (
	"encoding/json"
	"fmt"
	"main/Database"
	"main/Models"
	"main/Utilities"
	"net/http"
)

func AdminLogin(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	tmpl := Utilities.GetTemplate()

	if r.Method == "POST" {

		var signin Models.Signin
		json.NewDecoder(r.Body).Decode(&signin)
		fmt.Println(signin)
		//Getting Entry From Database
		var details Models.SignupAdmin

		//Database
		db := Database.DBConn()
		defer Database.Close(db)
		db.Where("email =?", signin.Email).First(&details)

		if details.Email == "" {
			fmt.Println("Error occurred here")
			var err Models.Error
			err = Utilities.SetError(err, "Invalid Email or Password")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

		//Password Decrypt
		chkpass := Utilities.UnHash(signin.Password, details.Password)
		if !chkpass {
			var err Models.Error
			err = Utilities.SetError(err, "Invalid Email or Password")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

		//Token Generation
		token, err := Utilities.GenerateJWT(details.Email, details.Role)
		if err != nil {
			var err Models.Error
			err = Utilities.SetError(err, "Failed to generate token")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

		var signinResponse Models.SigninResponse
		signinResponse.Token = token
		signinResponse.Role = details.Role
		pBytes, err := json.MarshalIndent(signinResponse, " ", " ")
		w.Write(pBytes)
		return
	}
	tmpl.ExecuteTemplate(w, "admin.html", nil)
}
