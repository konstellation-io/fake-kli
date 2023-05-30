package product

import (
	"fmt"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
)

// CreateProductCmd represents the create product command
var CreateProductCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new AIO product",
	Args:  cobra.ExactArgs(1),
	Run:   exec,
}

func exec(cmd *cobra.Command, args []string) {
	name := args[0]

	fmt.Println("Product successfully created!")

	initLocal, _ := cmd.Flags().GetBool("init-local")

	if initLocal {
		kaiFolderPath, _ := cmd.Flags().GetString("local-path")

		version, _ := cmd.Flags().GetString("version")
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
}

func init() {
	// Add subcommand to parent command
	ProductCmd.AddCommand(CreateProductCmd)

	// Add flags
	CreateProductCmd.Flags().String("version", "v0.0.1", "Set product version")
	CreateProductCmd.Flags().String("description", "", "Set product description")
	CreateProductCmd.Flags().String("remote", "", "The remote server name")
	CreateProductCmd.Flags().Bool("init-local", true, "Sync the created product locally")
	CreateProductCmd.Flags().String("local-path", ".kai", "Local path where to initialize the product")
}
