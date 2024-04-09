package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	haveTry  *bool
	haveTry2 *bool
)

func tryWithError(cmd *cobra.Command, args []string) error {
	if args[0] == "true" {
		if *haveTry == true {
			fmt.Println("successed")
		} else {
			fmt.Println("failed")
		}
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

func try2WithError(cmd *cobra.Command, args []string) error {
	if *haveTry2 == true {
		fmt.Println("successed")
	} else {
		fmt.Println("failed")
	}
	return nil
}

var cmdTry2 = &cobra.Command{
	Use:   "try2",
	Short: "try2 function",
	Long:  "this is try2() funcion with an error",
	Args:  cobra.MinimumNArgs(0),
	RunE:  try2WithError,
}

func init() {
	haveTry = cmdTry.Flags().BoolP("true", "t", false, "try usage")
	haveTry2 = cmdTry2.Flags().BoolP("true", "t", false, "try2 usage")
}
