package servercmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/fabiant7t/hobot/pkg/printer"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
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
			var p printer.RendererPrinter[*server.Server]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[*server.Server]{}
			case "yaml":
				p = &printer.YAMLPrinter[*server.Server]{}
			default:
				tp := &printer.TablePrinter[*server.Server]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.PrintAll(servers, os.Stdout); err != nil {
				return fmt.Errorf("error printing all servers: %w", err)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format. One of (table, json, yaml). Table also supports selecting custom fields using the syntax `table=Foo,Bar,Baz`.")
	return cmd
}
