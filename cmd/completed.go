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
		if len(completedTasks) == 0 {
			fmt.Printf("%sHey! Get to work today!%s\n", chalk.Magenta, chalk.Reset)
			os.Exit(1)
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------")
		if err != nil {
			fmt.Println(chalk.Red, "Some error occurred while reading the data,", err, chalk.Reset)
			os.Exit(1)
		}
		fmt.Println(chalk.Cyan.Color("Here are the tasks you have completed today:"))
		for _, t := range completedTasks {
			fmt.Printf("%s- %s%s\n", chalk.Cyan, t.Value, chalk.Reset)
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------")
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}
