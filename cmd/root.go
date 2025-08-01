package cmd

import (
	"cli-tool/cmd/sub"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "poc",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(sub.NameCmd)
	rootCmd.AddCommand(sub.CreateCmd)
}
