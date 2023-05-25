package workspace

import (
	"github.com/spf13/cobra"
)

// WorkspaceCmd represents the workspace plugin command
var WorkspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "workspace plugin",
}

func init() {
}
