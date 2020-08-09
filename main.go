package main

import (
	"integration/database/mysql"
	"integration/router"
)

func main() {

	db := mysql.ConnectDB()

	defer db.Close()

	router.Start()
}