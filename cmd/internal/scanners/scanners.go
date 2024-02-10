package scanners

import (
	"fmt"
	"net"
	"sort"
	"time"
)

func Scan(minPort uint16, maxPort uint16, host string, workers int) []uint16 {

	requests := make(chan uint16, workers)
	results := make(chan uint16)
	var openPorts []uint16

	for i := 0; i < workers; i++ {
		go worker(host, requests, results)
	}

	go func() {
		for i := minPort; i <= maxPort; i++ {
			requests <- i
		}
	}()

	for i := uint16(minPort); i <= maxPort; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		}
	}

	close(requests)
	close(results)
	sort.Slice(openPorts, func(i, j int) bool {
		return openPorts[i] < openPorts[j]
	})
	return openPorts
}

func worker(host string, ports <-chan uint16, results chan<- uint16) {
	for port := range ports {
		address := fmt.Sprintf("%s:%d", host, port)
		conn, err := net.DialTimeout("tcp", address, time.Second*3)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- port
	}
}
