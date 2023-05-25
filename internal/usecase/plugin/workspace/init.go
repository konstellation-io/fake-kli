package workspace

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// WorkspaceInitCmd represents the workspace init command
var WorkspaceInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Workspace init command",
	Args:  cobra.ExactArgs(1),
	Run:   initWorkspace,
}

func initWorkspace(cmd *cobra.Command, args []string) {
	productID := args[0]
	remote, _ := cmd.Flags().GetString("remote")
	withRunner, _ := cmd.Flags().GetString("with-runner")
	withRepository, _ := cmd.Flags().GetString("with-repository")

	if remote == "" {
		fmt.Println("Could not initialize the workspace, the remote does not exist")
		return
	}

	if productID == "" {
		fmt.Printf("Could not initialize the workspace for remote %q: The product does not exist\n", remote)
		return
	}

	if withRunner == "" || withRepository == "" {
		fmt.Printf("Could not initialize the workspace for remote %q: "+
			"Some required parameters are not give\n", remote)
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

	fmt.Printf("Workspace correctly initialized, copy the next configuration to your IDE:"+
		"\n%q\n", data)
}

func init() {
	//Add subcommand to parent command
	WorkspaceCmd.AddCommand(WorkspaceInitCmd)

	// Add flags
	WorkspaceInitCmd.Flags().String("remote", "", "The remote server name")
	WorkspaceInitCmd.Flags().String("with-runner", "", "The runner to be used in the workspace.")
	WorkspaceInitCmd.Flags().String("with-repository", "", "The repository URL to clone into the workspace.")

	// Add required flags
	WorkspaceInitCmd.MarkFlagRequired("with-runner")
	WorkspaceInitCmd.MarkFlagRequired("with-repository")
}
