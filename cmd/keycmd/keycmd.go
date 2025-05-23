package keycmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "key",
		Short: "Manage SSH keys",
		Long:  "Manage SSH keys",
	}
	cmd.AddCommand(NewCreateCommand())
	cmd.AddCommand(NewDeleteCommand())
	cmd.AddCommand(NewFingerprintCommand())
	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewListCommand())
	return cmd
}
