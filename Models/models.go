package Models

import "gorm.io/gorm"

type SignupAdmin struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobilenumber"`
	City         string `json:"city"`
	Role         string `json:"role"`
}
type SignupCustomer struct {
	gorm.Model
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber string `json:"mobilenumber"`
	City         string `json:"city"`
}
