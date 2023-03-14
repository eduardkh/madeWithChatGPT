package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"time"

	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	"github.com/spf13/viper"
)

func main() {
	// Load configuration from config file
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config file: %v", err)
	}

	// Get IP address, username, password, and SSH config file path from config file
	ip := viper.GetString("device")
	username := viper.GetString("username")
	password := viper.GetString("password")
	sshConfigFile := viper.GetString("ssh_config_file")

	p, err := platform.NewPlatform(
		"cisco_iosxe",
		ip,
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername(username),
		options.WithAuthPassword(password),
		options.WithSSHConfigFile(sshConfigFile),
	)
	if err != nil {
		log.Fatalf("Failed to create platform: %v", err)
	}

	d, err := p.GetNetworkDriver()
	if err != nil {
		log.Fatalf("Failed to fetch network driver from platform: %v", err)
	}

	if err := d.Open(); err != nil {
		log.Fatalf("Failed to open driver: %v", err)
	}
	defer d.Close()

	// Send "terminal length 0" command to disable pagination
	_, err = d.SendCommand("terminal length 0")
	if err != nil {
		log.Fatalf("Failed to send command: %v", err)
	}

	// Send "show running-config" command to retrieve configuration
	response, err := d.SendCommand("show running-config")
	if err != nil {
		log.Fatalf("Failed to send command: %v", err)
	}

	// Use regular expression to extract hostname from running config output
	hostnameRegex := regexp.MustCompile(`hostname\s+(\S+)`)
	hostnameMatch := hostnameRegex.FindStringSubmatch(response.Result)
	if len(hostnameMatch) != 2 {
		log.Fatalf("Failed to extract hostname from running config output")
	}
	hostname := hostnameMatch[1]

	// Get current date in YYYY-MM-DD format
	date := time.Now().Format("2006-01-02")

	// Save output to file with hostname and date in filename
	filename := fmt.Sprintf("%s_%s.ios", hostname, date)
	err = ioutil.WriteFile(filename, []byte(response.Result), 0644)
	if err != nil {
		log.Fatalf("Failed to write configuration to file: %v", err)
	}

	fmt.Printf("Backup saved to %s\n", filename)
}
