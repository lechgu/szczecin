package scanners

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func work(wg *sync.WaitGroup, host string, timeout time.Duration, ports <-chan uint16, results chan<- uint16) {
	go func() {
		defer wg.Done()
		for port := range ports {
			address := fmt.Sprintf("%s:%d", host, port)
			conn, err := net.DialTimeout("tcp", address, timeout)
			if err != nil {
				results <- 0
				continue
			}
			conn.Close()
			results <- port
		}
	}()
}
