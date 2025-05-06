package servercmd

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fabiant7t/hobot/internal/server"
	"github.com/spf13/cobra"
)

func NewNumberCommand() *cobra.Command {
	var (
		ip         string
		name       string
		ignoreCase bool
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
			},
			"\n",
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			// At least one of ip and name must be provided
			if ip == "" && name == "" {
				return errors.New("at least one of --ip or --name must be provided")
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
				return fmt.Errorf("error listing servers: %w", err)
			}

			var matchingSrvs []*server.Server
			for _, srv := range srvs {
				if matchServer(srv.ServerName, srv.IPs, name, ip, ignoreCase) {
					matchingSrvs = append(matchingSrvs, srv)
				}
			}
			switch n := len(matchingSrvs); n {
			case 0:
				return errors.New("no matching server found")
			case 1:
				fmt.Println(matchingSrvs[0].ServerNumber)
				return nil
			default:
				results := make([]string, n)
				for i, srv := range matchingSrvs {
					results[i] = fmt.Sprintf("%d", srv.ServerNumber)
				}
				return fmt.Errorf("more than one server matches: %s", strings.Join(results, ", "))
			}
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "Server Name")
	cmd.Flags().StringVar(&ip, "ip", "", "Server IP (can be IPv4 or IPv6)")
	cmd.Flags().BoolVarP(&ignoreCase, "ignore-case", "i", false, "Ignore case when matching (default is false)")
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
