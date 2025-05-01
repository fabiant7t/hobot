package servercmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:     "list",
	Short:   "List servers",
	Long:    "List servers",
	Example: "hobot server list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
		defer cancel()

		servers, err := server.ListServers(
			ctx,
			cmd.Context().Value("user").(string),
			cmd.Context().Value("password").(string),
			&http.Client{},
		)
		if err != nil {
			return fmt.Errorf("error listing servers: %w", err)
		}
		for _, srv := range servers {
			fmt.Println(srv.String())
		}
		return nil
	},
}
