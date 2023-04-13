package usecase

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

const kaiFolderPath = ".kai"
const kaiFolderPermissions = 0777

// InitCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize new Konstellation project",
	Run:   exec,
}

func exec(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	version, _ := cmd.Flags().GetString("version")
	if _, err := os.Stat(kaiFolderPath); !os.IsNotExist(err) {
		fmt.Printf("The project %q is already initialized\n", name)
		return
	}

	// Create kai server folder
	err := os.Mkdir(kaiFolderPath, kaiFolderPermissions)
	if err != nil {
		fmt.Printf("The project %q could not be initialized: %s\n", name, err)
		return
	}

	// Create a properties file
	viper.Set("krt_version", 2)
	viper.Set("name", name)
	viper.Set("version", version)

	// Save to file
	err = viper.SafeWriteConfig()
	if err != nil {
		fmt.Printf("The project %q could not be initialized: %s\n", name, err)
		return
	}

	fmt.Printf("Project %q correctly initialized\n", name)
}

func init() {
	// Add flags
	InitCmd.Flags().String("name", "", "Set project name")
	InitCmd.Flags().String("version", "", "Set project version")
	InitCmd.Flags().String("description", "", "Set project description")

	// Add required flags
	InitCmd.MarkFlagRequired("name")
	InitCmd.MarkFlagRequired("version")
}
