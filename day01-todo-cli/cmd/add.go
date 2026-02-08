package cmd

import (
	"fmt"

	"example.com/day01-todo-cli/jsonmanager"
	"example.com/day01-todo-cli/tasks"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"insertion"},
	Short:   "Add a task to the list",
	Long:    "Add a task to the list of tasks contained in the json file",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := tasks.New(jsonmanager.New("taskList.json"))
		tasks.ReadFromFile()
		if len(args) == 1 {
			tasks.Add(args[0])
		}
		tasks.Save()
		fmt.Println(tasks.List())
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
