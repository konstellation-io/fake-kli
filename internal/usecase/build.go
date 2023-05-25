package usecase

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// BuildCmd represents the build command
var BuildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new KRT file from the configuration",
	Run: func(cmd *cobra.Command, args []string) {
		outputPath, _ := cmd.Flags().GetString("output-dir")
		fileName, _ := cmd.Flags().GetString("name")

		if outputPath == "" {
			// set outputPath to the current path
			outputPath, _ = os.Getwd()
		}

		if fileName == "" {
			fileName = "myProduct"
		}

		fileName = strings.TrimSuffix(fileName, ".krt")

		err := os.MkdirAll(outputPath, 0755)
		if err != nil {
			fmt.Printf("Error creating Krt output folder: %s", err)
			return
		}

		fullpath := fmt.Sprintf("%s/%s.krt", outputPath, fileName)
		err = os.WriteFile(fullpath, []byte("KRT file contents will be here! :)"), 0660)
		if err != nil {
			log.Fatal(err)
		}
		if err != nil {
			fmt.Printf("Error creating Krt file: %s", err)
			return
		}

		fmt.Printf("KRT successfully built at: %q \n", fullpath)
	},
}

func init() {
	// Add flags to the subcommand
	BuildCmd.Flags().String("product", "", "Product name")
	BuildCmd.Flags().String("output-dir", "", "Output directory")
	BuildCmd.Flags().String("name", "", "Name to be given to the KRT file")

	// Add required flags
	BuildCmd.MarkFlagRequired("product")
}
