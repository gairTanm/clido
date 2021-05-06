package cmd

import (
	"fmt"
	"strconv"

	"clido/db"

	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do <idx>",
	Short: "Mark a task as done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		fmt.Println("--------------------------------------------------------------------------------------------------------")
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println(chalk.Red, "Failed to parse the argument,", arg, chalk.Reset)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println(chalk.Red, "Something went wrong", err, chalk.Reset)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println(chalk.Red, "Invalid task number", id, chalk.Reset)
				continue
			}
			task := tasks[id-1]
			err := db.DeleteTasks(task.Key)
			if err != nil {
				fmt.Printf("%sFailed to mark \"%d\" as completed. Error %s\n occurred%s", chalk.Red, id, err, chalk.Reset)
			} else {
				fmt.Printf("%sMarked \"%s\" as completed.%s\n", chalk.Magenta, task.Value, chalk.Reset)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
