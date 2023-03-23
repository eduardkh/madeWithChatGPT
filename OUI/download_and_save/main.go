package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Read the contents of the raw_oui.txt file
	fileContents, err := ioutil.ReadFile("raw_oui.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Compile the regular expression
	re := regexp.MustCompile(`([0-9A-Fa-f]{6})\s+\(base 16\)\s+(.*)\n\s+(.*)\n\s+(.*)\n\s+(.*)\n\s`)

	// Find all the matches of the regular expression
	matches := re.FindAllStringSubmatch(string(fileContents), -1)

	// Open the SQLite database
	db, err := sql.Open("sqlite3", "oui.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the OUI table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS oui (
        id INTEGER PRIMARY KEY,
        oui TEXT NOT NULL,
        organization TEXT NOT NULL,
        address TEXT NOT NULL,
        city TEXT NOT NULL,
        country TEXT NOT NULL
    )`)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare the SQL statement for inserting the data
	stmt, err := db.Prepare("INSERT INTO oui(oui, organization, address, city, country) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Loop over the matches and insert the data into the database
	for _, match := range matches {
		// Extract the fields from the matching groups
		oui := strings.ToLower(strings.ReplaceAll(match[1], "-", ""))
		organization := strings.TrimSpace(match[2])
		address := strings.TrimSpace(match[3])
		city := strings.TrimSpace(match[4])
		country := strings.TrimSpace(match[5])

		// Execute the SQL statement to insert the data
		_, err := stmt.Exec(oui, organization, address, city, country)
		if err != nil {
			log.Fatal(err)
		}
	}
}
