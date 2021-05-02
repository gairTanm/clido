package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("do called")
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
