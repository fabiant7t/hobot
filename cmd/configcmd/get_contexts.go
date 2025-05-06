package configcmd

import (
	"fmt"

	"github.com/fabiant7t/hobot/internal/configfile"
	"github.com/spf13/cobra"
)

var getContextsCommand = &cobra.Command{
	Use:     "get-contexts",
	Short:   "Describe one or more contexts",
	Long:    "Describe one or more contexts",
	Example: "hobot config get-contexts",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := cmd.Flags().GetString("config")
		if err != nil {
			cobra.CheckErr(fmt.Errorf("error: cannot get flag config: %w", err))
		}
		contexts, err := configfile.GetContexts(config)
		cobra.CheckErr(err)
		for _, contextName := range contexts {
			fmt.Println(contextName)
		}
	},
}
