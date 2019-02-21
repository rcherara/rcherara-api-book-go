package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"./controllers"
	"./models"
	"./service"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

// Array of books
var books []models.Book

// DataBase
var db *sql.DB

// ======== Init ========

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ======== Main ========

func main() {
	// Status
	log.Println("Start Service on port ", os.Getenv("PORT"))
	log.Println("Serving API Book on http://localhost", os.Getenv("PORT"))
	// fmt.Println()
	db = service.ConnectDB()

	controller := controllers.Controller{}
	// Init router
	router := mux.NewRouter()

	// Route handles & endpoints & Http-Actions

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	// ======== Server? ========

	s := http.Server{
		Addr:           os.Getenv("PORT"), // Server address
		Handler:        router,            // Route manager
		ReadTimeout:    10 * time.Second,  // Request reading time
		WriteTimeout:   10 * time.Second,  // Response recording time
		IdleTimeout:    10 * time.Second,  // Waiting time for the next request
		MaxHeaderBytes: 1 << 20,           // Maximum size of http header in bytes (1 * 2 ^ 20 = 128 kByte)
	}

	// Start server
	log.Fatal(s.ListenAndServe()) // If there are errors then display them
	fmt.Println("Listening on port :  ", os.Getenv("PORT"))

}
