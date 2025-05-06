package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func ListServers(ctx context.Context, user, password string, client *http.Client) ([]*Server, error) {
	if client == nil {
		client = &http.Client{}
	}
	servers := []*Server{}
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
	var serverWrappers []ServerWrapper
	err = json.Unmarshal(b, &serverWrappers)
	if err != nil {
		return servers, err
	}
	for _, item := range serverWrappers {
		servers = append(servers, &item.Server)
	}
	sort.Slice(servers, func(i, j int) bool {
		return servers[i].ServerName < servers[j].ServerName
	})
	return servers, nil
}

func GetServer(ctx context.Context, serverNumber int, user, password string, client *http.Client) (*DetailedServer, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://robot-ws.your-server.de/server/%d", serverNumber), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(user, password)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, fmt.Errorf("server %d not found", serverNumber)
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var detailedServerWrapper DetailedServerWrapper
	err = json.Unmarshal(b, &detailedServerWrapper)
	if err != nil {
		return nil, err
	}
	return &detailedServerWrapper.Server, nil
}

func RenameServer(ctx context.Context, serverNumber int, serverName, user, password string, client *http.Client) (*DetailedServer, error) {
	if client == nil {
		client = &http.Client{}
	}
	data := url.Values{}
	data.Set("server_name", serverName)
	data.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://robot-ws.your-server.de/server/%d", serverNumber), strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(user, password)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, fmt.Errorf("server %d not found", serverNumber)
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var detailedServerWrapper DetailedServerWrapper
	err = json.Unmarshal(b, &detailedServerWrapper)
	if err != nil {
		return nil, err
	}
	return &detailedServerWrapper.Server, nil
}
