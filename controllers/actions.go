package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../models"

	"github.com/gorilla/mux"
)

// Controller ...
type Controller struct{}

// ======== Init

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ======== Actions

// GetBooks ... Get all books
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("--- Server [net/http] method ", r.Method, " connection from ", r.RemoteAddr)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var book models.Book
		books = []models.Book{}
		// Status
		log.Println("Get all books")

		rows, err := db.Query("select * from books")
		logFatal(err)

		// defer rows.Close()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		for rows.Next() {
			err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
			logFatal(err)

			books = append(books, book)
			log.Println(" Book ------ >>> ", book)
		}

		//json.NewEncoder(w).Encode(books)
		response := models.Response{
			Message: books,
		}
		respJson, _ := json.Marshal(response)

		_, _ = w.Write(respJson)

	}

}

// -------------------------------------

// GetBook ... Get single book
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Not Working :(

		var response models.Response
		var status int

		bookId := strings.Replace(r.URL.Path, "/book/", "", 1)
		log.Println("Get info of a books  id = ", bookId)
		rows := db.QueryRow("SELECT * FROM books WHERE id = $1", bookId)
		//err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

		if rows == nil {
			status = http.StatusNotFound
			response.Error = fmt.Sprintf("Book with Id:%s not found!", bookId)
		} else {
			status = http.StatusOK
			response.Message = rows
		}

		w.WriteHeader(status)
		respJson, _ := json.Marshal(response)
		_, _ = w.Write(respJson)

	}
}

// -------------------------------------

// AddBook ... Add new book
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // This func works

		var book models.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)

		err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)
		logFatal(err)

		json.NewEncoder(w).Encode(bookID)

		// Info | log.Println(book)

		// Status
		log.Println("New book is added")
	}
}

// -------------------------------------

// UpdateBook ... Update book
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var book models.Book

		json.NewDecoder(r.Body).Decode(&book)

		res, err := db.Exec("UPDATE books SET title=$1, author=$2, year=$3 WHERE id=$4 RETURNING id", &book.Title, &book.Author, &book.Year, &book.ID)

		rowsUpdated, err := res.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)

		// Status
		log.Println("The book is updated")
	}
}

// -------------------------------------

// RemoveBook ... Delete book
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) { // Works

		params := mux.Vars(r)
		i, _ := strconv.Atoi(params["id"])

		res, err := db.Exec("DELETE FROM books WHERE id = $1", i)
		logFatal(err)

		rowsDeleted, err := res.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)

		// Status
		log.Println("The book is deleted")
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	name := strings.Replace(r.URL.Path, "/hello/", "", 1)
	response := models.Response{
		Message: fmt.Sprintf("Hello %s! Glad to see you again.", name),
	}

	respJson, _ := json.Marshal(response)
	_, _ = w.Write(respJson)
}
