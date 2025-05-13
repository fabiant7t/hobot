package configcmd

import (
	"fmt"
	"os"

	"github.com/fabiant7t/hobot/internal/configfile"
	"github.com/fabiant7t/hobot/internal/statefile"
	"github.com/spf13/cobra"
)

var useContextCommand = &cobra.Command{
	Use:     "use-context [CONTEXT_NAME]",
	Aliases: []string{"use"},
	Short:   "Use given context",
	Long:    "Sets the given context in the state file",
	Example: "hobot config use-context private",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contextName := args[0]
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			cobra.CheckErr(fmt.Errorf("error: cannot get config flag: %w", err))
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
			cobra.CheckErr(fmt.Errorf(`Context "%s" is not configured in "%s"`, contextName, config))
		}
		state, err := cmd.Flags().GetString("state")
		if err != nil {
			cobra.CheckErr(fmt.Errorf("error: cannot get state flag: %w", err))
		}
		cobra.CheckErr(statefile.SetContext(state, contextName))
		os.Stderr.WriteString(fmt.Sprintf("Switched to context \"%s\".\n", contextName))
	},
}
