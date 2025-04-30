package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func ListServers(ctx context.Context, user, password string, client *http.Client) ([]*Server, error) {
	servers := []*Server{}

	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://robot-ws.your-server.de/server", nil)
	if err != nil {
		return servers, err
	}
	req.SetBasicAuth(user, password)
	res, err := client.Do(req)
	if err != nil {
		return servers, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return servers, fmt.Errorf("error: API did not respond with status 200: %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return servers, err
	}
	var serverListItems []ServerListItem
	err = json.Unmarshal(b, &serverListItems)
	if err != nil {
		return servers, err
	}
	for _, item := range serverListItems {
		servers = append(servers, &item.Server)
	}
	sort.Slice(servers, func(i, j int) bool {
		return servers[i].ServerName < servers[j].ServerName
	})
	return servers, nil
}
