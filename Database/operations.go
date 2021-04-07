package Database

import "main/Models"

func SaveSalesPerson(details Models.SignupCustomer) {
	db := DBConn()
	defer Close(db)
	db.Create(&details)
}
