package rescuecmd

import (
	"github.com/spf13/cobra"
)

func New() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rescue",
		Short: "Manage server rescue",
		Long:  "Manage server rescue",
	}
	cmd.AddCommand(NewActivateCommand())
	cmd.AddCommand(NewOptionsCommand())
	cmd.AddCommand(NewStatusCommand())
	return cmd
}
