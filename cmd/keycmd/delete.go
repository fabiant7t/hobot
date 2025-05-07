package keycmd

import (
	"context"
	"net/http"
	"time"

	"github.com/fabiant7t/hobot/internal/key"
	"github.com/spf13/cobra"
)

func NewDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete [FINGERPRINT]",
		Short:   "Delete SSH key",
		Long:    "Delete SSH key",
		Args:    cobra.ExactArgs(1),
		Example: "hobot key delete [FINGERPRINT]",
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			fingerprint := args[0]

			err := key.DeleteKey(
				ctx,
				fingerprint,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			cobra.CheckErr(err)
		},
	}
	return cmd
}
