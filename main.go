package main

import (
	"go-gin-crud/setup"
)

func init() {
}

func main() {
	mode := "main"

	db := setup.ConnectToDB(mode)
	router := setup.Routers(mode, db)

	router.Run()
}
