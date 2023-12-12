package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// WhoisData represents the parsed WHOIS data.
type WhoisData struct {
	AS        string
	IP        string
	BGPPrefix string
	CC        string
	Registry  string
	Allocated string
	ASName    string
}

// parseWhoisData parses a single line of WHOIS data into a WhoisData struct.
func parseWhoisData(line string) (WhoisData, error) {
	parts := strings.Split(line, "|")
	if len(parts) != 7 {
		return WhoisData{}, fmt.Errorf("invalid data format")
	}

	return WhoisData{
		AS:        strings.TrimSpace(parts[0]),
		IP:        strings.TrimSpace(parts[1]),
		BGPPrefix: strings.TrimSpace(parts[2]),
		CC:        strings.TrimSpace(parts[3]),
		Registry:  strings.TrimSpace(parts[4]),
		Allocated: strings.TrimSpace(parts[5]),
		ASName:    strings.TrimSpace(parts[6]),
	}, nil
}

// lookup performs a WHOIS lookup for the given IP segment and returns parsed data.
func lookup(ipSegment string) (WhoisData, error) {
	conn, err := net.Dial("tcp", "whois.cymru.com:43")
	if err != nil {
		return WhoisData{}, err
	}
	defer conn.Close()

	query := fmt.Sprintf(" -v %s\n", ipSegment)
	_, err = conn.Write([]byte(query))
	if err != nil {
		return WhoisData{}, err
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip header line
		if strings.Contains(line, "AS      | IP               | BGP Prefix          | CC | Registry | Allocated  | AS Name") {
			continue
		}
		return parseWhoisData(line)
	}

	if err := scanner.Err(); err != nil {
		return WhoisData{}, err
	}

	return WhoisData{}, fmt.Errorf("no data received")
}

func main() {
	data, err := lookup("8.8.8.8")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("IP Address: %s\n", data.IP)
	fmt.Printf("BGP Prefix: %s\n", data.BGPPrefix)
	fmt.Printf("Autonomous System Number: %s\n", data.AS)
	fmt.Printf("Autonomous System Name: %s\n", data.ASName)
	fmt.Printf("Regional Internet registry (RIR): %s\n", data.Registry)
	fmt.Printf("Allocated: %s\n", data.Allocated)
	fmt.Printf("Country: %s\n", data.CC)
}
