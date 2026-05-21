package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/Zolet-hash/port-scanner/tcp"
)

// scanPort returns true if the TCP port is open on a given host
func scanPort(host string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	conn.Close()
	return true

}

func main() {
	//tinyscanner -u 20.78.06.364 -p 80
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Usage: portscanner [OPTIONS] [HOST] [PORT]")
		fmt.Fprintln(os.Stderr, "Options:")
		fmt.Fprintln(os.Stderr, "-p               Range of ports to scan")
		fmt.Fprintln(os.Stderr, "-u               UDP scan")
		fmt.Fprintln(os.Stderr, "-t               TCP scan")
		fmt.Fprint(os.Stderr, "--help           Show this message and exit.\n")
		return
	}

	host := os.Args[2]
	scanType := os.Args[1]
	//timeout := 1 * time.Second

	switch scanType {
	case "-t":
		fmt.Printf("Scanning %s ports 1-65535...\n", host) //scans 65,535 total available ports
		start := time.Now()

		openPorts := tcp.ScanHost(host, 1, 1024, 500, 500*time.Millisecond)

		elapsed := time.Since(start)

		fmt.Printf("Scan completed in %s\n", elapsed)
		fmt.Printf("Open ports on %s:\n", host)

		for _, port := range openPorts {
			fmt.Printf("  %d/tcp   open\n", port)
		}
	case "-u":
		fmt.Printf("Scanning %s ports 1-65535...\n", host)
		start := time.Now()

		openPorts := tcp.ScanUDPHost(host, 1, 1024, 500)

		elapsed := time.Since(start)

		fmt.Printf("Scan completed in %s\n", elapsed)
		fmt.Printf("Open ports on %s:\n", host)

		for _, port := range openPorts {
			fmt.Printf("  %d/udp   open\n", port)
		}

	}

	// for port := 1; port <= 1024; port++ {
	// 	if scanPort(host, port, timeout) {
	// 		fmt.Printf("Port: %d is OPEN\n", port)
	// 	}
	// 	fmt.Printf("Port: %d is CLOSED\n", port)
	// }
}
