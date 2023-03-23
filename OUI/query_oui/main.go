package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open the SQLite database
	db, err := sql.Open("sqlite3", "oui.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the SQL statement for querying the data
	stmt, err := db.Prepare("SELECT oui, organization FROM oui WHERE oui LIKE ? OR organization LIKE ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Get the query string from the command-line arguments
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run query_oui.go <query_string>")
	}
	query := "%" + os.Args[1] + "%"

	// Execute the SQL statement with the query string as the parameter
	rows, err := stmt.Query(query, query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop over the rows and print the OUI and organization
	for rows.Next() {
		var oui string
		var organization string
		err := rows.Scan(&oui, &organization)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\n", oui, organization)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
