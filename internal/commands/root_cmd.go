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
	host     string
	minPort  uint16
	maxPort  uint16
	workers  int
	timeout  int
	progress bool
)

var rootCmd = &cobra.Command{
	Use:   "szczecin",
	Short: "Port scanner",
	RunE:  scan,
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func scan(cmd *cobra.Command, args []string) error {
	results := scanners.Scan(minPort, maxPort, host, workers, time.Second*time.Duration(timeout))
	var bar *progressbar.ProgressBar
	if progress {
		bar = progressbar.Default(int64(maxPort - minPort + 1))
	}
	var openPorts []uint16
	i := 0
	for port := range results {
		if port != 0 {
			openPorts = append(openPorts, port)
		}
		i++
		if progress {
			_ = bar.Add(1)
		}
	}
	sort.Slice(openPorts, func(i, j int) bool {
		return openPorts[i] < openPorts[j]
	})
	lo.ForEach(openPorts, func(item uint16, _ int) {
		fmt.Println(item)
	})
	return nil
}

func init() {
	rootCmd.Flags().StringVar(&host, "host", "", "Host to scan")
	rootCmd.MarkFlagRequired("host")
	rootCmd.Flags().Uint16Var(&minPort, "min-port", 1, "starting port")
	rootCmd.Flags().Uint16Var(&maxPort, "max-port", 1024, "ending port")
	rootCmd.Flags().IntVar(&workers, "workers", 16, "Numbers of concurrent workers")
	rootCmd.Flags().IntVar(&timeout, "timeout", 10, "Connection timeout, in seconds")
	rootCmd.Flags().BoolVar(&progress, "progress", false, "Show progress")
}
