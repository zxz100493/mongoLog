package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "sslog",
	Long: "sslog cmd model",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sslog cli")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
