package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

func sendPing(listener *icmp.PacketConn, address string, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// Create an ICMP echo request message
	message := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: id, Seq: 1, // Custom ID for matching
			Data: []byte("HELLO"),
		},
	}
	bytes, err := message.Marshal(nil)
	if err != nil {
		fmt.Printf("Failed to marshal message: %v\n", err)
		return
	}

	// Send the ICMP message
	if _, err := listener.WriteTo(bytes, &net.IPAddr{IP: net.ParseIP(address)}); err != nil {
		fmt.Printf("Failed to write to listener: %v\n", err)
	}
}

func main() {
	const baseIP = "192.168.1."
	const timeout = 3 * time.Second
	var wg sync.WaitGroup

	// Listen for ICMP packets
	listener, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		fmt.Println("Failed to create listener:", err)
		return
	}
	defer listener.Close()

	// Send ICMP echo requests
	for i := 1; i <= 254; i++ {
		ip := fmt.Sprintf("%s%d", baseIP, i)
		wg.Add(1)
		go sendPing(listener, ip, &wg, i) // Use loop index as ID for uniqueness
	}

	// Close the listener when all sends are done
	go func() {
		wg.Wait()
		time.Sleep(timeout) // Wait a bit longer for any late replies
		listener.Close()
	}()
	// Read replies
	reply := make([]byte, 1500)
	for {
		n, peer, err := listener.ReadFrom(reply)
		if err != nil {
			break // Exit the loop when the listener is closed
		}

		// Parse the received message
		receivedMessage, err := icmp.ParseMessage(1, reply[:n])
		if err != nil {
			fmt.Printf("Failed to parse message: %v\n", err)
			continue
		}

		// Check if the message is an ICMP echo reply
		if receivedMessage.Type == ipv4.ICMPTypeEchoReply {
			// if pkt, ok := receivedMessage.Body.(*icmp.Echo); ok {
			// 	fmt.Printf("Host %v is up [ID=%d, Seq=%d]\n", peer, pkt.ID, pkt.Seq)
			// }
			fmt.Printf("Host %v is up\n", peer)
		}
	}
}
