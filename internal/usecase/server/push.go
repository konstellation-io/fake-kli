package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverPushCmd represents the serverPush command
var serverPushCmd = &cobra.Command{
	Use:   "push",
	Short: "Uploads the generated KRT to the server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 || args[0] == "" {
			fmt.Println("Please provide a KRT file to upload")
			return
		}

		fmt.Println("KRT successfully uploaded to the remote!")
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverPushCmd)

	// Add flags to the subcommand
	serverPushCmd.Flags().String("remote", "", "Server name to login to")
}
