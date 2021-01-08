package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var books []Book

// Book struct
type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Isbn   string  `json:"isbn"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(Book{})
}

func main() {
	path := os.Getenv("APIPATH")
	r := mux.NewRouter()

	book1 := Book{ID: "1", Title: "Book One", Isbn: "1221", Author: &Author{Firstname: "Vivek", Lastname: "Singh"}}
	book2 := Book{ID: "2", Title: "Book Two", Isbn: "2211", Author: &Author{Firstname: "Abhash", Lastname: "Kumar"}}

	books = append(books, book1)
	books = append(books, book2)

	r.HandleFunc(path+"/", getBooks).Methods("GET")
	r.HandleFunc(path+"/{id}", getBook).Methods("GET")

	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
