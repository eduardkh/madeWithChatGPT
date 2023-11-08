package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ipcalc <IP>/<CIDR>")
		os.Exit(1)
	}

	ip, ipnet, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		fmt.Println("Please provide a valid IP address and subnet or CIDR.")
		os.Exit(1)
	}

	mask := ipnet.Mask
	ones, bits := ipnet.Mask.Size()

	if ones == 31 {
		// Special case for /31 networks according to RFC 3021.
		fmt.Printf("Address:     %s\n", ip)
		fmt.Printf("Subnet Mask: %s\n", net.IP(mask).String())
		fmt.Printf("Usable Host IPs: %s - %s\n", ip, nextIP(ip))
		fmt.Println("According to RFC 3021, this is a /31 network without network and broadcast addresses.")
	} else if ones == 32 {
		// Special case for /32 networks.
		fmt.Printf("Address:     %s\n", ip)
		fmt.Println("Subnet Mask: 255.255.255.255")
		fmt.Println("This is a /32 network, the IP address is for a single host.")
	} else {
		network := ip.Mask(mask)
		broadcast := net.IP(make([]byte, 4))
		for i := range network {
			broadcast[i] = network[i] | ^mask[i]
		}

		hostBits := bits - ones
		hosts := (1 << hostBits) - 2 // Subtract network and broadcast addresses.

		fmt.Printf("Address:     %s\n", ip)
		fmt.Printf("Network:     %s\n", network)
		fmt.Printf("Broadcast:   %s\n", broadcast)
		fmt.Printf("Host Range:  %s - %s\n", nextIP(network), previousIP(broadcast))
		fmt.Printf("Subnet Mask: %s\n", net.IP(mask).String())
		fmt.Printf("Host Number: %d\n", hosts)
	}
}

func nextIP(ip net.IP) net.IP {
	next := net.IP(make([]byte, len(ip)))
	copy(next, ip)
	for j := len(next) - 1; j >= 0; j-- {
		next[j]++
		if next[j] > 0 {
			break
		}
	}
	return next
}

func previousIP(ip net.IP) net.IP {
	previous := net.IP(make([]byte, len(ip)))
	copy(previous, ip)
	for j := len(previous) - 1; j >= 0; j-- {
		previous[j]--
		if previous[j] < 255 {
			break
		}
	}
	return previous
}
