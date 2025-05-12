package configcmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var currentContextCommand = &cobra.Command{
	Use:     "current-context",
	Short:   "Display the current-context",
	Long:    "Display the current-context",
	Example: "hobot config current-context",
	Run: func(cmd *cobra.Command, args []string) {
		context, err := cmd.Flags().GetString("context")
		if err != nil {
			cobra.CheckErr(fmt.Errorf("error: cannot get context flag: %w", err))
		}
		fmt.Println(context)
	},
}
