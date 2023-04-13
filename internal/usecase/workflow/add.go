/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package workflow

import (
	"KliMock/internal/domain"

	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// workflowAddCmd represents the workflowAdd command
var workflowAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new workflow",
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
	workflowAddCmd.Flags().String("name", "", "Workflow name")
	workflowAddCmd.Flags().String("type", "", "Workflow type")

	// Add required flags
	workflowAddCmd.MarkFlagRequired("name")
	workflowAddCmd.MarkFlagRequired("type")
}
