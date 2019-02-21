package service

import (
	"database/sql"
	"log"
	"os"

	"github.com/lib/pq"
)

// ======== Init

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectDB ...
func ConnectDB() *sql.DB {
	// pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL_CLOUD"))
	pgURL, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL_LOCAL"))
	logFatal(err)

	db, err = sql.Open("postgres", pgURL)
	logFatal(err)

	err = db.Ping()
	logFatal(err)

	// Info | log.Println(pgURL)

	return db
}
