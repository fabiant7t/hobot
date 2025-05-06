package servercmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/fabiant7t/hobot/pkg/printer"
	"github.com/spf13/cobra"
)

func NewGetCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "get [SERVER_NUMBER]",
		Short: "Get server",
		Long:  "Get server",
		Example: strings.Join([]string{
			"hobot server get 123456",
			"hobot server get 123456 -o table --no-headers",
			"hobot server get 123456 -o table=ServerNumber,ServerIP,ServerName",
			"hobot server get 123456 -o json",
			"hobot server get 123456 -o yaml",
		}, "\n"),
		Args: cobra.ExactArgs(1),
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
			var p printer.RendererPrinter[server.DetailedServer]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[server.DetailedServer]{}
			case "yaml":
				p = &printer.YAMLPrinter[server.DetailedServer]{}
			default:
				tp := &printer.TablePrinter[server.DetailedServer]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.Print(*srv, os.Stdout); err != nil {
				return fmt.Errorf("error printing server: %w", err)
			}
			return nil
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", `Output format (table, json, yaml). Table also supports selcting custom fields using "table=Foo,Bar,Baz" syntax.`)
	return cmd
}
