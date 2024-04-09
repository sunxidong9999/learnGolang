package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
