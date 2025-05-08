package setup

import (
	"github.com/gin-gonic/gin"
	"go-gin-crud/controllers"
	"gorm.io/gorm"
)

func Routers(mode string, db *gorm.DB) *gin.Engine {

	var router *gin.Engine
	switch {
	case mode == "test":
		gin.SetMode(gin.TestMode)
		router = gin.New()
		router.Use(gin.Recovery())
	default:
		gin.SetMode(gin.ReleaseMode)
		router = gin.Default()
	}

	router.POST("/book", controllers.AddBook(db))
	router.PUT("/book/:id", controllers.UpdateBook(db))
	router.GET("/book/:id", controllers.GetBookByID(db))
	router.GET("/book", controllers.GetAllBooks(db))

	return router
}
