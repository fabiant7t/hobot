package servercmd

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/spf13/cobra"
)

var getCommand = &cobra.Command{
	Use:     "get [SERVER_NUMBER]",
	Short:   "Get server",
	Long:    "Get server",
	Example: "hobot server get 123456",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
		defer cancel()

		serverNumber, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)
		srv, err := server.GetServer(
			ctx,
			serverNumber,
			cmd.Context().Value("user").(string),
			cmd.Context().Value("password").(string),
			&http.Client{},
		)
		if err != nil {
			return fmt.Errorf("error getting server: %w", err)
		}
		fmt.Printf(srv.String())
		return nil
	},
}
