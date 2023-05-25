package process

import (
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// processDeleteCmd represents the processDelete command
var processDeleteCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a process",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The product is not initialized")
			return
		}

		processName := args[0]

		workflowName, _ := cmd.Flags().GetString("workflow")

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		for i, workflow := range workflows {
			if workflow.Name == workflowName {

				for j, process := range workflow.Processes {
					if process.Name == processName {
						workflows[i].Processes = append(workflow.Processes[:j], workflow.Processes[j+1:]...)

						viper.Set("workflows", workflows)

						err = viper.WriteConfig()
						if err != nil {
							fmt.Printf("%s", err)
						}
						fmt.Printf("Process %q removed from workflow %q\n", processName, workflow.Name)
						return
					}
				}

				fmt.Println("Nothing to remove")
				return
			}

		}
		fmt.Printf("The %q workflow does not exist\n", workflowName)
	},
}

func init() {
	// Add subcommand to the parent command
	ProcessCmd.AddCommand(processDeleteCmd)

	// Add flags
	processDeleteCmd.Flags().String("product_id", "", "Product name")
	processDeleteCmd.Flags().String("workflow_id", "", "Workflow name")

	// Define required flags
	processDeleteCmd.MarkFlagRequired("workflow_id")
}
