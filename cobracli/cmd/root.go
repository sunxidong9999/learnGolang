package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "app",
	// disable the default completion cmd
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func showVersion(cmd *cobra.Command, args []string) {
	fmt.Println("2024.04.01")
}

var cmdShowVersion = &cobra.Command{
	Use:   "version",
	Short: "Show version",
	Long:  "Show version of the program",
	Args:  cobra.MinimumNArgs(0),
	Run:   showVersion,
}

func tryWithError(cmd *cobra.Command, args []string) error {
	if args[0] == "true" {
		fmt.Println("successed")
		return nil
	} else {
		return fmt.Errorf("an error occurred")
	}
}

var cmdTry = &cobra.Command{
	Use:   "try [true or other]",
	Short: "try function",
	Long:  "this is try() funcion with an error",
	Args:  cobra.MinimumNArgs(1),
	RunE:  tryWithError,
}

func init() {
	RootCmd.AddCommand(cmdShowVersion)
	RootCmd.AddCommand(cmdTry)
}

func Execute() {
	RootCmd.Execute()
}
