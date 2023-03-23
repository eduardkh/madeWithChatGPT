package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

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
	stmt, err := db.Prepare("SELECT * FROM oui WHERE oui LIKE ? OR organization LIKE ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Get the query string from the command-line arguments
	query := strings.Join(os.Args[1:], " ")

	// Execute the SQL statement with the query string as the parameter
	rows, err := stmt.Query("%"+query+"%", "%"+query+"%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Loop over the rows and print the data
	for rows.Next() {
		var id int
		var oui string
		var organization string
		var address string
		var city string
		var country string
		err := rows.Scan(&id, &oui, &organization, &address, &city, &country)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\t%s\t%s\t%s\t%s\t%s\n", oui, organization, address, city, country)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
