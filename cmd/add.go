package cmd

import (
	"fmt"
	"strings"

	"clido/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Some error occurred", err.Error(), "can't continue")
			return
		}
		fmt.Println("Added", strings.Join(args, " "), "to clido list")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
