package setup

import (
	"github.com/glebarez/sqlite"
	"go-gin-crud/models"
	"gorm.io/gorm"
	"log"
)

func ConnectToDB(mode string) *gorm.DB {
	var db *gorm.DB
	var err error

	switch {
	case mode == "test":
		db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	default:
		db, err = gorm.Open(sqlite.Open("dreambase.db"), &gorm.Config{})
	}

	if err != nil {
		log.Fatal("Failed to connect to DB")
	}

	db.AutoMigrate(&models.Book{})

	return db

}
