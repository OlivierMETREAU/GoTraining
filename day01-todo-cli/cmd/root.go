package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tasks",
	Short: "tasks is a cli tool for managing tasks",
	Long:  "tasks is a cli tool for managing tasks with operations such as - add, delete, list, show, search.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing tasks '%s'\n", err)
		os.Exit(1)
	}
}
