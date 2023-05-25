package product

import (
	"github.com/konstellation-io/fake-kli/internal/usecase/product/version"
	"github.com/spf13/cobra"
)

// ProductCmd represents the product command
var ProductCmd = &cobra.Command{
	Use:   "product",
	Short: "A brief description of your command",
}

func init() {
	ProductCmd.AddCommand(version.VersionCmd)
}
