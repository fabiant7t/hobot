package servercmd

import (
	"github.com/fabiant7t/hobot/cmd/servercmd/rescuecmd"
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "Manage servers",
		Long:  "Manage servers",
	}
	cmd.AddCommand(NewGetCommand())
	cmd.AddCommand(NewListCommand())
	cmd.AddCommand(NewNumberCommand())
	cmd.AddCommand(NewRenameCommand())
	cmd.AddCommand(NewResetCommand())
	cmd.AddCommand(NewResetOptionsCommand())
	cmd.AddCommand(rescuecmd.New())
	return cmd
}
