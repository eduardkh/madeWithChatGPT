package main

import (
	"fmt"

	"github.com/vishvananda/netlink"
)

func main() {
	links, err := netlink.LinkList()
	if err != nil {
		fmt.Println("Error fetching interfaces:", err)
		return
	}

	for _, link := range links {
		fmt.Printf("Interface Name: %s, Type: %s, MAC: %s\n", link.Attrs().Name, link.Type(), link.Attrs().HardwareAddr)
	}
}
