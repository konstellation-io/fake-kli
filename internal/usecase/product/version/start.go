package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VersionStartCmd represents the product version start command
var VersionStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a product version",
	Args:  cobra.ExactArgs(2),
	Run:   start,
}

func start(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]

	fmt.Printf("%q - %q correctly started.\n", product, version)
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionStartCmd)

	// Add flags
	VersionStartCmd.Flags().String("remote", "", "The remote server name")
	VersionStartCmd.Flags().Bool("detached", false, "Execute the start command in detached mode")
}
