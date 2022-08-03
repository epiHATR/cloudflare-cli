/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/util"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.2"
var isDebug bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "CLOUDFLARE CLI version " + version,
	Long: `A compact CLI works with Cloudflare REST API at https://api.cloudflare.com/v4

Contributed at https://github.com/epiHATR/cloudflare-cli
Author: Hai Tran (hidetran@gmail.com)
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CLOUDFLARE CLI version " + version)
		fmt.Println(`A compact CLI works with Cloudflare REST API at https://api.cloudflare.com/v4

Contributed at https://github.com/epiHATR/cloudflare-cli	
Author: Hai Tran (hidetran@gmail.com)

Usages:
	cloudflare version	get cloudflare-cli module version
	cloudflare login	login to Cloudflare REST API

Flags:
	--help	display command help & instructions
		`)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = false
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "d", false, "show debugging information in output windows")
	rootCmd.PersistentFlags().BoolP("help", "h", false, "show command help for instructions and examples")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if isDebug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}
	util.LoadConfig()
}
