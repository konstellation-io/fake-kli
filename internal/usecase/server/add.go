/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverAddCmd represents the serverAdd command
var serverAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new server to the project",
	Run: func(cmd *cobra.Command, args []string) {

		url, _ := cmd.Flags().GetString("url")
		remote, _ := cmd.Flags().GetString("remote")

		fmt.Printf("Server with url %q added as %q", url, remote)
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverAddCmd)

	// Add flags to the subcommand
	serverAddCmd.Flags().String("url", "", "Server url")
	serverAddCmd.Flags().String("remote", "", "Server name")

	// Mark flags as required
	serverAddCmd.MarkFlagRequired("url")
	serverAddCmd.MarkFlagRequired("remote")
}
