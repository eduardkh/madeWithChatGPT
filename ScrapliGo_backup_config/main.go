package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"sync"

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

	// Get username, password, and SSH config file path from config file
	username := viper.GetString("username")
	password := viper.GetString("password")
	sshConfigFile := viper.GetString("ssh_config_file")

	// Get list of devices from config file
	devices := viper.GetStringSlice("devices")

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	// Iterate over devices concurrently
	for _, device := range devices {
		wg.Add(1)

		// Use a closure to pass the device to the goroutine
		go func(device string) {
			p, err := platform.NewPlatform(
				"cisco_iosxe",
				device,
				options.WithAuthNoStrictKey(),
				options.WithAuthUsername(username),
				options.WithAuthPassword(password),
				options.WithSSHConfigFile(sshConfigFile),
			)
			if err != nil {
				log.Printf("Failed to create platform for %s: %v", device, err)
				wg.Done()
				return
			}

			d, err := p.GetNetworkDriver()
			if err != nil {
				log.Printf("Failed to fetch network driver from platform for %s: %v", device, err)
				wg.Done()
				return
			}

			if err := d.Open(); err != nil {
				log.Printf("Failed to open driver for %s: %v", device, err)
				wg.Done()
				return
			}
			defer d.Close()

			// Send "terminal length 0" command to disable pagination
			_, err = d.SendCommand("terminal length 0")
			if err != nil {
				log.Printf("Failed to send command to %s: %v", device, err)
				wg.Done()
				return
			}

			// Send "show running-config" command to retrieve configuration
			response, err := d.SendCommand("show running-config")
			if err != nil {
				log.Printf("Failed to send command to %s: %v", device, err)
				wg.Done()
				return
			}

			// Use regular expression to extract hostname from running config output
			hostnameRegex := regexp.MustCompile(`hostname\s+(\S+)`)
			hostnameMatch := hostnameRegex.FindStringSubmatch(response.Result)
			if len(hostnameMatch) != 2 {
				log.Printf("Failed to extract hostname from running config output for %s", device)
				wg.Done()
				return
			}
			hostname := hostnameMatch[1]

			// Save output to file with hostname and date in filename
			filename := fmt.Sprintf("%s.ios", hostname)
			err = ioutil.WriteFile(filename, []byte(response.Result), 0644)
			if err != nil {
				log.Printf("Failed to write configuration to file for %s: %v", device, err)
				wg.Done()
				return
			}

			log.Printf("Backup saved to %s for %s\n", filename, device)

			wg.Done()
		}(device)
	}

	wg.Wait()

	log.Println("Done")
	os.Exit(0)
}
