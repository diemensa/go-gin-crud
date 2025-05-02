package tests

import (
	"encoding/json"
	"go-gin-crud/models"
	"go-gin-crud/setup"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestBookControllers(t *testing.T) {
	mode := "test"

	db := setup.ConnectToDB(mode)
	router := setup.SetupRouters(mode, db)

	want := models.Book{
		Title:  "1984",
		Author: "Orwell",
		Genre:  "Dystopia",
	}

	body := `{"title": "1984",
			  "author": "Orwell",
			  "genre": "Dystopia"}`

	t.Run("AddBook", func(t *testing.T) {

		w := httptest.NewRecorder()
		req, err := http.NewRequest("POST", "/book", strings.NewReader(body))

		if err != nil {
			t.Fatal("Error sending POST request")
		}

		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Fatal("Error adding book")
		}
	})

	t.Run("GetBookByID", func(t *testing.T) {

		var got models.Book
		w := httptest.NewRecorder()

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

		w := httptest.NewRecorder()

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
