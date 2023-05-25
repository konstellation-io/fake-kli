package plugin

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ListPluginsCmd represents the list plugins command
var ListPluginsCmd = &cobra.Command{
	Use:   "list",
	Short: "List all plugins",
	Args:  cobra.NoArgs,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		fmt.Println("Could not list the plugins, the remote does not exist")
		return
	}

	fmt.Printf("Plugins for remote %q:\n", remote)
	fmt.Println("Workspace", "MlFlow")
}

func init() {
	//Add subcommand to parent command
	PluginCmd.AddCommand(ListPluginsCmd)

	// Add flags
	ListPluginsCmd.Flags().String("remote", "", "The remote server name")
}
