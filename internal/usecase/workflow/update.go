/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package workflow

import (
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// workflowUpdateCmd represents the workflowUpdate command
var workflowUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a given workflow",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The project is not initialized")
			return
		}

		workflowName, _ := cmd.Flags().GetString("name")
		workflowType, _ := cmd.Flags().GetString("type")

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("error reading project metadata: %s", err)
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
	workflowUpdateCmd.Flags().String("name", "", "Workflow name")
	workflowUpdateCmd.Flags().String("type", "", "Workflow type")

	// Add required flags
	workflowUpdateCmd.MarkFlagRequired("name")
}
