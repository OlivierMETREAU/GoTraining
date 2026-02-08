package cmd

import (
	"fmt"
	"strconv"

	"example.com/day01-todo-cli/jsonmanager"
	"example.com/day01-todo-cli/tasks"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"insertion"},
	Short:   "Delete a task from the list",
	Long:    "Delete a task from the list of tasks contained in the json file",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := tasks.New(jsonmanager.New("taskList.json"))
		tasks.ReadFromFile()
		if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err == nil {
				tasks.Delete(id)
			}
		}
		tasks.Save()
		fmt.Println(tasks.List())
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
