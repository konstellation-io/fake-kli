package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

// VersionPublishCmd represents the product version publish command
var VersionPublishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Publish a product version",
	Args:  cobra.ExactArgs(2),
	Run:   publish,
}

func publish(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]

	url := fmt.Sprintf("https://kai.intelygenz.com/%s/%s", product, version)

	fmt.Printf("%q - %q correctly published on URL %q.\n", product, version, url)
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionPublishCmd)

	// Add flags
	VersionPublishCmd.Flags().String("remote", "", "The remote server name")
	VersionPublishCmd.Flags().Bool("detached", false, "Execute the start command in detached mode")
}
