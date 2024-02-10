package main

import (
	"fmt"

	"github.com/lechgu/szczecin/cmd/internal/scanners"
)

func main() {
	minPort := 1
	maxPort := 1024
	results := scanners.Scan(uint16(minPort), uint16(maxPort), "emily.local", 100)
	//i := 0
	for port := range results {
		if port != 0 {
			fmt.Printf("%d open\n", port)
		}
		// i++
		// if i == maxPort-minPort+1 {
		// 	break
		// }
	}
}
