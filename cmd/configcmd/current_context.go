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
	RunE: func(cmd *cobra.Command, args []string) error {
		context, err := cmd.Flags().GetString("context")
		if err != nil {
			return err
		}
		fmt.Println(context)
		return nil
	},
}
