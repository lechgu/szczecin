package version

import (
	"fmt"
	"runtime"
)

var (
	Major = 24
	Minor = 4
	Patch = 0
)

func Print() {
	fmt.Printf("szczecin version %02d.%02d.%d %s/%s\n", Major, Minor, Patch, runtime.GOOS, runtime.GOARCH)
}
