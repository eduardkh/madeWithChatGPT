package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func main() {
	err := getARPTable()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

// getARPTable retrieves and displays the ARP table entries.
// Modified to show only dynamic IP addresses.
func getARPTable() error {
	// Load the iphlpapi.dll library.
	modiphlpapi := windows.NewLazySystemDLL("iphlpapi.dll")
	// Get the procedure address for the GetIpNetTable function.
	procGetIpNetTable := modiphlpapi.NewProc("GetIpNetTable")

	// First, call the function with a null buffer to get the size needed.
	var size uint32
	ret, _, _ := procGetIpNetTable.Call(
		0,                              // Buffer pointer (0 to get the required size).
		uintptr(unsafe.Pointer(&size)), // Pointer to the size variable.
		0,                              // We don't want the table sorted.
	)
	if ret != uintptr(syscall.ERROR_INSUFFICIENT_BUFFER) {
		// If the return value isn't ERROR_INSUFFICIENT_BUFFER, an error occurred.
		return fmt.Errorf("GetIpNetTable failed: %v", syscall.Errno(ret))
	}

	// Allocate a buffer with the required size.
	buffer := make([]byte, size)
	// Second call to GetIpNetTable to get the actual ARP table data.
	ret, _, _ = procGetIpNetTable.Call(
		uintptr(unsafe.Pointer(&buffer[0])), // Buffer to receive the ARP table.
		uintptr(unsafe.Pointer(&size)),      // Size of the buffer.
		0,                                   // We don't want the table sorted.
	)
	if ret != 0 {
		// If the return value isn't 0 (NO_ERROR), an error occurred.
		return fmt.Errorf("GetIpNetTable failed: %v", syscall.Errno(ret))
	}

	// Parse the ARP table.
	// The first 4 bytes of the buffer contain the number of entries.
	entries := *(*uint32)(unsafe.Pointer(&buffer[0]))
	// Offset to the first MIB_IPNETROW structure in the buffer.
	offset := uintptr(unsafe.Sizeof(entries))

	fmt.Println("Dynamic ARP Table Entries:")
	for i := uint32(0); i < entries; i++ {
		// Calculate the pointer to the current MIB_IPNETROW.
		rowPtr := unsafe.Pointer(uintptr(unsafe.Pointer(&buffer[0])) + offset)
		// Dereference the pointer to get the row data.
		row := *(*MIB_IPNETROW)(rowPtr)

		// Convert the IP address from DWORD to net.IP.
		ipAddr := uint32ToIP(row.DwAddr)
		// Format the MAC address.
		macAddr := formatMACAddress(row.BPhysAddr[:row.DwPhysAddrLen])
		// Get the entry type (Dynamic, Static, etc.).
		entryType := netTypeToString(row.DwType)

		// Only display dynamic entries.
		if entryType == "Dynamic" {
			fmt.Printf("IP Address: %s, MAC Address: %s\n", ipAddr, macAddr)
		}

		// Move the offset to the next MIB_IPNETROW structure.
		offset += unsafe.Sizeof(row)
	}

	return nil
}

// MIB_IPNETROW represents an entry in the ARP table.
type MIB_IPNETROW struct {
	DwIndex       uint32  // Interface index.
	DwPhysAddrLen uint32  // Length of the physical address.
	BPhysAddr     [8]byte // Physical (MAC) address.
	DwAddr        uint32  // IP address.
	DwType        uint32  // Type of ARP entry.
}

// uint32ToIP converts a DWORD IP address to a net.IP type.
func uint32ToIP(ip uint32) net.IP {
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[:], ip)
	return net.IPv4(bytes[0], bytes[1], bytes[2], bytes[3])
}

// formatMACAddress formats a MAC address byte slice into a string.
func formatMACAddress(mac []byte) string {
	macAddr := net.HardwareAddr(mac)
	return macAddr.String()
}

// netTypeToString converts the ARP entry type to a string.
func netTypeToString(t uint32) string {
	switch t {
	case 3:
		return "Dynamic"
	case 4:
		return "Static"
	case 2:
		return "Invalid"
	default:
		return fmt.Sprintf("Unknown (%d)", t)
	}
}
