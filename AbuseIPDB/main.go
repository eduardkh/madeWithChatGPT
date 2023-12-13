package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// AbuseIPDBResponse represents the JSON response structure from AbuseIPDB API
type AbuseIPDBResponse struct {
	Data struct {
		IPAddress            string   `json:"ipAddress"`
		IsPublic             bool     `json:"isPublic"`
		IpVersion            int      `json:"ipVersion"`
		IsWhitelisted        *bool    `json:"isWhitelisted"` // Pointer to handle null
		AbuseConfidenceScore int      `json:"abuseConfidenceScore"`
		CountryCode          string   `json:"countryCode"`
		UsageType            string   `json:"usageType"`
		Isp                  string   `json:"isp"`
		Domain               string   `json:"domain"`
		Hostnames            []string `json:"hostnames"`
		IsTor                bool     `json:"isTor"`
		CountryName          string   `json:"countryName"`
		TotalReports         int      `json:"totalReports"`
		NumDistinctUsers     int      `json:"numDistinctUsers"`
		LastReportedAt       *string  `json:"lastReportedAt"` // Pointer to handle null
		Reports              []struct {
			ReportedAt          string `json:"reportedAt"`
			Comment             string `json:"comment"`
			Categories          []int  `json:"categories"`
			ReporterId          int    `json:"reporterId"`
			ReporterCountryCode string `json:"reporterCountryCode"`
			ReporterCountryName string `json:"reporterCountryName"`
		} `json:"reports"`
	} `json:"data"`
}

// ReadAPIKey reads the API key from a config file
func ReadAPIKey(configFilePath string) (string, error) {
	data, err := os.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// QueryAbuseIPDB makes a request to the AbuseIPDB API and returns the result
func QueryAbuseIPDB(ipAddress, apiKey string) (*AbuseIPDBResponse, error) {
	url := fmt.Sprintf("https://api.abuseipdb.com/api/v2/check?maxAgeInDays=90&verbose&ipAddress=%s", ipAddress)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Key", apiKey)
	req.Header.Add("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AbuseIPDBResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func main() {
	// Get the local application data directory
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		log.Println(err)
	}

	configFilePath := filepath.Join(appDataDir, "show", "show.config")

	apiKey, err := ReadAPIKey(configFilePath)
	if err != nil {
		fmt.Println("Error reading API key:", err)
		return
	}

	// ipAddress := "212.70.149.150" // Replace with dynamic input as needed
	ipAddress := "8.8.8.8" // Replace with dynamic input as needed
	response, err := QueryAbuseIPDB(ipAddress, apiKey)
	if err != nil {
		fmt.Println("Error querying AbuseIPDB:", err)
		return
	}

	fmt.Printf("IP Address            : %+v\n", response.Data.IPAddress)
	fmt.Printf("Abuse Confidence Score: %+v\n", response.Data.AbuseConfidenceScore)
	fmt.Printf("Country Code          : %+v\n", response.Data.CountryCode)
	fmt.Printf("Usage Type            : %+v\n", response.Data.UsageType)
	if len(response.Data.Hostnames) > 0 {
		fmt.Printf("Hostnames             : %+v\n", response.Data.Hostnames)
	}
	if response.Data.Domain != "" {
		fmt.Printf("Domain                : %+v\n", response.Data.Domain)
	}
	if response.Data.Isp != "" {
		fmt.Printf("ISP                   : %+v\n", response.Data.Isp)
	}
	fmt.Printf("Is TOR Server         : %+v\n", response.Data.IsTor)
	if response.Data.LastReportedAt != nil {
		fmt.Printf("Last Reported At      : %+v\n", *response.Data.LastReportedAt)
	}

	// For reports, it's best to loop over individual responses.
	if len(response.Data.Reports) > 0 {
		fmt.Println("")
		fmt.Println("Reports:")
		for _, report := range response.Data.Reports {
			fmt.Println("# Reported At     : ", report.ReportedAt)
			fmt.Println("# Reporter Country: ", report.ReporterCountryName)
			fmt.Println("# Comment         : ", report.Comment)
		}
	} else {
		fmt.Println("No reports available.")
	}
}
