package docscmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func New(rootCmd *cobra.Command, path string) *cobra.Command {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(path, os.ModePerm); err != nil {
				cobra.CheckErr(fmt.Errorf(`error creating directory "%s": %w`, path, err))
			}
		}
	}
	return &cobra.Command{
		Use:   "docs",
		Short: "Generate documentation files",
		RunE: func(cmd *cobra.Command, args []string) error {
			return doc.GenMarkdownTree(rootCmd, path)
		},
	}
}
