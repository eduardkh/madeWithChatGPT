package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
)

func main() {
	p, err := platform.NewPlatform(
		"cisco_iosxe",
		"192.168.99.160", // Replace with the IP address of your device
		options.WithAuthNoStrictKey(),
		options.WithAuthUsername("cisco"), // Replace with the username for your device
		options.WithAuthPassword("cisco"), // Replace with the password for your device
		options.WithSSHConfigFile("~/.ssh/config"),
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

	err = ioutil.WriteFile("backup.txt", []byte(response.Result), 0644)
	if err != nil {
		log.Fatalf("Failed to write configuration to file: %v", err)
	}

	fmt.Println("Backup saved to backup.txt")
}
