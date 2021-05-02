package cmd

import (
	"fmt"
	"strconv"

	"clido/db"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument,", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number", id)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error %s\n occurred", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
