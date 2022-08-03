/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api"
	"cloudflare/pkg/text"
	"cloudflare/pkg/util"
	"fmt"
	"log"
	"os"

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
	Long:  text.CmdLoginLongText + text.SubCmdHelpText,

	Run: func(cmd *cobra.Command, args []string) {
		// if token flag provided
		if cfToken != "" {
			log.Println("API Token was provided, ignored --email|-e and --key|-k")
			res := api.VerifyToken(cfToken)
			if res.Success {
				_ = util.SetToken(cfToken)
			} else {
				fmt.Fprintln(os.Stderr, res.Errors[0].Message)
			}
		} else {
			if cfEmail != "" && cfApiKey != "" {
				// if email & key flags provided
				res := api.VerifyKeyEmail(cfEmail, cfApiKey)
				if res.Success {
					_ = util.SetEmailKey(cfEmail, cfApiKey)
				} else {
					fmt.Fprintln(os.Stderr, res.Errors[0].Message)
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
						fmt.Fprintln(os.Stderr, res.Errors[0].Message)
					}
				} else {
					email := viper.GetString("auth.email")
					key := viper.GetString("auth.key")

					if email == "" || key == "" {
						fmt.Fprintln(os.Stderr, text.CmdLoginNoCredentialText, text.SubCmdHelpText)
						os.Exit(1)
					} else {
						log.Println("authenticating against using configuration values for auth.email & auth.key")
						res := api.VerifyKeyEmail(email, key)
						if res.Success {
							_ = util.SetEmailKey(email, key)
						} else {
							fmt.Fprintln(os.Stderr, res.Errors[0].Message)
						}
					}
				}
			} else {
				fmt.Fprintln(os.Stderr, "both of --email|-e and --key|-k are required to authenticate with Cloudflare API")
				os.Exit(1)
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
