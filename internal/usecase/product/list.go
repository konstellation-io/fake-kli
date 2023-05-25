package product

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ListProductCmd represents the list product command
var ListProductCmd = &cobra.Command{
	Use:   "list",
	Short: "List all products",
	Args:  cobra.NoArgs,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		fmt.Println("Could not list the products, the remote does not exist")
		return
	}

	fmt.Printf("Products for remote %q:\n", remote)
	fmt.Println("Product 1", "Product 2", "Product 3")
}

func init() {
	//Add subcommand to parent command
	ProductCmd.AddCommand(ListProductCmd)

	// Add flags
	ListProductCmd.Flags().String("remote", "", "The remote server name")
}
