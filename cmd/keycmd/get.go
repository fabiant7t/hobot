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

func NewGetCommand() *cobra.Command {
	var (
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "get [FINGERPRINT]",
		Short: "Get SSH key",
		Long:  "Get SSH key",
		Args:  cobra.ExactArgs(1),
		Example: strings.Join([]string{
			"hobot key get [FINGERPRINT]",
			"hobot key get [FINGERPRINT] --no-headers",
			"hobot key get [FINGERPRINT] -o table=CreatedAt,Name --no-headers",
			"hobot key get [FINGERPRINT] -o json",
			"hobot key get [FINGERPRINT] -o yaml",
		}, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			fingerprint := args[0]

			gotKey, err := key.GetKey(
				ctx,
				fingerprint,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			cobra.CheckErr(err)

			var p printer.RendererPrinter[*key.Key]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[*key.Key]{}
			case "yaml":
				p = &printer.YAMLPrinter[*key.Key]{}
			default:
				p = &printer.TablePrinter[*key.Key]{WithHeader: !noHeaders}
			}
			if err := p.Print(gotKey, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing key: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format. One of (table, json, yaml). Table also supports selecting custom fields using the syntax `table=Foo,Bar,Baz`.")
	return cmd
}
