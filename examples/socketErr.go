package main

import (
	"fmt"
	"net"
)

func main() {
	tcpAddr, resolveErr := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	if resolveErr != nil {
		// Handle the error gracefully here.
		fmt.Printf("Error resolving address: %s\n", resolveErr.Error())
		return
	}

	conn, connErr := net.DialTCP("tcp", nil, tcpAddr)
	if connErr != nil {
		// Handle the error gracefully here.
		fmt.Printf("Conn err: %s\n", connErr.Error())
		return
	}
	// END MAIN OMIT
	buffer := make([]byte, 1024)
	conn.Read(buffer)
}
