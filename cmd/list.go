package cmd

import (
	"fmt"
	"os"

	"clido/db"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Some error occurred", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks left, add one, maybe? ")
			return
		}
		fmt.Println("Here are all the tasks still left:")
		for idx, task := range tasks {
			fmt.Printf("%d. %s\n", idx+1, string(task.Value))

		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
