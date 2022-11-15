package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(startUpCmd)
	rootCmd.AddCommand(scheduleCmd)
	rootCmd.AddCommand(databaseMigrateCmd)
	rootCmd.AddCommand(maketoken)
	rootCmd.AddCommand(operatorCmd)
	rootCmd.AddCommand(appCmd)
}

func Main() {
	if err := Execute(); err != nil {
		panic(err)
	}
}
