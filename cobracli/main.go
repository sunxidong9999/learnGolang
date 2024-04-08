package main

import (
	"clitest/cmd"
)

/*
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
	fmt.Println("in the try() funcion")
	if args[0] == "true" {
		return nil
	} else {
		return fmt.Errorf("an error occurred")
	}
}

var cmdTry = &cobra.Command{
	Use:   "try [true or false]",
	Short: "try function",
	Long:  "this is try() funcion with an error",
	Args:  cobra.MinimumNArgs(1),
	RunE:  tryWithError,
}
*/

func main() {
	/*
		var rootCmd = &cobra.Command{
			Use: "app",
			// disable the default completion cmd
			CompletionOptions: cobra.CompletionOptions{
				DisableDefaultCmd: true,
			},
		}
		rootCmd.AddCommand(cmdShowVersion)
		rootCmd.AddCommand(cmdTry)
		rootCmd.Execute()
	*/
	cmd.Execute()
}
