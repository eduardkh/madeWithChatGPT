package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"regexp"
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
	var stmt *sql.Stmt
	var sqlQuery string

	// Check if the organization flag is set
	var organizationFlag = flag.Bool("organization", false, "Query by organization")
	flag.Parse()

	if *organizationFlag {
		sqlQuery = "SELECT * FROM oui WHERE organization LIKE ?"
	} else {
		sqlQuery = "SELECT * FROM oui WHERE oui LIKE ?"
	}

	stmt, err = db.Prepare(sqlQuery)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Get the query string from the command-line arguments
	if len(flag.Args()) != 1 {
		log.Fatal("Usage: go run query_mac_address.go [--organization] <query>")
	}
	query := strings.ToLower(flag.Arg(0))

	// Remove any non-alphanumeric characters from the query
	re := regexp.MustCompile(`[^0-9A-Za-z ]`)
	query = re.ReplaceAllString(query, "")

	// Execute the SQL statement with the query string as the parameter
	rows, err := stmt.Query("%" + query + "%")
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
		fmt.Printf("OUI:\t\t%s\nOrganization:\t%s\nAddress:\t%s\nCity:\t\t%s\nCountry:\t%s\n\n", strings.ToUpper(oui), organization, address, city, country)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
