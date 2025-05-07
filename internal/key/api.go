package key

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

func CreateKey(ctx context.Context, name, authKey, user, password string, client *http.Client) (*Key, error) {
	if client == nil {
		client = &http.Client{}
	}
	data := url.Values{}
	data.Set("name", name)
	data.Set("data", authKey)
	data.Encode()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://robot-ws.your-server.de/key", strings.NewReader(data.Encode()))
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
	case http.StatusBadRequest:
		return nil, fmt.Errorf("key cannot be created")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusConflict:
		return nil, errors.New("key already exists")
	case http.StatusCreated: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var keyWrapper KeyWrapper
	err = json.Unmarshal(b, &keyWrapper)
	if err != nil {
		return nil, err
	}
	return &keyWrapper.Key, nil
}

func GetFingerprint(ctx context.Context, name, user, password string, client *http.Client) (string, error) {
	if client == nil {
		client = &http.Client{}
	}
	keys, err := ListKeys(ctx, user, password, client)
	if err != nil {
		return "", fmt.Errorf("error listing keys: %w", err)
	}
	var matchingKeys []*Key
	for _, k := range keys {
		if k.Name == name {
			matchingKeys = append(matchingKeys, k)
		}
	}
	if n := len(matchingKeys); n == 0 {
		return "", errors.New("Key not found")
	} else if n > 1 {
		return "", errors.New("More than one key found")
	}
	return matchingKeys[0].Fingerprint, nil
}
