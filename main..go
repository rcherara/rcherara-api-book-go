package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	// The first comma below is to separate the name ID from the omitempty tag
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Year   string `json:"year,omitempty"`
}

var book []Book

// rcherara-api-book
func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
	log.Println("Hello World")

}

func getBooks(w http.ResponseWriter, r *http.Request) {

	log.Println("Gets books")
}
func getBook(w http.ResponseWriter, r *http.Request) {

	log.Println("Get a book")
}
func addBook(w http.ResponseWriter, r *http.Request) {

	log.Println("Adds a Book")
}
func updateBook(w http.ResponseWriter, r *http.Request) {

	log.Println("Updates a book")
}
func removeBook(w http.ResponseWriter, r *http.Request) {

	log.Println("Removes a book")
}
