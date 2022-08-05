/*
Copyright © 2022 2022 Hai.Tran (github.com/epiHATR)

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

var dnsListFlagRecordType = ""
var dnsListFlagZoneId = ""

// dnsCmd represents the dns command
var dnsCmd = &cobra.Command{
	Use:   "dns list",
	Short: "list all DNS records in a Cloudflare zone",
	Long:  text.ZoneDNSListCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		if dnsListFlagZoneId == "" {
			fmt.Fprintln(os.Stderr, "Error: Please specify a Cloudflare zone Id to retrieve data.")
			os.Exit(1)
		}
		response := dns.GetZoneDns(dnsListFlagZoneId, dnsListFlagRecordType)
		if !response.Success {
			fmt.Fprintln(os.Stderr, "Error:", response.Errors[0].Message)
			os.Exit(1)
		}

		switch flagOutput {

		case "json":
			fmt.Println(output.ToPrettyJson(response.Result, flagQuery))

		case "yaml":
			fmt.Println(output.ToPrettyYaml(response.Result, flagQuery))

		default:
			fmt.Println(output.ToPrettyJson(response.Result, flagQuery))
		}
	},
}

func init() {
	zoneCmd.AddCommand(dnsCmd)
	dnsCmd.Flags().StringVarP(&dnsListFlagZoneId, "zone-id", "", "", "cloudlfare zone ID")
	dnsCmd.Flags().StringVarP(&dnsListFlagRecordType, "type", "t", "", "type of record need to be filtered (CNAME, NS, TXT, ...)")

}
