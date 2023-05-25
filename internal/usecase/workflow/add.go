package workflow

import (
	"github.com/konstellation-io/fake-kli/internal/domain"

	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// workflowAddCmd represents the workflowAdd command
var workflowAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new workflow",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		workflowType := args[0]
		workflowName := args[1]

		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The product is not initialized")
			return
		}

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("error reading product metadata: %s", err)
			return
		}

		for _, workflow := range workflows {
			if workflow.Name == workflowName {
				fmt.Printf("The %q workflow already exist\n", workflowName)
				return
			}
		}
		workflows = append(workflows, domain.Workflow{Name: workflowName, WorkflowType: workflowType})
		viper.Set("workflows", workflows)

		err = viper.WriteConfig()
		if err != nil {
			fmt.Printf("%s", err)
		}
		fmt.Printf("Workflow %q successfully added\n", workflowName)
	},
}

func init() {
	// Add subcommand to the parent command
	WorkflowCmd.AddCommand(workflowAddCmd)

	// Add flags
	workflowAddCmd.Flags().String("product-id", "", "Product in where to add the workflow")
}
