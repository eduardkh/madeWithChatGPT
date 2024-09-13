package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	err := getARPTable()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// getARPTable reads the ARP table from /proc/net/arp and displays only reachable entries.
func getARPTable() error {
	// Open the /proc/net/arp file.
	file, err := os.Open("/proc/net/arp")
	if err != nil {
		return fmt.Errorf("failed to open ARP table: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skip the first line (header).
	if !scanner.Scan() {
		return fmt.Errorf("failed to read ARP table header")
	}

	fmt.Println("Reachable ARP Table Entries:")
	// Process each subsequent line in the ARP table.
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Ensure that the line has the correct number of fields.
		if len(fields) < 6 {
			continue
		}

		ipAddr := fields[0]
		flags := fields[2]
		macAddr := fields[3]
		device := fields[5]

		// Filter entries where Flags is 0x2 (reachable).
		if flags == "0x2" {
			fmt.Printf("IP Address: %s, MAC Address: %s, Device: %s\n", ipAddr, macAddr, device)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading ARP table: %v", err)
	}

	return nil
}
