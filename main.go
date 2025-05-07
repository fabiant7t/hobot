/*
Copyright © 2025 Fabian Topfstedt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/fabiant7t/hobot/cmd"
	"github.com/fabiant7t/hobot/cmd/configcmd"
	"github.com/fabiant7t/hobot/cmd/docscmd"
	"github.com/fabiant7t/hobot/cmd/servercmd"
	"github.com/fabiant7t/hobot/internal/configfile"
	"github.com/fabiant7t/hobot/internal/statefile"
	"github.com/spf13/cobra"
)

var (
	configFile  string
	stateFile   string
	contextName string
)

func main() {
	defaultConfigFile, err := configfile.DefaultLocation()
	cobra.CheckErr(err)
	defaultStateFile, err := statefile.DefaultLocation()
	cobra.CheckErr(err)
	defaultContextName := "default"

	rootCmd := &cobra.Command{
		Use:   "hobot COMMAND SUBCOMMAND [options]",
		Short: "Hetzner Robot API CLI",
		Long:  "A CLI to interact with the Hetzner Robot API - Copyright 2025 Fabian Topfstedt",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// Set user and password as context values for sub commands to use
			configFile, err := cmd.Flags().GetString("config")
			cobra.CheckErr(err)
			contextName, err := cmd.Flags().GetString("context")
			cobra.CheckErr(err)
			credentials, err := configfile.GetCredentials(configFile, contextName)
			cobra.CheckErr(err)
			cmd.SetContext(
				context.WithValue(
					context.WithValue(cmd.Context(), "user", credentials.User),
					"password",
					credentials.Password,
				),
			)
		},
	}
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", fmt.Sprintf(`config file (default is "%s")`, defaultConfigFile))
	rootCmd.PersistentFlags().StringVar(&stateFile, "state", "", fmt.Sprintf(`state file (default is "%s")`, defaultStateFile))
	rootCmd.PersistentFlags().StringVar(&contextName, "context", "", fmt.Sprintf(`default is read from state file or "%s"`, defaultContextName))

	if configFile == "" {
		configFile = defaultConfigFile
		created, err := createIfNotExists(configFile, configfile.Create)
		cobra.CheckErr(err)
		if created {
			os.Stderr.WriteString(fmt.Sprintf("Created \"%s\" for you. Edit the file and enter your real credentials!\n", configFile))
		}
	}
	if stateFile == "" {
		stateFile = defaultStateFile
		_, err := createIfNotExists(stateFile, statefile.Create)
		cobra.CheckErr(err)
	}
	if contextName == "" {
		savedContext, err := statefile.GetContext(stateFile)
		cobra.CheckErr(err)
		if savedContext == "" {
			contextName = defaultContextName
		} else {
			contextName = savedContext
		}
	}

	// Mount commands
	rootCmd.AddCommand(configcmd.New())
	rootCmd.AddCommand(servercmd.New())
	rootCmd.AddCommand(cmd.VersionCmd)
	rootCmd.AddCommand(docscmd.New(rootCmd, "./docs"))

	if err = rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func createIfNotExists(path string, f func(string) error) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := f(path)
			if err != nil {
				return false, err
			}
			return true, nil
		}
	}
	return false, nil
}
