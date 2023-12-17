package main

import (
	"log"

	"github.com/bi-zone/wmi"
)

type DNSClientCache struct {
	Name       string
	Entry      string
	Data       string
	Type       uint16
	TimeToLive uint32
}

func main() {
	var dnsCacheEntries []DNSClientCache

	query := "SELECT Name, Data, Type, TimeToLive FROM MSFT_DNSClientCache"
	err := wmi.QueryNamespace(query, &dnsCacheEntries, "root/StandardCimv2")
	if err != nil {
		log.Fatalf("Cimv2 query failed: %v", err)
	}

	for _, entry := range dnsCacheEntries {
		// Filter for A (IPv4 - Type 1) and AAAA (IPv6 - Type 28) records
		if (entry.Type == 1 || entry.Type == 28) && entry.Data != "" {

			log.Printf(`%s
Entry (A / AAAA)       : %s
Data (IP)              : %s
TTL (Seconds)          : %d

`, entry.Name, entry.Entry, entry.Data, entry.TimeToLive)
		}
	}
}
