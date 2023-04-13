/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverDeleteCmd represents the serverDelete command
var serverDeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serverDelete called")
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverDeleteCmd)
}
