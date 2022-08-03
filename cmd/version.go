/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionNumber = "0.0.3"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show cloudflare-cli version",
	Long: `Show current version of cloudflare-cli and the latest REST API version of Cloudflare
	`,
	Run: func(cmd *cobra.Command, args []string) {
		shortTag, _ := cmd.Flags().GetBool("short")
		if shortTag {
			fmt.Println(version)
		} else {
			fmt.Println(`cloudflare-cli version ` + version)
			fmt.Println(`Cloudflare REST API v4 at https://api.cloudflare.com/v4`)
		}

	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	//versionCmd.InitDefaultHelpFlag()
	//versionCmd.Flags().MarkHidden("help")
	versionCmd.Flags().BoolP("short", "s", false, "display short description for current cloudflare-cli version")
}
