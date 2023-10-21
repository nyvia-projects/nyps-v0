/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package net

import (
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

var (
	urlPath string
	client  = http.Client{
		Timeout: time.Second * 2,
	}
)

func ping(domain string) (int, error) {
	url := domain
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return 0, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	resp.Body.Close()
	return resp.StatusCode, nil
}

var pingCmd = &cobra.Command{
	Use:   "ping",
	Short: "pings remote server",
	Long:  `ping from network : `,
	Run: func(cmd *cobra.Command, args []string) {

		if resp, err := ping(urlPath); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(resp)
		}
	},
}

func init() {
	pingCmd.Flags().StringVarP(&urlPath, "url", "u", "", "url to ping")
	if error := pingCmd.MarkFlagRequired("url"); error != nil {
		fmt.Println(error)
	}
	NetCmd.AddCommand(pingCmd)

}
