package tcp

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ScanPortConcurrently scans ports concurrently with go routines
func ScanPortConcurrently(host string, port int, timeout time.Duration,
	results chan<- int, wg *sync.WaitGroup, sem chan struct{}) {

	defer wg.Done()

	defer func() { <-sem }()

	address := fmt.Sprintf("%s %d", host, port)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	if err != nil {
		return
	}
	conn.Close()
	results <- port
}

//ScanHost
func ScanHost(host string, startPort,endPort,maxConcurrent int, timeout time.Duration) []int {
	
}