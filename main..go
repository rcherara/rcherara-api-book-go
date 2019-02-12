package main

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

type Book struct {
	// The first comma below is to separate the name ID from the omitempty tag
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Year   string `json:"year,omitempty"`
}

var books []Book

// rcherara-api-book-go
func main() {
	router := mux.NewRouter()

	books = append(books,
		Book{ID: 1, Title: "Spring 5.0 By Example: Grasp the fundamentals of Spring 5.0 to build modern, robust, and scalable Java applications.", Author: "Mr. Claudio Eduardo de Oliveira", Year: "2010"},
		Book{ID: 2, Title: "Network Security with OpenSSL", Author: "Mr. John Viega.", Year: "2099"},
		Book{ID: 3, Title: "Learn AutoCAD!", Author: "Mr. David Martin", Year: "2016"},
		Book{ID: 4, Title: "Ansible for DevOps", Author: "Mr. Jeff Geerling", Year: "2015"},
	)

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

	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get a book")

	params := mux.Vars(r)
	log.Println(params)
	log.Println(reflect.TypeOf(params["id"]))

	i, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
	}
	log.Println(reflect.TypeOf(i))

	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
}
func addBook(w http.ResponseWriter, r *http.Request) {

	log.Println("Adds a Book")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	log.Println(book)
	json.NewEncoder(w).Encode(books)

}
func updateBook(w http.ResponseWriter, r *http.Request) {

	log.Println("Updates a book")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}

	}
	json.NewEncoder(w).Encode(books)

}
func removeBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Println(err)
	}
	log.Println("Removes a book id=", id)
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}

}
