/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package process

import (
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// processUpdateCmd represents the processUpdate command
var processUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a given process",
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The project is not initialized")
			return
		}

		workflowName, _ := cmd.Flags().GetString("workflow")
		processName, _ := cmd.Flags().GetString("name")
		processType, _ := cmd.Flags().GetString("type")
		processBaseImage, _ := cmd.Flags().GetString("img")
		processSourceCode, _ := cmd.Flags().GetString("src")

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

						if processType == "" {
							processType = process.ProcessType
						}

						if processBaseImage == "" {
							processBaseImage = process.Image
						}

						if processSourceCode == "" {
							processSourceCode = process.Src
						}

						workflows[i].Processes[j] = domain.Process{
							Name:        processName,
							ProcessType: processType,
							Image:       processBaseImage,
							Src:         processSourceCode,
						}

						viper.Set("workflows", workflows)

						err = viper.WriteConfig()
						if err != nil {
							fmt.Printf("%s", err)
						}
						fmt.Printf("Process %q updated\n", processName)
						return
					}
				}
			}
			fmt.Printf("The %q process does not exist\n", processName)
			return
		}
		fmt.Printf("The %q workflow does not exist\n", workflowName)
	},
}

func init() {
	// Add subcommand to the parent command
	ProcessCmd.AddCommand(processUpdateCmd)

	// Add flags
	processUpdateCmd.Flags().String("name", "", "Process name")
	processUpdateCmd.Flags().String("type", "", "Process type")
	processUpdateCmd.Flags().String("img", "", "Base image")
	processUpdateCmd.Flags().String("src", "", "Source code")

	// Define required flags
	processUpdateCmd.MarkFlagRequired("name")
}
