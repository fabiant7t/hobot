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

func NewActivateCommand() *cobra.Command {
	var (
		noHeaders      bool
		outputFormat   string
		keyboardLayout string
		authorizedKeys []string
		osType         string
	)
	cmd := &cobra.Command{
		Use:   "activate [SERVER_NUMBER]",
		Short: "List activate options",
		Long:  "List activate options",
		Example: strings.Join([]string{
			"hobot server rescue activate 123456",
			"hobot server rescue activate 123456 -o table=ServerNumber,Active,Password",
			"hobot server rescue activate 123456 -o table=Password --no-headers",
			"hobot server rescue activate 123456 -o json",
			"hobot server rescue activate 123456 -o yaml",
			"hobot server rescue activate 123456 --os lkvm",
			"hobot server rescue activate 123456 --keyboard-layout de",
		}, "\n"),
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			serverNumber, err := strconv.Atoi(args[0])
			if err != nil {
				cobra.CheckErr(fmt.Errorf("Cannot convert server number %s to int: %w", args[0], err))
			}

			rescueActivated, err := server.ActivateRescue(
				ctx,
				serverNumber,
				osType,
				authorizedKeys,
				keyboardLayout,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("error getting rescue option: %w", err))
			}

			var p printer.RendererPrinter[server.RescueActivated]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[server.RescueActivated]{}
			case "yaml":
				p = &printer.YAMLPrinter[server.RescueActivated]{}
			default:
				tp := &printer.TablePrinter[server.RescueActivated]{WithHeader: !noHeaders}
				if after, found := strings.CutPrefix(outputFormat, "table="); found {
					tp.SetFieldNames(strings.Split(after, ","))
				}
				p = tp
			}
			if err := p.Print(*rescueActivated, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing rescue activated data: %w", err))
			}
		},
	}
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format. One of (table, json, yaml). Table also supports selecting custom fields using the syntax `table=Foo,Bar,Baz`.")
	cmd.Flags().StringVarP(&keyboardLayout, "keyboard-layout", "l", "us", "Keyboard layout, like us, de, fr ...")
	cmd.Flags().StringVar(&osType, "os", "linux", "The OS type (like linux, linuxold or vkvm)")
	cmd.Flags().StringSliceVarP(&authorizedKeys, "authorized-key", "k", []string{}, "Fingerprint of an authorized key. Can be passed zero to many times.")
	return cmd
}
