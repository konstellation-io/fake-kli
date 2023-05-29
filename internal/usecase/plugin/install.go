package plugin

import (
	"fmt"
	"github.com/spf13/cobra"
)

// InstallPluginCmd represents the install plugin command
var InstallPluginCmd = &cobra.Command{
	Use:   "install",
	Short: "Install new plugin",
	Args:  cobra.ExactArgs(1),
	Run:   install,
}

func install(cmd *cobra.Command, args []string) {
	pluginName := args[0]
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		remote = "origin"
	}

	if pluginName != "workspace" && pluginName != "mlflow" {
		fmt.Printf("Could not install the plugin, the plugin <sanitized plugin name> "+
			"is not available for the remote %q\n", remote)
	}

	fmt.Printf("Plugin %q was succesfully installed in remote %q\n", pluginName, remote)
}

func init() {
	//Add subcommand to parent command
	PluginCmd.AddCommand(InstallPluginCmd)

	// Add flags
	InstallPluginCmd.Flags().String("remote", "", "The remote server name")
}
