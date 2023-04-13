/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverPushCmd represents the serverPush command
var serverPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Uploads the generated KRT to the server",
	Run: func(cmd *cobra.Command, args []string) {
		remote, _ := cmd.Flags().GetString("remote")
		krtFile, _ := cmd.Flags().GetString("krt-file")

		fmt.Printf("KRT file %q uploaded to server %q", krtFile, remote)
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverPushCmd)

	// Add flags to the subcommand
	serverPushCmd.Flags().String("remote", "", "Server name to login to")
	serverPushCmd.Flags().String("krt-file", "", "Krt file to upload")

	// Mark flags as required
	serverPushCmd.MarkFlagRequired("remote")
	serverPushCmd.MarkFlagRequired("krt-file")
}
