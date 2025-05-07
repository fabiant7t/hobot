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

func NewFingerprintCommand() *cobra.Command {
	var (
		noHeaders       bool
		outputDirectory string
		outputFormat    string
		comment         string
		passphrase      string
	)
	cmd := &cobra.Command{
		Use:   "fingerprint [NAME]",
		Short: "Get fingerprint of SSH key",
		Long:  "Get fingerprint of SSH key",
		Args:  cobra.ExactArgs(1),
		Example: strings.Join([]string{
			"hobot key fingerprint topfstedt",
			"hobot key fingerprint topfstedt --no-headers",
			"hobot key fingerprint topfstedt -o json",
			"hobot key fingerprint topfstedt -o yaml",
		}, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			name := strings.ReplaceAll(args[0], " ", "-")

			fingerprint, err := key.GetFingerprint(
				ctx,
				name,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			cobra.CheckErr(err)

			type FingerprintResponse struct {
				Fingerprint string `json:"fingerprint" yaml:"fingerprint"`
			}
			fingerprintResponse := FingerprintResponse{Fingerprint: fingerprint}

			var p printer.RendererPrinter[*FingerprintResponse]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[*FingerprintResponse]{}
			case "yaml":
				p = &printer.YAMLPrinter[*FingerprintResponse]{}
			default:
				p = &printer.TablePrinter[*FingerprintResponse]{WithHeader: !noHeaders}
			}
			if err := p.Print(&fingerprintResponse, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing fingerprint: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment to add to the key (optional)")
	cmd.Flags().StringVarP(&outputDirectory, "directory", "d", "", "Output directory where the keys are stored (default is current working directory)")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", `Output format (table, json, yaml).`)
	cmd.Flags().StringVar(&passphrase, "passphrase", "", "Passphrase to encrypt the key with (optional)")
	return cmd
}
