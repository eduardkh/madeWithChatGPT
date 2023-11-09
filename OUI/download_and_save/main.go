package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	err := DownloadOUIFile()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// DownloadOUIFile downloads the OUI CSV file from the IEEE website and stores it in the local application data directory.
func DownloadOUIFile() error {
	// Determine the local application data directory.
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("error getting user config directory: %w", err)
	}

	// Create a subdirectory for your app.
	appDataDir = filepath.Join(appDataDir, "MyOUIApp")
	if err := os.MkdirAll(appDataDir, 0755); err != nil {
		return fmt.Errorf("error creating application data directory: %w", err)
	}

	// Define the URL and the local file path.
	ouiURL := "http://standards-oui.ieee.org/oui/oui.csv"
	localFilePath := filepath.Join(appDataDir, "oui.csv")

	// Download the OUI file.
	err = downloadFile(localFilePath, ouiURL)
	if err != nil {
		return fmt.Errorf("error downloading the OUI file: %w", err)
	}

	fmt.Printf("OUI file downloaded successfully to %s\n", localFilePath)
	return nil
}

// downloadFile downloads the file from the given URL to the given local path.
func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
