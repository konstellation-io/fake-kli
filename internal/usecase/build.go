/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package usecase

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new KRT file from the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		outputPath, _ := cmd.Flags().GetString("out")

		if outputPath == "" {
			// set outputPath to the current path
			outputPath, _ = os.Getwd()
		}

		err := os.MkdirAll(outputPath, 0755)
		if err != nil {
			fmt.Printf("Error creating Krt output folder: %s", err)
			return
		}

		fullpath := fmt.Sprintf("%s/%s", outputPath, "myProject.krt")
		err = os.WriteFile(fullpath, []byte("KRT file contents will be here! :)"), 0660)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			fmt.Printf("Error creating Krt file: %s", err)
			return
		}

		fmt.Println("Krt file created at", fullpath)
	},
}

func init() {
	// Add flags to the subcommand
	BuildCmd.Flags().String("out", "", "Output path for the KRT file")
}
