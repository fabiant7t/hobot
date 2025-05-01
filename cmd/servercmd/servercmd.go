package servercmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Manage servers",
		Long:  "Manage servers",
	}
	cmd.AddCommand(listCommand)
	cmd.AddCommand(getCommand)
	return cmd
}
