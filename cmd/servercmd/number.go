package servercmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/fabiant7t/hobot/pkg/printer"
	"github.com/spf13/cobra"
)

func NewNumberCommand() *cobra.Command {
	var (
		ignoreCase   bool
		ip           string
		name         string
		noHeaders    bool
		outputFormat string
	)
	cmd := &cobra.Command{
		Use:   "number",
		Short: "Get server number",
		Long:  "Get server number (the identifier)",
		Example: strings.Join(
			[]string{
				`hobot server number --ip 12.34.56.78`,
				`hobot server number --name marvin`,
				`hobot server number --name macos --ignore-case`,
				`hobot server number --name "Deep Thought"`,
				`hobot server number --name marvin --no-headers`,
				`hobot server number --name marvin -o json`,
				`hobot server number --name marvin -o yaml`,
			},
			"\n",
		),
		Run: func(cmd *cobra.Command, args []string) {
			// At least one of ip and name must be provided
			if ip == "" && name == "" {
				cobra.CheckErr(errors.New("at least one of --ip or --name must be provided"))
			}

			ctx, cancel := context.WithTimeout(cmd.Context(), 30*time.Second)
			defer cancel()

			srvs, err := server.ListServers(
				ctx,
				cmd.Context().Value("user").(string),
				cmd.Context().Value("password").(string),
				&http.Client{},
			)
			if err != nil {
				cobra.CheckErr(fmt.Errorf("error listing servers: %w", err))
			}

			var matchingSrvs []*server.Server
			for _, srv := range srvs {
				if matchServer(srv.ServerName, srv.IPs, name, ip, ignoreCase) {
					matchingSrvs = append(matchingSrvs, srv)
				}
			}
			if n := len(matchingSrvs); n == 0 { // not found
				cobra.CheckErr(errors.New("no matching server found"))
			} else if n > 1 { // ambiguous
				results := make([]string, n)
				for i, srv := range matchingSrvs {
					results[i] = fmt.Sprintf("%d", srv.ServerNumber)
				}
				cobra.CheckErr(fmt.Errorf("more than one server matches: %s", strings.Join(results, ", ")))
			}
			type MatchingServer struct {
				ServerNumber int `json:"server_number"`
			}
			result := &MatchingServer{ServerNumber: matchingSrvs[0].ServerNumber}

			var p printer.RendererPrinter[*MatchingServer]
			switch outputFormat {
			case "json":
				p = &printer.JSONPrinter[*MatchingServer]{}
			case "yaml":
				p = &printer.YAMLPrinter[*MatchingServer]{}
			default:
				p = &printer.TablePrinter[*MatchingServer]{WithHeader: !noHeaders}
			}
			if err := p.Print(result, os.Stdout); err != nil {
				cobra.CheckErr(fmt.Errorf("error printing all servers: %w", err))
			}

		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Server Name")
	cmd.Flags().StringVar(&ip, "ip", "", "Server IP (can be IPv4 or IPv6)")
	cmd.Flags().BoolVarP(&ignoreCase, "ignore-case", "i", false, "Ignore case when matching (default is false)")
	cmd.Flags().StringVarP(&outputFormat, "output", "o", "table", "Output format. One of (table, json, yaml).")
	cmd.Flags().BoolVar(&noHeaders, "no-headers", false, "Do not print headers in the output")
	return cmd
}

func matchServer(srvName string, srvIPs []string, queryName, queryIP string, ignoreCase bool) bool {
	// either name or ip should be set (caller should error out)
	if queryName == "" && queryIP == "" {
		return false
	}
	if ignoreCase {
		srvName, queryName = strings.ToLower(srvName), strings.ToLower(queryName)
	}
	if queryName != "" && srvName != queryName {
		return false
	}
	if queryIP != "" {
		queryIP = strings.ToLower(queryIP)
		ipMatch := false
		for _, srvIP := range srvIPs {
			if strings.ToLower(srvIP) == queryIP {
				ipMatch = true
				break
			}
		}
		if !ipMatch {
			return false
		}
	}
	return true
}
