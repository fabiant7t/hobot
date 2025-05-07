package keycmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fabiant7t/hobot/internal/key"
	"github.com/fabiant7t/hobot/pkg/printer"
	"github.com/spf13/cobra"
)

func NewListCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List SSH keys",
		Long:  "List SSH keys",
		Example: strings.Join([]string{
			"hobot key list",
			"hobot key list --no-headers",
			"hobot key list -o table=Name,Fingerprint,CreatedAt",
			"hobot key list -o json",
			"hobot key list -o yaml",
		}, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			keys, err := key.ListKeys(
				ctx,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("error listing keys: %w", err))
			}
			var p printer.RendererPrinter[*key.Key]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[*key.Key]{}
			case "yaml":
				p = &printer.YAMLPrinter[*key.Key]{}
			default:
				tp := &printer.TablePrinter[*key.Key]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.PrintAll(keys, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing all keys: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format. One of (table, json, yaml). Table also supports selecting custom fields using the syntax `table=Foo,Bar,Baz`.")
	return cmd
}
