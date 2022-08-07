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
	"strings"

	"github.com/spf13/cobra"
)

var dnsShowCmdFlagZoneId = ""
var dnsShowCmdFlagRecordId = ""

// dnsCmd represents the dns command
var dnsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "show cloudflare DNS record details",
	Long:  text.ZoneDnsShowLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {

		errText := []string{}
		if dnsShowCmdFlagZoneId == "" {
			errText = append(errText, "--zone-id")
		}

		if dnsShowCmdFlagRecordId == "" {
			errText = append(errText, "--id|-i")
		}

		if dnsShowCmdFlagRecordId == "" || dnsShowCmdFlagZoneId == "" {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		res := dns.GetZoneDnsDetail(dnsShowCmdFlagZoneId, dnsShowCmdFlagRecordId)
		if !res.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to get DNS record details. The error is", res.Errors[0].Message)
			os.Exit(1)
		} else {
			output.PrintOut(res.Result, flagQuery, flagOutput)
		}
	},
}

func init() {
	dnsCmd.AddCommand(dnsShowCmd)
	dnsShowCmd.Flags().StringVarP(&dnsShowCmdFlagZoneId, "zone-id", "", "", "cloudlfare zone ID")
	dnsShowCmd.Flags().StringVarP(&dnsShowCmdFlagRecordId, "id", "i", "", "Id of the cloudflare DNS record")
}
