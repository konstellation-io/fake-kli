/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverUpdateCmd represents the serverUpdate command
var serverUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a server",
	Run: func(cmd *cobra.Command, args []string) {
		remote, _ := cmd.Flags().GetString("remote")
		url, _ := cmd.Flags().GetString("url")

		fmt.Printf("Server %q updated with url %q", remote, url)
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverUpdateCmd)

	// Add flags to the subcommand
	serverUpdateCmd.Flags().String("remote", "", "Server name to update")
	serverUpdateCmd.Flags().String("url", "", "Server url")

	// Mark flags as required
	serverUpdateCmd.MarkFlagRequired("remote")
	serverUpdateCmd.MarkFlagRequired("url")
}
