package cmd

import (
	"fmt"

	"example.com/day01-todo-cli/jsonmanager"
	"example.com/day01-todo-cli/tasks"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"insertion"},
	Short:   "List the tasks",
	Long:    "List the tasks listed in the json file.",
	Args:    cobra.MinimumNArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := tasks.New(jsonmanager.New("taskList.json"))
		tasks.ReadFromFile()
		fmt.Println(tasks.List())
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
