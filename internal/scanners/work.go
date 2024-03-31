package scanners

import (
	"net"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, timeout time.Duration, adresses <-chan string, results chan<- string) {
	go func() {
		defer wg.Done()
		for address := range adresses {
			conn, err := net.DialTimeout("tcp", address, timeout)
			if err != nil {
				results <- ""
				continue
			}
			conn.Close()
			results <- address
		}
	}()
}
