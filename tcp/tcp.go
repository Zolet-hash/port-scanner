package tcp

import (
	"fmt"
	"net"
	"sort"
	"sync"
	"time"
)

// ScanPortConcurrently scans ports concurrently with go routines
func ScanPortConcurrently(host string, port int, timeout time.Duration,
	results chan<- int, wg *sync.WaitGroup, sem chan struct{}) {

	defer wg.Done()

	defer func() { <-sem }()

	address := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp4", address, timeout)
	if err != nil {
		return
	}
	conn.Close()
	results <- port
}

//ScanHost
func ScanHost(host string, startPort,endPort,maxConcurrent int, timeout time.Duration) []int {
	results := make(chan int, endPort-startPort+1)
	sem := make(chan struct{}, maxConcurrent) //Limit concurrent connection attempts

	var wg sync.WaitGroup

	for port := startPort; port <= endPort; port++ {
		sem <-struct{}{} //Acquire a semaphore slot before launching the groutine
		wg.Add(1)
		go ScanPortConcurrently(host, port, timeout, results, &wg, sem,)
	}

	// wait for all goroutines to finish then close all channels
	go func() {
		wg.Wait()
		close(results)
	}()
	
	var openPorts []int
	for port := range results{
		openPorts = append(openPorts, port)
	}
	sort.Ints(openPorts)
	return openPorts
}