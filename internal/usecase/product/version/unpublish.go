package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VersionUnpublishCmd represents the product version unpublish command
var VersionUnpublishCmd = &cobra.Command{
	Use:   "unpublish",
	Short: "Unpublish a product version",
	Args:  cobra.ExactArgs(2),
	Run:   unpublish,
}

func unpublish(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]

	fmt.Printf("%q - %q correctly unpublished.\n", product, version)
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionUnpublishCmd)

	// Add flags
	VersionUnpublishCmd.Flags().String("remote", "", "The remote server name")
	VersionUnpublishCmd.Flags().Bool("detached", false, "Execute the start command in detached mode")
}
