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
	Short: "Remove a server from the project",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serverDelete called")
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverDeleteCmd)

	// Add flags to the subcommand
	serverDeleteCmd.Flags().String("remote", "", "Server name to remove")
}
