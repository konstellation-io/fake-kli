package product

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// SyncProductCmd represents the sync product command
var SyncProductCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new AIO product",
	Args:  cobra.MinimumNArgs(1),
	Run:   sync,
}

func sync(cmd *cobra.Command, args []string) {
	name := args[0]
	kaiFolderPath := ".kai"

	if len(args) == 2 {
		kaiFolderPath = args[1]
	}

	version := "v0.0.1"
	if _, err := os.Stat(kaiFolderPath); !os.IsNotExist(err) {
		fmt.Printf("The product %q is already initialized in the local folder\n", name)
		return
	}

	// Create kai server folder
	err := os.Mkdir(kaiFolderPath, kaiFolderPermissions)
	if err != nil {
		fmt.Printf("The product %q could not be initialized: %s\n", name, err)
		return
	}

	// Create a properties file
	viper.Set("krt_version", 2)
	viper.Set("name", name)
	viper.Set("version", version)

	// Save to file
	err = viper.SafeWriteConfig()
	if err != nil {
		fmt.Printf("Could not create the product: %s\n", err)
		return
	}

	fmt.Println("Product successfully synced locally!")
}

func init() {
	// Add subcommand to parent command
	ProductCmd.AddCommand(SyncProductCmd)
}
