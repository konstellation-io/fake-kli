/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
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
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The project is not initialized")
			return
		}

		workflowName, _ := cmd.Flags().GetString("workflow")

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
}
