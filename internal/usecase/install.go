package usecase

import (
	"fmt"
	"github.com/spf13/cobra"
)

// InstallCmd fake KLI installation
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Initialize new KAI product",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("KLI correctly installed\n")
	},
}
