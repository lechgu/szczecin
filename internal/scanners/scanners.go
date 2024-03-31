package scanners

import (
	"fmt"
	"sync"
	"time"

	"github.com/samber/lo"
)

func Scan(targets []string, ports []uint16, workers int, timeout time.Duration) <-chan string {

	requests := make(chan string, workers)
	results := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		work(&wg, timeout, requests, results)
	}

	go func() {
		lo.ForEach(targets, func(target string, _ int) {
			lo.ForEach(ports, func(port uint16, _ int) {
				address := fmt.Sprintf("%s:%d", target, port)
				requests <- address
			})
		})
		close(requests)
		wg.Wait()
		close(results)
	}()
	return results
}
