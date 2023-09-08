package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "unknown"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print hostker-ddns version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hostker-DDNS version: ", version)
	},
}
