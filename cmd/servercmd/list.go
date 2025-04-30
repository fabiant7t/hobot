package servercmd

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/fabiant7t/hobot/pkg/ini"
	"github.com/spf13/cobra"
)

var listCommand = &cobra.Command{
	Use:     "list",
	Short:   "List servers",
	Long:    "List servers",
	Example: "hobot server list",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		context, err := cmd.Flags().GetString("context")
		if err != nil {
			return fmt.Errorf("error getting context: %w", err)
		}
		configPath, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("error getting config path: %w", err)
		}
		config, err := ini.NewFromFile(configPath)
		if err != nil {
			return fmt.Errorf("error parsing config file: %w", err)
		}
		section := config.Section(context)
		user := section.Get("user")
		password := section.Get("password")

		servers, err := server.ListServers(ctx, user, password, &http.Client{})
		if err != nil {
			return fmt.Errorf("error listing servers: %w", err)
		}

		for _, srv := range servers {
			name := srv.ServerName
			if name == "" {
				name = "[unnamed]"
			}
			fmt.Printf("%-41s %-17s %-10s %-8s\n", name, srv.ServerIP, srv.DC, fmt.Sprintf("%d", srv.ServerNumber))
		}
		return nil
	},
}
