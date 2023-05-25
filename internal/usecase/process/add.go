package process

import (
	"fmt"
	"github.com/konstellation-io/fake-kli/internal/domain"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

// processAddCmd represents the processAdd command
var processAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new process",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Println("The product is not initialized")
			return
		}

		processType := args[0]
		processName := args[1]

		workflowName, _ := cmd.Flags().GetString("workflow")
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

				for _, process := range workflow.Processes {
					if process.Name == processName {
						fmt.Printf("The process %q already exists\n", processName)
						return
					}
				}

				workflows[i].Processes = append(workflow.Processes, domain.Process{Name: processName, ProcessType: processType, Image: processBaseImage, Src: processSourceCode})

				viper.Set("workflows", workflows)

				err = viper.WriteConfig()
				if err != nil {
					fmt.Printf("%s", err)
				}
				fmt.Printf("Process %q added to workflow %q\n", processName, workflow.Name)
				return
			}

		}
		fmt.Printf("The %q workflow does not exist\n", workflowName)
	},
}

func init() {
	// Add subcommand to the parent command
	ProcessCmd.AddCommand(processAddCmd)

	// Add flags
	processAddCmd.Flags().String("product-id", "", "The product ID in where the process will be added")
	processAddCmd.Flags().String("img", "", "Base image")
	processAddCmd.Flags().String("src", "", "Source code")

	// Define required flags
	processAddCmd.MarkFlagRequired("img")
	processAddCmd.MarkFlagRequired("src")
}
