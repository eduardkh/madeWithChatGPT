package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
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
	stmt, err := db.Prepare("SELECT * FROM oui WHERE oui LIKE ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Get the MAC address from the command-line arguments
	if len(os.Args) != 2 {
		log.Fatal("Usage: go run query_mac_address.go <mac_address>")
	}
	macAddress := os.Args[1]

	// Remove any non-alphanumeric characters from the MAC address
	re := regexp.MustCompile(`[^0-9A-Fa-f]`)
	macAddress = re.ReplaceAllString(macAddress, "")

	// Convert the MAC address to uppercase
	macAddress = strings.ToLower(macAddress)

	// Check if the MAC address is valid
	if len(macAddress) > 12 || len(macAddress) < 6 {
		log.Fatal("Invalid MAC address")
	}

	// Pad the MAC address with zeros to make it 12 characters long
	if len(macAddress) < 12 {
		macAddress = fmt.Sprintf("%s%s", macAddress, strings.Repeat("0", 12-len(macAddress)))
	}

	// Extract the OUI from the MAC address
	oui := macAddress[0:6]

	// Execute the SQL statement with the OUI as the parameter
	rows, err := stmt.Query(oui + "%")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Print the result
	found := false
	for rows.Next() {
		var id int
		var organization string
		var address string
		var city string
		var country string
		var oui string
		err := rows.Scan(&id, &oui, &organization, &address, &city, &country)
		if err != nil {
			log.Fatal(err)
		}
		if strings.HasPrefix(strings.ToLower(oui), oui) {
			fmt.Printf("Organization:\t%s\nAddress:\t%s\nCity:\t\t%s\nCountry:\t%s\nOUI:\t\t%s\n", organization, address, city, country, strings.ToUpper(oui))
			found = true
		}
	}
	if !found {
		fmt.Printf("No results found for %s\n", macAddress)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
