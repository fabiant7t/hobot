package configfile

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fabiant7t/hobot/pkg/ini"
)

type Credentials struct {
	User     string
	Password string
}

func DefaultLocation() (string, error) {
	confHome := "~/.config"
	if xdgConfHome := os.Getenv("XDG_CONFIG_HOME"); xdgConfHome != "" {
		confHome = xdgConfHome
	}
	absConfHome, err := abs(confHome)
	if err != nil {
		return "", err
	}
	return filepath.Join(absConfHome, "hobot", "config.ini"), nil
}

func Create(path string) error {
	dirPath := filepath.Dir(path)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("error creating missing config directory %s: %w", dirPath, err)
	}
	if err := SetCredentials(path, "default", Credentials{
		User:     "your-robot-user",
		Password: "your-robot-password",
	}); err != nil {
		return err
	}
	os.Chmod(path, 0600) // ignore error
	return nil
}

func GetContexts(path string) ([]string, error) {
	contexts := []string{}
	config, err := ini.NewFromFile(path)
	if err != nil {
		return contexts, fmt.Errorf("error loading config file: %w", err)
	}
	for _, contextName := range config.SectionNames() {
		if contextName != "" {
			contexts = append(contexts, contextName)
		}
	}
	return contexts, nil
}

func SetCredentials(path, context string, credentials Credentials) error {
	config, err := ini.NewFromFile(path)
	if err != nil {
		config = ini.New()
	}
	section := config.Section(context)
	section.Set("user", credentials.User)
	section.Set("password", credentials.Password)
	if err := config.SaveToFile(path); err != nil {
		return fmt.Errorf("error saving config file %s: %w", path, err)
	}
	return nil
}

func GetCredentials(path, context string) (*Credentials, error) {
	config, err := ini.NewFromFile(path)
	if err != nil {
		return nil, fmt.Errorf("error loading config file %s: %w", path, err)
	}
	section := config.Section(context)
	return &Credentials{
		User:     section.Get("user"),
		Password: section.Get("password"),
	}, nil
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
