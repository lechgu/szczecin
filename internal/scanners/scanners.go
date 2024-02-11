package scanners

import (
	"fmt"
	"sync"
	"time"
)

func Scan(minPort uint16, maxPort uint16, host string, workers int, timeout time.Duration) <-chan string {

	requests := make(chan string, workers)
	results := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		work(&wg, timeout, requests, results)
	}

	go func() {
		for i := minPort; i <= maxPort; i++ {
			address := fmt.Sprintf("%s:%d", host, i)
			requests <- address
		}
		close(requests)
		wg.Wait()
		close(results)
	}()
	return results
}
