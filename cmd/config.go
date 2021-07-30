/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/mrgarelli/net/system"
	"github.com/mrgarelli/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "show the config file contents",
	Long:  `show the config file contents`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := system.Cat([]string{viper.ConfigFileUsed()})
		if err != nil {
			return err
		}
		return nil
	},
}

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "set the network interface for the program",
	Long:  `set the network interface for the program`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if interfaces, err := system.Interfaces(); err != nil {
			return err
		} else {
			selection := tui.LaunchSelection(interfaces)
			viper.Set("interface", selection)
			err := viper.WriteConfig()
			if err != nil {
				return err
			}
			fmt.Println("set interface in config file to: ", selection)
		}
		return nil
	},
}

// configCmd represents the config command
var configCmd = func() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "command set for viewing and editing program configuration",
		Long:  `command set for viewing and editing program configuration`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(showCmd)
	cmd.AddCommand(interfaceCmd)
	return cmd
}()

func init() {
	rootCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
