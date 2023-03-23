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
	// Parse the command-line arguments
	orgFlag := flag.Bool("organization", false, "Query by organization")
	flag.Parse()

	// Open the SQLite database
	db, err := sql.Open("sqlite3", "oui.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Prepare the SQL statement for querying the data
	var stmt *sql.Stmt
	var query string
	if *orgFlag {
		stmt, err = db.Prepare("SELECT * FROM oui WHERE organization LIKE ?")
		if err != nil {
			log.Fatal(err)
		}
		query = strings.Join(flag.Args(), " ")
	} else {
		stmt, err = db.Prepare("SELECT * FROM oui WHERE oui LIKE ?")
		if err != nil {
			log.Fatal(err)
		}
		// Get the MAC address from the command-line arguments
		if len(flag.Args()) != 1 {
			log.Fatal("Usage: go run query_mac_address.go <mac_address> | [-organization <organization>]")
		}
		macAddress := flag.Args()[0]

		// Remove any non-alphanumeric characters from the MAC address
		re := regexp.MustCompile(`[^0-9A-Fa-f]`)
		macAddress = re.ReplaceAllString(macAddress, "")

		// Convert the MAC address to lowercase
		macAddress = strings.ToLower(macAddress)

		// Extract the OUI from the MAC address
		oui := macAddress[0:6]

		query = oui + "%"
	}
	defer stmt.Close()

	// Execute the SQL statement with the query string as the parameter
	rows, err := stmt.Query("%" + query + "%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the result
	found := false
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
		if *orgFlag {
			fmt.Printf("OUI:\t\t%s\nOrganization:\t%s\nAddress:\t%s\nCity:\t\t%s\nCountry:\t%s\n", oui, organization, address, city, country)
		} else {
			fmt.Printf("Organization:\t%s\nAddress:\t%s\nCity:\t\t%s\nCountry:\t%s\nOUI:\t\t%s\n", organization, address, city, country, strings.ToUpper(oui))
		}
		found = true
	}
	if !found {
		if *orgFlag {
			fmt.Printf("No results found for organization: %s\n", query)
		} else {
			fmt.Printf("No results found for MAC address: %s\n", query)
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
