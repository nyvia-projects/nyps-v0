/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"

	"github.com/spf13/cobra"
)

// net/pingCmd represents the net/ping command
var PingCmd = &cobra.Command{
	Use:   "ping",
	Short: "pings remote server",
	Long:  `ping from network : `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("net/ping called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// net/pingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// net/pingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
