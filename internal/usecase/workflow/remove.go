package workflow

import (
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// workflowDeleteCmd represents the workflowDelete command
var workflowDeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove workflow",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The product is not initialized")
			return
		}

		workflowName := args[0]

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("error reading product metadata: %s", err)
			return
		}

		for i, workflow := range workflows {
			if workflow.Name == workflowName {
				workflows = append(workflows[:i], workflows[i+1:]...)

				viper.Set("workflows", workflows)

				err = viper.WriteConfig()
				if err != nil {
					fmt.Printf("%s", err)
				}
				fmt.Printf("Workflow %q successfully removed\n", workflowName)
				return
			}
		}
		fmt.Println("Nothing to remove")
	},
}

func init() {
	// Add subcommand to the parent command
	WorkflowCmd.AddCommand(workflowDeleteCmd)

	// Add flags
	workflowDeleteCmd.Flags().String("product-id", "", "Product name")
}
