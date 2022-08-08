/*
Copyright © 2022 Hai.Tran (github.com/epiHATR)

*/
package cmd

import (
	"cloudflare/pkg/api/zone/dns"
	"cloudflare/pkg/consts/text"
	"cloudflare/pkg/util/output"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var dnsListCmdZoneId = ""
var dnsListCmdRecordType = ""

// dnsCmd represents the dns command
var dnsListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all DNS records in a Cloudflare zone",
	Long:  text.ZoneDNSListCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if dnsListCmdZoneId == "" {
			fmt.Fprintln(os.Stderr, "Error: Please specify a Cloudflare zone Id to retrieve data.")
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}
		response := dns.GetZoneDns(dnsListCmdZoneId, dnsListCmdRecordType)
		if !response.Success {
			fmt.Fprintln(os.Stderr, "Error:", response.Errors[0].Message)
			os.Exit(1)
		}
		output.PrintOut(response.Result, flagQuery, flagOutput)
	},
}

func init() {
	dnsCmd.AddCommand(dnsListCmd)
	dnsListCmd.Flags().StringVarP(&dnsListCmdZoneId, "zone-id", "", "", "cloudlfare zone ID")
	dnsListCmd.Flags().StringVarP(&dnsListCmdRecordType, "type", "t", "", "type of record need to be filtered (CNAME, NS, TXT, ...)")
}
