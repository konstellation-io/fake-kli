package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VersionStopCmd represents the product version stop command
var VersionStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop a product version",
	Args:  cobra.ExactArgs(2),
	Run:   stop,
}

func stop(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]

	fmt.Printf("%q - %q correctly stopped.\n", product, version)
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionStopCmd)

	// Add flags
	VersionStopCmd.Flags().String("remote", "", "The remote server name")
	VersionStopCmd.Flags().Bool("detached", false, "Execute the start command in detached mode")
}
