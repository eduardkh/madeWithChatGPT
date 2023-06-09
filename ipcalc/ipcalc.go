package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/brotherpowers/ipsubnet"
)

func nextIPAddress(ip string) string {
	parsedIP := net.ParseIP(ip)
	for i := len(parsedIP) - 1; i >= 0; i-- {
		if parsedIP[i] < 255 {
			parsedIP[i]++
			break
		} else {
			parsedIP[i] = 0
		}
	}
	return parsedIP.String()
}

func previousIPAddress(ip string) string {
	parsedIP := net.ParseIP(ip)
	for i := len(parsedIP) - 1; i >= 0; i-- {
		if parsedIP[i] > 0 {
			parsedIP[i]--
			break
		} else {
			parsedIP[i] = 255
		}
	}
	return parsedIP.String()
}

func main() {
	// Check if IP address and subnet mask/CIDR are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: <IP>/<CIDR> or <IP> <Subnet Mask>")
		os.Exit(1)
	}

	var ip string
	var subnet int

	// Check if CIDR notation is used
	if strings.Contains(os.Args[1], "/") {
		// Split IP address and CIDR
		ipCIDR := strings.Split(os.Args[1], "/")
		ip = ipCIDR[0]
		fmt.Sscanf(ipCIDR[1], "%d", &subnet)
	} else {
		// Assume IP address and subnet mask are provided separately
		ip = os.Args[1]
		mask := net.IPMask(net.ParseIP(os.Args[2]).To4())
		_, subnet = mask.Size()
	}

	// Initialize the subnet calculator
	sub := ipsubnet.SubnetCalculator(ip, subnet)

	IPRange := sub.GetIPAddressRange()
	network := sub.GetNetworkPortion()
	broadcast := sub.GetBroadcastAddress()

	usableIPRange := ""
	hostInSubnet := ""

	switch subnet {
	case 31:
		usableIPRange = fmt.Sprintf("%s - %s (RFC 3021)", IPRange[0], IPRange[1])
		network = "N/A"
		broadcast = "N/A"
		hostInSubnet = "2 (RFC 3021)"
	case 32:
		usableIPRange = fmt.Sprintf("%s (Single host IP)", IPRange[0])
		network = "N/A"
		broadcast = "N/A"
		hostInSubnet = "1 (Single host IP)"
	default:
		usableIPRange = fmt.Sprintf("%s - %s", nextIPAddress(IPRange[0]), previousIPAddress(IPRange[1]))
		hostInSubnet = fmt.Sprint(sub.GetNumberAddressableHosts())
	}

	// Calculate and print the network information
	fmt.Printf("Address         : %s/%d\n", ip, subnet)
	fmt.Println("Network         :", network)
	fmt.Println("Usable IP Range :", usableIPRange)
	fmt.Println("Broadcast       :", broadcast)
	fmt.Println("SubnetMask      :", sub.GetSubnetMask())
	fmt.Println("Host in Subnet  :", hostInSubnet)
}
