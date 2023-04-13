/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package process

import (
	"github.com/spf13/cobra"
)

// ProcessCmd represents the process command
var ProcessCmd = &cobra.Command{
	Use:   "process",
	Short: "Manage workflow processes",
}

func init() {
	// Add flags
	ProcessCmd.PersistentFlags().String("workflow", "", "The workflow to interact with")

	// Define required flags
	ProcessCmd.MarkFlagRequired("workflow")
}
