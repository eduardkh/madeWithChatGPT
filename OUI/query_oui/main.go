package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Define the CLI flag
	macFlag := flag.String("mac", "", "Full or partial MAC address to search for.")

	// Parse the CLI flags
	flag.Parse()

	// Concatenate all remaining arguments to handle unquoted dotted notation
	macAddress := strings.Join(flag.Args(), "")

	// If the mac flag is set, use it; otherwise, use the concatenated arguments
	if *macFlag != "" {
		macAddress = *macFlag
	}

	// Validate the MAC address input
	if macAddress == "" {
		fmt.Println("Please provide a MAC address using the -mac flag or as a regular argument.")
		return
	}

	// Call the function to print OUI information
	err := PrintOUIInfo(macAddress)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// PrintOUIInfo searches the OUI CSV file for the given MAC address and prints its information.
func PrintOUIInfo(macAddress string) error {
	// Standardize the MAC address by removing delimiters and converting to uppercase
	macAddress = strings.ToUpper(macAddress)
	macAddress = strings.NewReplacer(":", "", "-", "", ".", "").Replace(macAddress)

	// Now that we've cleaned the MAC address, check if it's at least 6 characters long
	if len(macAddress) < 6 {
		return fmt.Errorf("MAC address must be at least 6 characters after standardizing")
	}

	// Trim the MAC address to the first 6 characters (OUI portion)
	ouiPortion := macAddress[:6]

	// Get the local application data directory
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	// Construct the path to the oui.csv file
	csvFilePath := filepath.Join(appDataDir, "MyOUIApp", "oui.csv")

	// Open the CSV file
	file, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read and ignore the header line
	if _, err := reader.Read(); err != nil {
		return err
	}

	// Iterate through the CSV records and print the matching OUI information
	found := false
	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Check if the current record contains the OUI portion of the MAC address
		if strings.HasPrefix(strings.ToUpper(record[1]), ouiPortion) {
			fmt.Printf("MAC Address: \"%s\"\n", macAddress)
			fmt.Printf("Vendor: \"%s\"\n", record[2])
			found = true
			break // Stop after finding the first match
		}
	}

	if !found {
		fmt.Println("No matching records found for the given MAC address.")
	}

	return nil
}
