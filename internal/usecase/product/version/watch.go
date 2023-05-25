package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"gopkg.in/loremipsum.v1"
	"math/rand"
	"strings"
)

// VersionWatchCmd represents the product version watch command
var VersionWatchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch product version logs",
	Args:  cobra.ExactArgs(2),
	Run:   watch,
}

func watch(cmd *cobra.Command, args []string) {
	product := args[0]
	version := args[1]
	logLevel, _ := cmd.Flags().GetString("log-level")

	if logLevel == "" {
		levels := []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
		logLevel = levels[rand.Intn(len(levels))]
	}

	logLevel = strings.ToUpper(logLevel)

	loremIpsumGenerator := loremipsum.New()

	for true {
		fmt.Printf("[%s][%s] - %s %s\n", product, version, logLevel, loremIpsumGenerator.Sentence())
	}

}

func init() {
	// Add subcommand to parent command
	VersionCmd.AddCommand(VersionWatchCmd)

	// Add flags
	VersionWatchCmd.Flags().String("remote", "", "The remote server name")
	VersionWatchCmd.Flags().String("log-level", "", "The log-level to use")
}
