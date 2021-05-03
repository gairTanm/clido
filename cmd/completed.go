package cmd

import (
	"fmt"
	"os"

	"clido/db"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List all the tasks completed in the last 24 hours",
	Run: func(cmd *cobra.Command, args []string) {
		completedTasks, err := db.CompletedTasks()
		if err != nil {
			fmt.Println(chalk.Red, "Some error occurred while reading the data,", err, chalk.Reset)
			os.Exit(1)
		}
		fmt.Println(chalk.Cyan.Color("Here are the tasks you have completed:"))
		for _, t := range completedTasks {
			fmt.Printf("%s- %s%s\n", chalk.Cyan, t.Value, chalk.Reset)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
