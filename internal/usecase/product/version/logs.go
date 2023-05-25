package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/loremipsum.v1"
	"math/rand"
	"strings"
)

// VersionLogsCmd represents the product version logs command
var VersionLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "Watch product version logs",
	Args:  cobra.ExactArgs(2),
	Run:   logs,
}

func logs(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]
	logLevel, _ := cmd.Flags().GetString("log-level")

	if logLevel == "" {
		levels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
		logLevel = levels[rand.Intn(len(levels))]
	}

	logLevel = strings.ToUpper(logLevel)

	loremIpsumGenerator := loremipsum.New()

	randLoop := rand.Intn(5) + 2

	for i := 0; i < randLoop; i++ {
		fmt.Printf("[%s][%s] - %s %s\n", product, version, logLevel, loremIpsumGenerator.Sentence())
	}
}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionLogsCmd)

	// Add flags
	VersionLogsCmd.Flags().String("remote", "", "The remote server name")
	VersionLogsCmd.Flags().String("log-level", "", "The log-level to use")
	VersionLogsCmd.Flags().String("from", "", "Log from filter")
	VersionLogsCmd.Flags().String("to", "", "Log to filter")
	VersionLogsCmd.Flags().String("out", "", "Output format")
}
