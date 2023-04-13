/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package usecase

import (
	"fmt"

	"github.com/spf13/cobra"
)

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new KRT file from the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("build called")
	},
}

func init() {
}
