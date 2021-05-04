package cmd

import (
	"github.com/Delta456/box-cli-maker"
	"github.com/spf13/cobra"
)

var RootCmd = cobra.Command{
	Use:   "clido",
	Short: "CLI task manager",
	Run: func(cmd *cobra.Command, args []string) {
		Box := box.New(box.Config{Px: 2, Py: 3, Type: "Round", TitlePos: "Top", Color: "Cyan"})
		Box.Print("clido", "Batteries included task manager for the CLI")
	},
}
