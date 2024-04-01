package main

import (
	"example.com/crud_api_tutorial/db"
	"example.com/crud_api_tutorial/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
