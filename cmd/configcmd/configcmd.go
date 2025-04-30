package configcmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Configure the context",
		Long:  "Configure the context",
	}
	cmd.AddCommand(currentContextCommand)
	cmd.AddCommand(getContextsCommand)
	cmd.AddCommand(useContextCommand)
	return cmd
}
