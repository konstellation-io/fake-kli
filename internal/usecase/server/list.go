/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"
	"github.com/spf13/cobra"
)

// serverListCmd represents the serverList command
var serverListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all servers",
	Run: func(cmd *cobra.Command, args []string) {
		// list a list of fake servers with their names and urls
		fmt.Printf("Server %q: %q\n", "server1", "https://server1.com")
		fmt.Printf("Server %q: %q\n", "server2", "https://server2.com")
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverListCmd)
}
