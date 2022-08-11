/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/config"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	version     string = "v0.1"
	build       string = "0"
	commit      string = "sha"
	releaseDate string = "2022-11-00"
)

var isDebug bool = false
var flagOutput = "json"
var flagQuery = ""

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cloudflare",
	Short: "CLOUDFLARE CLI version " + version,
	Long:  text.RootLongText,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(os.Stderr, text.AdditionalText)
		os.Exit(1)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.DisableSuggestions = false
	rootCmd.PersistentFlags().SortFlags = false

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
	rootCmd.PersistentFlags().StringVarP(&flagQuery, "query", "q", "", "query in result using JMESpath query")
	rootCmd.PersistentFlags().StringVarP(&flagOutput, "output", "o", "", "show output format in json, yaml, table, ...")
	rootCmd.PersistentFlags().BoolVarP(&isDebug, "debug", "", false, "show debugging information in output windows")
	rootCmd.PersistentFlags().BoolP("help", "", false, "show command help for instructions and examples")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if isDebug {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(ioutil.Discard)
	}
	config.LoadConfig()
}
