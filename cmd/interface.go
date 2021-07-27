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

	"github.com/jroimartin/gocui"
	"github.com/spf13/cobra"
)

// interfaceCmd represents the interface command
var interfaceCmd = &cobra.Command{
	Use:   "interface",
	Short: "show the network interfaces created corresponding to drivers",
	Long: `show the network interfaces created corresponding to drivers
	should also show relevant ip addresses`,
	RunE: func(cmd *cobra.Command, args []string) error {
		g, err := gocui.NewGui(gocui.OutputNormal)
		if err != nil {
			return err
		}
		defer g.Close()

		g.SetManagerFunc(layout)

		if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
			return err
		}

		if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
			return err
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

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Hello world!")
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
