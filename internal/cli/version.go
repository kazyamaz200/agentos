package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of AgentOS",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("AgentOS v0.1.0")
	},
}
