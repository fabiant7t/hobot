package rescuecmd

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

func NewDeactivateCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "deactivate [SERVER_NUMBER]",
		Short: "Deactivate rescue boot",
		Long:  "Deactivate rescue boot",
		Example: strings.Join([]string{
			"hobot server rescue deactivate 123456",
			"hobot server rescue deactivate 123456 -o table=ServerNumber,Active",
			"hobot server rescue deactivate 123456 -o json",
			"hobot server rescue deactivate 123456 -o yaml",
		}, "\n"),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			serverNumber, err := strconv.Atoi(args[0])
			if err != nil {
				cobra.CheckErr(fmt.Errorf("Cannot convert server number %s to int: %w", args[0], err))
			}

			rescueSetting, err := server.DeactivateRescue(
				ctx,
				serverNumber,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("error getting rescue setting: %w", err))
			}

			var p printer.RendererPrinter[server.RescueSetting]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[server.RescueSetting]{}
			case "yaml":
				p = &printer.YAMLPrinter[server.RescueSetting]{}
			default:
				tp := &printer.TablePrinter[server.RescueSetting]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.Print(*rescueSetting, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing rescue status: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format. One of (table, json, yaml). Table also supports selecting custom fields using the syntax `table=Foo,Bar,Baz`.")
	return cmd
}
