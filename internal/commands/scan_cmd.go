package commands

import (
	"fmt"
	"sort"
	"time"

	"github.com/lechgu/szczecin/internal/portformat"
	"github.com/lechgu/szczecin/internal/scanners"
	"github.com/samber/lo"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

var (
	targets   []string
	portQuery string
	workers   int
	timeout   int
	progress  bool
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan the host for open ports",
	RunE:  scan,
}

func scan(cmd *cobra.Command, args []string) error {

	ports, err := portformat.Parse(portQuery)
	if err != nil {
		return err
	}

	results := scanners.Scan(targets, ports, workers, time.Second*time.Duration(timeout))
	var bar *progressbar.ProgressBar
	if progress {
		bar = progressbar.Default(int64(len(targets) * len(ports)))
	}
	var opened []string
	i := 0
	for result := range results {
		if result != "" {
			opened = append(opened, result)
		}
		i++
		if progress {
			_ = bar.Add(1)
		}
	}
	opened = lo.Uniq(opened)

	sort.Slice(opened, func(i, j int) bool {
		return opened[i] < opened[j]
	})
	lo.ForEach(opened, func(result string, _ int) {
		fmt.Println(result)
	})
	return nil
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringSliceVarP(&targets, "target", "t", nil, "Target host to scan")
	_ = scanCmd.MarkFlagRequired("target")
	scanCmd.Flags().StringVarP(&portQuery, "ports", "p", "", "Ports to scan")
	_ = scanCmd.MarkFlagRequired("ports")
	scanCmd.Flags().IntVar(&workers, "workers", 16, "Number of concurrent workers")
	scanCmd.Flags().IntVar(&timeout, "timeout", 10, "Connection timeout, in seconds")
	scanCmd.Flags().BoolVar(&progress, "progress", false, "Show progress")
}
