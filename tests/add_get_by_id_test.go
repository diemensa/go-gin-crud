package tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go-gin-crud/config"
	"go-gin-crud/controllers"
	"go-gin-crud/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd_and_GetByID(t *testing.T) {
	config.ConnectToDB()
	db := config.DB
	db.AutoMigrate(&models.Book{})

	router := gin.Default()
	router.GET("/book/:id", controllers.GetBookByID)

	want := models.Book{
		Title:  "1984",
		Author: "Orwell",
		Genre:  "Dystopia",
	}

	create_res := db.Create(&want)
	if create_res.Error != nil {
		t.Fatal("Error adding book")

	}
	req, err := http.NewRequest("GET", "/book/1", nil)
	if err != nil {
		t.Fatal("Error creating a request")
	}

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %v", w.Result())
	}

	var got models.Book

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

	db.Exec("DELETE FROM books")

}
