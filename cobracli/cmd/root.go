package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "app",
	// disable the default completion cmd
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func init() {
	RootCmd.AddCommand(cmdShowVersion)
	RootCmd.AddCommand(cmdTry)
	RootCmd.AddCommand(cmdTry2)
}

func Execute() {
	RootCmd.Execute()
}
