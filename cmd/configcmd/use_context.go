package configcmd

import (
	"errors"
	"fmt"

	"github.com/fabiant7t/hobot/pkg/ini"
	"github.com/spf13/cobra"
)

var useContextCommand = &cobra.Command{
	Use:     "use-context CONTEXT_NAME",
	Aliases: []string{"use"},
	Short:   "Use given context",
	Long:    "Sets the given context in the config file",
	Example: "hobot config use-context private",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("CONTEXT_NAME missing")
		}
		contextName := args[0]
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("error: cannot get flag config: %w", err)
		}
		cfg, err := ini.NewFromFile(config)
		if err != nil {
			return fmt.Errorf("error: cannot load config file: %w", err)
		}
		if !cfg.HasSection(contextName) {
			return fmt.Errorf("error: no context exists with the name: %s", contextName)
		}
		cfg.DefaultSection().Set("context", contextName)
		if err := cfg.SaveToFile(config); err != nil {
			return fmt.Errorf("error: cannot write config file: %w", err)
		}
		fmt.Printf("Switched to context \"%s\".\n", contextName)
		return nil
	},
}
