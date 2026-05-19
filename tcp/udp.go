package tcp

import (
	"fmt"
	"net"
)

// ScanUDPPort scans open ports using UDP port scanning technique
func ScanUDPPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("udp", address)

	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	conn.Close()
	return true

}
