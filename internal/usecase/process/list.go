package process

import (
	"github.com/konstellation-io/fake-kli/internal/domain"

	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// processListCmd represents the processList command
var processListCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all processes for a workflow",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		workflowName, _ := cmd.Flags().GetString("workflow-id")

		if workflowName == "" {
			fmt.Printf("Could not get the processes, the workflow %q does not exist.\n", workflowName)
		}

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The product is not initialized")
			return
		}

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		for _, workflow := range workflows {
			if workflow.Name == workflowName {

				b, err := json.MarshalIndent(workflow.Processes, "", "  ")
				if err != nil {
					fmt.Println("error:", err)
				}
				fmt.Printf(`%s\n`, string(b))
				return
			}

		}
		fmt.Printf("The %q workflow does not exist\n", workflowName)
	},
}

func init() {
	// Add subcommand to the parent command
	ProcessCmd.AddCommand(processListCmd)

	// Add flags
	processListCmd.Flags().String("workflow-id", "", "Workflow from where to list the processes")
	processListCmd.Flags().String("product-id", "", "Product from where to list the processes")

	// Add required flags
	_ = processListCmd.MarkFlagRequired("workflow-id")
}
