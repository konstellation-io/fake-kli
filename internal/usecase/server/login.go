/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverLoginCmd represents the serverLogin command
var serverLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the given server",
	Run: func(cmd *cobra.Command, args []string) {

		remote, _ := cmd.Flags().GetString("remote")
		fmt.Printf("Correctly logged in to %q", remote)
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverLoginCmd)

	// Add flags to the subcommand
	serverLoginCmd.Flags().String("remote", "", "Server name to login to")
	serverLoginCmd.Flags().String("user", "", "Username")
	serverLoginCmd.Flags().String("password", "", "Password")

	// Mark flags as required
	serverLoginCmd.MarkFlagRequired("remote")
	serverLoginCmd.MarkFlagRequired("user")
	serverLoginCmd.MarkFlagRequired("password")
}
