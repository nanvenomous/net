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
)

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "show the network interfaces created corresponding to drivers",
	Long: `show the network interfaces created corresponding to drivers
	should also show relevant ip addresses`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if interfaces, err := system.Interfaces(); err != nil {
			return err
		} else {
			selection := tui.LaunchSelection(interfaces)
			fmt.Printf("%q\n", selection)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(interfaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// interfaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// interfaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
