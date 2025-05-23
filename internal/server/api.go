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

	switch res.StatusCode {
	case http.StatusNotFound:
		return nil, errors.New("no servers found")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
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

func GetResetOption(ctx context.Context, serverNumber int, user, password string, client *http.Client) (*ResetOption, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://robot-ws.your-server.de/reset/%d", serverNumber), nil)
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
		return nil, fmt.Errorf("server %d not found or no reset option available", serverNumber)
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
	var resetOptionWrapper ResetOptionWrapper
	err = json.Unmarshal(b, &resetOptionWrapper)
	if err != nil {
		return nil, err
	}
	return &resetOptionWrapper.ResetOption, nil
}

func ResetServer(ctx context.Context, serverNumber int, resetType string, user, password string, client *http.Client) (*Reset, error) {
	if client == nil {
		client = &http.Client{}
	}
	data := url.Values{}
	data.Set("type", resetType)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://robot-ws.your-server.de/reset/%d", serverNumber), strings.NewReader(data.Encode()))
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
		return nil, fmt.Errorf(`invalid reset type "%s" for server %d`, resetType, serverNumber)
	case http.StatusNotFound:
		return nil, fmt.Errorf("server %d not found or no reset option available", serverNumber)
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusConflict:
		return nil, errors.New("currently performing a manual reset")
	case http.StatusInternalServerError:
		return nil, errors.New("reset failed")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resetWrapper ResetWrapper
	err = json.Unmarshal(b, &resetWrapper)
	if err != nil {
		return nil, err
	}
	return &resetWrapper.Reset, nil
}

func GetRescueOption(ctx context.Context, serverNumber int, user, password string, client *http.Client) (*RescueOption, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://robot-ws.your-server.de/boot/%d/rescue", serverNumber), nil)
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
		return nil, fmt.Errorf("server %d not found or no boot option available", serverNumber)
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
	var rescueOptionWrapper RescueOptionWrapper
	err = json.Unmarshal(b, &rescueOptionWrapper)
	if err != nil {
		return nil, err
	}
	return &rescueOptionWrapper.RescueOption, nil
}

func ActivateRescue(ctx context.Context, serverNumber int, osType string, authorizedKeys []string, keyboardLayout, user, password string, client *http.Client) (*RescueSetting, error) {
	if client == nil {
		client = &http.Client{}
	}
	data := url.Values{}
	data.Set("keyboard", keyboardLayout)
	data.Set("os", osType)
	for _, authorizedKey := range authorizedKeys {
		data.Add("authorized_key[]", authorizedKey)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, fmt.Sprintf("https://robot-ws.your-server.de/boot/%d/rescue", serverNumber), strings.NewReader(data.Encode()))
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
		return nil, fmt.Errorf("invalid input")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusNotFound:
		return nil, fmt.Errorf("server %d not found or no boot option available", serverNumber)
	case http.StatusInternalServerError:
		return nil, errors.New("boot activation failed due to internal error")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rescueSettingWrapper RescueSettingWrapper
	err = json.Unmarshal(b, &rescueSettingWrapper)
	if err != nil {
		return nil, err
	}
	return &rescueSettingWrapper.RescueSettings, nil
}

func RescueStatus(ctx context.Context, serverNumber int, user, password string, client *http.Client) (*RescueSetting, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://robot-ws.your-server.de/boot/%d/rescue", serverNumber), nil)
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
	case http.StatusBadRequest:
		return nil, fmt.Errorf("invalid input")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusNotFound:
		return nil, fmt.Errorf("server %d not found or no boot option available", serverNumber)
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rescueSettingWrapper RescueSettingWrapper
	err = json.Unmarshal(b, &rescueSettingWrapper)
	if err != nil {
		return nil, err
	}
	return &rescueSettingWrapper.RescueSettings, nil
}

func DeactivateRescue(ctx context.Context, serverNumber int, user, password string, client *http.Client) (*RescueSetting, error) {
	if client == nil {
		client = &http.Client{}
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, fmt.Sprintf("https://robot-ws.your-server.de/boot/%d/rescue", serverNumber), nil)
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
	case http.StatusBadRequest:
		return nil, fmt.Errorf("invalid input")
	case http.StatusUnauthorized:
		return nil, errors.New("unauthorized: check your your credentials")
	case http.StatusNotFound:
		return nil, fmt.Errorf("server %d not found or no boot option available", serverNumber)
	case http.StatusInternalServerError:
		return nil, errors.New("boot deactivation failed due to internal server error")
	case http.StatusOK: // happy path, NOOP
	default: // unexpected status code
		return nil, fmt.Errorf("API responded with HTTP status code %d", res.StatusCode)
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var rescueSettingWrapper RescueSettingWrapper
	err = json.Unmarshal(b, &rescueSettingWrapper)
	if err != nil {
		return nil, err
	}
	return &rescueSettingWrapper.RescueSettings, nil
}
