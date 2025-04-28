package config

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("./database/dreambase.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}
}
