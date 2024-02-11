package commands

import (
	"fmt"
	"sort"
	"time"

	"github.com/lechgu/szczecin/internal/scanners"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var (
	targets  []string
	minPort  uint16
	maxPort  uint16
	workers  int
	timeout  int
	progress bool
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the host for open ports",
	RunE:  scan,
}

func scan(cmd *cobra.Command, args []string) error {
	lo.ForEach(targets, func(target string, _ int) {
		scanOne(target)
	})
	return nil
}

func scanOne(target string) error {
	results := scanners.Scan(minPort, maxPort, target, workers, time.Second*time.Duration(timeout))
	var bar *progressbar.ProgressBar
	if progress {
		bar = progressbar.Default(int64(maxPort - minPort + 1))
	}
	var ports []uint16
	i := 0
	for port := range results {

		ports = append(ports, port)
		i++
		if progress {
			_ = bar.Add(1)
		}
	}
	ports = lo.Filter(ports, func(item uint16, _ int) bool {
		return item != 0
	})
	sort.Slice(ports, func(i, j int) bool {
		return ports[i] < ports[j]
	})
	lo.ForEach(ports, func(port uint16, _ int) {
		fmt.Printf("%s:%d\n", target, port)
	})
	return nil
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringSliceVarP(&targets, "target", "t", nil, "Target host to scan")
	scanCmd.MarkFlagRequired("host")
	scanCmd.Flags().Uint16Var(&minPort, "min-port", 1, "starting port")
	scanCmd.Flags().Uint16Var(&maxPort, "max-port", 1024, "ending port")
	scanCmd.Flags().IntVar(&workers, "workers", 16, "Numbers of concurrent workers")
	scanCmd.Flags().IntVar(&timeout, "timeout", 10, "Connection timeout, in seconds")
	scanCmd.Flags().BoolVar(&progress, "progress", false, "Show progress")
}
