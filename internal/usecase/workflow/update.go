package workflow

import (
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"
	"strings"

	"github.com/spf13/cobra"
)

// workflowUpdateCmd represents the workflowUpdate command
var workflowUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a given workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		workflowName, _ := cmd.Flags().GetString("name")

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The product is not initialized")
			return
		}

		workflowType, _ := cmd.Flags().GetString("workflow-type")

		if workflowType == "" || (strings.ToLower(workflowType) != "data" &&
			strings.ToLower(workflowType) != "training" && strings.ToLower(workflowType) != "feedback" &&
			strings.ToLower(workflowType) != "serving") {
			fmt.Println("The workflow type is required and must be one of: data, training, feedback or serving")
			return

		}

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("error reading product metadata: %s", err)
			return
		}

		for i, workflow := range workflows {
			if workflow.Name == workflowName {
				if workflowType == "" {
					workflowType = workflow.WorkflowType
				}

				workflows[i] = domain.Workflow{Name: workflowName, WorkflowType: workflowType, Processes: workflow.Processes}

				viper.Set("workflows", workflows)

				err = viper.WriteConfig()
				if err != nil {
					fmt.Printf("%s", err)
				}
				fmt.Printf("Workflow %q successfully updated\n", workflowName)
				return
			}
		}
		fmt.Printf("The %q workflow does not exist\n", workflowName)
	},
}

func init() {
	// Add subcommand to the parent command
	WorkflowCmd.AddCommand(workflowUpdateCmd)

	// Add flags
	workflowUpdateCmd.Flags().String("product-id", "", "Product from where to update the workflow")
	workflowUpdateCmd.Flags().String("workflow-type", "", "Workflow type")
}
