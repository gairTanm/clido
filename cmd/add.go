package cmd

import (
	"fmt"
	"strings"

	"clido/db"

	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		priority, _ := cmd.Flags().GetFloat64("priority")
		fmt.Println(priority)
		_, err := db.CreateTask(task, priority)
		if err != nil {
			fmt.Println(chalk.Red, "Some error occurred", err.Error(), "can't continue", chalk.Reset)
			return
		}
		fmt.Printf("%sAdded %s to clido's list, with priority %f%s\n", chalk.Cyan, strings.Join(args, " "), priority, chalk.Reset)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	addCmd.Flags().Float64("priority", 0.0, "Add a priority value, can be negative, to the task, higher the better")
}
