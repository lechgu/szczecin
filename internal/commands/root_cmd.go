package commands

import (
	"fmt"

	"github.com/lechgu/szczecin/internal/version"
	"github.com/spf13/cobra"
)

var (
	ver bool
)

var rootCmd = &cobra.Command{
	Use:   "szczecin",
	Short: "Port scanner",
	RunE: func(cmd *cobra.Command, args []string) error {
		if ver {
			version.Print()
			return nil
		}
		return cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	usage := fmt.Sprintf("%s version", rootCmd.Use)
	rootCmd.Flags().BoolVarP(&ver, "version", "v", false, usage)
}
