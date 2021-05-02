package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "List all of the tasks not done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			var id, err = strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("%v is not a valid number\n", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
