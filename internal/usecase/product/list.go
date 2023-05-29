package product

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ListProductCmd represents the list product command
var ListProductCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all products",
	Args:  cobra.NoArgs,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		remote = "origin"
	}

	fmt.Printf("Products for remote %q:\n", remote)
	fmt.Println("- Product 1", "\n- Product 2", "\n- Product 3")
}

func init() {
	//Add subcommand to parent command
	ProductCmd.AddCommand(ListProductCmd)

	// Add flags
	ListProductCmd.Flags().String("remote", "", "The remote server name")
}
