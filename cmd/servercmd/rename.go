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

func NewRenameCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "rename [SERVER_NUMBER] [NAME]",
		Short: "Rename server",
		Long:  "Rename server",
		Example: strings.Join([]string{
			"hobot server rename 123456 marvin",
			`hobot server rename 123456 "Deep Thought"`,
			`hobot server rename 123456 marvin -o json`,
			`hobot server rename 123456 marvin -o yaml`,
		}, "\n"),
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			serverNumber, err := strconv.Atoi(args[0])
			cobra.CheckErr(err)
			serverName := args[1]
			srv, err := server.RenameServer(
				ctx,
				serverNumber,
				serverName,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			cobra.CheckErr(err)

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
				cobra.CheckErr(fmt.Errorf("error printing server: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", `Output format (table, json, yaml). Table also supports selcting custom fields using "table=Foo,Bar,Baz" syntax.`)
	return cmd
}
