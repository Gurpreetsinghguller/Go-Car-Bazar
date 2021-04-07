package Utilities

import (
	"encoding/json"
	"fmt"
	"main/Models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func IsAuthorized(handle func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There is an error")
				}
				return mySigningKey, nil
			})
			if err != nil {
				var err Models.Error
				// err.IsError = true
				err = SetError(err, "Your Token has been expired")
				json.NewEncoder(w).Encode(err)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if token.Valid && claims["role"] == "salesperson" && ok {
				r.Header.Set("Role", "salesperson")
				handle(w, r)
			} else if token.Valid && claims["role"] == "admin" && ok {
				r.Header.Set("Role", "admin")
				handle(w, r)
			}
		} else {

			w.Write([]byte("Error in Token"))
		}
	})
}
