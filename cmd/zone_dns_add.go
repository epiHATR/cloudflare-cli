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

var dnsAddCmdZoneId = ""
var dnsAddCmdData = ""

// dnsCmd represents the dns command
var dnsAdd = &cobra.Command{
	Use:   "add",
	Short: "add a DNS record to a Cloudflare zone",
	Long:  text.ZoneDNSAddCmdLongText + text.SubCmdHelpText,
	Run: func(cmd *cobra.Command, args []string) {
		errText := []string{}
		if dnsAddCmdZoneId == "" {
			errText = append(errText, "--zone-id")
		}

		if dnsAddCmdData == "" {
			errText = append(errText, "--data|-d")
		}

		if len(errText) > 0 {
			fmt.Fprintln(os.Stderr, "Error: Missing mandatory inputs, following flags are required: ", strings.Join(errText, ", "))
			fmt.Fprintln(os.Stderr, text.SubCmdHelpText)
			os.Exit(1)
		}

		response := dns.AddDNSRecord(dnsAddCmdZoneId, dnsAddCmdData)
		if !response.Success {
			fmt.Fprintln(os.Stderr, "Error: failed to add DNS record. The error is", response.Errors[0].Message)
			os.Exit(1)
		} else {
			output.PrintOut(response.Result, flagQuery, flagOutput)
		}

	},
}

func init() {
	dnsCmd.AddCommand(dnsAdd)
	dnsAdd.Flags().StringVarP(&dnsAddCmdZoneId, "zone-id", "", "", "cloudlfare zone ID")
	dnsAdd.Flags().StringVarP(&dnsAddCmdData, "data", "d", "", "JSON format of a DNS record.")
}
