package commands

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

const (
	version = "0.1"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display szczecin version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("szczecin version %s %s/%s\n", version, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
