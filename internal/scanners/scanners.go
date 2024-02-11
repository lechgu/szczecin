package scanners

import (
	"sync"
	"time"
)

func Scan(minPort uint16, maxPort uint16, host string, workers int, timeout time.Duration) <-chan string {

	requests := make(chan uint16, workers)
	results := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		work(&wg, host, timeout, requests, results)
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
