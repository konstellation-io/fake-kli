package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VersionListCmd represents the product version list command
var VersionListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List product versions",
	Args:  cobra.ExactArgs(1),
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	product := args[0]

	fmt.Printf("Versions of the product %q:\n", product)
	fmt.Println("Version 1 - Started", "Version 2 - Stopped", "Version 3 - Started", "Version 4 - Published")
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionListCmd)

	// Add flags
	VersionListCmd.Flags().String("remote", "", "The remote server name")
}
