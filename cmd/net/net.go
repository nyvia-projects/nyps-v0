/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"github.com/spf13/cobra"
)

// public
// netCmd represents the net command
var NetCmd = &cobra.Command{
	Use:   "net",
	Short: "network commands",
	Long: `net is a group of network commands

text text text`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

}
