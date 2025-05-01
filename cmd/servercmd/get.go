package servercmd

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/fabiant7t/hobot/internal/configfile"
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
		serverNumber, err := strconv.Atoi(args[0])
		cobra.CheckErr(err)

		contextFlag, err := cmd.Flags().GetString("context")
		cobra.CheckErr(err)
		configFlag, err := cmd.Flags().GetString("config")
		cobra.CheckErr(err)

		credentials, err := configfile.GetCredentials(configFlag, contextFlag)
		cobra.CheckErr(err)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv, err := server.GetServer(ctx, serverNumber, credentials.User, credentials.Password, &http.Client{})
		if err != nil {
			return fmt.Errorf("error getting server: %w", err)
		}

		fmt.Printf(srv.String())
		return nil
	},
}
