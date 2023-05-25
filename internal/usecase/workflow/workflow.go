package workflow

import (
	"github.com/spf13/cobra"
)

// WorkflowCmd represents the workflow command
var WorkflowCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Manage your product workflows",
}

func init() {
}
