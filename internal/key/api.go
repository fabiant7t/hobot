package key

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sort"
)

func ListKeys(ctx context.Context, user, password string, client *http.Client) ([]*Key, error) {
	if client == nil {
		client = &http.Client{}
	}
	keys := []*Key{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://robot-ws.your-server.de/key", nil)
	if err != nil {
		return keys, err
	}
	req.SetBasicAuth(user, password)
	res, err := client.Do(req)
	if err != nil {
		return keys, err
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New("no keys found")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return keys, err
	}
	var keyWrappers []KeyWrapper
	err = json.Unmarshal(b, &keyWrappers)
	if err != nil {
		return keys, err
	}
	for _, item := range keyWrappers {
		keys = append(keys, &item.Key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Name < keys[j].Name
	})
	return keys, nil
}
