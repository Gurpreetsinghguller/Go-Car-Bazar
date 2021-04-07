package main

import (
	"flag"
	"main/Database"
	"main/Routes"
)

func main() {
	flagval := flag.Bool("migrate", false, "Insert data to database")
	flag.Parse()
	if *flagval {
		Database.AutoMigrateAdmin()
	}
	Routes.CreateRouter()
	Routes.InitializeRouters()
	Routes.StartServer()

}
