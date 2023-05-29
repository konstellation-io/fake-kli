package runner

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ListRunnerCmd represents the list runners command
var ListRunnerCmd = &cobra.Command{
	Use:   "ls",
	Short: "List all runners",
	Args:  cobra.NoArgs,
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	remote, _ := cmd.Flags().GetString("remote")

	if remote == "" {
		remote = "origin"
	}

	fmt.Printf("Runners for remote %q:\n", remote)
	fmt.Println("- Python3.9 - V1.1", "\n- Python3.9 CUDA - V1.1", "\n- Python3.9 CUDA 11.6.1 - V1.1",
		"\n- Python3.10 CUDA 11.6.1 - Development - V1.1")
}

func init() {
	//Add subcommand to parent command
	RunnerCmd.AddCommand(ListRunnerCmd)

	// Add flags
	ListRunnerCmd.Flags().String("remote", "", "The remote server name")
}
