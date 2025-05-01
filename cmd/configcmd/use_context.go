package configcmd

import (
	"errors"
	"fmt"

	"github.com/fabiant7t/hobot/internal/configfile"
	"github.com/fabiant7t/hobot/internal/statefile"
	"github.com/spf13/cobra"
)

var useContextCommand = &cobra.Command{
	Use:     "use-context CONTEXT_NAME",
	Aliases: []string{"use"},
	Short:   "Use given context",
	Long:    "Sets the given context in the state file",
	Example: "hobot config use-context private",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("CONTEXT_NAME missing")
		}
		contextName := args[0]
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			return fmt.Errorf("error: cannot get state flag: %w", err)
		}
		configuredContexts, err := configfile.GetContexts(config)
		cobra.CheckErr(err)
		isValid := false
		for _, cn := range configuredContexts {
			if cn == contextName {
				isValid = true
			}
		}
		if !isValid {
			return fmt.Errorf(`Context "%s" is not configured in "%s"`, contextName, config)
		}
		state, err := cmd.Flags().GetString("state")
		if err != nil {
			return fmt.Errorf("error: cannot get state flag: %w", err)
		}
		err = statefile.SetContext(state, contextName)
		cobra.CheckErr(err)
		fmt.Printf("Switched to context \"%s\".\n", contextName)
		return nil
	},
}
