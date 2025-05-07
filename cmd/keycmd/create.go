package keycmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fabiant7t/hobot/internal/key"
	"github.com/fabiant7t/hobot/pkg/keypair"
	"github.com/fabiant7t/hobot/pkg/printer"
	"github.com/spf13/cobra"
)

func NewCreateCommand() *cobra.Command {
	var (
		noHeaders       bool
		outputDirectory string
		outputFormat    string
		comment         string
		passphrase      string
	)
	cmd := &cobra.Command{
		Use:   "create [NAME]",
		Short: "Create SSH key",
		Long:  "Create SSH key",
		Args:  cobra.ExactArgs(1),
		Example: strings.Join([]string{
			"hobot key create topfstedt",
			"hobot key create topfstedt -o table=Name,Fingerprint",
			"hobot key create topfstedt -o table=Fingerprint --no-headers",
			"hobot key create topfstedt -o json",
			"hobot key create topfstedt -o yaml",
		}, "\n"),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			name := strings.ReplaceAll(args[0], " ", "-")

			authKey, privPem, err := keypair.NewEd25519Keypair(comment, passphrase)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("error creating keypair: %w", err))
			}

			privKeyPath := filepath.Join(outputDirectory, fmt.Sprintf("%s.pem", name))
			if _, err := os.Stat(privKeyPath); err == nil {
				cobra.CheckErr(fmt.Errorf(`SSH key "%s" already exists locally`, privKeyPath))
			}
			if err := os.WriteFile(privKeyPath, privPem, 0600); err != nil {
				cobra.CheckErr(fmt.Errorf(`Error writing "%s"`, privKeyPath))
			}
			authKeyPath := filepath.Join(outputDirectory, fmt.Sprintf("%s.pub", name))
			if err := os.WriteFile(authKeyPath, authKey, 0600); err != nil {
				cobra.CheckErr(fmt.Errorf(`Error writing "%s"`, authKeyPath))
			}

			createdKey, err := key.CreateKey(
				ctx,
				name,
				string(authKey),
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
				tp := &printer.TablePrinter[*key.Key]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.Print(createdKey, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing key: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVar(&comment, "comment", "", "Comment to add to the key (optional)")
	cmd.Flags().StringVarP(&outputDirectory, "directory", "d", "", "Output directory where the keys are stored (default is current working directory)")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", `Output format (table, json, yaml). Table also supports selcting custom fields using "table=Foo,Bar,Baz" syntax.`)
	cmd.Flags().StringVar(&passphrase, "passphrase", "", "Passphrase to encrypt the key with (optional)")
	return cmd
}
