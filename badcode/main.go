package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq" // inject PQ database driver
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE"))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	http.HandleFunc("/counter", handleCounter)
	panic(http.ListenAndServe(":"+os.Getenv("PORT"), http.DefaultServeMux))
}

func handleCounter(w http.ResponseWriter, r *http.Request) {
	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error starting transaction: " + err.Error()))
		return
	}

	row := tx.QueryRow("UPDATE counter SET value = value + 1 WHERE 1=1 RETURNING value")
	var count int

	if err := row.Scan(&count); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error scanning row: " + err.Error()))
		return
	}

	if err := tx.Commit(); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error committing: " + err.Error()))
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%d", count)))
}
