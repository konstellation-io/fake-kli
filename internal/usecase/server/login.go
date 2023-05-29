package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverLoginCmd represents the serverLogin command
var serverLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to the given server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		remote := args[0] // remote

		if remote == "" {
			fmt.Println("There was an error trying to authenticate to the server: The remote does not exist")
			return
		}

		user, _ := cmd.Flags().GetString("user")
		password, _ := cmd.Flags().GetString("password")
		token, _ := cmd.Flags().GetString("token")

		if user == "" && password == "" && token != "" {
			fmt.Println("Successfully logged in!")
			return
		}

		if user != "" && password != "" && token == "" {
			fmt.Println("Successfully logged in!")
			return
		}

		fmt.Println("Unauthorized!")
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverLoginCmd)

	// Add flags to the subcommand
	serverLoginCmd.Flags().String("user", "", "Username")
	serverLoginCmd.Flags().String("password", "", "Password")
	serverLoginCmd.Flags().String("token", "", "Session token")
}
