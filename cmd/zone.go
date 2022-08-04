/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/text"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var output = ""
var query = ""

// zoneCmd represents the zone command
var zoneCmd = &cobra.Command{
	Use:   "zone",
	Short: "manage Cloudflare zones",
	Long:  text.ZoneCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Fprintln(os.Stderr, text.EmptyArgsText+text.SubCmdHelpText)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(zoneCmd)
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "show output format in json, yaml, table, ...")
	rootCmd.PersistentFlags().StringVarP(&query, "query", "q", "", "query in result using JMESpath query")
}
