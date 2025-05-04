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
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewGetCommand())
	return cmd
}
