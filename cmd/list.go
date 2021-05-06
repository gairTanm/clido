package cmd

import (
	"fmt"
	"os"
	"sort"
	"time"

	"clido/db"

	"github.com/jedib0t/go-pretty/table"

	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		lateFlag, _ := cmd.Flags().GetBool("late")

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)

		var late []db.Task
		if err != nil {
			fmt.Println(chalk.Red.Color("Some error occurred"), chalk.Red.Color(err.Error()))
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println(chalk.Cyan, "You have no tasks left, add one, maybe?", chalk.Reset, "😳")
			return
		}
		sort.Sort(db.ByPriority(tasks))
		for _, task := range tasks {
			if task.Start.YearDay() != time.Now().YearDay() {
				late = append(late, task)
			}
		}
		if lateFlag {
			fmt.Printf("%sHere are all the tasks you're running late on!%s\n", chalk.Red, chalk.Reset)
			t.AppendHeader(table.Row{"#", "Task Name", "Due Date"})
			for _, lateTask := range late {
				t.AppendRow(table.Row{"-", lateTask.Value, fmt.Sprintf("%s %d", lateTask.Start.Month(), lateTask.Start.Day())})
				//fmt.Printf("%s- %s\t %s %d%s\n", chalk.Red, lateTask.Value, lateTask.Start.Month(), lateTask.Start.Day(), chalk.Reset)
			}
		} else {
			fmt.Println(chalk.Green.Color("Here are all the tasks still left:"))
			t.AppendHeader(table.Row{"#", "Task Name", "Priority", "Due Date"})
			for idx, task := range tasks {
				t.AppendRow(table.Row{idx + 1, task.Value, task.Priority, fmt.Sprintf("%s %d", task.Start.Month(), task.Start.Day())})
				//fmt.Printf("%s%d.\t%s\t\t%f\t %s %d%s\n", chalk.Green, idx+1, task.Value, task.Priority, task.Start.Month(), task.Start.Day(), chalk.Reset)
			}
		}
		t.SetStyle(table.StyleColoredBright)
		t.Render()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
	listCmd.Flags().Bool("late", false, "See all the tasks which are running late")
}
