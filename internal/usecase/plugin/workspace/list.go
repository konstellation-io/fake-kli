package workspace

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ListCmd represents the workspace list command
var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Workspace list command",
	Args:  cobra.NoArgs,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		fmt.Println("Could not initialize the workspace, the remote does not exist")
		return
	}

	fmt.Printf("Workspace on the remote %q:\n", remote)
	fmt.Println("Workspace 1: Product 1", "Workspace 2: Product 2", "Workspace 3: Product 2")
}

func init() {
	//Add subcommand to parent command
	WorkspaceCmd.AddCommand(ListCmd)

	// Add flags
	ListCmd.Flags().String("remote", "", "The remote server name")
}
