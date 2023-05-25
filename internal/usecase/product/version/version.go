package version

import (
	"github.com/spf13/cobra"
)

// VersionCmd represents the product version command
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Product version management",
}

func init() {
}
