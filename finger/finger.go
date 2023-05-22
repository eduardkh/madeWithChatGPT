package main

import (
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <filename>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	// Read the certificate file
	certPEM, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		os.Exit(1)
	}

	// Parse the PEM-encoded certificate
	block, _ := pem.Decode(certPEM)
	if block == nil {
		fmt.Println("Failed to parse certificate")
		os.Exit(1)
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("Failed to parse certificate: %v\n", err)
		os.Exit(1)
	}

	// Compute the SHA256 fingerprint
	fingerprint := sha256.Sum256(cert.Raw)

	// Print the lowercase hexadecimal fingerprint
	for i, byte := range fingerprint {
		if i > 0 {
			fmt.Print("")
		}
		fmt.Printf("%02x", byte)
	}
	fmt.Println()
}
