package internal

import (
	"KliMock/internal/usecase"
	"KliMock/internal/usecase/process"
	"KliMock/internal/usecase/server"
	"KliMock/internal/usecase/workflow"
	"github.com/spf13/viper"
	"os"

	"github.com/spf13/cobra"
	_ "github.com/spf13/viper"
)

// This variable is modified on build for setting the correct value,
// the defined value is just an example and is not meant to be the current version
var version = "0.0.1"

// Configuration
const configPath = ".kai"
const configName = "krt"
const configType = "yaml"

// rootCmd represents the base command when called without any subcommands
var kliCmd = &cobra.Command{
	Use:     "kli",
	Version: version,
	Short:   "Kli mock instance",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	// Initialize configuration
	initConfig()

	// Execute kliCMD
	err := kliCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig() {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
}

func init() {
	// Add subcommands to the root command
	kliCmd.AddCommand(usecase.InitCmd)
	kliCmd.AddCommand(usecase.BuildCmd)
	kliCmd.AddCommand(workflow.WorkflowCmd)
	kliCmd.AddCommand(process.ProcessCmd)
	kliCmd.AddCommand(server.ServerCmd)
}
