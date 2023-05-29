package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serverAddCmd represents the serverAdd command
var serverAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new server to the product",
	Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	// ValidArgs: []string{"url", "remote"},
	Run: func(cmd *cobra.Command, args []string) {

		remote := args[0] // remote
		url := args[1]    // url

		fmt.Printf("Server %q added successfully as %q.\n", url, remote)
	},
}

func init() {
	// Add subcommand to the parent command
	ServerCmd.AddCommand(serverAddCmd)

	// Add flags to the subcommand
	serverAddCmd.Flags().Bool("default", false, "Define current server as the default server")
}
