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

func NewResetOptionsCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "resetoptions [SERVER_NUMBER]",
		Short: "Reset options of server",
		Long:  "Reset options of server",
		Example: strings.Join([]string{
			"hobot server resetoptions 123456",
			"hobot server resetoptions 123456 -o table=ServerNumber,Types,OperatingStatus",
			"hobot server resetoptions 123456 -o table=ServerNumber,Types,OperatingStatus --no-headers",
			"hobot server resetoptions 123456 -o json",
			"hobot server resetoptions 123456 -o yaml",
		}, "\n"),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			serverNumber, err := strconv.Atoi(args[0])
			cobra.CheckErr(err)

			resetOption, err := server.GetResetOption(ctx,
				serverNumber,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("error getting reset option: %w", err))
			}
			var p printer.RendererPrinter[server.ResetOption]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[server.ResetOption]{}
			case "yaml":
				p = &printer.YAMLPrinter[server.ResetOption]{}
			default:
				tp := &printer.TablePrinter[server.ResetOption]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.Print(*resetOption, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing reset options: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", `Output format (table, json, yaml). Table also supports selcting custom fields using "table=Foo,Bar,Baz" syntax.`)
	return cmd
}
