package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:  "tests",
	Long: "sslog cmd test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sslog test")
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
