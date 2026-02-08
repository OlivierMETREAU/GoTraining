package cmd

import (
	"fmt"
	"strconv"

	"example.com/day01-todo-cli/jsonmanager"
	"example.com/day01-todo-cli/tasks"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"insertion"},
	Short:   "Change the state of a task to Done",
	Long:    "Change the state of a task contained in the json file to Done",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		tasks := tasks.New(jsonmanager.New("taskList.json"))
		tasks.ReadFromFile()
		if len(args) == 1 {
			id, err := strconv.Atoi(args[0])
			if err == nil {
				tasks.Done(id)
			}
		}
		tasks.Save()
		fmt.Println(tasks.List())
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
