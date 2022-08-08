/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/consts/text"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns",
	Short: "manage DNS for cloudflare zone",
	Long:  text.ZoneDnsLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) <= 0 {
			fmt.Fprintln(os.Stderr, text.EmptyArgsText+text.SubCmdHelpText)
			os.Exit(1)
		}
		fmt.Print(cmd.Aliases)
	},
}

func init() {
	zoneCmd.AddCommand(dnsCmd)
}
