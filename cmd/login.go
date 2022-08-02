/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"cloudflare/pkg/api"
	"cloudflare/pkg/color"
	"cloudflare/pkg/util"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfEmail = ""
var cfApiKey = ""
var cfToken = ""

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login into Cloudflare REST API",
	Long: `This command let you authenticate against Cloudflare REST API and store credential to local file
	
Examples:

cloudflare login --email <youremail@example.com> --key <your valid api key>
cloudflare login -t <your valid Api Token>

export CF_AUTH_TOKEN=<your valid Cloudflare API token>
cloudflare login`,

	Run: func(cmd *cobra.Command, args []string) {
		// if token flag provided
		if cfToken != "" {
			log.Println(color.Green, "API Token was provided, ignored --email|-e and --key|-k", color.Reset)
			res := api.VerifyToken(cfToken)
			if res.Success {
				_ = util.SetToken(cfToken)
			} else {
				fmt.Println(color.Red, res.Errors[0].Message, color.Reset)
			}
		} else {
			if cfEmail != "" && cfApiKey != "" {
				// if email & key flags provided
				res := api.VerifyKeyEmail(cfEmail, cfApiKey)
				if res.Success {
					_ = util.SetEmailKey(cfEmail, cfApiKey)
				} else {
					fmt.Println(color.Red, res.Errors[0].Message, color.Reset)
				}
			} else if cfEmail == "" && cfApiKey == "" {
				// if both of them was not provided

				//get token from viper config file
				token := viper.GetString("auth.token")
				if token != "" {
					//use token as authenticate key
					log.Println("authenticating using API token using environment variable CF_AUTH_TOKEN")
					res := api.VerifyToken(token)
					if res.Success {
						_ = util.SetToken(token)
					} else {
						fmt.Println(color.Red, res.Errors[0].Message, color.Reset)
					}
				} else {
					email := viper.GetString("auth.email")
					key := viper.GetString("auth.key")

					if email == "" || key == "" {
						fmt.Println(color.Red, "no credentials provided, consider using flags or environment variables", color.Reset)
						fmt.Println("\n\r")
						fmt.Println(color.Green, "Examples", color.Reset)
						fmt.Println(color.Green, "cloudflare login --email user@example.com --key asdf121n10dbm390d0@@123mdk11j133d132", color.Reset)
					} else {
						log.Println("authenticating against using environment variables for CF_AUTH_EMAIL and CF_AUTH_KEY")
						res := api.VerifyKeyEmail(email, key)
						if res.Success {
							_ = util.SetEmailKey(email, key)
						} else {
							fmt.Println(color.Red, res.Errors[0].Message, color.Reset)
						}
					}
				}
			} else {
				fmt.Println(color.Red, "both of --email|-e and --key|-k are required to authenticate with Cloudflare API", color.Reset)
			}
		}
	},
}

func init() {
	loginCmd.InitDefaultHelpFlag()
	loginCmd.Flags().MarkHidden("help")
	loginCmd.Flags().StringVarP(&cfEmail, "email", "e", "", "your Cloudflare email address(required if --key|-k)")
	loginCmd.Flags().StringVarP(&cfApiKey, "key", "k", "", "your Cloudflare Api key(required if --email|-e)")
	loginCmd.Flags().StringVarP(&cfToken, "token", "t", "", "your Cloudflare Api token")

	rootCmd.AddCommand(loginCmd)
}
