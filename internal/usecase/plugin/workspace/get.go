package workspace

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// WorkspaceGetCmd represents the workspace get command
var WorkspaceGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Workspace get command",
	Args:  cobra.ExactArgs(1),
	Run:   getWorkspace,
}

func getWorkspace(cmd *cobra.Command, args []string) {
	productID := args[0]
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		fmt.Println("Could not get the workspace, the remote does not exist")
		return
	}

	if productID == "" {
		fmt.Printf("Could not get the workspace for remote %q: The product does not exist\n", remote)
		return
	}

	yfile, err := os.ReadFile("kubeconf.yaml")

	if err != nil {

		log.Fatal(err)
	}

	data := make(map[interface{}]interface{})

	err2 := yaml.Unmarshal(yfile, &data)
	if err2 != nil {

		log.Fatal(err2)
	}

	fmt.Printf("Workspace correctly retrieved, copy the next configuration to your IDE:"+
		"\n%q\n", data)
}

func init() {
	//Add subcommand to parent command
	WorkspaceCmd.AddCommand(WorkspaceGetCmd)

	// Add flags
	WorkspaceGetCmd.Flags().String("remote", "", "The remote server name")
}
