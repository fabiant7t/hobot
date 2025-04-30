package configfile

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/fabiant7t/hobot/pkg/ini"
)

func GetContexts(path string) ([]string, error) {
	contexts := []string{}
	config, err := ini.NewFromFile(path)
	if err != nil {
		return contexts, fmt.Errorf("error: cannot load config file: %w", err)
	}
	for _, contextName := range config.SectionNames() {
		if contextName != "" {
			contexts = append(contexts, contextName)
		}
	}
	return contexts, nil
}

func Create(path string) error {
	dirPath := filepath.Dir(path)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("error: could not create missing config directory %s", dirPath)
	}
	config := ini.New()
	config.DefaultSection().Set("context", "default")
	section := config.Section("default")
	section.Set("user", "your-robot-user")
	section.Set("password", "your-robot-password")
	if err := config.SaveToFile(path); err != nil {
		return fmt.Errorf("error: could not create missing config file %s", path)
	}
	os.Chmod(path, 0600) // ignore error
	return nil
}

func CurrentContext(path string) (string, error) {
	config, err := ini.NewFromFile(path)
	if err != nil {
		return "", fmt.Errorf("error: cannot load config file \"%s\": %w", path, err)
	}
	context := config.DefaultSection().Get("context")
	return context, nil
}
