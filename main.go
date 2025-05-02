package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud/config"
	"go-gin-crud/controllers"
	"go-gin-crud/models"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
	config.ConnectToDB()
	config.DB.AutoMigrate(&models.Book{})
}

func main() {
	router := gin.Default()

	router.POST("/book", controllers.AddBook(config.DB))
	router.GET("/book/:id", controllers.GetBookByID(config.DB))
	router.GET("/book", controllers.GetAllBooks(config.DB))

	router.Run()
}
