package commands

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "szczecin",
	Short: "Port scanner",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}
