package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hellow")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
