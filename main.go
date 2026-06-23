package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	var err error
	db, err := sql.Open("postgres", "user=postgres dbname=expense-tracker sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/expenses", handleExpenses)
	http.ListenAndServe(":3000", nil)

	expenseDatabase()

}

func handleExpenses(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
