package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hobot %s\n", Version)
	},
}
