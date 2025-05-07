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
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewCreateCommand())
	cmd.AddCommand(NewFingerprintCommand())
	return cmd
}
