package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
)

// VersionStatusCmd represents the product version status command
var VersionStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Status of a product version",
	Args:  cobra.ExactArgs(2),
	Run:   status,
}

func status(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]

	statuses := []string{"Stopping", "Stopped", "Starting", "Started", "Failed", "Published"}

	fmt.Printf("%q - %q status is: %q.\n", product, version, statuses[rand.Intn(len(statuses))])
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionStatusCmd)

	// Add flags
	VersionStatusCmd.Flags().String("remote", "", "The remote server name")
}
