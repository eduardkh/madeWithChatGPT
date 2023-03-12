package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Read configuration from file
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Error reading configuration file: %s\n", err)
		return
	}

	// Get directories to scan from configuration
	dirs := viper.GetStringSlice("directories")

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "./file_fingerprints.db")
	if err != nil {
		log.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	// Create a table to store the file fingerprints
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS fingerprints (path TEXT, filename TEXT, md5 TEXT, sha1 TEXT, sha256 TEXT)")
	if err != nil {
		log.Println("Error creating table:", err)
		return
	}

	// Prepare a statement to insert data into the table
	stmt, err := db.Prepare("INSERT INTO fingerprints (path, filename, md5, sha1, sha256) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return
	}
	defer stmt.Close()

	// Walk through each directory and calculate the fingerprints for each file
	for _, dir := range dirs {
		err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Skip directories
			if info.IsDir() {
				return nil
			}
			// Split the file path and filename
			filepath, filename := filepath.Split(path)
			// Open the file for reading
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			// Calculate the MD5, SHA1, and SHA256 fingerprints
			hash_md5 := md5.New()
			hash_sha1 := sha1.New()
			hash_sha256 := sha256.New()

			if _, err := io.Copy(hash_md5, file); err != nil {
				return err
			}
			if _, err := file.Seek(0, 0); err != nil {
				return err
			}
			if _, err := io.Copy(hash_sha1, file); err != nil {
				return err
			}
			if _, err := file.Seek(0, 0); err != nil {
				return err
			}
			if _, err := io.Copy(hash_sha256, file); err != nil {
				return err
			}

			// Convert the fingerprints to hex strings
			md5_string := fmt.Sprintf("%x", hash_md5.Sum(nil))
			sha1_string := fmt.Sprintf("%x", hash_sha1.Sum(nil))
			sha256_string := fmt.Sprintf("%x", hash_sha256.Sum(nil))

			// Bind the data to the prepared statement and execute the statement
			_, err = stmt.Exec(filepath, filename, md5_string, sha1_string, sha256_string)
			if err != nil {
				return err
			}
			return nil
		})

	}
}
