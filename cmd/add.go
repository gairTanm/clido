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
		fmt.Println("--------------------------------------------------------------------------------------------------------")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println(chalk.Red, "Some error occurred", err.Error(), "can't continue", chalk.Reset)
			return
		}
		fmt.Println(chalk.Cyan, "Added", strings.Join(args, " "), "to clido's list", chalk.Reset)
		fmt.Println("--------------------------------------------------------------------------------------------------------")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
