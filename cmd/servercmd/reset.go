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

func NewResetCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
		resetType    string
	)
	cmd := &cobra.Command{
		Use:   "reset [SERVER_NUMBER]",
		Short: "Reset server",
		Long:  "Reset server",
		Example: strings.Join([]string{
			"hobot server reset 123456",
			`hobot server reset 123456 --type sw`,
			`hobot server reset 123456 --type hw`,
			`hobot server reset 123456 --type man`,
			`hobot server reset 123456 --type power`,
			`hobot server reset 123456 --type power_long`,
		}, "\n"),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			serverNumber, err := strconv.Atoi(args[0])
			cobra.CheckErr(err)
			reset, err := server.ResetServer(
				ctx,
				serverNumber,
				resetType,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			cobra.CheckErr(err)

			var p printer.RendererPrinter[server.Reset]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[server.Reset]{}
			case "yaml":
				p = &printer.YAMLPrinter[server.Reset]{}
			default:
				tp := &printer.TablePrinter[server.Reset]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.Print(*reset, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing reset details: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", `Output format (table, json, yaml). Table also supports selcting custom fields using "table=Foo,Bar,Baz" syntax.`)
	cmd.Flags().StringVarP(&resetType, "type", "t", "hw", `Reset option type.`)
	return cmd
}
