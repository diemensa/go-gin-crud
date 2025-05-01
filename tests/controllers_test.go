package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-gin-crud/config"
	"go-gin-crud/controllers"
	"go-gin-crud/models"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func setupTestRouters() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(gin.Recovery())

	router.POST("/book", controllers.AddBook)
	router.GET("/book/:id", controllers.GetBookByID)
	router.GET("/book", controllers.GetAllBooks)

	return router
}

func TestBookControllers(t *testing.T) {
	config.ConnectToDB()
	db := config.DB
	db.AutoMigrate(&models.Book{})

	router := setupTestRouters()
	want := models.Book{
		Title:  "1984",
		Author: "Orwell",
		Genre:  "Dystopia",
	}

	w := httptest.NewRecorder()

	t.Run("AddBook", func(t *testing.T) {

		create_res := db.Create(&want)
		if create_res.Error != nil {
			t.Fatal("Error adding book")

		}
	})

	t.Run("GetBookByID", func(t *testing.T) {
		var got models.Book
		req, err := http.NewRequest("GET", "/book/1", nil)
		if err != nil {
			t.Fatal("Error creating a request")
		}

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %v", w.Result())
		}

		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			t.Fatalf("Couldn't unmarshal JSON into the struct")
		}

		switch {
		case got.Title != want.Title:
			t.Errorf("Expected title %s, got %s", want.Title, got.Title)

		case got.Author != want.Author:
			t.Errorf("Expected author %s, got %s", want.Author, got.Author)

		case got.Genre != want.Genre:
			t.Errorf("Expected genre %s, got %s", want.Genre, got.Genre)

		}
	})

	t.Run("GetAllBooks", func(t *testing.T) {
		var got []models.Book
		want := []models.Book{{
			ID:     1,
			Title:  "1984",
			Author: "Orwell",
			Genre:  "Dystopia",
		}}

		req, err := http.NewRequest("GET", "/book", nil)
		if err != nil {
			t.Fatal("Error creating a request")
		}

		router.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("Expected status 200, got %v", w.Result())
		}

		if err := json.NewDecoder(w.Body).Decode(&got); err != nil {
			t.Fatalf("Couldn't unmarshal JSON into the struct")
		}

		if !reflect.DeepEqual(got, want) {
			log.Fatalf("GetAll doesn't return all the books:\n got %v\n want %v", got, want)

		}
	})

}
