package commands

import (
	"fmt"
	"runtime"

	"github.com/lechgu/szczecin/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays szczecin version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("szczecin version %02d.%02d.%d %s/%s\n", version.Major, version.Minor, version.Patch, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
