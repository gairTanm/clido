package cmd

import (
	"fmt"
	"os"
	"time"

	"clido/db"

	"github.com/ttacon/chalk"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		fmt.Println("--------------------------------------------------------------------------------------------------------")
		lateFlag, _ := cmd.Flags().GetBool("late")

		var onTime, late []db.Task
		if err != nil {
			fmt.Println(chalk.Red.Color("Some error occurred"), chalk.Red.Color(err.Error()))
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println(chalk.Cyan, "You have no tasks left, add one, maybe?", chalk.Reset, "😳")
			return
		}
		fmt.Println(chalk.Green.Color("Here are all the tasks still left:"))
		for idx, task := range tasks {
			if task.Start.YearDay() == time.Now().YearDay() {
				onTime = append(onTime, task)
				fmt.Printf("%s%d. %s: %s %d%s\n", chalk.Green, idx+1, task.Value, task.Start.Month(), task.Start.Day(), chalk.Reset)
			} else {
				late = append(late, task)
				fmt.Printf("%s%d. %s: %s %d%s\n", chalk.Red, idx+1, task.Value, task.Start.Month(), task.Start.Day(), chalk.Reset)
			}
		}
		if lateFlag {
			fmt.Println("--------------------------------------------------------------------------------------------------------")
			fmt.Printf("%sHere are all the tasks you're running late on!%s\n", chalk.Red, chalk.Reset)
			for _, lateTask := range late {
				fmt.Printf("%s- %s: %s %d%s\n", chalk.Red, lateTask.Value, lateTask.Start.Month(), lateTask.Start.Day(), chalk.Reset)
			}
		}
		fmt.Println("--------------------------------------------------------------------------------------------------------")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("late", false, "See all the tasks which are running late")
}
