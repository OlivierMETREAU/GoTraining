package cmd

import (
	"fmt"
	"strconv"

	"example.com/day01-todo-cli/jsonmanager"
	"example.com/day01-todo-cli/tasks"
	"github.com/spf13/cobra"
)

var progressCmd = &cobra.Command{
	Use:     "progress",
	Aliases: []string{"insertion"},
	Short:   "Change the state of a task to in progress",
	Long:    "Change the state of a task contained in the json file to in progress",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := tasks.New(jsonmanager.New("taskList.json"))
		tasks.ReadFromFile()
		if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err == nil {
				tasks.Progress(id)
			}
		}
		tasks.Save()
		fmt.Println(tasks.List())
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)
}
