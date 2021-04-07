package Database

import (
	"fmt"
	"main/Models"
	"main/Utilities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConn() *gorm.DB {
	connection, err := gorm.Open(postgres.Open("postgres://admin:1234@localhost/CarAdmin?sslmode=disable"), &gorm.Config{})
	Utilities.CheckErr(err)
	sqldb, err := connection.DB()
	Utilities.CheckErr(err)
	err = sqldb.Ping()
	Utilities.CheckErr(err)
	fmt.Println("connected to database")
	return connection
}
func Close(connection *gorm.DB) {
	sqldb, err := connection.DB()
	Utilities.CheckErr(err)
	sqldb.Close()
}

func AutoMigrateAdmin() {
	db := DBConn()
	db.AutoMigrate(&Models.SignupAdmin{})
}
func AutoMigrateCustomer() {
	db := DBConn()
	db.AutoMigrate(&Models.SignupCustomer{})
}
