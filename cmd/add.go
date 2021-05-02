package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Added", strings.Join(args, " "), "to clido list")
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
