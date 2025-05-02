package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-gin-crud/models"
	"gorm.io/gorm"
	"net/http"
)

var validate *validator.Validate = validator.New()

func GetBookByID(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		id := c.Param("id")

		var book models.Book

		res := db.First(&book, id)

		if res.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Book with this ID doesn't exist",
			})
			return
		}

		c.JSON(http.StatusOK, &book)
	}
}

func GetAllBooks(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var books []models.Book

		res := db.Find(&books)

		if res.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Something's wrong",
			})
			return
		} else if len(books) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"Error": "Library is empty :(",
			})
			return
		}
		c.JSON(http.StatusOK, books)
	}

}

func AddBook(db *gorm.DB) func(*gin.Context) {

	return func(c *gin.Context) {
		var request models.Book

		err := c.ShouldBindJSON(&request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Check the types of what you send",
			})
			return
		}

		if err := validate.Struct(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Couldn't validate the structure",
			})
			return
		}

		request.ID = 0 // защита от попытки передать айди в запросе

		res := db.Create(&request)

		if res.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": res.Error.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"Success": fmt.Sprintf("Book \"%s\" added successfuly!", request.Title),
		})
	}
}
