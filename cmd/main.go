package main

import (
	"fmt"

	"github.com/lechgu/szczecin/cmd/internal/scanners"
)

func main() {
	results := scanners.Scan(1, 1024, "emily.local", 100)
	for _, port := range results {
		fmt.Printf("%d open\n", port)
	}
}
