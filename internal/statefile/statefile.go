package statefile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fabiant7t/hobot/pkg/ini"
)

func DefaultLocation() (string, error) {
	stateHome := "~/.local/state"
	if xdgStateHome := os.Getenv("XDG_STATE_HOME"); xdgStateHome != "" {
		stateHome = xdgStateHome
	}
	absStateHome, err := abs(stateHome)
	if err != nil {
		return "", err
	}
	return filepath.Join(absStateHome, "hobot", "state.ini"), nil
}

func Create(path string) error {
	dirPath := filepath.Dir(path)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating state directory %s: %w", dirPath, err)
	}
	SetContext(path, "default")
	os.Chmod(path, 0600) // ignore error
	return nil
}

func GetContext(path string) (string, error) {
	state, err := ini.NewFromFile(path)
	if err != nil {
		return "", fmt.Errorf("error loading state file \"%s\": %w", path, err)
	}
	context := state.DefaultSection().Get("context")
	return context, nil
}

func SetContext(path, context string) error {
	state := ini.New()
	state.DefaultSection().Set("context", context)
	if err := state.SaveToFile(path); err != nil {
		return fmt.Errorf("error setting context in state file: %w", err)
	}
	return nil
}

// abs replaces tilde by user home directory and returns an absolute path
func abs(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("error getting user home dir: %w", err)
		}
		path = strings.Replace(path, "~", home, 1)
	}
	return filepath.Abs(path)
}
