/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cloudflare/pkg/api"
	"fmt"

	"github.com/spf13/cobra"
)

var cfEmail = ""
var cfApiKey = ""
var token = ""

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login into Cloudflare REST API",
	Long: `This command let you authenticate against Cloudflare REST API
	
Examples:

cloudflare login --email <youremail@example.com> --key <your valid api key>`,
	Run: func(cmd *cobra.Command, args []string) {
		if token != "" {
			result := api.VerifyToken(token)
			fmt.Println(result)
		} else {
			if cfEmail == "" || cfApiKey == "" {
				fmt.Println("You must provide both of --email|-e and --key|-k to authenticate against Cloudflare API")
			} else {
				fmt.Println(api.VerifyKeyEmail(cfEmail, cfApiKey))
			}
		}
	},
}

func init() {
	loginCmd.InitDefaultHelpFlag()
	loginCmd.Flags().MarkHidden("help")

	loginCmd.Flags().StringVarP(&cfEmail, "email", "e", "", "your Cloudflare email address(required if --key|-k)")
	loginCmd.Flags().StringVarP(&cfApiKey, "key", "k", "", "your Cloudflare Api key(required if --email|-e)")
	loginCmd.Flags().StringVarP(&token, "token", "t", "", "your Cloudflare Api token")
	rootCmd.AddCommand(loginCmd)
}
