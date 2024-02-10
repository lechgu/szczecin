package scanners

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func Scan(minPort uint16, maxPort uint16, host string, workers int) <-chan uint16 {

	requests := make(chan uint16, workers)
	results := make(chan uint16)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		work(&wg, host, requests, results)
	}

	go func() {
		for i := minPort; i <= maxPort; i++ {
			requests <- i
		}
		close(requests)
		wg.Wait()
		close(results)
	}()

	return results
}

func work(wg *sync.WaitGroup, host string, ports <-chan uint16, results chan<- uint16) {
	go func() {
		defer wg.Done()
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
	}()
}
