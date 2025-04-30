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
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fabiant7t/hobot/cmd"
	"github.com/fabiant7t/hobot/cmd/configcmd"
	"github.com/fabiant7t/hobot/cmd/servercmd"
	"github.com/fabiant7t/hobot/internal/configfile"
	"github.com/spf13/cobra"
)

var (
	config  string
	context string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hobot COMMAND SUBCOMMAND [options]",
		Short: "Hetzner Robot API CLI",
		Long:  "A CLI to interact with the Hetzner Robot API - Copyright 2025 Fabian Topfstedt",
	}
	rootCmd.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.config/hobot/config.ini)")
	rootCmd.PersistentFlags().StringVar(&context, "context", "", "default")

	// Configure config
	if config == "" { // default to $HOME/.config/hobot/config.ini
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		config = filepath.Join(home, ".config", "hobot", "config.ini")
	} else if strings.HasPrefix(config, "~") { // resolve ~
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		config = strings.Replace(config, "~", home, 1)
	}
	_, err := os.Stat(config)
	if err != nil {
		if os.IsNotExist(err) {
			err := configfile.Create(config)
			cobra.CheckErr(err)
			fmt.Println("The config file was missing and got created for you!")
			fmt.Printf("TODO: Edit %s and set your robot credentials!\n", config)
		}
	}

	// Configure context
	if context == "" {
		currentContext, err := configfile.CurrentContext(config)
		cobra.CheckErr(err)
		context = currentContext
	}

	// Mount commands
	rootCmd.AddCommand(configcmd.New())
	rootCmd.AddCommand(servercmd.New())
	rootCmd.AddCommand(cmd.VersionCmd)

	rootCmd.Execute()
}
