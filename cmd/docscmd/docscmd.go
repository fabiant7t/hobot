package docscmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

func New(rootCmd *cobra.Command, path string) *cobra.Command {
	return &cobra.Command{
		Use:   "docs",
		Short: "Generate documentation files",
		RunE: func(cmd *cobra.Command, args []string) error {
			if _, err := os.Stat(path); err != nil {
				if os.IsNotExist(err) {
					if err := os.MkdirAll(path, os.ModePerm); err != nil {
						return fmt.Errorf(`error creating directory "%s": %w`, path, err)
					}
				}
			}
			return doc.GenMarkdownTree(rootCmd, path)
		},
	}
}
