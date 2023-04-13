/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package workflow

import (
	"encoding/json"
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// workflowListCmd represents the workflowList command
var workflowListCmd = &cobra.Command{
	Use:   "ls",
	Short: "list all workflows",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The project is not initialized")
			return
		}

		var workflows []domain.Workflow

		err = viper.UnmarshalKey("workflows", &workflows)
		if err != nil {
			fmt.Printf("%s", err)
			return
		}

		b, err := json.MarshalIndent(workflows, "", "  ")

		if err != nil {
			fmt.Println("error:", err)
		}
		fmt.Printf(`%s\n`, string(b))
		return
	},
}

func init() {
	// Add subcommand to the parent command
	WorkflowCmd.AddCommand(workflowListCmd)
}
