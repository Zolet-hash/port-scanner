package tcp

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

// ScanUDPPort scans open ports using UDP port scanning technique
func ScanUDPPort(host string, port int) bool {
	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.Dial("udp", address)

	if err != nil {
		fmt.Println("Error:", err)
		return false
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(1 * time.Second))

	_, err = conn.Write([]byte("ping"))
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		//timeout or ICMP unreachable
		fmt.Println("Error:", err)
		return false
	}

	return true

}

//ScanUDPPortConcurrently scans UDP post concurently
func ScanUDPPortConcurrently(
	host string,
	port int,
	results chan<- int,
	wg *sync.WaitGroup,
	sem chan struct{},
) {
	defer wg.Done()
	defer func() { <-sem }()

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("udp", address, 1*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	_ = conn.SetDeadline(time.Now().Add(1 * time.Second))

	// Send probe packet
	_, err = conn.Write([]byte("ping"))
	if err != nil {
		return
	}

	// Wait for response
	buffer := make([]byte, 1024)

	_, err = conn.Read(buffer)
	if err != nil {
		// Most UDP services won't respond,
		// so timeout does NOT necessarily mean closed.
		return
	}

	results <- port
}

// ScanUDPHost scans host using udp protocol
func ScanUDPHost(host string, startPort, endPort, maxConcurrent int) []int {
	result := make(chan int, endPort-startPort+1)
	sem := make(chan struct{}, maxConcurrent) // Limit concurrent connection attempts at a time

	var wg sync.WaitGroup
	for port := startPort; port <= endPort; port++ {
		sem <- struct{}{}
		wg.Add(1)

		go ScanUDPPortConcurrently(host, port, result, &wg, sem)
	}

	//Wait for go routines to finish then close all channels
	go func() {
		wg.Wait()
		close(result)
	}()

	var openPorts []int
	for port := range result {
		openPorts = append(openPorts, port)
	}
	sort.Ints(openPorts)
	return openPorts

}
